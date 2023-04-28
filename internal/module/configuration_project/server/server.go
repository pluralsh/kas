package server

import (
	"context"
	"fmt"
	"regexp"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/api"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/gitaly"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/agent_configuration"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/configuration_project/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modshared"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/git"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/logz"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// https://gitlab.com/gitlab-org/cluster-integration/gitlab-agent/-/blob/master/doc/identity_and_auth.md#agent-identity-and-name
	agentNameRegex = `[a-z0-9](?:[-a-z0-9]*[a-z0-9])?`
)

var (
	agentConfigFileRegex = regexp.MustCompile(fmt.Sprintf("^%s/(%s)/%s$",
		regexp.QuoteMeta(agent_configuration.Directory),
		agentNameRegex,
		regexp.QuoteMeta(agent_configuration.FileName),
	))
)

type server struct {
	rpc.UnimplementedConfigurationProjectServer
	gitaly gitaly.PoolInterface
}

func (s *server) ListAgentConfigFiles(ctx context.Context, req *rpc.ListAgentConfigFilesRequest) (*rpc.ListAgentConfigFilesResponse, error) {
	rpcApi := modserver.RpcApiFromContext(ctx)
	pf, err := s.gitaly.PathFetcher(ctx, &api.GitalyInfo{
		Address: req.GitalyAddress.Address,
		Token:   req.GitalyAddress.Token,
		//Features: nil, // TODO
	})
	if err != nil {
		rpcApi.HandleProcessingError(rpcApi.Log(), modshared.NoAgentId, "PathFetcher", err)
		return nil, status.Error(codes.Unavailable, "Unavailable")
	}
	r := &gitalypb.Repository{
		StorageName:                   req.Repository.StorageName,
		RelativePath:                  req.Repository.RelativePath,
		GitObjectDirectory:            req.Repository.GitObjectDirectory,
		GitAlternateObjectDirectories: req.Repository.GitAlternateObjectDirectories,
		GlRepository:                  req.Repository.GlRepository,
		GlProjectPath:                 req.Repository.GlProjectPath,
	}
	v := &configVisitor{}
	ref := git.ExplicitRefOrHead(req.DefaultBranch)
	err = pf.Visit(ctx, r, []byte(ref), []byte(agent_configuration.Directory), true, v)
	if err != nil {
		log := rpcApi.Log().With(logz.ProjectId(req.Repository.GlProjectPath))
		rpcApi.HandleProcessingError(log, modshared.NoAgentId, "PathFetcher", err)
		return nil, status.Error(codes.Unavailable, "Unavailable")
	}
	return &rpc.ListAgentConfigFilesResponse{
		ConfigFiles: v.resp,
	}, nil
}

type configVisitor struct {
	resp []*rpc.AgentConfigFile
}

func (c *configVisitor) Entry(entry *gitalypb.TreeEntry) (bool /* download? */, int64 /* max size */, error) {
	submatch := agentConfigFileRegex.FindSubmatch(entry.Path)
	if submatch == nil {
		return false, 0, nil
	}
	c.resp = append(c.resp, &rpc.AgentConfigFile{
		Name:      string(entry.Path),
		AgentName: string(submatch[1]),
	})
	return false, 0, nil
}

func (c *configVisitor) StreamChunk(path []byte, data []byte) (bool, error) {
	return false, nil
}

func (c *configVisitor) EntryDone(entry *gitalypb.TreeEntry, err error) {}
