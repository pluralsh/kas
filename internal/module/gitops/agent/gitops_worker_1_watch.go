package agent

import (
	"context"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/gitops/rpc"
)

func (w *defaultGitopsWorker) watch(ctx context.Context, desiredState chan<- rpc.ObjectsToSynchronizeData) {
	req := &rpc.ObjectsToSynchronizeRequest{
		ProjectId: w.project.Id,
		Paths:     w.project.Paths,
	}
	w.objWatcher.Watch(ctx, req, func(ctx context.Context, data rpc.ObjectsToSynchronizeData) {
		select {
		case <-ctx.Done():
		case desiredState <- data:
		}
	})
}
