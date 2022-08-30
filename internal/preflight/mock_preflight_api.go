// Code generated by MockGen. DO NOT EDIT.
// Source: preflight.go

// Package preflight is a generated GoMock package.
package preflight

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/qbarrand/oot-operator/api/v1beta1"
)

// MockPreflightAPI is a mock of PreflightAPI interface.
type MockPreflightAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPreflightAPIMockRecorder
}

// MockPreflightAPIMockRecorder is the mock recorder for MockPreflightAPI.
type MockPreflightAPIMockRecorder struct {
	mock *MockPreflightAPI
}

// NewMockPreflightAPI creates a new mock instance.
func NewMockPreflightAPI(ctrl *gomock.Controller) *MockPreflightAPI {
	mock := &MockPreflightAPI{ctrl: ctrl}
	mock.recorder = &MockPreflightAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPreflightAPI) EXPECT() *MockPreflightAPIMockRecorder {
	return m.recorder
}

// PreflightUpgradeCheck mocks base method.
func (m *MockPreflightAPI) PreflightUpgradeCheck(ctx context.Context, mod *v1beta1.Module, kernelVersion string) (bool, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreflightUpgradeCheck", ctx, mod, kernelVersion)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// PreflightUpgradeCheck indicates an expected call of PreflightUpgradeCheck.
func (mr *MockPreflightAPIMockRecorder) PreflightUpgradeCheck(ctx, mod, kernelVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreflightUpgradeCheck", reflect.TypeOf((*MockPreflightAPI)(nil).PreflightUpgradeCheck), ctx, mod, kernelVersion)
}
