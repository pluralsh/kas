package agent

import (
	"context"
	"time"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/gitops/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/logz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/retry"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/pkg/agentcfg"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/cli-runtime/pkg/resource"
	"sigs.k8s.io/cli-utils/pkg/apply"
	"sigs.k8s.io/cli-utils/pkg/apply/event"
	"sigs.k8s.io/cli-utils/pkg/common"
	"sigs.k8s.io/cli-utils/pkg/inventory"
)

const (
	dryRunStrategyNone   = "none"
	dryRunStrategyClient = "client"
	dryRunStrategyServer = "server"

	prunePropagationPolicyOrphan     = "orphan"
	prunePropagationPolicyBackground = "background"
	prunePropagationPolicyForeground = "foreground"

	inventoryPolicyMustMatch          = "must_match"
	inventoryPolicyAdoptIfNoInventory = "adopt_if_no_inventory"
	inventoryPolicyAdoptAll           = "adopt_all"

	defaultReapplyInterval = 5 * time.Minute
)

var (
	dryRunStrategyMapping = map[string]common.DryRunStrategy{
		dryRunStrategyNone:   common.DryRunNone,
		dryRunStrategyClient: common.DryRunClient,
		dryRunStrategyServer: common.DryRunServer,
	}
	prunePropagationPolicyMapping = map[string]metav1.DeletionPropagation{
		prunePropagationPolicyOrphan:     metav1.DeletePropagationOrphan,
		prunePropagationPolicyBackground: metav1.DeletePropagationBackground,
		prunePropagationPolicyForeground: metav1.DeletePropagationForeground,
	}
	inventoryPolicyMapping = map[string]inventory.InventoryPolicy{
		inventoryPolicyMustMatch:          inventory.InventoryPolicyMustMatch,
		inventoryPolicyAdoptIfNoInventory: inventory.AdoptIfNoInventory,
		inventoryPolicyAdoptAll:           inventory.AdoptAll,
	}
)

type Applier interface {
	Run(ctx context.Context, invInfo inventory.InventoryInfo, objects []*unstructured.Unstructured, options apply.Options) <-chan event.Event
}

type GitopsWorkerFactory interface {
	New(int64, *agentcfg.ManifestProjectCF) GitopsWorker
}

type GitopsWorker interface {
	Run(context.Context)
}

type defaultGitopsWorker struct {
	objWatcher rpc.ObjectsToSynchronizeWatcherInterface
	synchronizerConfig
}

func (w *defaultGitopsWorker) Run(ctx context.Context) {
	var wg wait.Group
	defer wg.Wait()
	desiredState := make(chan rpc.ObjectsToSynchronizeData)
	defer close(desiredState)
	wg.Start(func() {
		s := newSynchronizer(w.synchronizerConfig)
		s.run(desiredState)
	})
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

type defaultGitopsWorkerFactory struct {
	log               *zap.Logger
	applier           Applier
	restMapper        meta.RESTMapper
	restClientGetter  resource.RESTClientGetter
	gitopsClient      rpc.GitopsClient
	watchPollConfig   retry.PollConfigFactory
	applierPollConfig retry.PollConfigFactory
}

func (f *defaultGitopsWorkerFactory) New(agentId int64, project *agentcfg.ManifestProjectCF) GitopsWorker {
	l := f.log.With(logz.ProjectId(project.Id), logz.AgentId(agentId))
	return &defaultGitopsWorker{
		objWatcher: &rpc.ObjectsToSynchronizeWatcher{
			Log:          l,
			GitopsClient: f.gitopsClient,
			PollConfig:   f.watchPollConfig,
		},
		synchronizerConfig: synchronizerConfig{
			log:               l,
			agentId:           agentId,
			project:           project,
			applier:           f.applier,
			restMapper:        f.restMapper,
			restClientGetter:  f.restClientGetter,
			applierPollConfig: f.applierPollConfig(),
			applyOptions: apply.Options{
				ServerSideOptions: common.ServerSideOptions{
					// It's supported since Kubernetes 1.16, so there should be no reason not to use it.
					// https://kubernetes.io/docs/reference/using-api/server-side-apply/
					ServerSideApply: true,
					// GitOps repository is the source of truth and that's what we are applying, so overwrite any conflicts.
					// https://kubernetes.io/docs/reference/using-api/server-side-apply/#conflicts
					ForceConflicts: true,
					// https://kubernetes.io/docs/reference/using-api/server-side-apply/#field-management
					FieldManager: "agentk",
				},
				ReconcileTimeout:       project.ReconcileTimeout.AsDuration(),
				PollInterval:           0, // use default value
				EmitStatusEvents:       true,
				NoPrune:                !project.GetPrune(),
				DryRunStrategy:         f.mapDryRunStrategy(project.DryRunStrategy),
				PrunePropagationPolicy: f.mapPrunePropagationPolicy(project.PrunePropagationPolicy),
				PruneTimeout:           project.PruneTimeout.AsDuration(),
				InventoryPolicy:        f.mapInventoryPolicy(project.InventoryPolicy),
			},
		},
	}
}

func (f *defaultGitopsWorkerFactory) mapDryRunStrategy(strategy string) common.DryRunStrategy {
	ret, ok := dryRunStrategyMapping[strategy]
	if !ok {
		// This shouldn't happen because we've checked the value in DefaultAndValidateConfiguration().
		// Just being extra cautious.
		f.log.Sugar().Errorf("Invalid dry-run strategy: %q, using client dry-run for safety - NO CHANGES WILL BE APPLIED!", strategy)
		ret = common.DryRunClient
	}
	return ret
}

func (f *defaultGitopsWorkerFactory) mapPrunePropagationPolicy(policy string) metav1.DeletionPropagation {
	ret, ok := prunePropagationPolicyMapping[policy]
	if !ok {
		// This shouldn't happen because we've checked the value in DefaultAndValidateConfiguration().
		// Just being extra cautious.
		f.log.Sugar().Errorf("Invalid prune propagation policy: %q, defaulting to %s", policy, metav1.DeletePropagationForeground)
		ret = metav1.DeletePropagationForeground
	}
	return ret
}

func (f *defaultGitopsWorkerFactory) mapInventoryPolicy(policy string) inventory.InventoryPolicy {
	ret, ok := inventoryPolicyMapping[policy]
	if !ok {
		// This shouldn't happen because we've checked the value in DefaultAndValidateConfiguration().
		// Just being extra cautious.
		f.log.Sugar().Errorf("Invalid inventory policy: %q, defaulting to 'must match'", policy)
		ret = inventory.InventoryPolicyMustMatch
	}
	return ret
}
