// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/modserver (interfaces: Api,RpcApi)

// Package mock_modserver is a generated GoMock package.
package mock_modserver

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/api"
	retry "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/retry"
	zap "go.uber.org/zap"
)

// MockApi is a mock of Api interface.
type MockApi struct {
	ctrl     *gomock.Controller
	recorder *MockApiMockRecorder
}

// MockApiMockRecorder is the mock recorder for MockApi.
type MockApiMockRecorder struct {
	mock *MockApi
}

// NewMockApi creates a new mock instance.
func NewMockApi(ctrl *gomock.Controller) *MockApi {
	mock := &MockApi{ctrl: ctrl}
	mock.recorder = &MockApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApi) EXPECT() *MockApiMockRecorder {
	return m.recorder
}

// HandleProcessingError mocks base method.
func (m *MockApi) HandleProcessingError(arg0 context.Context, arg1 *zap.Logger, arg2 int64, arg3 string, arg4 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleProcessingError", arg0, arg1, arg2, arg3, arg4)
}

// HandleProcessingError indicates an expected call of HandleProcessingError.
func (mr *MockApiMockRecorder) HandleProcessingError(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleProcessingError", reflect.TypeOf((*MockApi)(nil).HandleProcessingError), arg0, arg1, arg2, arg3, arg4)
}

// MockRpcApi is a mock of RpcApi interface.
type MockRpcApi struct {
	ctrl     *gomock.Controller
	recorder *MockRpcApiMockRecorder
}

// MockRpcApiMockRecorder is the mock recorder for MockRpcApi.
type MockRpcApiMockRecorder struct {
	mock *MockRpcApi
}

// NewMockRpcApi creates a new mock instance.
func NewMockRpcApi(ctrl *gomock.Controller) *MockRpcApi {
	mock := &MockRpcApi{ctrl: ctrl}
	mock.recorder = &MockRpcApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRpcApi) EXPECT() *MockRpcApiMockRecorder {
	return m.recorder
}

// GetAgentInfo mocks base method.
func (m *MockRpcApi) GetAgentInfo(arg0 context.Context, arg1 *zap.Logger) (*api.AgentInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentInfo", arg0, arg1)
	ret0, _ := ret[0].(*api.AgentInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentInfo indicates an expected call of GetAgentInfo.
func (mr *MockRpcApiMockRecorder) GetAgentInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentInfo", reflect.TypeOf((*MockRpcApi)(nil).GetAgentInfo), arg0, arg1)
}

// HandleProcessingError mocks base method.
func (m *MockRpcApi) HandleProcessingError(arg0 *zap.Logger, arg1 int64, arg2 string, arg3 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleProcessingError", arg0, arg1, arg2, arg3)
}

// HandleProcessingError indicates an expected call of HandleProcessingError.
func (mr *MockRpcApiMockRecorder) HandleProcessingError(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleProcessingError", reflect.TypeOf((*MockRpcApi)(nil).HandleProcessingError), arg0, arg1, arg2, arg3)
}

// HandleSendError mocks base method.
func (m *MockRpcApi) HandleSendError(arg0 *zap.Logger, arg1 string, arg2 error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleSendError", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleSendError indicates an expected call of HandleSendError.
func (mr *MockRpcApiMockRecorder) HandleSendError(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSendError", reflect.TypeOf((*MockRpcApi)(nil).HandleSendError), arg0, arg1, arg2)
}

// PollWithBackoff mocks base method.
func (m *MockRpcApi) PollWithBackoff(arg0 retry.PollConfig, arg1 retry.PollWithBackoffFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollWithBackoff", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PollWithBackoff indicates an expected call of PollWithBackoff.
func (mr *MockRpcApiMockRecorder) PollWithBackoff(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollWithBackoff", reflect.TypeOf((*MockRpcApi)(nil).PollWithBackoff), arg0, arg1)
}
