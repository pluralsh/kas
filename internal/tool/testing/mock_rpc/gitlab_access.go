// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/gitlab_access/rpc (interfaces: GitlabAccess_MakeRequestServer)

// Package mock_rpc is a generated GoMock package.
package mock_rpc

import (
	context "context"
	reflect "reflect"

	grpctool "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/grpctool"
	gomock "go.uber.org/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
)

// MockGitlabAccess_MakeRequestServer is a mock of GitlabAccess_MakeRequestServer interface.
type MockGitlabAccess_MakeRequestServer struct {
	ctrl     *gomock.Controller
	recorder *MockGitlabAccess_MakeRequestServerMockRecorder
}

// MockGitlabAccess_MakeRequestServerMockRecorder is the mock recorder for MockGitlabAccess_MakeRequestServer.
type MockGitlabAccess_MakeRequestServerMockRecorder struct {
	mock *MockGitlabAccess_MakeRequestServer
}

// NewMockGitlabAccess_MakeRequestServer creates a new mock instance.
func NewMockGitlabAccess_MakeRequestServer(ctrl *gomock.Controller) *MockGitlabAccess_MakeRequestServer {
	mock := &MockGitlabAccess_MakeRequestServer{ctrl: ctrl}
	mock.recorder = &MockGitlabAccess_MakeRequestServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitlabAccess_MakeRequestServer) EXPECT() *MockGitlabAccess_MakeRequestServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) Recv() (*grpctool.HttpRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*grpctool.HttpRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) Send(arg0 *grpctool.HttpResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).SetTrailer), arg0)
}
