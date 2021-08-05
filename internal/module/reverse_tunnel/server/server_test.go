package server

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/api"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/reverse_tunnel/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/testing/mock_modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/testing/mock_reverse_tunnel"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/testing/mock_reverse_tunnel_rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/testing/testhelpers"
	"go.uber.org/zap/zaptest"
)

var (
	_ rpc.ReverseTunnelServer = &server{}
)

func TestConnectAllowsValidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	h := mock_reverse_tunnel.NewMockTunnelHandler(ctrl)
	mockRpcApi := mock_modserver.NewMockRpcApi(ctrl)
	s := &server{
		tunnelHandler: h,
	}
	agentInfo := testhelpers.AgentInfoObj()
	ctx := api.InjectAgentMD(context.Background(), &api.AgentMD{Token: testhelpers.AgentkToken})
	ctx = grpctool.InjectLogger(ctx, zaptest.NewLogger(t))
	ctx = grpctool.AddMaxConnectionAgeContext(ctx, context.Background())
	ctx = grpctool.InjectRpcApi(ctx, mockRpcApi)
	connectServer := mock_reverse_tunnel_rpc.NewMockReverseTunnel_ConnectServer(ctrl)
	connectServer.EXPECT().
		Context().
		Return(ctx).
		MinTimes(1)
	gomock.InOrder(
		mockRpcApi.EXPECT().
			GetAgentInfo(gomock.Any(), gomock.Any()).
			Return(agentInfo, nil),
		h.EXPECT().
			HandleTunnel(gomock.Any(), agentInfo, connectServer),
	)
	err := s.Connect(connectServer)
	require.NoError(t, err)
}

func TestConnectRejectsInvalidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	h := mock_reverse_tunnel.NewMockTunnelHandler(ctrl)
	mockRpcApi := mock_modserver.NewMockRpcApi(ctrl)
	s := &server{
		tunnelHandler: h,
	}
	ctx := api.InjectAgentMD(context.Background(), &api.AgentMD{Token: "invalid"})
	ctx = grpctool.InjectLogger(ctx, zaptest.NewLogger(t))
	ctx = grpctool.AddMaxConnectionAgeContext(ctx, context.Background())
	ctx = grpctool.InjectRpcApi(ctx, mockRpcApi)
	connectServer := mock_reverse_tunnel_rpc.NewMockReverseTunnel_ConnectServer(ctrl)
	connectServer.EXPECT().
		Context().
		Return(ctx).
		MinTimes(1)
	gomock.InOrder(
		mockRpcApi.EXPECT().
			GetAgentInfo(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("expected err")),
	)
	err := s.Connect(connectServer)
	assert.EqualError(t, err, "expected err")
}
