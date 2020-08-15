// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/cluster-integration/gitlab-agent/pkg/gitlab (interfaces: ClientInterface)

// Package mock_gitlab is a generated GoMock package.
package mock_gitlab

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/pkg/api"
)

// MockClientInterface is a mock of ClientInterface interface.
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface.
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance.
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// GetAgentInfo mocks base method.
func (m *MockClientInterface) GetAgentInfo(arg0 context.Context, arg1 *api.AgentMeta) (*api.AgentInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentInfo", arg0, arg1)
	ret0, _ := ret[0].(*api.AgentInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentInfo indicates an expected call of GetAgentInfo.
func (mr *MockClientInterfaceMockRecorder) GetAgentInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentInfo", reflect.TypeOf((*MockClientInterface)(nil).GetAgentInfo), arg0, arg1)
}

// GetProjectInfo mocks base method.
func (m *MockClientInterface) GetProjectInfo(arg0 context.Context, arg1 *api.AgentMeta, arg2 string) (*api.ProjectInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.ProjectInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectInfo indicates an expected call of GetProjectInfo.
func (mr *MockClientInterfaceMockRecorder) GetProjectInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectInfo", reflect.TypeOf((*MockClientInterface)(nil).GetProjectInfo), arg0, arg1, arg2)
}
