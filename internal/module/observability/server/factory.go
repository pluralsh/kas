package server

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modshared"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/observability"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/tlstool"
)

type Factory struct {
	Gatherer prometheus.Gatherer
}

func (f *Factory) New(config *modserver.Config) (modserver.Module, error) {
	listenCfg := config.Config.Observability.Listen
	var listener func() (net.Listener, error)

	tlsConfig, err := tlstool.MaybeDefaultServerTLSConfig(listenCfg.GetCertificateFile(), listenCfg.GetKeyFile())
	if err != nil {
		return nil, err
	}
	if tlsConfig != nil {
		listener = func() (net.Listener, error) {
			return tls.Listen(*listenCfg.Network, listenCfg.Address, tlsConfig)
		}
	} else {
		listener = func() (net.Listener, error) {
			return net.Listen(*listenCfg.Network, listenCfg.Address)
		}
	}
	return &module{
		log:           config.Log,
		api:           config.Api,
		cfg:           config.Config.Observability,
		listener:      listener,
		gatherer:      f.Gatherer,
		registerer:    config.Registerer,
		serverName:    fmt.Sprintf("%s/%s/%s", config.KasName, config.Version, config.CommitId),
		probeRegistry: config.ProbeRegistry,
	}, nil
}

func (f *Factory) Name() string {
	return observability.ModuleName
}

func (f *Factory) StartStopPhase() modshared.ModuleStartStopPhase {
	return modshared.ModuleStartBeforeServers
}
