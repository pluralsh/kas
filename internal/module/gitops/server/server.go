package server

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/api"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/gitaly"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/gitlab"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/gitops/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/usage_metrics"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/errz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/logz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	gitOpsManifestMaxChunkSize = 128 * 1024
)

var (
	// globPrefix captures glob prefix that does not contain any special characters, recognized by doublestar.Match.
	// See https://github.com/bmatcuk/doublestar#about and
	// https://pkg.go.dev/github.com/bmatcuk/doublestar/v2#Match for globbing rules.
	globPrefix = regexp.MustCompile(`^([^\\*?[\]{}]+)/(.*)$`)
)

type server struct {
	rpc.UnimplementedGitopsServer
	gitalyPool                  gitaly.PoolInterface
	projectInfoClient           *projectInfoClient
	syncCount                   usage_metrics.Counter
	gitOpsPollIntervalHistogram prometheus.Histogram
	getObjectsPollConfig        retry.PollConfigFactory
	maxManifestFileSize         int64
	maxTotalManifestFileSize    int64
	maxNumberOfPaths            uint32
	maxNumberOfFiles            uint32
}

func (s *server) GetObjectsToSynchronize(req *rpc.ObjectsToSynchronizeRequest, server rpc.Gitops_GetObjectsToSynchronizeServer) error {
	err := s.validateGetObjectsToSynchronizeRequest(req)
	if err != nil {
		return err // no wrap
	}
	var (
		lastPoll  time.Time
		agentInfo *api.AgentInfo
	)
	ctx := server.Context()
	rpcApi := modserver.AgentRpcApiFromContext(server.Context())
	agentToken := rpcApi.AgentToken()
	log := rpcApi.Log().With(logz.ProjectId(req.ProjectId))
	return rpcApi.PollWithBackoff(s.getObjectsPollConfig(), func() (error, retry.AttemptResult) {
		if agentInfo == nil { // executed only once (if successful)
			agentInfo, err = rpcApi.AgentInfo(ctx, log)
			if err != nil {
				if status.Code(err) == codes.Unavailable {
					return nil, retry.Backoff
				}
				return err, retry.Done // no wrap
			}
			log = log.With(logz.AgentId(agentInfo.Id))
		}
		// This call is made on each poll because:
		// - it checks that the agent's token is still valid
		// - repository location in Gitaly might have changed
		projectInfo, err := s.getProjectInfo(ctx, log, rpcApi, agentInfo.Id, agentToken, req.ProjectId)
		if err != nil {
			if status.Code(err) == codes.Unavailable {
				return nil, retry.Backoff
			}
			return err, retry.Done // no wrap
		}

		var commitToSynchronize string
		synchronizingCommit := req.GetRef().GetCommit() != ""
		// Declare a new logger to no append the same field in every poll.
		log := log // nolint:govet
		if synchronizingCommit {
			// no need to poll, because we already have an immutable commit sha
			commitToSynchronize = req.GetRef().GetCommit()
			log = log.With(logz.CommitId(commitToSynchronize))
		} else {
			// Resolve the ref provided by the agent to a full ref which can be used to unambiguously find a Git ref.
			fullRefName := req.GetRef().GetResolvedRef()
			info, err := s.poll(ctx, projectInfo, req.CommitId, fullRefName) // nolint:govet
			if err != nil {
				rpcApi.HandleProcessingError(log, agentInfo.Id, "GitOps: repository poll failed", err)
				return nil, retry.Backoff
			}
			commitToSynchronize = info.CommitId
			log = log.With(logz.CommitId(commitToSynchronize), logz.GitRef(fullRefName))

			if info.EmptyRepository {
				log.Debug("GitOps: empty repository")
				return nil, retry.Continue
			}
		}
		s.trackPollInterval(&lastPoll)

		// If the commit to synchronize is the same as the previously synchronized commit, we don't need to do anything
		if commitToSynchronize == req.CommitId {
			log.Debug("GitOps: no updates")

			if !synchronizingCommit { // the actual ref may change, and we need to poll the given ref.
				return nil, retry.Continue
			} else { // we are synchronizing a commit and don't need to poll again, because it doesn't change.
				log.Debug("blocking GetObjectsToSynchronize because the request wants to synchronize a commit which already has been synced")
				<-ctx.Done()
				return nil, retry.Done
			}
		}

		log.Info("GitOps: new commit")
		err = s.sendObjectsToSynchronizeHeader(server, commitToSynchronize, projectInfo.ProjectId)
		if err != nil {
			return rpcApi.HandleIoError(log, "GitOps: failed to send header for objects to synchronize", err), retry.Done
		}
		filesVisited, filesSent, err := s.sendObjectsToSynchronizeBody(log, rpcApi, req, server, agentInfo.Id, projectInfo, commitToSynchronize) // nolint: contextcheck
		if err != nil {
			return err, retry.Done // no wrap
		}
		err = s.sendObjectsToSynchronizeTrailer(server)
		if err != nil {
			return rpcApi.HandleIoError(log, "GitOps: failed to send trailer for objects to synchronize", err), retry.Done
		}
		log.Info("GitOps: fetched files", logz.NumberOfFilesVisited(filesVisited), logz.NumberOfFilesSent(filesSent))
		s.syncCount.Inc()
		return nil, retry.Done
	})
}

