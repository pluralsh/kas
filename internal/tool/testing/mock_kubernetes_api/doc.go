package mock_kubernetes_api

//go:generate mockgen.sh -destination "rpc.go" -package "mock_kubernetes_api" "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/kubernetes_api/rpc" "KubernetesApiClient,KubernetesApi_MakeRequestClient"
