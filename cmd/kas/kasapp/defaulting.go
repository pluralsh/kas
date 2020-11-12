package kasapp

import (
	"time"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/tools/protodefault"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/pkg/kascfg"
)

const (
	defaultListenNetwork kascfg.ListenNetworkEnum = 0 // whatever is 0 is the default value

	defaultListenAgentAddress = "127.0.0.1:8150"
	defaultGitLabAddress      = "http://localhost:8080"

	defaultAgentConfigurationPollPeriod = 20 * time.Second

	defaultAgentInfoCacheTTL      = 5 * time.Minute
	defaultAgentInfoCacheErrorTTL = 1 * time.Minute

	defaultAgentLimitsConnectionsPerTokenPerMinute   = 100
	defaultAgentLimitsRedisKeyPrefix                 = "kas:agent_limits"
	defaultAgentLimitsMaxConfigurationFileSize       = 128 * 1024
	defaultAgentLimitsMaxGitopsManifestFileSize      = 1024 * 1024
	defaultAgentLimitsMaxGitopsTotalManifestFileSize = 2 * 1024 * 1024
	defaultAgentLimitsMaxGitopsNumberOfPaths         = 100
	defaultAgentLimitsMaxGitopsNumberOfFiles         = 1000
	defaultAgentConnectionMaxAge                     = 30 * time.Minute

	defaultGitOpsPollPeriod               = 20 * time.Second
	defaultGitOpsProjectInfoCacheTTL      = 5 * time.Minute
	defaultGitOpsProjectInfoCacheErrorTTL = 1 * time.Minute

	defaultUsageReportingPeriod       = 1 * time.Minute
	defaultObservabilityListenAddress = "0.0.0.0:8151"
	defaultPrometheusUrlPath          = "/metrics"
	defaultLivenessProbeUrlPath       = "/liveness"
	defaultReadinessProbeUrlPath      = "/readiness"

	defaultGitalyGlobalApiRefillRate    = 10.0
	defaultGitalyGlobalApiBucketSize    = 50
	defaultGitalyPerServerApiRate       = 5.0
	defaultGitalyPerServerApiBucketSize = 10

	defaultRedisMaxIdle      = 1
	defaultRedisMaxActive    = 1
	defaultRedisReadTimeout  = 1 * time.Second
	defaultRedisWriteTimeout = 1 * time.Second
	defaultRedisKeepAlive    = 5 * time.Minute
)

func ApplyDefaultsToKasConfigurationFile(cfg *kascfg.ConfigurationFile) {
	protodefault.NotNil(&cfg.ListenAgent)
	defaultListenAgent(cfg.ListenAgent)

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

func defaultListenAgent(l *kascfg.ListenAgentCF) {
	protodefault.String(&l.Address, defaultListenAgentAddress)
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
	protodefault.Uint32(&a.Limits.ConnectionsPerTokenPerMinute, defaultAgentLimitsConnectionsPerTokenPerMinute)
	protodefault.String(&a.Limits.RedisKeyPrefix, defaultAgentLimitsRedisKeyPrefix)
	protodefault.Uint32(&a.Limits.MaxConfigurationFileSize, defaultAgentLimitsMaxConfigurationFileSize)
	protodefault.Uint32(&a.Limits.MaxGitopsManifestFileSize, defaultAgentLimitsMaxGitopsManifestFileSize)
	protodefault.Uint32(&a.Limits.MaxGitopsTotalManifestFileSize, defaultAgentLimitsMaxGitopsTotalManifestFileSize)
	protodefault.Uint32(&a.Limits.MaxGitopsNumberOfPaths, defaultAgentLimitsMaxGitopsNumberOfPaths)
	protodefault.Uint32(&a.Limits.MaxGitopsNumberOfFiles, defaultAgentLimitsMaxGitopsNumberOfFiles)
	protodefault.Duration(&a.Limits.ConnectionMaxAge, defaultAgentConnectionMaxAge)
}

func defaultObservability(o *kascfg.ObservabilityCF) {
	protodefault.Duration(&o.UsageReportingPeriod, defaultUsageReportingPeriod)

	protodefault.NotNil(&o.Listen)
	protodefault.String(&o.Listen.Address, defaultObservabilityListenAddress)

	protodefault.NotNil(&o.Prometheus)
	protodefault.String(&o.Prometheus.UrlPath, defaultPrometheusUrlPath)

	protodefault.NotNil(&o.Tracing)

	protodefault.NotNil(&o.Sentry)

	protodefault.NotNil(&o.Logging)

	protodefault.NotNil(&o.GoogleProfiler)

	protodefault.NotNil(&o.LivenessProbe)
	protodefault.String(&o.LivenessProbe.UrlPath, defaultLivenessProbeUrlPath)

	protodefault.NotNil(&o.ReadinessProbe)
	protodefault.String(&o.ReadinessProbe.UrlPath, defaultReadinessProbeUrlPath)
}

func defaultGitaly(g *kascfg.GitalyCF) {
	protodefault.NotNil(&g.GlobalApiRateLimit)
	protodefault.Float64(&g.GlobalApiRateLimit.RefillRatePerSecond, defaultGitalyGlobalApiRefillRate)
	protodefault.Uint32(&g.GlobalApiRateLimit.BucketSize, defaultGitalyGlobalApiBucketSize)

	protodefault.NotNil(&g.PerServerApiRateLimit)
	protodefault.Float64(&g.PerServerApiRateLimit.RefillRatePerSecond, defaultGitalyPerServerApiRate)
	protodefault.Uint32(&g.PerServerApiRateLimit.BucketSize, defaultGitalyPerServerApiBucketSize)
}

func defaultRedis(r *kascfg.RedisCF) {
	protodefault.Uint32(&r.MaxIdle, defaultRedisMaxIdle)
	protodefault.Uint32(&r.MaxActive, defaultRedisMaxActive)
	protodefault.Duration(&r.ReadTimeout, defaultRedisReadTimeout)
	protodefault.Duration(&r.WriteTimeout, defaultRedisWriteTimeout)
	protodefault.Duration(&r.Keepalive, defaultRedisKeepAlive)
}
