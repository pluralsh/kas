package kasapp

import (
	"time"

	agent_configuration_server "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/agent_configuration/server"
	gitops_server "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/gitops/server"
	google_profiler_server "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/google_profiler/server"
	kubernetes_api_server "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/kubernetes_api/server"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/modserver"
	observability_server "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/observability/server"
	usage_metrics_server "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/usage_metrics/server"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/prototool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/pkg/kascfg"
)

const (
	defaultGitLabApiRateLimitRefillRate = 10.0
	defaultGitLabApiRateLimitBucketSize = 50

	defaultAgentInfoCacheTTL         = 5 * time.Minute
	defaultAgentInfoCacheErrorTTL    = 1 * time.Minute
	defaultAgentRedisConnInfoTTL     = 5 * time.Minute
	defaultAgentRedisConnInfoRefresh = 4 * time.Minute
	defaultAgentRedisConnInfoGC      = 10 * time.Minute

	defaultAgentListenAddress                      = "127.0.0.1:8150"
	defaultAgentListenConnectionsPerTokenPerMinute = 10000
	defaultAgentListenMaxConnectionAge             = 30 * time.Minute

	defaultGitalyGlobalApiRefillRate    = 30.0
	defaultGitalyGlobalApiBucketSize    = 70
	defaultGitalyPerServerApiRate       = 15.0
	defaultGitalyPerServerApiBucketSize = 40

	defaultRedisPoolSize     = 5
	defaultRedisDialTimeout  = 5 * time.Second
	defaultRedisReadTimeout  = 1 * time.Second
	defaultRedisWriteTimeout = 1 * time.Second
	defaultRedisIdleTimeout  = 50 * time.Second // must be smaller than server's setting
	defaultRedisKeyPrefix    = "gitlab-kas"
	defaultRedisNetwork      = "tcp"

	defaultApiListenAddress          = "127.0.0.1:8153"
	defaultApiListenMaxConnectionAge = 30 * time.Minute

	defaultPrivateApiListenMaxConnectionAge = 30 * time.Minute
)

var (
	defaulters = []modserver.ApplyDefaults{
		observability_server.ApplyDefaults,
		google_profiler_server.ApplyDefaults,
		agent_configuration_server.ApplyDefaults,
		usage_metrics_server.ApplyDefaults,
		gitops_server.ApplyDefaults,
		kubernetes_api_server.ApplyDefaults,
	}
)

func ApplyDefaultsToKasConfigurationFile(cfg *kascfg.ConfigurationFile) {
	prototool.NotNil(&cfg.Gitlab)
	defaultGitLab(cfg.Gitlab)

	prototool.NotNil(&cfg.Agent)
	defaultAgent(cfg.Agent)

	prototool.NotNil(&cfg.Gitaly)
	defaultGitaly(cfg.Gitaly)

	// TODO this should become required
	if cfg.Redis != nil {
		defaultRedis(cfg.Redis)
	}

	// TODO this should become required
	if cfg.Api != nil {
		defaultApi(cfg.Api)
	}

	// TODO this should become required
	if cfg.PrivateApi != nil {
		defaultPrivateApi(cfg.PrivateApi)
	}

	for _, defaulter := range defaulters {
		defaulter(cfg)
	}
}

func defaultApi(api *kascfg.ApiCF) {
	prototool.NotNil(&api.Listen)
	prototool.String(&api.Listen.Address, defaultApiListenAddress)
	prototool.Duration(&api.Listen.MaxConnectionAge, defaultApiListenMaxConnectionAge)
}

func defaultPrivateApi(api *kascfg.PrivateApiCF) {
	prototool.NotNil(&api.Listen)
	prototool.Duration(&api.Listen.MaxConnectionAge, defaultPrivateApiListenMaxConnectionAge)
}

func defaultGitLab(g *kascfg.GitLabCF) {
	prototool.NotNil(&g.ApiRateLimit)
	prototool.Float64(&g.ApiRateLimit.RefillRatePerSecond, defaultGitLabApiRateLimitRefillRate)
	prototool.Uint32(&g.ApiRateLimit.BucketSize, defaultGitLabApiRateLimitBucketSize)
}

func defaultAgent(a *kascfg.AgentCF) {
	prototool.NotNil(&a.Listen)
	prototool.String(&a.Listen.Address, defaultAgentListenAddress)
	prototool.Uint32(&a.Listen.ConnectionsPerTokenPerMinute, defaultAgentListenConnectionsPerTokenPerMinute)
	prototool.Duration(&a.Listen.MaxConnectionAge, defaultAgentListenMaxConnectionAge)

	prototool.Duration(&a.InfoCacheTtl, defaultAgentInfoCacheTTL)
	prototool.Duration(&a.InfoCacheErrorTtl, defaultAgentInfoCacheErrorTTL)
	prototool.Duration(&a.RedisConnInfoTtl, defaultAgentRedisConnInfoTTL)
	prototool.Duration(&a.RedisConnInfoRefresh, defaultAgentRedisConnInfoRefresh)
	prototool.Duration(&a.RedisConnInfoGc, defaultAgentRedisConnInfoGC)
}

func defaultGitaly(g *kascfg.GitalyCF) {
	prototool.NotNil(&g.GlobalApiRateLimit)
	prototool.Float64(&g.GlobalApiRateLimit.RefillRatePerSecond, defaultGitalyGlobalApiRefillRate)
	prototool.Uint32(&g.GlobalApiRateLimit.BucketSize, defaultGitalyGlobalApiBucketSize)

	prototool.NotNil(&g.PerServerApiRateLimit)
	prototool.Float64(&g.PerServerApiRateLimit.RefillRatePerSecond, defaultGitalyPerServerApiRate)
	prototool.Uint32(&g.PerServerApiRateLimit.BucketSize, defaultGitalyPerServerApiBucketSize)
}

func defaultRedis(r *kascfg.RedisCF) {
	prototool.Uint32(&r.PoolSize, defaultRedisPoolSize)
	prototool.Duration(&r.DialTimeout, defaultRedisDialTimeout)
	prototool.Duration(&r.ReadTimeout, defaultRedisReadTimeout)
	prototool.Duration(&r.WriteTimeout, defaultRedisWriteTimeout)
	prototool.Duration(&r.IdleTimeout, defaultRedisIdleTimeout)
	prototool.String(&r.KeyPrefix, defaultRedisKeyPrefix)
	prototool.String(&r.Network, defaultRedisNetwork)
}
