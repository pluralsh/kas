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

func TestGetProjectInfo(t *testing.T) {
	const (
		projectId = "bla/bla"
	)
	ctx, traceId := testhelpers.CtxWithSpanContext(t)
	response := &GetProjectInfoResponse{
		ProjectId: 234,
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
	gitLabClient := mock_gitlab.SetupClient(t, ProjectInfoApiPath, func(w http.ResponseWriter, r *http.Request) {
		testhelpers.AssertRequestMethod(t, r, http.MethodGet)
		testhelpers.AssertGetJsonRequestIsCorrect(t, r, traceId)
		assert.Equal(t, projectId, r.URL.Query().Get(ProjectIdQueryParam))

		testhelpers.RespondWithJSON(t, w, response)
	})

	projInfo, err := GetProjectInfo(ctx, gitLabClient, testhelpers.AgentkToken, projectId)
	require.NoError(t, err)

	assert.Equal(t, response.ProjectId, projInfo.ProjectId)
	AssertGitalyInfo(t, response.GitalyInfo, projInfo.GitalyInfo)
	AssertGitalyRepository(t, response.GitalyRepository, projInfo.Repository)
	assert.Equal(t, response.DefaultBranch, projInfo.DefaultBranch)
}
