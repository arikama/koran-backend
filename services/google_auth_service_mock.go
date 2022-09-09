// Code generated by MockGen. DO NOT EDIT.
// Source: ./google_auth_service.go

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// GoogleAuthServiceMock is a mock of GoogleAuthService interface.
type GoogleAuthServiceMock struct {
	ctrl     *gomock.Controller
	recorder *GoogleAuthServiceMockMockRecorder
}

// GoogleAuthServiceMockMockRecorder is the mock recorder for GoogleAuthServiceMock.
type GoogleAuthServiceMockMockRecorder struct {
	mock *GoogleAuthServiceMock
}

// NewGoogleAuthServiceMock creates a new mock instance.
func NewGoogleAuthServiceMock(ctrl *gomock.Controller) *GoogleAuthServiceMock {
	mock := &GoogleAuthServiceMock{ctrl: ctrl}
	mock.recorder = &GoogleAuthServiceMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *GoogleAuthServiceMock) EXPECT() *GoogleAuthServiceMockMockRecorder {
	return m.recorder
}

// AuthUserCode mocks base method.
func (m *GoogleAuthServiceMock) AuthUserCode(userAuthCode string) (*GoogleUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUserCode", userAuthCode)
	ret0, _ := ret[0].(*GoogleUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthUserCode indicates an expected call of AuthUserCode.
func (mr *GoogleAuthServiceMockMockRecorder) AuthUserCode(userAuthCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUserCode", reflect.TypeOf((*GoogleAuthServiceMock)(nil).AuthUserCode), userAuthCode)
}
