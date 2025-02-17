// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application (interfaces: ApplicationAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	params "github.com/juju/juju/apiserver/params"
)

// MockApplicationAPI is a mock of ApplicationAPI interface.
type MockApplicationAPI struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationAPIMockRecorder
}

// MockApplicationAPIMockRecorder is the mock recorder for MockApplicationAPI.
type MockApplicationAPIMockRecorder struct {
	mock *MockApplicationAPI
}

// NewMockApplicationAPI creates a new mock instance.
func NewMockApplicationAPI(ctrl *gomock.Controller) *MockApplicationAPI {
	mock := &MockApplicationAPI{ctrl: ctrl}
	mock.recorder = &MockApplicationAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationAPI) EXPECT() *MockApplicationAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockApplicationAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockApplicationAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockApplicationAPI)(nil).Close))
}

// Get mocks base method.
func (m *MockApplicationAPI) Get(arg0, arg1 string) (*params.ApplicationGetResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*params.ApplicationGetResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockApplicationAPIMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApplicationAPI)(nil).Get), arg0, arg1)
}

// SetConfig mocks base method.
func (m *MockApplicationAPI) SetConfig(arg0, arg1, arg2 string, arg3 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockApplicationAPIMockRecorder) SetConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockApplicationAPI)(nil).SetConfig), arg0, arg1, arg2, arg3)
}

// UnsetApplicationConfig mocks base method.
func (m *MockApplicationAPI) UnsetApplicationConfig(arg0, arg1 string, arg2 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsetApplicationConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsetApplicationConfig indicates an expected call of UnsetApplicationConfig.
func (mr *MockApplicationAPIMockRecorder) UnsetApplicationConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsetApplicationConfig", reflect.TypeOf((*MockApplicationAPI)(nil).UnsetApplicationConfig), arg0, arg1, arg2)
}
