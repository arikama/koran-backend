// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_dao.go

// Package daos is a generated GoMock package.
package daos

import (
	reflect "reflect"

	models "github.com/arikama/koran-backend/models"
	gomock "github.com/golang/mock/gomock"
)

// UserDaoMock is a mock of UserDao interface.
type UserDaoMock struct {
	ctrl     *gomock.Controller
	recorder *UserDaoMockMockRecorder
}

// UserDaoMockMockRecorder is the mock recorder for UserDaoMock.
type UserDaoMockMockRecorder struct {
	mock *UserDaoMock
}

// NewUserDaoMock creates a new mock instance.
func NewUserDaoMock(ctrl *gomock.Controller) *UserDaoMock {
	mock := &UserDaoMock{ctrl: ctrl}
	mock.recorder = &UserDaoMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *UserDaoMock) EXPECT() *UserDaoMockMockRecorder {
	return m.recorder
}

// QueryUser mocks base method.
func (m *UserDaoMock) QueryUser(token string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUser", token)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUser indicates an expected call of QueryUser.
func (mr *UserDaoMockMockRecorder) QueryUser(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUser", reflect.TypeOf((*UserDaoMock)(nil).QueryUser), token)
}

// UpsertUser mocks base method.
func (m *UserDaoMock) UpsertUser(user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertUser", user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertUser indicates an expected call of UpsertUser.
func (mr *UserDaoMockMockRecorder) UpsertUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertUser", reflect.TypeOf((*UserDaoMock)(nil).UpsertUser), user)
}