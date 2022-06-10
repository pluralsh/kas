package agent

import (
	"context"
	"fmt"
	"net"

	"github.com/ash2k/stager"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/modshared"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/observability"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/logz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/prototool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/pkg/agentcfg"
	"go.uber.org/zap"
)

type module struct {
	log                 *zap.Logger
	logLevel            zap.AtomicLevel
	grpcLogLevel        zap.AtomicLevel
	defaultGrpcLogLevel agentcfg.LogLevelEnum
	api                 modshared.Api
	serverName          string
}

const (
	listenAddress         = ":8080"
	prometheusUrlPath     = "/metrics"
	livenessProbeUrlPath  = "/liveness"
	readinessProbeUrlPath = "/readiness"
)

func (m *module) Run(ctx context.Context, cfg <-chan *agentcfg.AgentConfiguration) error {
	return stager.RunStages(ctx,
		func(stage stager.Stage) {
			// Listen for config changes and apply to logger
			stage.Go(func(ctx context.Context) error {
				for config := range cfg {
					err := m.setConfigurationLogging(config.Observability.Logging)
					if err != nil {
						m.log.Error("Failed to apply logging configuration", logz.Error(err))
						continue
					}
				}
				return nil
			})
			// Start metrics server
			stage.Go(func(ctx context.Context) error {
				lis, err := net.Listen("tcp", listenAddress) // nolint:gosec
				if err != nil {
					return fmt.Errorf("Observability listener failed to start: %w", err)
				}
				// Error is ignored because metricSrv.Run() closes the listener and
				// a second close always produces an error.
				defer lis.Close() //nolint:errcheck

				m.log.Info("Observability endpoint is up",
					logz.NetNetworkFromAddr(lis.Addr()),
					logz.NetAddressFromAddr(lis.Addr()),
				)

				metricSrv := observability.MetricServer{
					Log:                   m.log,
					Api:                   m.api,
					Name:                  m.serverName,
					Listener:              lis,
					PrometheusUrlPath:     prometheusUrlPath,
					LivenessProbeUrlPath:  livenessProbeUrlPath,
					ReadinessProbeUrlPath: readinessProbeUrlPath,
					Gatherer:              prometheus.DefaultGatherer,
					Registerer:            prometheus.DefaultRegisterer,
					LivenessProbe:         observability.NoopProbe,
					ReadinessProbe:        observability.NoopProbe,
				}

				return metricSrv.Run(ctx)
			})
		},
	)
}

func (m *module) DefaultAndValidateConfiguration(config *agentcfg.AgentConfiguration) error {
	prototool.NotNil(&config.Observability)
	prototool.NotNil(&config.Observability.Logging)
	err := m.defaultAndValidateLogging(config.Observability.Logging)
	if err != nil {
		return fmt.Errorf("logging: %w", err)
	}
	return nil
}

func (m *module) Name() string {
	return observability.ModuleName
}

func (m *module) defaultAndValidateLogging(logging *agentcfg.LoggingCF) error {
	if logging.GrpcLevel == nil {
		logging.GrpcLevel = &m.defaultGrpcLogLevel
	}
	_, err := logz.LevelFromString(logging.Level.String())
	if err != nil {
		return err
	}
	_, err = logz.LevelFromString(logging.GrpcLevel.String())
	if err != nil {
		return err
	}
	return nil
}

func (m *module) setConfigurationLogging(logging *agentcfg.LoggingCF) error {
	err := setLogLevel(m.logLevel, logging.Level)
	if err != nil {
		return err
	}

	return setLogLevel(m.grpcLogLevel, *logging.GrpcLevel) // not nil after defaulting
}

func setLogLevel(logLevel zap.AtomicLevel, val agentcfg.LogLevelEnum) error {
	level, err := logz.LevelFromString(val.String())
	if err != nil {
		return err
	}
	logLevel.SetLevel(level)
	return nil
}