func (s *server) poll(ctx context.Context, projectInfo *api.ProjectInfo, commitId string, fullRefName string) (*gitaly.PollInfo, error) {
	p, err := s.gitalyPool.Poller(ctx, &projectInfo.GitalyInfo)
	if err != nil {
		return nil, err
	}
	return p.Poll(ctx, projectInfo.Repository, commitId, fullRefName)
}

func (s *server) validateGetObjectsToSynchronizeRequest(req *rpc.ObjectsToSynchronizeRequest) error {
	numberOfPaths := uint32(len(req.Paths))
	if numberOfPaths > s.maxNumberOfPaths {
		// TODO validate config in GetConfiguration too and send it somewhere the user can see it https://gitlab.com/gitlab-org/gitlab/-/issues/277323
		// This check must be here, but there too.
		return status.Errorf(codes.InvalidArgument, "maximum number of GitOps paths per manifest project is %d, but %d was requested", s.maxNumberOfPaths, numberOfPaths)
	}
	return nil
}

func (s *server) trackPollInterval(lastPoll *time.Time) {
	now := time.Now()

	if !lastPoll.IsZero() {
		pollInterval := now.Sub(*lastPoll).Seconds()
		s.gitOpsPollIntervalHistogram.Observe(pollInterval)
	}

	*lastPoll = now
}

func (s *server) sendObjectsToSynchronizeHeader(server rpc.Gitops_GetObjectsToSynchronizeServer, commitId string, projectId int64) error {
	return server.Send(&rpc.ObjectsToSynchronizeResponse{
		Message: &rpc.ObjectsToSynchronizeResponse_Header_{
			Header: &rpc.ObjectsToSynchronizeResponse_Header{
				CommitId:  commitId,
				ProjectId: projectId,
			},
		},
	})
}

