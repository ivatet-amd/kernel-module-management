// Code generated by MockGen. DO NOT EDIT.
// Source: helper.go

// Package ca is a generated GoMock package.
package ca

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockHelper is a mock of Helper interface.
type MockHelper struct {
	ctrl     *gomock.Controller
	recorder *MockHelperMockRecorder
}

// MockHelperMockRecorder is the mock recorder for MockHelper.
type MockHelperMockRecorder struct {
	mock *MockHelper
}

// NewMockHelper creates a new mock instance.
func NewMockHelper(ctrl *gomock.Controller) *MockHelper {
	mock := &MockHelper{ctrl: ctrl}
	mock.recorder = &MockHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelper) EXPECT() *MockHelperMockRecorder {
	return m.recorder
}

// GetClusterCA mocks base method.
func (m *MockHelper) GetClusterCA(ctx context.Context, namespace string) (*ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterCA", ctx, namespace)
	ret0, _ := ret[0].(*ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterCA indicates an expected call of GetClusterCA.
func (mr *MockHelperMockRecorder) GetClusterCA(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterCA", reflect.TypeOf((*MockHelper)(nil).GetClusterCA), ctx, namespace)
}

// GetServiceCA mocks base method.
func (m *MockHelper) GetServiceCA(ctx context.Context, namespace string) (*ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceCA", ctx, namespace)
	ret0, _ := ret[0].(*ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceCA indicates an expected call of GetServiceCA.
func (mr *MockHelperMockRecorder) GetServiceCA(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceCA", reflect.TypeOf((*MockHelper)(nil).GetServiceCA), ctx, namespace)
}

// Sync mocks base method.
func (m *MockHelper) Sync(ctx context.Context, namespace string, owner client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", ctx, namespace, owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockHelperMockRecorder) Sync(ctx, namespace, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockHelper)(nil).Sync), ctx, namespace, owner)
}