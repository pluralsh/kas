package kasapp

import (
	"io"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/logz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/prototool"
	"go.uber.org/zap"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (r *router) RouteToAgentStreamHandler(srv interface{}, stream grpc.ServerStream) error {
	ctx := stream.Context()
	md, _ := metadata.FromIncomingContext(ctx)
	agentId, err := agentIdFromMeta(md)
	if err != nil {
		return err
	}
	sts := grpc.ServerTransportStreamFromContext(ctx)
	service, method := grpctool.SplitGrpcMethod(sts.Method())
	wrappedStream := grpc_middleware.WrapServerStream(stream)
	// Overwrite incoming MD with sanitized MD
	wrappedStream.WrappedContext = metadata.NewIncomingContext(
		wrappedStream.WrappedContext,
		removeHopMeta(md),
	)
	stream = wrappedStream
	rpcApi := modserver.RpcApiFromContext(ctx)
	log := rpcApi.Log().With(logz.AgentId(agentId))
	tunnelFound, findHandle := r.tunnelFinder.FindTunnel(agentId, service, method)
	defer findHandle.Done()
	if !tunnelFound && isNoTunnelSupported(md) {
		err = stream.SendMsg(&GatewayKasResponse{
			Msg: &GatewayKasResponse_NoTunnel_{
				NoTunnel: &GatewayKasResponse_NoTunnel{},
			},
		})
		if err != nil {
			return rpcApi.HandleIoError(log, "SendMsg(GatewayKasResponse_NoTunnel) failed", err)
		}
	}
	tunnel, err := findHandle.Get(wrappedStream.WrappedContext)
	if err != nil {
		return err
	}
	defer tunnel.Done()
	err = stream.SendMsg(&GatewayKasResponse{
		Msg: &GatewayKasResponse_TunnelReady_{
			TunnelReady: &GatewayKasResponse_TunnelReady{},
		},
	})
	if err != nil {
		return rpcApi.HandleIoError(log, "SendMsg(GatewayKasResponse_TunnelReady) failed", err)
	}
	var start StartStreaming
	err = stream.RecvMsg(&start)
	if err != nil {
		if err == io.EOF { // nolint:errorlint
			// Routing kas decided not to proceed
			return nil
		}
		return err
	}
	return tunnel.ForwardStream(log, rpcApi, stream, newWrappingCallback(log, rpcApi, stream))
}

func isNoTunnelSupported(md metadata.MD) bool {
	return len(md.Get(modserver.RoutingFeatureNoTunnel)) > 0
}

func removeHopMeta(md metadata.MD) metadata.MD {
	md = md.Copy()
	for k := range md {
		if strings.HasPrefix(k, modserver.RoutingHopPrefix) {
			delete(md, k)
		}
	}
	return md
}

type wrappingCallback struct {
	log    *zap.Logger
	rpcApi modserver.RpcApi
	stream grpc.ServerStream
	// msgResp is an embedded Message response to allocate it only once, not for each message
	msgResp GatewayKasResponse
	data    *[]byte
}

func newWrappingCallback(log *zap.Logger, rpcApi modserver.RpcApi, stream grpc.ServerStream) *wrappingCallback {
	message := &GatewayKasResponse_Message{}
	return &wrappingCallback{
		log:    log,
		rpcApi: rpcApi,
		stream: stream,
		msgResp: GatewayKasResponse{
			Msg: &GatewayKasResponse_Message_{
				Message: message,
			},
		},
		data: &message.Data,
	}
}

func (c *wrappingCallback) Header(md map[string]*prototool.Values) error {
	return c.sendMsg("SendMsg(GatewayKasResponse_Header) failed", &GatewayKasResponse{
		Msg: &GatewayKasResponse_Header_{
			Header: &GatewayKasResponse_Header{
				Meta: md,
			},
		},
	})
}

func (c *wrappingCallback) Message(data []byte) error {
	*c.data = data
	return c.sendMsg("SendMsg(GatewayKasResponse_Message) failed", &c.msgResp)
}

func (c *wrappingCallback) Trailer(md map[string]*prototool.Values) error {
	return c.sendMsg("SendMsg(GatewayKasResponse_Trailer) failed", &GatewayKasResponse{
		Msg: &GatewayKasResponse_Trailer_{
			Trailer: &GatewayKasResponse_Trailer{
				Meta: md,
			},
		},
	})
}

func (c *wrappingCallback) Error(stat *statuspb.Status) error {
	return c.sendMsg("SendMsg(GatewayKasResponse_Error) failed", &GatewayKasResponse{
		Msg: &GatewayKasResponse_Error_{
			Error: &GatewayKasResponse_Error{
				Status: stat,
			},
		},
	})
}

func (c *wrappingCallback) sendMsg(errMsg string, msg *GatewayKasResponse) error {
	err := c.stream.SendMsg(msg)
	if err != nil {
		return c.rpcApi.HandleIoError(c.log, errMsg, err)
	}
	return nil
}
