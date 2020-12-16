package kascfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
)

type validatable interface {
	Validate() error
}

func TestValidation_Valid(t *testing.T) {
	tests := []struct {
		name  string
		valid validatable
	}{
		{
			name: "minimal",
			valid: &ConfigurationFile{
				Gitlab: &GitLabCF{
					Address:                  "http://localhost:8080",
					AuthenticationSecretFile: "/some/file",
				},
			},
		},
		{
			name: "GitopsCF",
			valid: &GitopsCF{
				ProjectInfoCacheTtl:      durationpb.New(0), // zero means "disabled"
				MaxManifestFileSize:      0,                 // zero means "use default value"
				MaxTotalManifestFileSize: 0,                 // zero means "use default value"
				MaxNumberOfPaths:         0,                 // zero means "use default value"
				MaxNumberOfFiles:         0,                 // zero means "use default value"
			},
		},
		{
			name: "AgentCF",
			valid: &AgentCF{
				InfoCacheTtl: durationpb.New(0), // zero means "disabled"
			},
		},
		{
			name: "ObservabilityCF",
			valid: &ObservabilityCF{
				UsageReportingPeriod: durationpb.New(0), // zero means "disabled"
			},
		},
		{
			name: "TokenBucketRateLimitCF",
			valid: &TokenBucketRateLimitCF{
				RefillRatePerSecond: 0, // zero means "use default value"
				BucketSize:          0, // zero means "use default value"
			},
		},
		{
			name: "RedisCF",
			valid: &RedisCF{
				Url:       "unix:///tmp/redis.sock", // some value - it's a required field
				MaxIdle:   0,                        // zero means "use default value"
				MaxActive: 0,                        // zero means "use default value"
				Keepalive: durationpb.New(0),        // zero means "disabled"
			},
		},
		{
			name: "AgentConfigurationCF",
			valid: &AgentConfigurationCF{
				MaxConfigurationFileSize: 0, // zero means "use default value"
			},
		},
		{
			name: "ListenAgentCF",
			valid: &ListenAgentCF{
				ConnectionsPerTokenPerMinute: 0, // zero means "use default value"
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) { // nolint: scopelint
			assert.NoError(t, tc.valid.Validate()) // nolint: scopelint
		})
	}
}

