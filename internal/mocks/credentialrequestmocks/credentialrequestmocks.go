// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
//

// Code generated by MockGen. DO NOT EDIT.
// Source: go.pinniped.dev/internal/registry/credentialrequest (interfaces: CertIssuer,TokenCredentialRequestAuthenticator)

// Package credentialrequestmocks is a generated GoMock package.
package credentialrequestmocks

import (
	context "context"
	pkix "crypto/x509/pkix"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	login "go.pinniped.dev/generated/1.19/apis/concierge/login"
	user "k8s.io/apiserver/pkg/authentication/user"
)

// MockCertIssuer is a mock of CertIssuer interface
type MockCertIssuer struct {
	ctrl     *gomock.Controller
	recorder *MockCertIssuerMockRecorder
}

// MockCertIssuerMockRecorder is the mock recorder for MockCertIssuer
type MockCertIssuerMockRecorder struct {
	mock *MockCertIssuer
}

// NewMockCertIssuer creates a new mock instance
func NewMockCertIssuer(ctrl *gomock.Controller) *MockCertIssuer {
	mock := &MockCertIssuer{ctrl: ctrl}
	mock.recorder = &MockCertIssuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertIssuer) EXPECT() *MockCertIssuerMockRecorder {
	return m.recorder
}

// IssuePEM mocks base method
func (m *MockCertIssuer) IssuePEM(arg0 pkix.Name, arg1 []string, arg2 time.Duration) ([]byte, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssuePEM", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IssuePEM indicates an expected call of IssuePEM
func (mr *MockCertIssuerMockRecorder) IssuePEM(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssuePEM", reflect.TypeOf((*MockCertIssuer)(nil).IssuePEM), arg0, arg1, arg2)
}

// MockTokenCredentialRequestAuthenticator is a mock of TokenCredentialRequestAuthenticator interface
type MockTokenCredentialRequestAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenCredentialRequestAuthenticatorMockRecorder
}

// MockTokenCredentialRequestAuthenticatorMockRecorder is the mock recorder for MockTokenCredentialRequestAuthenticator
type MockTokenCredentialRequestAuthenticatorMockRecorder struct {
	mock *MockTokenCredentialRequestAuthenticator
}

// NewMockTokenCredentialRequestAuthenticator creates a new mock instance
func NewMockTokenCredentialRequestAuthenticator(ctrl *gomock.Controller) *MockTokenCredentialRequestAuthenticator {
	mock := &MockTokenCredentialRequestAuthenticator{ctrl: ctrl}
	mock.recorder = &MockTokenCredentialRequestAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenCredentialRequestAuthenticator) EXPECT() *MockTokenCredentialRequestAuthenticatorMockRecorder {
	return m.recorder
}

// AuthenticateTokenCredentialRequest mocks base method
func (m *MockTokenCredentialRequestAuthenticator) AuthenticateTokenCredentialRequest(arg0 context.Context, arg1 *login.TokenCredentialRequest) (user.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateTokenCredentialRequest", arg0, arg1)
	ret0, _ := ret[0].(user.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateTokenCredentialRequest indicates an expected call of AuthenticateTokenCredentialRequest
func (mr *MockTokenCredentialRequestAuthenticatorMockRecorder) AuthenticateTokenCredentialRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateTokenCredentialRequest", reflect.TypeOf((*MockTokenCredentialRequestAuthenticator)(nil).AuthenticateTokenCredentialRequest), arg0, arg1)
}
