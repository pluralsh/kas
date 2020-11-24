package gitops_agent

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"testing"
	"time"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/modules/modclient"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/tools/testing/mock_engine"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/tools/testing/mock_modclient"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/pkg/agentcfg"
	"go.uber.org/zap/zaptest"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	_ modclient.Module  = &Module{}
	_ modclient.Factory = &Factory{}
)

func TestStartsWorkersAccordingToConfiguration(t *testing.T) {
	for i, config := range testConfigurations() {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			var wg wait.Group
			defer wg.Wait()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			expectedNumberOfWorkers := len(config.GetGitops().GetManifestProjects()) // nolint: scopelint
			m, mockCtrl, factory, _ := setupModule(t)
			for i := 0; i < expectedNumberOfWorkers; i++ {
				engine := mock_engine.NewMockGitOpsEngine(mockCtrl)
				engine.EXPECT().
					Run().
					Return(nil, errors.New("i'm not ok, but that's ok")).
					MinTimes(1)
				factory.EXPECT().
					New(gomock.Any(), gomock.Any()).
					Return(engine)
			}
			wg.Start(func() {
				err := m.Run(ctx)
				assert.NoError(t, err)
			})
			require.NoError(t, m.DefaultAndValidateConfiguration(config)) // nolint: scopelint
			require.NoError(t, m.SetConfiguration(config))                // nolint: scopelint
			cancel()
			wg.Wait()
			assertWorkersMatchConfiguration(t, m, config) // nolint: scopelint
		})
	}
}

func TestUpdatesWorkersAccordingToConfiguration(t *testing.T) {
	increasingOrder := sortableConfigs(testConfigurations())
	sort.Stable(increasingOrder)
	decreasingOrder := sortableConfigs(testConfigurations())
	sort.Sort(sort.Reverse(decreasingOrder))
	tests := []struct {
		name    string
		configs []*agentcfg.AgentConfiguration
	}{
		{
			name:    "increasing order",
			configs: increasingOrder,
		},
		{
			name:    "decreasing order",
			configs: decreasingOrder,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m, mockCtrl, factory, _ := setupModule(t)
			var wg wait.Group
			defer wg.Wait()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			engine := mock_engine.NewMockGitOpsEngine(mockCtrl)
			engine.EXPECT().
				Run().
				Return(nil, errors.New("i'm not ok, but that's ok")).
				AnyTimes()
			factory.EXPECT().
				New(gomock.Any(), gomock.Any()).
				Return(engine).
				AnyTimes()
			wg.Start(func() {
				err := m.Run(ctx)
				assert.NoError(t, err)
			})
			for _, config := range tc.configs { // nolint: scopelint
				require.NoError(t, m.DefaultAndValidateConfiguration(config))
				require.NoError(t, m.SetConfiguration(config))
			}
			cancel()
			wg.Wait()
			assertWorkersMatchConfiguration(t, m, tc.configs[len(tc.configs)-1]) // nolint: scopelint
		})
	}
}

func setupModule(t *testing.T) (*Module, *gomock.Controller, *mock_engine.MockGitOpsEngineFactory, *mock_modclient.MockAgentAPI) { // nolint: unparam
	mockCtrl := gomock.NewController(t)
	engFactory := mock_engine.NewMockGitOpsEngineFactory(mockCtrl)
	configFlags := genericclioptions.NewTestConfigFlags()
	factory := Factory{
		Log: zaptest.NewLogger(t),
		EngineFactory: &mock_engine.ThreadSafeGitOpsEngineFactory{
			EngineFactory: engFactory,
		},
		K8sClientGetter:                    configFlags,
		GetObjectsToSynchronizeRetryPeriod: 10 * time.Second,
	}
	mockApi := mock_modclient.NewMockAgentAPI(mockCtrl)
	module := factory.New(mockApi)
	return module.(*Module), mockCtrl, engFactory, mockApi
}

func testConfigurations() []*agentcfg.AgentConfiguration {
	const (
		project1 = "bla1/project1"
		project2 = "bla1/project2"
	)
	return []*agentcfg.AgentConfiguration{
		{},
		{
			Gitops: &agentcfg.GitopsCF{
				ManifestProjects: []*agentcfg.ManifestProjectCF{
					{
						Id: project1,
					},
				},
			},
		},
		{
			Gitops: &agentcfg.GitopsCF{
				ManifestProjects: []*agentcfg.ManifestProjectCF{
					{
						Id:                 project1,
						ResourceInclusions: defaultResourceExclusions, // update config
						ResourceExclusions: defaultResourceExclusions, // update config
					},
					{
						Id: project2,
					},
				},
			},
		},
		{
			Gitops: &agentcfg.GitopsCF{
				ManifestProjects: []*agentcfg.ManifestProjectCF{
					{
						Id: "bla3/project3",
					},
					{
						Id:                 project2,
						ResourceInclusions: defaultResourceExclusions, // update config
						ResourceExclusions: defaultResourceExclusions, // update config
					},
				},
			},
		},
	}
}

func assertWorkersMatchConfiguration(t *testing.T, p *Module, config *agentcfg.AgentConfiguration) bool { // nolint: unparam
	projects := config.GetGitops().GetManifestProjects()
	if !assert.Len(t, p.workers, len(projects)) {
		return false
	}
	success := true
	for _, project := range projects {
		if !assert.Contains(t, p.workers, project.Id) {
			success = false
			continue
		}
		success = assert.Empty(t, cmp.Diff(p.workers[project.Id].worker.projectConfiguration, project, protocmp.Transform())) || success
	}
	return success
}

type sortableConfigs []*agentcfg.AgentConfiguration

func (r sortableConfigs) Len() int {
	return len(r)
}

func (r sortableConfigs) Less(i, j int) bool {
	return i < j
}

func (r sortableConfigs) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
