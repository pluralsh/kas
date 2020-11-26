package agentk

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/agentrpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/agentrpc/mock_agentrpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/module/modagent"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/tool/testing/mock_grpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/tool/testing/mock_modagent"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/pkg/agentcfg"
	"go.uber.org/zap/zaptest"
)

const (
	revision1 = "rev12341234_1"
	revision2 = "rev12341234_2"
)

func TestConfigurationIsApplied(t *testing.T) {
	cfg1 := &agentcfg.AgentConfiguration{}
	cfg2 := &agentcfg.AgentConfiguration{
		Gitops: &agentcfg.GitopsCF{
			ManifestProjects: []*agentcfg.ManifestProjectCF{
				{
					Id: "bla",
				},
			},
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mockCtrl := gomock.NewController(t)
	watcher := mock_agentrpc.NewMockConfigurationWatcherInterface(mockCtrl)
	f := mock_modagent.NewMockFactory(mockCtrl)
	m := mock_modagent.NewMockModule(mockCtrl)
	m.EXPECT().
		Run(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		<-ctx.Done()
		return nil
	})
	gomock.InOrder(
		m.EXPECT().
			DefaultAndValidateConfiguration(cfg1),
		m.EXPECT().
			SetConfiguration(cfg1),
		m.EXPECT().
			DefaultAndValidateConfiguration(cfg2),
		m.EXPECT().
			SetConfiguration(cfg2),
	)
	gomock.InOrder(
		f.EXPECT().
			New(gomock.Any()).
			Return(m),
		watcher.EXPECT().
			Watch(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, callback agentrpc.ConfigurationCallback) {
				callback(ctx, agentrpc.ConfigurationData{CommitId: revision1, Config: cfg1})
				callback(ctx, agentrpc.ConfigurationData{CommitId: revision2, Config: cfg2})
				cancel()
			}),
	)
	a := &Agent{
		Log:                  zaptest.NewLogger(t),
		KasConn:              mock_grpc.NewMockClientConnInterface(mockCtrl),
		ConfigurationWatcher: watcher,
		ModuleFactories:      []modagent.Factory{f},
	}
	err := a.Run(ctx)
	require.NoError(t, err)
}
