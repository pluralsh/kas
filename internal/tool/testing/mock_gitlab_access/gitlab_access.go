// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/gitlab_access/rpc (interfaces: GitlabAccessClient,GitlabAccess_MakeRequestClient)
//
// Generated by this command:
//
//	mockgen -typed -destination gitlab_access.go -package mock_gitlab_access gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/gitlab_access/rpc GitlabAccessClient,GitlabAccess_MakeRequestClient
//
// Package mock_gitlab_access is a generated GoMock package.
package mock_gitlab_access

import (
	context "context"
	reflect "reflect"

	rpc "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/gitlab_access/rpc"
	grpctool "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/grpctool"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockGitlabAccessClient is a mock of GitlabAccessClient interface.
type MockGitlabAccessClient struct {
	ctrl     *gomock.Controller
	recorder *MockGitlabAccessClientMockRecorder
}

// MockGitlabAccessClientMockRecorder is the mock recorder for MockGitlabAccessClient.
type MockGitlabAccessClientMockRecorder struct {
	mock *MockGitlabAccessClient
}

// NewMockGitlabAccessClient creates a new mock instance.
func NewMockGitlabAccessClient(ctrl *gomock.Controller) *MockGitlabAccessClient {
	mock := &MockGitlabAccessClient{ctrl: ctrl}
	mock.recorder = &MockGitlabAccessClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitlabAccessClient) EXPECT() *MockGitlabAccessClientMockRecorder {
	return m.recorder
}

// MakeRequest mocks base method.
func (m *MockGitlabAccessClient) MakeRequest(arg0 context.Context, arg1 ...grpc.CallOption) (rpc.GitlabAccess_MakeRequestClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MakeRequest", varargs...)
	ret0, _ := ret[0].(rpc.GitlabAccess_MakeRequestClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeRequest indicates an expected call of MakeRequest.
func (mr *MockGitlabAccessClientMockRecorder) MakeRequest(arg0 any, arg1 ...any) *GitlabAccessClientMakeRequestCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRequest", reflect.TypeOf((*MockGitlabAccessClient)(nil).MakeRequest), varargs...)
	return &GitlabAccessClientMakeRequestCall{Call: call}
}

// GitlabAccessClientMakeRequestCall wrap *gomock.Call
type GitlabAccessClientMakeRequestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccessClientMakeRequestCall) Return(arg0 rpc.GitlabAccess_MakeRequestClient, arg1 error) *GitlabAccessClientMakeRequestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccessClientMakeRequestCall) Do(f func(context.Context, ...grpc.CallOption) (rpc.GitlabAccess_MakeRequestClient, error)) *GitlabAccessClientMakeRequestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccessClientMakeRequestCall) DoAndReturn(f func(context.Context, ...grpc.CallOption) (rpc.GitlabAccess_MakeRequestClient, error)) *GitlabAccessClientMakeRequestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockGitlabAccess_MakeRequestClient is a mock of GitlabAccess_MakeRequestClient interface.
type MockGitlabAccess_MakeRequestClient struct {
	ctrl     *gomock.Controller
	recorder *MockGitlabAccess_MakeRequestClientMockRecorder
}

// MockGitlabAccess_MakeRequestClientMockRecorder is the mock recorder for MockGitlabAccess_MakeRequestClient.
type MockGitlabAccess_MakeRequestClientMockRecorder struct {
	mock *MockGitlabAccess_MakeRequestClient
}

// NewMockGitlabAccess_MakeRequestClient creates a new mock instance.
func NewMockGitlabAccess_MakeRequestClient(ctrl *gomock.Controller) *MockGitlabAccess_MakeRequestClient {
	mock := &MockGitlabAccess_MakeRequestClient{ctrl: ctrl}
	mock.recorder = &MockGitlabAccess_MakeRequestClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitlabAccess_MakeRequestClient) EXPECT() *MockGitlabAccess_MakeRequestClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockGitlabAccess_MakeRequestClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockGitlabAccess_MakeRequestClientMockRecorder) CloseSend() *GitlabAccess_MakeRequestClientCloseSendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockGitlabAccess_MakeRequestClient)(nil).CloseSend))
	return &GitlabAccess_MakeRequestClientCloseSendCall{Call: call}
}

// GitlabAccess_MakeRequestClientCloseSendCall wrap *gomock.Call
type GitlabAccess_MakeRequestClientCloseSendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestClientCloseSendCall) Return(arg0 error) *GitlabAccess_MakeRequestClientCloseSendCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestClientCloseSendCall) Do(f func() error) *GitlabAccess_MakeRequestClientCloseSendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestClientCloseSendCall) DoAndReturn(f func() error) *GitlabAccess_MakeRequestClientCloseSendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Context mocks base method.
func (m *MockGitlabAccess_MakeRequestClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockGitlabAccess_MakeRequestClientMockRecorder) Context() *GitlabAccess_MakeRequestClientContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockGitlabAccess_MakeRequestClient)(nil).Context))
	return &GitlabAccess_MakeRequestClientContextCall{Call: call}
}

// GitlabAccess_MakeRequestClientContextCall wrap *gomock.Call
type GitlabAccess_MakeRequestClientContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestClientContextCall) Return(arg0 context.Context) *GitlabAccess_MakeRequestClientContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestClientContextCall) Do(f func() context.Context) *GitlabAccess_MakeRequestClientContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestClientContextCall) DoAndReturn(f func() context.Context) *GitlabAccess_MakeRequestClientContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Header mocks base method.
func (m *MockGitlabAccess_MakeRequestClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockGitlabAccess_MakeRequestClientMockRecorder) Header() *GitlabAccess_MakeRequestClientHeaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockGitlabAccess_MakeRequestClient)(nil).Header))
	return &GitlabAccess_MakeRequestClientHeaderCall{Call: call}
}

