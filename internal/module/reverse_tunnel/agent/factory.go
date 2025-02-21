package agent

import (
	"time"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modagent"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modshared"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/reverse_tunnel"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/reverse_tunnel/info"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/reverse_tunnel/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/retry"
	"google.golang.org/grpc"
)

const (
	minIdleConnections = 2
	maxConnections     = 500
	maxIdleTime        = time.Minute
	// scaleUpStep defines how many new connections are started when there is not enough idle connections.
	scaleUpStep = 10

	connectionInitBackoff   = 1 * time.Second
	connectionMaxBackoff    = 20 * time.Second
	connectionResetDuration = 25 * time.Second
	connectionBackoffFactor = 1.6
	connectionJitter        = 0.2
)

type Factory struct {
	InternalServerConn grpc.ClientConnInterface
}

func (f *Factory) IsProducingLeaderModules() bool {
	return false
}

func (f *Factory) New(config *modagent.Config) (modagent.Module, error) {
	sv, err := grpctool.NewStreamVisitor(&rpc.ConnectResponse{})
	if err != nil {
		return nil, err
	}
	client := rpc.NewReverseTunnelClient(config.KasConn)
	pollConfig := retry.NewPollConfigFactory(0, retry.NewExponentialBackoffFactory(
		connectionInitBackoff,
		connectionMaxBackoff,
		connectionResetDuration,
		connectionBackoffFactor,
		connectionJitter,
	))
	return &module{
		server:             config.Server,
		minIdleConnections: minIdleConnections,
		maxConnections:     maxConnections,
		scaleUpStep:        scaleUpStep,
		maxIdleTime:        maxIdleTime,
		connectionFactory: func(descriptor *info.AgentDescriptor, onActive, onIdle func(c connectionInterface)) connectionInterface {
			return &connection{
				log:                config.Log,
				descriptor:         descriptor,
				client:             client,
				internalServerConn: f.InternalServerConn,
				streamVisitor:      sv,
				pollConfig:         pollConfig,
				onActive:           onActive,
				onIdle:             onIdle,
			}
		},
	}, nil
}

func (f *Factory) Name() string {
	return reverse_tunnel.ModuleName
}

func (f *Factory) StartStopPhase() modshared.ModuleStartStopPhase {
	return modshared.ModuleStartAfterServers
}
