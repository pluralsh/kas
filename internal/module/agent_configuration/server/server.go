package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"path"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/api"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/gitaly"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/gitlab"
	gapi "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/gitlab/api"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/agent_configuration"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/agent_configuration/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/agent_tracker"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/errz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/logz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/mathz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/retry"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/syncz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/pkg/agentcfg"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/pkg/event"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/yaml"
)

type server struct {
	rpc.UnimplementedAgentConfigurationServer
	serverApi                  modserver.Api
	gitaly                     gitaly.PoolInterface
	gitLabClient               gitlab.ClientInterface
	agentRegisterer            agent_tracker.Registerer
	maxConfigurationFileSize   int64
	getConfigurationPollConfig retry.PollConfigFactory
	gitLabExternalUrl          string
}

func (s *server) GetConfiguration(req *rpc.ConfigurationRequest, server rpc.AgentConfiguration_GetConfigurationServer) error {
	connectedAgentInfo := &agent_tracker.ConnectedAgentInfo{
		AgentMeta:    req.AgentMeta,
		ConnectedAt:  timestamppb.Now(),
		ConnectionId: mathz.Int63(),
	}
	ctx := server.Context()
	rpcApi := modserver.AgentRpcApiFromContext(ctx)
	log := rpcApi.Log()
	defer s.maybeUnregisterAgent(log, rpcApi, connectedAgentInfo, req.SkipRegister)

	pollCfg := s.getConfigurationPollConfig()

	wh := syncz.NewComparableWorkerHolder[string](
		func(projectId string) syncz.Worker {
			return syncz.WorkerFunc(func(ctx context.Context) {
				s.serverApi.OnGitPushEvent(ctx, func(ctx context.Context, e *event.GitPushEvent) {
					// NOTE: yes, the req.ProjectId is NOT a project id, but a full project path ...
					if e.Project.FullPath == projectId {
						pollCfg.Poke()
					}
				})
			})
		},
	)
	defer wh.StopAndWait()

	lastProcessedCommitId := req.CommitId
	return rpcApi.PollWithBackoff(pollCfg, func() (error, retry.AttemptResult) {
		// This call is made on each poll because:
		// - it checks that the agent's token is still valid
		// - repository location in Gitaly might have changed
		agentInfo, err := rpcApi.AgentInfo(ctx, log)
		if err != nil {
			if status.Code(err) == codes.Unavailable {
				return nil, retry.Backoff
			}
			return err, retry.Done
		}
		wh.ApplyConfig(ctx, agentInfo.Repository.GlProjectPath)
		// re-define log to avoid accidentally using the old one
		log := log.With(logz.AgentId(agentInfo.Id), logz.ProjectId(agentInfo.Repository.GlProjectPath)) // nolint:govet
		s.maybeRegisterAgent(ctx, log, rpcApi, connectedAgentInfo, agentInfo, req.SkipRegister)
		info, err := s.poll(ctx, agentInfo, lastProcessedCommitId)
		if err != nil {
			switch gitaly.ErrorCodeFromError(err) { // nolint:exhaustive
			case gitaly.NotFound: // ref not found
				return status.Errorf(codes.NotFound, "Config: repository poll failed: %v", err), retry.Done
			default:
				rpcApi.HandleProcessingError(log, agentInfo.Id, "Config: repository poll failed", err)
				return nil, retry.Backoff
			}
		}
		if info.EmptyRepository {
			log.Debug("Config: empty repository")
			return nil, retry.Continue
		}
		if !info.UpdateAvailable {
			log.Debug("Config: no updates", logz.CommitId(lastProcessedCommitId))
			return nil, retry.Continue
		}
		log.Info("Config: new commit", logz.CommitId(info.CommitId))
		configFile, err := s.fetchConfiguration(ctx, agentInfo, info.CommitId)
		if err != nil {
			rpcApi.HandleProcessingError(log, agentInfo.Id, "Config: failed to fetch", err)
			var ue errz.UserError
			if errors.As(err, &ue) {
				// return the error to the client because it's a user error
				return status.Errorf(codes.FailedPrecondition, "Config: %v", err), retry.Done
			}
			return nil, retry.Backoff
		}
		var wg wait.Group
		defer wg.Wait()
		wg.Start(func() {
			err := gapi.PostAgentConfiguration(ctx, s.gitLabClient, agentInfo.Id, configFile) // nolint:govet
			switch {
			case err == nil:
			case gitlab.IsNotFound(err):
				// Agent has been deleted from DB, but it's still running in the cluster. Don't need to send this error
				// to Sentry.
				log.Debug("Failed to notify GitLab of new agent configuration. Deleted agent?", logz.Error(err))
			default:
				rpcApi.HandleProcessingError(log, agentInfo.Id, "Failed to notify GitLab of new agent configuration", err)
			}
		})
		err = s.sendConfigResponse(server, agentInfo, configFile, info.CommitId)
		if err != nil {
			return rpcApi.HandleIoError(log, "Config: failed to send config", err), retry.Done
		}
		lastProcessedCommitId = info.CommitId
		return nil, retry.Continue
	})
}

