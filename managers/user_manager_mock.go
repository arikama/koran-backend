// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_manager.go

// Package managers is a generated GoMock package.
package managers

import (
	reflect "reflect"

	beans "github.com/arikama/koran-backend/beans"
	gomock "github.com/golang/mock/gomock"
)

// UserManagerMock is a mock of UserManager interface.
type UserManagerMock struct {
	ctrl     *gomock.Controller
	recorder *UserManagerMockMockRecorder
}

// UserManagerMockMockRecorder is the mock recorder for UserManagerMock.
type UserManagerMockMockRecorder struct {
	mock *UserManagerMock
}

// NewUserManagerMock creates a new mock instance.
func NewUserManagerMock(ctrl *gomock.Controller) *UserManagerMock {
	mock := &UserManagerMock{ctrl: ctrl}
	mock.recorder = &UserManagerMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *UserManagerMock) EXPECT() *UserManagerMockMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *UserManagerMock) CreateUser(email, token string) (*beans.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", email, token)
	ret0, _ := ret[0].(*beans.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *UserManagerMockMockRecorder) CreateUser(email, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*UserManagerMock)(nil).CreateUser), email, token)
}

// GetUser mocks base method.
func (m *UserManagerMock) GetUser(token string) (*beans.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", token)
	ret0, _ := ret[0].(*beans.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *UserManagerMockMockRecorder) GetUser(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*UserManagerMock)(nil).GetUser), token)
}
