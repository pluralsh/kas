// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/gitops/agent/manifestops (interfaces: Applier)

// Package manifestops is a generated GoMock package.
package manifestops

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	apply "sigs.k8s.io/cli-utils/pkg/apply"
	event "sigs.k8s.io/cli-utils/pkg/apply/event"
	inventory "sigs.k8s.io/cli-utils/pkg/inventory"
	object "sigs.k8s.io/cli-utils/pkg/object"
)

// MockApplier is a mock of Applier interface.
type MockApplier struct {
	ctrl     *gomock.Controller
	recorder *MockApplierMockRecorder
}

// MockApplierMockRecorder is the mock recorder for MockApplier.
type MockApplierMockRecorder struct {
	mock *MockApplier
}

// NewMockApplier creates a new mock instance.
func NewMockApplier(ctrl *gomock.Controller) *MockApplier {
	mock := &MockApplier{ctrl: ctrl}
	mock.recorder = &MockApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplier) EXPECT() *MockApplierMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockApplier) Run(arg0 context.Context, arg1 inventory.Info, arg2 object.UnstructuredSet, arg3 apply.ApplierOptions) <-chan event.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan event.Event)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockApplierMockRecorder) Run(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockApplier)(nil).Run), arg0, arg1, arg2, arg3)
}
