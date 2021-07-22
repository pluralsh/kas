module gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14

go 1.16

require (
	cloud.google.com/go v0.86.0
	github.com/ash2k/stager v0.2.1
	github.com/bmatcuk/doublestar/v2 v2.0.4
	github.com/cilium/cilium v1.9.6
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1.0.20200107205605-c66185887605
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/go-logr/zapr v0.4.0
	github.com/go-redis/redis/v8 v8.11.0
	github.com/go-redis/redismock/v8 v8.0.6
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.6
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.1-0.20200507082539-9abf3eb82b4a
	github.com/opentracing/opentracing-go v1.2.0
	github.com/piotrkowalczuk/promgrpc/v4 v4.0.4
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	gitlab.com/gitlab-org/gitaly/v14 v14.1.0-rc3
	gitlab.com/gitlab-org/labkit v1.5.1-0.20210709085637-98de0925bac8
	go.uber.org/zap v1.18.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/time v0.0.0-20210611083556-38a9dc6acbc6
	golang.org/x/tools v0.1.4
	google.golang.org/api v0.50.0
	google.golang.org/genproto v0.0.0-20210701191553-46259e63a0a9
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/cli-runtime v0.21.2
	k8s.io/client-go v0.21.2
	k8s.io/klog/v2 v2.9.0
	k8s.io/kubectl v0.21.2
	nhooyr.io/websocket v1.8.7
	sigs.k8s.io/cli-utils v0.25.1-0.20210630230537-da80fcc751ec
	sigs.k8s.io/yaml v1.2.0
)

replace (
	// Cilium depends on v0.19.9 which is incompatible with v0.19.3 which Kubernetes uses. v0.19.8 works fine
	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.8
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b
	// Use a fork to avoid all the dependencies https://github.com/spf13/cobra/issues/1240#issuecomment-874387919
	// https://github.com/ash2k/cobra/commits/remove-cli
	github.com/spf13/cobra => github.com/ash2k/cobra v1.2.2-0.20210706005132-e91bfee91527
	// same version as used by rules_go to maintain compatibility with patches - see the WORKSPACE file
	golang.org/x/tools => golang.org/x/tools v0.1.4

	// https://github.com/kubernetes/kubernetes/issues/79384#issuecomment-505627280
	k8s.io/api => k8s.io/api v0.21.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.21.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.21.2
	k8s.io/apiserver => k8s.io/apiserver v0.21.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.21.2
	k8s.io/client-go => k8s.io/client-go v0.21.2
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.21.2
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.21.2
	k8s.io/code-generator => k8s.io/code-generator v0.21.2
	k8s.io/component-base => k8s.io/component-base v0.21.2
	k8s.io/component-helpers => k8s.io/component-helpers v0.21.2
	k8s.io/controller-manager => k8s.io/controller-manager v0.21.2
	k8s.io/cri-api => k8s.io/cri-api v0.21.2
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.21.2
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.21.2
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.21.2
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.21.2
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.21.2
	k8s.io/kubectl => k8s.io/kubectl v0.21.2
	k8s.io/kubelet => k8s.io/kubelet v0.21.2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.21.2
	k8s.io/metrics => k8s.io/metrics v0.21.2
	k8s.io/mount-utils => k8s.io/mount-utils v0.21.2
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.21.2
)
