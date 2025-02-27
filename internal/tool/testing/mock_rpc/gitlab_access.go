// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/gitlab_access/rpc (interfaces: GitlabAccess_MakeRequestServer)
//
// Generated by this command:
//
//	mockgen -typed -destination gitlab_access.go -package mock_rpc gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/gitlab_access/rpc GitlabAccess_MakeRequestServer
//
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
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) Context() *GitlabAccess_MakeRequestServerContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).Context))
	return &GitlabAccess_MakeRequestServerContextCall{Call: call}
}

// GitlabAccess_MakeRequestServerContextCall wrap *gomock.Call
type GitlabAccess_MakeRequestServerContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestServerContextCall) Return(arg0 context.Context) *GitlabAccess_MakeRequestServerContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestServerContextCall) Do(f func() context.Context) *GitlabAccess_MakeRequestServerContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestServerContextCall) DoAndReturn(f func() context.Context) *GitlabAccess_MakeRequestServerContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
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
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) Recv() *GitlabAccess_MakeRequestServerRecvCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).Recv))
	return &GitlabAccess_MakeRequestServerRecvCall{Call: call}
}

// GitlabAccess_MakeRequestServerRecvCall wrap *gomock.Call
type GitlabAccess_MakeRequestServerRecvCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestServerRecvCall) Return(arg0 *grpctool.HttpRequest, arg1 error) *GitlabAccess_MakeRequestServerRecvCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestServerRecvCall) Do(f func() (*grpctool.HttpRequest, error)) *GitlabAccess_MakeRequestServerRecvCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestServerRecvCall) DoAndReturn(f func() (*grpctool.HttpRequest, error)) *GitlabAccess_MakeRequestServerRecvCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RecvMsg mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) RecvMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) RecvMsg(arg0 any) *GitlabAccess_MakeRequestServerRecvMsgCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).RecvMsg), arg0)
	return &GitlabAccess_MakeRequestServerRecvMsgCall{Call: call}
}

// GitlabAccess_MakeRequestServerRecvMsgCall wrap *gomock.Call
type GitlabAccess_MakeRequestServerRecvMsgCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestServerRecvMsgCall) Return(arg0 error) *GitlabAccess_MakeRequestServerRecvMsgCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestServerRecvMsgCall) Do(f func(any) error) *GitlabAccess_MakeRequestServerRecvMsgCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestServerRecvMsgCall) DoAndReturn(f func(any) error) *GitlabAccess_MakeRequestServerRecvMsgCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Send mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) Send(arg0 *grpctool.HttpResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) Send(arg0 any) *GitlabAccess_MakeRequestServerSendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).Send), arg0)
	return &GitlabAccess_MakeRequestServerSendCall{Call: call}
}

// GitlabAccess_MakeRequestServerSendCall wrap *gomock.Call
type GitlabAccess_MakeRequestServerSendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestServerSendCall) Return(arg0 error) *GitlabAccess_MakeRequestServerSendCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestServerSendCall) Do(f func(*grpctool.HttpResponse) error) *GitlabAccess_MakeRequestServerSendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestServerSendCall) DoAndReturn(f func(*grpctool.HttpResponse) error) *GitlabAccess_MakeRequestServerSendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendHeader mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) SendHeader(arg0 any) *GitlabAccess_MakeRequestServerSendHeaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).SendHeader), arg0)
	return &GitlabAccess_MakeRequestServerSendHeaderCall{Call: call}
}

// GitlabAccess_MakeRequestServerSendHeaderCall wrap *gomock.Call
type GitlabAccess_MakeRequestServerSendHeaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestServerSendHeaderCall) Return(arg0 error) *GitlabAccess_MakeRequestServerSendHeaderCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestServerSendHeaderCall) Do(f func(metadata.MD) error) *GitlabAccess_MakeRequestServerSendHeaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestServerSendHeaderCall) DoAndReturn(f func(metadata.MD) error) *GitlabAccess_MakeRequestServerSendHeaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendMsg mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) SendMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) SendMsg(arg0 any) *GitlabAccess_MakeRequestServerSendMsgCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).SendMsg), arg0)
	return &GitlabAccess_MakeRequestServerSendMsgCall{Call: call}
}

// GitlabAccess_MakeRequestServerSendMsgCall wrap *gomock.Call
type GitlabAccess_MakeRequestServerSendMsgCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestServerSendMsgCall) Return(arg0 error) *GitlabAccess_MakeRequestServerSendMsgCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestServerSendMsgCall) Do(f func(any) error) *GitlabAccess_MakeRequestServerSendMsgCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestServerSendMsgCall) DoAndReturn(f func(any) error) *GitlabAccess_MakeRequestServerSendMsgCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetHeader mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) SetHeader(arg0 any) *GitlabAccess_MakeRequestServerSetHeaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).SetHeader), arg0)
	return &GitlabAccess_MakeRequestServerSetHeaderCall{Call: call}
}

// GitlabAccess_MakeRequestServerSetHeaderCall wrap *gomock.Call
type GitlabAccess_MakeRequestServerSetHeaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestServerSetHeaderCall) Return(arg0 error) *GitlabAccess_MakeRequestServerSetHeaderCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestServerSetHeaderCall) Do(f func(metadata.MD) error) *GitlabAccess_MakeRequestServerSetHeaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestServerSetHeaderCall) DoAndReturn(f func(metadata.MD) error) *GitlabAccess_MakeRequestServerSetHeaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetTrailer mocks base method.
func (m *MockGitlabAccess_MakeRequestServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockGitlabAccess_MakeRequestServerMockRecorder) SetTrailer(arg0 any) *GitlabAccess_MakeRequestServerSetTrailerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockGitlabAccess_MakeRequestServer)(nil).SetTrailer), arg0)
	return &GitlabAccess_MakeRequestServerSetTrailerCall{Call: call}
}

// GitlabAccess_MakeRequestServerSetTrailerCall wrap *gomock.Call
type GitlabAccess_MakeRequestServerSetTrailerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestServerSetTrailerCall) Return() *GitlabAccess_MakeRequestServerSetTrailerCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestServerSetTrailerCall) Do(f func(metadata.MD)) *GitlabAccess_MakeRequestServerSetTrailerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestServerSetTrailerCall) DoAndReturn(f func(metadata.MD)) *GitlabAccess_MakeRequestServerSetTrailerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