func TestValidation_Invalid(t *testing.T) {
	tests := []struct {
		name      string
		errString string
		invalid   validatable
	}{
		{
			name:      "zero GitopsCF.PollPeriod",
			errString: "invalid GitopsCF.PollPeriod: value must be greater than 0s",
			invalid: &GitopsCF{
				PollPeriod: durationpb.New(0),
			},
		},
		{
			name:      "negative GitopsCF.PollPeriod",
			errString: "invalid GitopsCF.PollPeriod: value must be greater than 0s",
			invalid: &GitopsCF{
				PollPeriod: durationpb.New(-1),
			},
		},
		{
			name:      "negative GitopsCF.ProjectInfoCacheTtl",
			errString: "invalid GitopsCF.ProjectInfoCacheTtl: value must be greater than or equal to 0s",
			invalid: &GitopsCF{
				ProjectInfoCacheTtl: durationpb.New(-1),
			},
		},
		{
			name:      "zero GitopsCF.ProjectInfoCacheErrorTtl",
			errString: "invalid GitopsCF.ProjectInfoCacheErrorTtl: value must be greater than 0s",
			invalid: &GitopsCF{
				ProjectInfoCacheErrorTtl: durationpb.New(0),
			},
		},
		{
			name:      "negative GitopsCF.ProjectInfoCacheErrorTtl",
			errString: "invalid GitopsCF.ProjectInfoCacheErrorTtl: value must be greater than 0s",
			invalid: &GitopsCF{
				ProjectInfoCacheErrorTtl: durationpb.New(-1),
			},
		},
		{
			name:      "negative AgentCF.InfoCacheTtl",
			errString: "invalid AgentCF.InfoCacheTtl: value must be greater than or equal to 0s",
			invalid: &AgentCF{
				InfoCacheTtl: durationpb.New(-1),
			},
		},
		{
			name:      "zero AgentCF.InfoCacheErrorTtl",
			errString: "invalid AgentCF.InfoCacheErrorTtl: value must be greater than 0s",
			invalid: &AgentCF{
				InfoCacheErrorTtl: durationpb.New(0),
			},
		},
		{
			name:      "negative AgentCF.InfoCacheErrorTtl",
			errString: "invalid AgentCF.InfoCacheErrorTtl: value must be greater than 0s",
			invalid: &AgentCF{
				InfoCacheErrorTtl: durationpb.New(-1),
			},
		},
		{
			name:      "zero AgentConfigurationCF.PollPeriod",
			errString: "invalid AgentConfigurationCF.PollPeriod: value must be greater than 0s",
			invalid: &AgentConfigurationCF{
				PollPeriod: durationpb.New(0),
			},
		},
		{
			name:      "negative AgentConfigurationCF.PollPeriod",
			errString: "invalid AgentConfigurationCF.PollPeriod: value must be greater than 0s",
			invalid: &AgentConfigurationCF{
				PollPeriod: durationpb.New(-1),
			},
		},
		{
			name:      "negative ObservabilityCF.UsageReportingPeriod",
			errString: "invalid ObservabilityCF.UsageReportingPeriod: value must be greater than or equal to 0s",
			invalid: &ObservabilityCF{
				UsageReportingPeriod: durationpb.New(-1),
			},
		},
		{
			name:      "negative TokenBucketRateLimitCF.RefillRatePerSecond",
			errString: "invalid TokenBucketRateLimitCF.RefillRatePerSecond: value must be greater than or equal to 0",
			invalid: &TokenBucketRateLimitCF{
				RefillRatePerSecond: -1,
			},
		},
		{
			name:      "zero RedisCF.ReadTimeout",
			errString: "invalid RedisCF.ReadTimeout: value must be greater than 0s",
			invalid: &RedisCF{
				Url:         "unix:///tmp/redis.sock",
				ReadTimeout: durationpb.New(0),
			},
		},
		{
			name:      "negative RedisCF.ReadTimeout",
			errString: "invalid RedisCF.ReadTimeout: value must be greater than 0s",
			invalid: &RedisCF{
				Url:         "unix:///tmp/redis.sock",
				ReadTimeout: durationpb.New(-1),
			},
		},
		{
			name:      "zero RedisCF.WriteTimeout",
			errString: "invalid RedisCF.WriteTimeout: value must be greater than 0s",
			invalid: &RedisCF{
				Url:          "unix:///tmp/redis.sock",
				WriteTimeout: durationpb.New(0),
			},
		},
		{
			name:      "negative RedisCF.WriteTimeout",
			errString: "invalid RedisCF.WriteTimeout: value must be greater than 0s",
			invalid: &RedisCF{
				Url:          "unix:///tmp/redis.sock",
				WriteTimeout: durationpb.New(-1),
			},
		},
		{
			name:      "negative RedisCF.Keepalive",
			errString: "invalid RedisCF.Keepalive: value must be greater than or equal to 0s",
			invalid: &RedisCF{
				Url:       "unix:///tmp/redis.sock",
				Keepalive: durationpb.New(-1),
			},
		},
		{
			name:      "empty RedisCF.Url",
			errString: "invalid RedisCF.Url: value length must be at least 1 runes",
			invalid:   &RedisCF{},
		},
		{
			name:      "relative RedisCF.Url",
			errString: "invalid RedisCF.Url: value must be absolute",
			invalid: &RedisCF{
				Url: "/redis.sock",
			},
		},
		{
			name:      "zero ListenAgentCF.MaxConnectionAge",
			errString: "invalid ListenAgentCF.MaxConnectionAge: value must be greater than 0s",
			invalid: &ListenAgentCF{
				MaxConnectionAge: durationpb.New(0),
			},
		},
		{
			name:      "negative ListenAgentCF.MaxConnectionAge",
			errString: "invalid ListenAgentCF.MaxConnectionAge: value must be greater than 0s",
			invalid: &ListenAgentCF{
				MaxConnectionAge: durationpb.New(-1),
			},
		},
		{
			name:      "empty GitLabCF.Address",
			errString: "invalid GitLabCF.Address: value length must be at least 1 runes",
			invalid: &GitLabCF{
				AuthenticationSecretFile: "/some/file",
			},
		},
		{
			name:      "relative GitLabCF.Address",
			errString: "invalid GitLabCF.Address: value must be absolute",
			invalid: &GitLabCF{
				Address:                  "/path",
				AuthenticationSecretFile: "/some/file",
			},
		},
		{
			name:      "empty GitLabCF.AuthenticationSecretFile",
			errString: "invalid GitLabCF.AuthenticationSecretFile: value length must be at least 1 runes",
			invalid: &GitLabCF{
				Address: "http://localhost:8080",
			},
		},
		// TODO uncomment when Redis becomes a hard dependency
		//{
		//	name:      "missing ConfigurationFile.Redis",
		//	errString: "invalid ConfigurationFile.Redis: value is required",
		//	invalid: &ConfigurationFile{
		//		Gitlab: &GitLabCF{
		//			Address:                  "http://localhost:8080",
		//			AuthenticationSecretFile: "/some/file",
		//		},
		//	},
		//},
		{
			name:      "missing ConfigurationFile.Gitlab",
			errString: "invalid ConfigurationFile.Gitlab: value is required",
			invalid:   &ConfigurationFile{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) { // nolint: scopelint
			err := tc.invalid.Validate() // nolint: scopelint
			require.Error(t, err)
			assert.EqualError(t, err, tc.errString) // nolint: scopelint
		})
	}
}
