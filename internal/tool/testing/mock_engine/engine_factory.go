// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/cluster-integration/gitlab-agent/internal/module/gitops/gitops_agent (interfaces: GitOpsEngineFactory)

// Package mock_engine is a generated GoMock package.
package mock_engine

import (
	reflect "reflect"

	cache "github.com/argoproj/gitops-engine/pkg/cache"
	engine "github.com/argoproj/gitops-engine/pkg/engine"
	gomock "github.com/golang/mock/gomock"
)

// MockGitOpsEngineFactory is a mock of GitOpsEngineFactory interface
type MockGitOpsEngineFactory struct {
	ctrl     *gomock.Controller
	recorder *MockGitOpsEngineFactoryMockRecorder
}

// MockGitOpsEngineFactoryMockRecorder is the mock recorder for MockGitOpsEngineFactory
type MockGitOpsEngineFactoryMockRecorder struct {
	mock *MockGitOpsEngineFactory
}

// NewMockGitOpsEngineFactory creates a new mock instance
func NewMockGitOpsEngineFactory(ctrl *gomock.Controller) *MockGitOpsEngineFactory {
	mock := &MockGitOpsEngineFactory{ctrl: ctrl}
	mock.recorder = &MockGitOpsEngineFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGitOpsEngineFactory) EXPECT() *MockGitOpsEngineFactoryMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockGitOpsEngineFactory) New(arg0 []engine.Option, arg1 []cache.UpdateSettingsFunc) engine.GitOpsEngine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0, arg1)
	ret0, _ := ret[0].(engine.GitOpsEngine)
	return ret0
}

// New indicates an expected call of New
func (mr *MockGitOpsEngineFactoryMockRecorder) New(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockGitOpsEngineFactory)(nil).New), arg0, arg1)
}
