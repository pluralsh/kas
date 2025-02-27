package server

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/reverse_tunnel/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/retry"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/testing/mock_modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/testing/mock_reverse_tunnel_rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/testing/mock_reverse_tunnel_tunnel"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/testing/testhelpers"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	_ rpc.ReverseTunnelServer = (*server)(nil)
)

func TestConnectAllowsValidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	h := mock_reverse_tunnel_tunnel.NewMockHandler(ctrl)
	mockRpcApi := mock_modserver.NewMockAgentRpcApiWithMockPoller(ctrl, 1)
	mockRpcApi.EXPECT().
		Log().
		Return(zaptest.NewLogger(t)).
		AnyTimes()
	s := &server{
		tunnelHandler:          h,
		getAgentInfoPollConfig: retry.NewPollConfigFactory(0, defaultRetry()),
	}
	agentInfo := testhelpers.AgentInfoObj()
	ctx := grpctool.AddMaxConnectionAgeContext(context.Background(), context.Background())
	ctx = modserver.InjectAgentRpcApi(ctx, mockRpcApi)
	connectServer := mock_reverse_tunnel_rpc.NewMockReverseTunnel_ConnectServer(ctrl)
	connectServer.EXPECT().
		Context().
		Return(ctx).
		MinTimes(1)
	gomock.InOrder(
		mockRpcApi.EXPECT().
			AgentInfo(gomock.Any(), gomock.Any()).
			Return(agentInfo, nil),
		h.EXPECT().
			HandleTunnel(gomock.Any(), agentInfo, connectServer),
	)
	err := s.Connect(connectServer)
	require.NoError(t, err)
}

func TestConnectRejectsInvalidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	h := mock_reverse_tunnel_tunnel.NewMockHandler(ctrl)
	mockRpcApi := mock_modserver.NewMockAgentRpcApiWithMockPoller(ctrl, 1)
	mockRpcApi.EXPECT().
		Log().
		Return(zaptest.NewLogger(t)).
		AnyTimes()
	s := &server{
		tunnelHandler:          h,
		getAgentInfoPollConfig: retry.NewPollConfigFactory(0, defaultRetry()),
	}
	ctx := grpctool.AddMaxConnectionAgeContext(context.Background(), context.Background())
	ctx = modserver.InjectAgentRpcApi(ctx, mockRpcApi)
	connectServer := mock_reverse_tunnel_rpc.NewMockReverseTunnel_ConnectServer(ctrl)
	connectServer.EXPECT().
		Context().
		Return(ctx).
		MinTimes(1)
	mockRpcApi.EXPECT().
		AgentInfo(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("expected err"))
	err := s.Connect(connectServer)
	assert.EqualError(t, err, "expected err")
}

func TestConnectRetriesFailedAgentInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	h := mock_reverse_tunnel_tunnel.NewMockHandler(ctrl)
	mockRpcApi := mock_modserver.NewMockAgentRpcApiWithMockPoller(ctrl, 2)
	mockRpcApi.EXPECT().
		Log().
		Return(zaptest.NewLogger(t)).
		AnyTimes()
	s := &server{
		tunnelHandler:          h,
		getAgentInfoPollConfig: retry.NewPollConfigFactory(0, defaultRetry()),
	}
	ctx := grpctool.AddMaxConnectionAgeContext(context.Background(), context.Background())
	ctx = modserver.InjectAgentRpcApi(ctx, mockRpcApi)
	connectServer := mock_reverse_tunnel_rpc.NewMockReverseTunnel_ConnectServer(ctrl)
	connectServer.EXPECT().
		Context().
		Return(ctx).
		MinTimes(1)
	agentInfo := testhelpers.AgentInfoObj()
	gomock.InOrder(
		mockRpcApi.EXPECT().
			AgentInfo(gomock.Any(), gomock.Any()).
			Return(nil, status.Error(codes.Unavailable, "unavailable")),
		mockRpcApi.EXPECT().
			AgentInfo(gomock.Any(), gomock.Any()).
			Return(agentInfo, nil),
		h.EXPECT().
			HandleTunnel(gomock.Any(), agentInfo, connectServer),
	)
	err := s.Connect(connectServer)
	assert.NoError(t, err)
}

func defaultRetry() retry.BackoffManagerFactory {
	return retry.NewExponentialBackoffFactory(
		getAgentInfoInitBackoff,
		getAgentInfoMaxBackoff,
		getAgentInfoResetDuration,
		getAgentInfoBackoffFactor,
		getAgentInfoJitter,
	)
}
