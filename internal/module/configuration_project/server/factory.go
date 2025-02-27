package server

import (
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/configuration_project"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/configuration_project/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modshared"
)

type Factory struct {
}

func (f *Factory) New(config *modserver.Config) (modserver.Module, error) {
	rpc.RegisterConfigurationProjectServer(config.ApiServer, &server{
		gitaly: config.Gitaly,
	})
	return &module{}, nil
}

func (f *Factory) Name() string {
	return configuration_project.ModuleName
}

func (f *Factory) StartStopPhase() modshared.ModuleStartStopPhase {
	return modshared.ModuleStartBeforeServers
}
