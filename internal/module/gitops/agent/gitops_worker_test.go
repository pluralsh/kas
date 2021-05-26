package agent

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/module/gitops/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/tool/testing/kube_testing"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/tool/testing/matcher"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/tool/testing/mock_rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/tool/testing/testhelpers"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/pkg/agentcfg"
	"go.uber.org/zap/zaptest"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	cmdtesting "k8s.io/kubectl/pkg/cmd/testing"
	"sigs.k8s.io/cli-utils/pkg/apply"
	"sigs.k8s.io/cli-utils/pkg/apply/event"
	"sigs.k8s.io/cli-utils/pkg/inventory"
)

const (
	projectId        = "bla123/bla-1"
	revision         = "rev12341234"
	defaultNamespace = "testing1"
)

var (
	_ ApplierFactory      = &defaultApplierFactory{}
	_ GitopsWorker        = &defaultGitopsWorker{}
	_ GitopsWorkerFactory = &defaultGitopsWorkerFactory{}
)

func TestRunHappyPathNoObjects(t *testing.T) {
	w, applier, watcher := setupWorker(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req := &rpc.ObjectsToSynchronizeRequest{
		ProjectId: projectId,
		Paths:     w.project.Paths,
	}
	gomock.InOrder(
		watcher.EXPECT().
			Watch(gomock.Any(), matcher.ProtoEq(t, req), gomock.Any()).
			Do(func(ctx context.Context, req *rpc.ObjectsToSynchronizeRequest, callback rpc.ObjectsToSynchronizeCallback) {
				callback(ctx, rpc.ObjectsToSynchronizeData{
					CommitId: revision,
				})
				<-ctx.Done()
			}),
		applier.EXPECT().
			Run(gomock.Any(), gomock.Any(), gomock.Len(0), gomock.Any()).
			DoAndReturn(func(ctx context.Context, invInfo inventory.InventoryInfo, objects []*unstructured.Unstructured, options apply.Options) <-chan event.Event {
				cancel() // all good, stop run()
				c := make(chan event.Event)
				close(c)
				return c
			}),
	)
	w.Run(ctx)
}

func TestRunHappyPath(t *testing.T) {
	w, applier, watcher := setupWorker(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req := &rpc.ObjectsToSynchronizeRequest{
		ProjectId: projectId,
		Paths:     w.project.Paths,
	}
	objs := []*unstructured.Unstructured{
		kube_testing.ToUnstructured(t, testMap1()),
		kube_testing.ToUnstructured(t, testNs1()),
		kube_testing.ToUnstructured(t, testMap2()),
	}
	gomock.InOrder(
		watcher.EXPECT().
			Watch(gomock.Any(), matcher.ProtoEq(t, req), gomock.Any()).
			Do(func(ctx context.Context, req *rpc.ObjectsToSynchronizeRequest, callback rpc.ObjectsToSynchronizeCallback) {
				callback(ctx, rpc.ObjectsToSynchronizeData{
					CommitId: revision,
					Sources: []rpc.ObjectSource{
						{
							Name: "obj1.yaml",
							Data: kube_testing.ObjsToYAML(t, objs[0]),
						},
						{
							Name: "obj2.yaml",
							Data: kube_testing.ObjsToYAML(t, objs[1], objs[2]),
						},
					},
				})
				<-ctx.Done()
			}),
		applier.EXPECT().
			Run(gomock.Any(), gomock.Any(), matcher.K8sObjectEq(t, objs), gomock.Any()).
			DoAndReturn(func(ctx context.Context, invInfo inventory.InventoryInfo, objects []*unstructured.Unstructured, options apply.Options) <-chan event.Event {
				cancel() // all good, stop run()
				c := make(chan event.Event)
				close(c)
				return c
			}),
	)
	w.Run(ctx)
}

func TestRunHappyPathSyncCancellation(t *testing.T) {
	w, applier, watcher := setupWorker(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req := &rpc.ObjectsToSynchronizeRequest{
		ProjectId: projectId,
		Paths:     w.project.Paths,
	}
	objs := []*unstructured.Unstructured{
		kube_testing.ToUnstructured(t, testMap1()),
		kube_testing.ToUnstructured(t, testNs1()),
		kube_testing.ToUnstructured(t, testMap2()),
	}
	job1started := make(chan struct{})
	watcher.EXPECT().
		Watch(gomock.Any(), matcher.ProtoEq(t, req), gomock.Any()).
		Do(func(ctx context.Context, req *rpc.ObjectsToSynchronizeRequest, callback rpc.ObjectsToSynchronizeCallback) {
			callback(ctx, rpc.ObjectsToSynchronizeData{
				CommitId: revision,
				Sources: []rpc.ObjectSource{
					{
						Name: "obj1.yaml",
						Data: kube_testing.ObjsToYAML(t, objs[0]),
					},
					{
						Name: "obj2.yaml",
						Data: kube_testing.ObjsToYAML(t, objs[1], objs[2]),
					},
				},
			})
			<-job1started
			callback(ctx, rpc.ObjectsToSynchronizeData{
				CommitId: revision,
			})
			<-ctx.Done()
		})
	gomock.InOrder(
		applier.EXPECT().
			Run(gomock.Any(), gomock.Any(), matcher.K8sObjectEq(t, objs), gomock.Any()).
			DoAndReturn(func(ctx context.Context, invInfo inventory.InventoryInfo, objects []*unstructured.Unstructured, options apply.Options) <-chan event.Event {
				close(job1started) // signal that this job has been started
				c := make(chan event.Event)
				go func() {
					<-ctx.Done() // block until the job is cancelled
					close(c)
				}()
				return c
			}),
		applier.EXPECT().
			Run(gomock.Any(), gomock.Any(), gomock.Len(0), gomock.Any()).
			DoAndReturn(func(ctx context.Context, invInfo inventory.InventoryInfo, objects []*unstructured.Unstructured, options apply.Options) <-chan event.Event {
				cancel() // all good, stop run()
				c := make(chan event.Event)
				close(c)
				return c
			}),
	)
	w.Run(ctx)
}

func setupWorker(t *testing.T) (*defaultGitopsWorker, *MockApplier, *mock_rpc.MockObjectsToSynchronizeWatcherInterface) {
	ctrl := gomock.NewController(t)
	applier := NewMockApplier(ctrl)
	applierFactory := NewMockApplierFactory(ctrl)
	watcher := mock_rpc.NewMockObjectsToSynchronizeWatcherInterface(ctrl)
	gomock.InOrder(
		applierFactory.EXPECT().
			New().
			Return(applier),
		applier.EXPECT().
			Initialize(),
	)
	w := &defaultGitopsWorker{
		objWatcher:     watcher,
		applierFactory: applierFactory,
		applierBackoff: testhelpers.NewBackoff(),
		synchronizerConfig: synchronizerConfig{
			log: zaptest.NewLogger(t),
			project: &agentcfg.ManifestProjectCF{
				Id:               projectId,
				DefaultNamespace: defaultNamespace, // as if user didn't specify configuration so it's the default value
				Paths: []*agentcfg.PathCF{
					{
						Glob: "*.yaml",
					},
				},
			},
			k8sUtilFactory: cmdtesting.NewTestFactory(),
		},
	}
	return w, applier, watcher
}

func testMap1() *corev1.ConfigMap {
	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "map1",
			Namespace: "test1",
			Annotations: map[string]string{
				"k1": "v1",
			},
		},
		Data: map[string]string{
			"key1": "value1",
		},
	}
}

func testMap2() *corev1.ConfigMap {
	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "map2",
			Namespace: "test2",
			Annotations: map[string]string{
				"k2": "v2",
			},
		},
		Data: map[string]string{
			"key2": "value2",
		},
	}
}

func testNs1() *corev1.Namespace {
	return &corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "ns1",
			Annotations: map[string]string{
				"k3": "v3",
			},
		},
	}
}
