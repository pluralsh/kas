// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/module/reverse_tunnel (interfaces: TunnelConnectionHandler)

// Package mock_reverse_tunnel is a generated GoMock package.
package mock_reverse_tunnel

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/api"
	rpc "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/module/reverse_tunnel/rpc"
)

// MockTunnelConnectionHandler is a mock of TunnelConnectionHandler interface.
type MockTunnelConnectionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTunnelConnectionHandlerMockRecorder
}

// MockTunnelConnectionHandlerMockRecorder is the mock recorder for MockTunnelConnectionHandler.
type MockTunnelConnectionHandlerMockRecorder struct {
	mock *MockTunnelConnectionHandler
}

// NewMockTunnelConnectionHandler creates a new mock instance.
func NewMockTunnelConnectionHandler(ctrl *gomock.Controller) *MockTunnelConnectionHandler {
	mock := &MockTunnelConnectionHandler{ctrl: ctrl}
	mock.recorder = &MockTunnelConnectionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTunnelConnectionHandler) EXPECT() *MockTunnelConnectionHandlerMockRecorder {
	return m.recorder
}

// HandleTunnelConnection mocks base method.
func (m *MockTunnelConnectionHandler) HandleTunnelConnection(arg0 *api.AgentInfo, arg1 rpc.ReverseTunnel_ConnectServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleTunnelConnection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleTunnelConnection indicates an expected call of HandleTunnelConnection.
func (mr *MockTunnelConnectionHandlerMockRecorder) HandleTunnelConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTunnelConnection", reflect.TypeOf((*MockTunnelConnectionHandler)(nil).HandleTunnelConnection), arg0, arg1)
}
