// Mocks for Gitaly.
package mock_gitaly

//go:generate mockgen.sh -destination "gitaly.go" -package "mock_gitaly" "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/gitaly/vendored/gitalypb" "CommitServiceClient,CommitService_TreeEntryClient,SmartHTTPServiceClient,SmartHTTPService_InfoRefsUploadPackClient,CommitService_GetTreeEntriesClient"
