package kasapp

import (
	"time"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/tools/protodefault"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/pkg/kascfg"
	"go.uber.org/zap"
)

const (
	defaultListenNetwork = "tcp"
	defaultListenAddress = "127.0.0.1:8150"
	defaultGitLabAddress = "http://localhost:8080"

	defaultAgentConfigurationPollPeriod = 20 * time.Second

	defaultAgentInfoCacheTTL      = 5 * time.Minute
	defaultAgentInfoCacheErrorTTL = 1 * time.Minute

	defaultAgentLimitsRedisKeyPrefix               = "kas:agent_limits"
	defaultAgentLimitsConnectionsPerTokenPerMinute = 100

	defaultGitOpsPollPeriod               = 20 * time.Second
	defaultGitOpsProjectInfoCacheTTL      = 5 * time.Minute
	defaultGitOpsProjectInfoCacheErrorTTL = 1 * time.Minute

	defaultUsageReportingPeriod    = 1 * time.Minute
	defaultPrometheusListenNetwork = "tcp"
	defaultPrometheusListenAddress = "127.0.0.1:8151"
	defaultPrometheusListenUrlPath = "/metrics"
	defaultLoggingLevel            = zap.InfoLevel

	defaultGitalyGlobalApiRefillRate    float64 = 10.0 // type matches protobuf type from kascfg.TokenBucketRateLimitCF
	defaultGitalyGlobalApiBucketSize    int32   = 50   // type matches protobuf type from kascfg.TokenBucketRateLimitCF
	defaultGitalyPerServerApiRate       float64 = 5.0  // type matches protobuf type from kascfg.TokenBucketRateLimitCF
	defaultGitalyPerServerApiBucketSize int32   = 10   // type matches protobuf type from kascfg.TokenBucketRateLimitCF

	defaultRedisMaxIdle      = 1
	defaultRedisMaxActive    = 1
	defaultRedisReadTimeout  = 1 * time.Second
	defaultRedisWriteTimeout = 1 * time.Second
	defaultRedisKeepAlive    = 5 * time.Minute
)

func ApplyDefaultsToKasConfigurationFile(cfg *kascfg.ConfigurationFile) {
	protodefault.NotNil(&cfg.Listen)
	defaultListen(cfg.Listen)

	protodefault.NotNil(&cfg.Gitlab)
	defaultGitLab(cfg.Gitlab)

	protodefault.NotNil(&cfg.Agent)
	defaultAgent(cfg.Agent)

	protodefault.NotNil(&cfg.Observability)
	defaultObservability(cfg.Observability)

	protodefault.NotNil(&cfg.Gitaly)
	defaultGitaly(cfg.Gitaly)

	if cfg.Redis != nil {
		defaultRedis(cfg.Redis)
	}
}

func defaultListen(l *kascfg.ListenCF) {
	protodefault.String(&l.Network, defaultListenNetwork)
	protodefault.String(&l.Address, defaultListenAddress)
}

func defaultGitLab(g *kascfg.GitLabCF) {
	protodefault.String(&g.Address, defaultGitLabAddress)
}

func defaultAgent(a *kascfg.AgentCF) {
	protodefault.NotNil(&a.Configuration)
	protodefault.Duration(&a.Configuration.PollPeriod, defaultAgentConfigurationPollPeriod)

	protodefault.NotNil(&a.Gitops)
	protodefault.Duration(&a.Gitops.PollPeriod, defaultGitOpsPollPeriod)
	protodefault.Duration(&a.Gitops.ProjectInfoCacheTtl, defaultGitOpsProjectInfoCacheTTL)
	protodefault.Duration(&a.Gitops.ProjectInfoCacheErrorTtl, defaultGitOpsProjectInfoCacheErrorTTL)

	protodefault.Duration(&a.InfoCacheTtl, defaultAgentInfoCacheTTL)
	protodefault.Duration(&a.InfoCacheErrorTtl, defaultAgentInfoCacheErrorTTL)

	protodefault.NotNil(&a.Limits)
	protodefault.String(&a.Limits.RedisKeyPrefix, defaultAgentLimitsRedisKeyPrefix)
	protodefault.Uint32(&a.Limits.ConnectionsPerTokenPerMinute, defaultAgentLimitsConnectionsPerTokenPerMinute)
}

func defaultObservability(o *kascfg.ObservabilityCF) {
	protodefault.Duration(&o.UsageReportingPeriod, defaultUsageReportingPeriod)

	protodefault.NotNil(&o.Listen)
	protodefault.String(&o.Listen.Network, defaultPrometheusListenNetwork)
	protodefault.String(&o.Listen.Address, defaultPrometheusListenAddress)

	protodefault.NotNil(&o.Prometheus)
	protodefault.String(&o.Prometheus.UrlPath, defaultPrometheusListenUrlPath)

	protodefault.NotNil(&o.Tracing)

	protodefault.NotNil(&o.Sentry)

	protodefault.NotNil(&o.Logging)
	protodefault.String(&o.Logging.Level, defaultLoggingLevel.String())

	protodefault.NotNil(&o.GoogleProfiler)

	protodefault.NotNil(&o.Pprof)
}

func defaultGitaly(g *kascfg.GitalyCF) {
	protodefault.NotNil(&g.GlobalApiRateLimit)
	protodefault.Float64(&g.GlobalApiRateLimit.RefillRatePerSecond, defaultGitalyGlobalApiRefillRate)
	protodefault.Int32(&g.GlobalApiRateLimit.BucketSize, defaultGitalyGlobalApiBucketSize)

	protodefault.NotNil(&g.PerServerApiRateLimit)
	protodefault.Float64(&g.PerServerApiRateLimit.RefillRatePerSecond, defaultGitalyPerServerApiRate)
	protodefault.Int32(&g.PerServerApiRateLimit.BucketSize, defaultGitalyPerServerApiBucketSize)
}

func defaultRedis(r *kascfg.RedisCF) {
	protodefault.Int32(&r.MaxIdle, defaultRedisMaxIdle)
	protodefault.Int32(&r.MaxActive, defaultRedisMaxActive)
	protodefault.Duration(&r.ReadTimeout, defaultRedisReadTimeout)
	protodefault.Duration(&r.WriteTimeout, defaultRedisWriteTimeout)
	protodefault.Duration(&r.Keepalive, defaultRedisKeepAlive)
}
