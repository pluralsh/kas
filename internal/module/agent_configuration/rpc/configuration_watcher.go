package rpc

import (
	"context"
	"io"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/logz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/retry"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/pkg/agentcfg"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/pkg/entity"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ConfigurationData struct {
	CommitId string
	Config   *agentcfg.AgentConfiguration
}

type ConfigurationCallback func(context.Context, ConfigurationData)

// ConfigurationWatcherInterface abstracts ConfigurationWatcher.
type ConfigurationWatcherInterface interface {
	Watch(context.Context, ConfigurationCallback)
}

type ConfigurationWatcher struct {
	Log                *zap.Logger
	AgentMeta          *entity.AgentMeta
	Client             AgentConfigurationClient
	PollConfig         retry.PollConfigFactory
	ConfigPreProcessor func(ConfigurationData) error
}

func (w *ConfigurationWatcher) Watch(ctx context.Context, callback ConfigurationCallback) {
	var lastProcessedCommitId string
	_ = retry.PollWithBackoff(ctx, w.PollConfig(), func(ctx context.Context) (error, retry.AttemptResult) {
		ctx, cancel := context.WithCancel(ctx) // nolint:govet
		defer cancel()                         // ensure streaming call is canceled
		res, err := w.Client.GetConfiguration(ctx, &ConfigurationRequest{
			CommitId:     lastProcessedCommitId,
			AgentMeta:    w.AgentMeta,
			SkipRegister: true,
		}, grpc.WaitForReady(true))
		if err != nil {
			if !grpctool.RequestCanceledOrTimedOut(err) {
				w.Log.Warn("GetConfiguration failed", logz.Error(err))
			}
			return nil, retry.Backoff
		}
		for {
			config, err := res.Recv()
			if err != nil {
				switch {
				case err == io.EOF: // nolint:errorlint
					return nil, retry.ContinueImmediately // immediately reconnect after a clean close
				case grpctool.RequestCanceledOrTimedOut(err):
				default:
					w.Log.Warn("GetConfiguration.Recv failed", logz.Error(err))
				}
				return nil, retry.Backoff
			}
			data := ConfigurationData{
				CommitId: config.CommitId,
				Config:   config.Configuration,
			}
			err = w.ConfigPreProcessor(data)
			if err != nil {
				w.Log.Error("Failed to preprocess configuration", logz.Error(err))
				continue
			}
			callback(ctx, data)
			lastProcessedCommitId = config.CommitId
		}
	})
}
