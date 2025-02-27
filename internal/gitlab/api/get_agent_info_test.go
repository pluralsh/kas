package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/testing/mock_gitlab"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/testing/testhelpers"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/pkg/entity"
)

func TestGetAgentInfo(t *testing.T) {
	ctx, traceId := testhelpers.CtxWithSpanContext(t)
	response := &GetAgentInfoResponse{
		ProjectId: 234,
		AgentId:   555,
		AgentName: "agent-x",
		GitalyInfo: &entity.GitalyInfo{
			Address: "example.com",
			Token:   "123123",
			Features: map[string]string{
				"a": "b",
			},
		},
		GitalyRepository: &entity.GitalyRepository{
			StorageName:                   "234",
			RelativePath:                  "123",
			GitObjectDirectory:            "sfasdf",
			GitAlternateObjectDirectories: []string{"a", "b"},
			GlRepository:                  "254634",
			GlProjectPath:                 "64662",
		},
		DefaultBranch: "main",
	}
	c := mock_gitlab.SetupClient(t, AgentInfoApiPath, func(w http.ResponseWriter, r *http.Request) {
		testhelpers.AssertGetJsonRequestIsCorrect(t, r, traceId)
		testhelpers.RespondWithJSON(t, w, response)
	})
	agentInfo, err := GetAgentInfo(ctx, c, testhelpers.AgentkToken)
	require.NoError(t, err)

	assert.Equal(t, response.ProjectId, agentInfo.ProjectId)
	assert.Equal(t, response.AgentId, agentInfo.Id)
	assert.Equal(t, response.AgentName, agentInfo.Name)

	AssertGitalyInfo(t, response.GitalyInfo, agentInfo.GitalyInfo)
	AssertGitalyRepository(t, response.GitalyRepository, agentInfo.Repository)
	assert.Equal(t, response.DefaultBranch, agentInfo.DefaultBranch)
}
