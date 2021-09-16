package kasapp

import (
	"context"

	"github.com/getsentry/sentry-go"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/modshared"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/logz"
	"gitlab.com/gitlab-org/labkit/correlation"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverRpcApi struct {
	modshared.RpcApiStub
	Hub SentryHub
}

func (a *serverRpcApi) HandleProcessingError(log *zap.Logger, agentId int64, msg string, err error) {
	handleProcessingError(a.StreamCtx, a.Hub, log, agentId, msg, err)
}

func (a *serverRpcApi) HandleSendError(log *zap.Logger, msg string, err error) error {
	// The problem is almost certainly with the client's connection.
	// Still log it on Debug.
	log.Debug(msg, logz.Error(err))
	c := codes.Canceled
	if grpctool.RequestTimedOut(err) {
		c = codes.DeadlineExceeded
	}
	return status.Error(c, "gRPC send failed")
}

type serverRpcApiFactory struct {
	log       *zap.Logger
	sentryHub *sentry.Hub
}

func (f *serverRpcApiFactory) New(ctx context.Context, fullMethodName string) modserver.RpcApi {
	service, method := grpctool.SplitGrpcMethod(fullMethodName)
	correlationId := correlation.ExtractFromContext(ctx)
	return &serverRpcApi{
		RpcApiStub: modshared.RpcApiStub{
			Logger: f.log.With(
				logz.CorrelationId(correlationId),
				logz.GrpcService(service),
				logz.GrpcMethod(method),
			),
			StreamCtx: ctx,
		},
		Hub: f.hub(service, method, correlationId),
	}
}

func (f *serverRpcApiFactory) hub(service, method, correlationId string) SentryHub {
	hub := f.sentryHub.Clone()
	scope := hub.Scope()
	scope.SetTag(modserver.GrpcServiceSentryField, service)
	scope.SetTag(modserver.GrpcMethodSentryField, method)
	transaction := service + "::" + method              // Like in Gitaly
	scope.SetTransaction(transaction)                   // Like in Gitaly
	scope.SetFingerprint([]string{"grpc", transaction}) // Like in Gitaly
	if correlationId != "" {
		scope.SetTag(modserver.CorrelationIdSentryField, correlationId)
	}
	return hub
}