// GitlabAccess_MakeRequestClientHeaderCall wrap *gomock.Call
type GitlabAccess_MakeRequestClientHeaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestClientHeaderCall) Return(arg0 metadata.MD, arg1 error) *GitlabAccess_MakeRequestClientHeaderCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestClientHeaderCall) Do(f func() (metadata.MD, error)) *GitlabAccess_MakeRequestClientHeaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestClientHeaderCall) DoAndReturn(f func() (metadata.MD, error)) *GitlabAccess_MakeRequestClientHeaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Recv mocks base method.
func (m *MockGitlabAccess_MakeRequestClient) Recv() (*grpctool.HttpResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*grpctool.HttpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockGitlabAccess_MakeRequestClientMockRecorder) Recv() *GitlabAccess_MakeRequestClientRecvCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockGitlabAccess_MakeRequestClient)(nil).Recv))
	return &GitlabAccess_MakeRequestClientRecvCall{Call: call}
}

// GitlabAccess_MakeRequestClientRecvCall wrap *gomock.Call
type GitlabAccess_MakeRequestClientRecvCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestClientRecvCall) Return(arg0 *grpctool.HttpResponse, arg1 error) *GitlabAccess_MakeRequestClientRecvCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestClientRecvCall) Do(f func() (*grpctool.HttpResponse, error)) *GitlabAccess_MakeRequestClientRecvCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestClientRecvCall) DoAndReturn(f func() (*grpctool.HttpResponse, error)) *GitlabAccess_MakeRequestClientRecvCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RecvMsg mocks base method.
func (m *MockGitlabAccess_MakeRequestClient) RecvMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockGitlabAccess_MakeRequestClientMockRecorder) RecvMsg(arg0 any) *GitlabAccess_MakeRequestClientRecvMsgCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockGitlabAccess_MakeRequestClient)(nil).RecvMsg), arg0)
	return &GitlabAccess_MakeRequestClientRecvMsgCall{Call: call}
}

// GitlabAccess_MakeRequestClientRecvMsgCall wrap *gomock.Call
type GitlabAccess_MakeRequestClientRecvMsgCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestClientRecvMsgCall) Return(arg0 error) *GitlabAccess_MakeRequestClientRecvMsgCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestClientRecvMsgCall) Do(f func(any) error) *GitlabAccess_MakeRequestClientRecvMsgCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestClientRecvMsgCall) DoAndReturn(f func(any) error) *GitlabAccess_MakeRequestClientRecvMsgCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Send mocks base method.
func (m *MockGitlabAccess_MakeRequestClient) Send(arg0 *grpctool.HttpRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockGitlabAccess_MakeRequestClientMockRecorder) Send(arg0 any) *GitlabAccess_MakeRequestClientSendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockGitlabAccess_MakeRequestClient)(nil).Send), arg0)
	return &GitlabAccess_MakeRequestClientSendCall{Call: call}
}

// GitlabAccess_MakeRequestClientSendCall wrap *gomock.Call
type GitlabAccess_MakeRequestClientSendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestClientSendCall) Return(arg0 error) *GitlabAccess_MakeRequestClientSendCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestClientSendCall) Do(f func(*grpctool.HttpRequest) error) *GitlabAccess_MakeRequestClientSendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestClientSendCall) DoAndReturn(f func(*grpctool.HttpRequest) error) *GitlabAccess_MakeRequestClientSendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendMsg mocks base method.
func (m *MockGitlabAccess_MakeRequestClient) SendMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockGitlabAccess_MakeRequestClientMockRecorder) SendMsg(arg0 any) *GitlabAccess_MakeRequestClientSendMsgCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockGitlabAccess_MakeRequestClient)(nil).SendMsg), arg0)
	return &GitlabAccess_MakeRequestClientSendMsgCall{Call: call}
}

// GitlabAccess_MakeRequestClientSendMsgCall wrap *gomock.Call
type GitlabAccess_MakeRequestClientSendMsgCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestClientSendMsgCall) Return(arg0 error) *GitlabAccess_MakeRequestClientSendMsgCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestClientSendMsgCall) Do(f func(any) error) *GitlabAccess_MakeRequestClientSendMsgCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestClientSendMsgCall) DoAndReturn(f func(any) error) *GitlabAccess_MakeRequestClientSendMsgCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Trailer mocks base method.
func (m *MockGitlabAccess_MakeRequestClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockGitlabAccess_MakeRequestClientMockRecorder) Trailer() *GitlabAccess_MakeRequestClientTrailerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockGitlabAccess_MakeRequestClient)(nil).Trailer))
	return &GitlabAccess_MakeRequestClientTrailerCall{Call: call}
}

// GitlabAccess_MakeRequestClientTrailerCall wrap *gomock.Call
type GitlabAccess_MakeRequestClientTrailerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GitlabAccess_MakeRequestClientTrailerCall) Return(arg0 metadata.MD) *GitlabAccess_MakeRequestClientTrailerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GitlabAccess_MakeRequestClientTrailerCall) Do(f func() metadata.MD) *GitlabAccess_MakeRequestClientTrailerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GitlabAccess_MakeRequestClientTrailerCall) DoAndReturn(f func() metadata.MD) *GitlabAccess_MakeRequestClientTrailerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