func (s *server) sendObjectsToSynchronizeBody(
	log *zap.Logger,
	rpcApi modserver.RpcApi,
	req *rpc.ObjectsToSynchronizeRequest,
	server rpc.Gitops_GetObjectsToSynchronizeServer,
	agentId int64,
	projectInfo *api.ProjectInfo,
	commitId string,
) (uint32 /* files visited */, uint32 /* files sent */, error) {
	ctx := server.Context()
	pf, err := s.gitalyPool.PathFetcher(ctx, &projectInfo.GitalyInfo)
	if err != nil {
		rpcApi.HandleProcessingError(log, agentId, "GitOps: PathFetcher", err)
		return 0, 0, status.Error(codes.Unavailable, "GitOps: PathFetcher")
	}
	v := &objectsToSynchronizeVisitor{
		server:        server,
		fileSizeLimit: s.maxManifestFileSize,
	}
	var delegate gitaly.FetchVisitor = v
	delegate = gitaly.NewChunkingFetchVisitor(delegate, gitOpsManifestMaxChunkSize)
	delegate = gitaly.NewTotalSizeLimitingFetchVisitor(delegate, s.maxTotalManifestFileSize)
	delegate = gitaly.NewDuplicateFileDetectingVisitor(delegate, gitaly.DupError)
	delegate = gitaly.NewHiddenDirFilteringFetchVisitor(delegate)
	vGlob := gitaly.NewGlobFilteringFetchVisitor(delegate, "")
	vCounting := gitaly.NewEntryCountLimitingFetchVisitor(vGlob, s.maxNumberOfFiles)
	for _, p := range req.Paths {
		globNoSlash := strings.TrimPrefix(p.Glob, "/") // original glob without the leading slash
		repoPath, recursive := globToGitaly(globNoSlash)
		vGlob.Glob = globNoSlash // set new glob for each path
		err = pf.Visit(ctx, projectInfo.Repository, []byte(commitId), repoPath, recursive, vCounting)
		if err != nil {
			switch {
			case v.sendFailed:
				return vCounting.FilesVisited, vCounting.FilesSent, rpcApi.HandleIoError(log, "GitOps: failed to send objects to synchronize", err)
			case isUserError(err):
				err = errz.NewUserErrorWithCause(err, "manifest file")
				rpcApi.HandleProcessingError(log, agentId, "GitOps: failed to get objects to synchronize", err)
				// return the error to the client because it's a user error
				return vCounting.FilesVisited, vCounting.FilesSent, status.Errorf(codes.FailedPrecondition, "GitOps: failed to get objects to synchronize: %v", err)
			case grpctool.RequestCanceled(err):
				return vCounting.FilesVisited, vCounting.FilesSent, status.Errorf(codes.Canceled, "GitOps: failed to get objects to synchronize")
			case grpctool.RequestTimedOut(err):
				return vCounting.FilesVisited, vCounting.FilesSent, status.Errorf(codes.DeadlineExceeded, "GitOps: failed to get objects to synchronize")
			default:
				rpcApi.HandleProcessingError(log, agentId, "GitOps: failed to get objects to synchronize", err)
				return vCounting.FilesVisited, vCounting.FilesSent, status.Error(codes.Unavailable, "GitOps: failed to get objects to synchronize")
			}
		}
	}
	return vCounting.FilesVisited, vCounting.FilesSent, nil
}

func (s *server) sendObjectsToSynchronizeTrailer(server rpc.Gitops_GetObjectsToSynchronizeServer) error {
	return server.Send(&rpc.ObjectsToSynchronizeResponse{
		Message: &rpc.ObjectsToSynchronizeResponse_Trailer_{
			Trailer: &rpc.ObjectsToSynchronizeResponse_Trailer{},
		},
	})
}

// getProjectInfo returns an error with code Unavailable if there was a retriable error.
func (s *server) getProjectInfo(ctx context.Context, log *zap.Logger, rpcApi modserver.RpcApi, agentId int64,
	agentToken api.AgentToken, projectId string) (*api.ProjectInfo, error) {
	projectInfo, err := s.projectInfoClient.GetProjectInfo(ctx, agentToken, projectId)
	switch {
	case err == nil:
		return projectInfo, nil
	case errors.Is(err, context.Canceled):
		err = status.Error(codes.Canceled, err.Error())
	case errors.Is(err, context.DeadlineExceeded):
		err = status.Error(codes.DeadlineExceeded, err.Error())
	case gitlab.IsForbidden(err):
		err = status.Error(codes.PermissionDenied, "forbidden")
	case gitlab.IsUnauthorized(err):
		err = status.Error(codes.Unauthenticated, "unauthenticated")
	case gitlab.IsNotFound(err):
		err = status.Error(codes.NotFound, "project not found")
	default:
		rpcApi.HandleProcessingError(log, agentId, "GetProjectInfo()", err)
		err = status.Error(codes.Unavailable, "unavailable")
	}
	return nil, err
}

func isUserError(err error) bool {
	switch err.(type) { // nolint:errorlint
	case *gitaly.GlobMatchFailedError, *gitaly.MaxNumberOfFilesError, *gitaly.DuplicatePathFoundError:
		return true
	}
	switch gitaly.ErrorCodeFromError(err) { // nolint:exhaustive
	case gitaly.NotFound, gitaly.FileTooBig, gitaly.UnexpectedTreeEntryType:
		return true
	}
	return false
}

// globToGitaly accepts a glob without a leading slash!
func globToGitaly(glob string) ([]byte /* repoPath */, bool /* recursive */) {
	var repoPath []byte
	matches := globPrefix.FindStringSubmatch(glob)
	if matches == nil {
		repoPath = []byte{'.'}
	} else {
		repoPath = []byte(matches[1])
		glob = matches[2]
	}
	recursive := strings.ContainsAny(glob, "[/") || // cannot determine if recursive or not because character class may contain ranges, etc
		strings.Contains(glob, "**") // contains directory match
	return repoPath, recursive
}