func (s *server) poll(ctx context.Context, agentInfo *api.AgentInfo, lastProcessedCommitId string) (*gitaly.PollInfo, error) {
	p, err := s.gitaly.Poller(ctx, agentInfo.GitalyInfo)
	if err != nil {
		return nil, err
	}
	return p.Poll(ctx, agentInfo.Repository, lastProcessedCommitId, "refs/heads/"+agentInfo.DefaultBranch)
}

func (s *server) sendConfigResponse(server rpc.AgentConfiguration_GetConfigurationServer,
	agentInfo *api.AgentInfo, configFile *agentcfg.ConfigurationFile, commitId string) error {
	return server.Send(&rpc.ConfigurationResponse{
		Configuration: &agentcfg.AgentConfiguration{
			Gitops:            configFile.Gitops,
			Observability:     configFile.Observability,
			AgentId:           agentInfo.Id,
			ProjectId:         agentInfo.ProjectId,
			ProjectPath:       agentInfo.Repository.GlProjectPath,
			CiAccess:          configFile.CiAccess,
			ContainerScanning: configFile.ContainerScanning,
			RemoteDevelopment: configFile.RemoteDevelopment,
			Flux:              configFile.Flux,
			GitlabExternalUrl: s.gitLabExternalUrl,
		},
		CommitId: commitId,
	})
}

// fetchConfiguration fetches agent's configuration from a corresponding repository.
// Assumes configuration is stored in ".gitlab/agents/<agent id>/config.yaml" file.
// fetchConfiguration returns a wrapped context.Canceled, context.DeadlineExceeded or gRPC error if ctx signals done and interrupts a running gRPC call.
func (s *server) fetchConfiguration(ctx context.Context, agentInfo *api.AgentInfo, commitId string) (*agentcfg.ConfigurationFile, error) {
	pf, err := s.gitaly.PathFetcher(ctx, agentInfo.GitalyInfo)
	if err != nil {
		return nil, fmt.Errorf("PathFetcher: %w", err) // wrap
	}
	filename := path.Join(agent_configuration.Directory, agentInfo.Name, agent_configuration.FileName)
	configYAML, err := pf.FetchFile(ctx, agentInfo.Repository, []byte(commitId), []byte(filename), s.maxConfigurationFileSize)
	if err != nil {
		switch gitaly.ErrorCodeFromError(err) { // nolint:exhaustive
		case gitaly.NotFound:
			configYAML = nil // Missing config is the same as empty config
		case gitaly.FileTooBig, gitaly.UnexpectedTreeEntryType:
			return nil, errz.NewUserErrorWithCause(err, "agent configuration file")
		default:
			return nil, fmt.Errorf("fetch agent configuration: %w", err) // wrap
		}
	}
	configFile, err := parseYAMLToConfiguration(configYAML)
	if err != nil {
		return nil, errz.NewUserErrorWithCause(err, "failed to parse agent configuration")
	}
	err = configFile.ValidateAll()
	if err != nil {
		return nil, errz.NewUserErrorWithCause(err, "invalid agent configuration")
	}
	return configFile, nil
}

func (s *server) maybeRegisterAgent(ctx context.Context, log *zap.Logger, rpcApi modserver.AgentRpcApi,
	connectedAgentInfo *agent_tracker.ConnectedAgentInfo, agentInfo *api.AgentInfo, skipRegister bool) {
	// Skip registering agent if skipRegister is true. The agent will call "Register" gRPC method instead.
	if skipRegister {
		return
	}

	if connectedAgentInfo.AgentId != 0 {
		return
	}
	connectedAgentInfo.AgentId = agentInfo.Id
	connectedAgentInfo.ProjectId = agentInfo.ProjectId
	err := s.agentRegisterer.RegisterConnection(ctx, connectedAgentInfo)
	if err != nil {
		rpcApi.HandleProcessingError(log, agentInfo.Id, "Failed to register agent", err)
	}
}

func (s *server) maybeUnregisterAgent(log *zap.Logger, rpcApi modserver.AgentRpcApi,
	connectedAgentInfo *agent_tracker.ConnectedAgentInfo, skipRegister bool) {
	// Skip unregistering agent if skipRegister is true. GC will clean up the agent from the storage.
	if skipRegister {
		return
	}

	if connectedAgentInfo.AgentId == 0 {
		return
	}
	err := s.agentRegisterer.UnregisterConnection(context.Background(), connectedAgentInfo)
	if err != nil {
		rpcApi.HandleProcessingError(log, connectedAgentInfo.AgentId, "Failed to unregister agent", err)
	}
}

func parseYAMLToConfiguration(configYAML []byte) (*agentcfg.ConfigurationFile, error) {
	configJSON, err := yaml.YAMLToJSON(configYAML)
	if err != nil {
		return nil, fmt.Errorf("YAMLToJSON: %w", err)
	}
	configFile := &agentcfg.ConfigurationFile{}
	if bytes.Equal(configJSON, []byte("null")) {
		// Empty config
		return configFile, nil
	}
	err = protojson.Unmarshal(configJSON, configFile)
	if err != nil {
		return nil, fmt.Errorf("protojson.Unmarshal: %w", err)
	}
	return configFile, nil
}
