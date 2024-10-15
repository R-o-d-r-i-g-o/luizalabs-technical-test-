// Code generated by MockGen. DO NOT EDIT.
// Source: internal/features/auth/service.go

// Package mock is a generated GoMock package.
package mock

import (
	auth "luizalabs-technical-test/internal/features/auth"
	entity "luizalabs-technical-test/internal/pkg/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServiceImp is a mock of ServiceImp interface.
type MockServiceImp struct {
	ctrl     *gomock.Controller
	recorder *MockServiceImpMockRecorder
}

// MockServiceImpMockRecorder is the mock recorder for MockServiceImp.
type MockServiceImpMockRecorder struct {
	mock *MockServiceImp
}

// NewMockServiceImp creates a new mock instance.
func NewMockServiceImp(ctrl *gomock.Controller) *MockServiceImp {
	mock := &MockServiceImp{ctrl: ctrl}
	mock.recorder = &MockServiceImpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceImp) EXPECT() *MockServiceImpMockRecorder {
	return m.recorder
}

// AuthenticateUser mocks base method.
func (m *MockServiceImp) AuthenticateUser(input auth.AuthenticateUserInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateUser", input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateUser indicates an expected call of AuthenticateUser.
func (mr *MockServiceImpMockRecorder) AuthenticateUser(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateUser", reflect.TypeOf((*MockServiceImp)(nil).AuthenticateUser), input)
}

// RegisterUser mocks base method.
func (m *MockServiceImp) RegisterUser(user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockServiceImpMockRecorder) RegisterUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockServiceImp)(nil).RegisterUser), user)
}
