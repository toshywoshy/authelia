// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/authelia/authelia/v4/internal/handlers (interfaces: U2FVerifier)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	u2f "github.com/tstranex/u2f"
)

// MockU2FVerifier is a mock of U2FVerifier interface.
type MockU2FVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockU2FVerifierMockRecorder
}

// MockU2FVerifierMockRecorder is the mock recorder for MockU2FVerifier.
type MockU2FVerifierMockRecorder struct {
	mock *MockU2FVerifier
}

// NewMockU2FVerifier creates a new mock instance.
func NewMockU2FVerifier(ctrl *gomock.Controller) *MockU2FVerifier {
	mock := &MockU2FVerifier{ctrl: ctrl}
	mock.recorder = &MockU2FVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockU2FVerifier) EXPECT() *MockU2FVerifierMockRecorder {
	return m.recorder
}

// Verify mocks base method.
func (m *MockU2FVerifier) Verify(arg0, arg1 []byte, arg2 u2f.SignResponse, arg3 u2f.Challenge) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockU2FVerifierMockRecorder) Verify(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockU2FVerifier)(nil).Verify), arg0, arg1, arg2, arg3)
}
