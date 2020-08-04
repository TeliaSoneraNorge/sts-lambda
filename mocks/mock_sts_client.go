// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/telia-oss/concourse-sts-lambda (interfaces: STSClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	sts "github.com/aws/aws-sdk-go/service/sts"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSTSClient is a mock of STSClient interface
type MockSTSClient struct {
	ctrl     *gomock.Controller
	recorder *MockSTSClientMockRecorder
}

// MockSTSClientMockRecorder is the mock recorder for MockSTSClient
type MockSTSClientMockRecorder struct {
	mock *MockSTSClient
}

// NewMockSTSClient creates a new mock instance
func NewMockSTSClient(ctrl *gomock.Controller) *MockSTSClient {
	mock := &MockSTSClient{ctrl: ctrl}
	mock.recorder = &MockSTSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSTSClient) EXPECT() *MockSTSClientMockRecorder {
	return m.recorder
}

// AssumeRole mocks base method
func (m *MockSTSClient) AssumeRole(arg0 *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRole", arg0)
	ret0, _ := ret[0].(*sts.AssumeRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRole indicates an expected call of AssumeRole
func (mr *MockSTSClientMockRecorder) AssumeRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRole", reflect.TypeOf((*MockSTSClient)(nil).AssumeRole), arg0)
}

// AssumeRoleRequest mocks base method
func (m *MockSTSClient) AssumeRoleRequest(arg0 *sts.AssumeRoleInput) (*request.Request, *sts.AssumeRoleOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.AssumeRoleOutput)
	return ret0, ret1
}

// AssumeRoleRequest indicates an expected call of AssumeRoleRequest
func (mr *MockSTSClientMockRecorder) AssumeRoleRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleRequest", reflect.TypeOf((*MockSTSClient)(nil).AssumeRoleRequest), arg0)
}

// AssumeRoleWithContext mocks base method
func (m *MockSTSClient) AssumeRoleWithContext(arg0 context.Context, arg1 *sts.AssumeRoleInput, arg2 ...request.Option) (*sts.AssumeRoleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssumeRoleWithContext", varargs...)
	ret0, _ := ret[0].(*sts.AssumeRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithContext indicates an expected call of AssumeRoleWithContext
func (mr *MockSTSClientMockRecorder) AssumeRoleWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithContext", reflect.TypeOf((*MockSTSClient)(nil).AssumeRoleWithContext), varargs...)
}

// AssumeRoleWithSAML mocks base method
func (m *MockSTSClient) AssumeRoleWithSAML(arg0 *sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleWithSAML", arg0)
	ret0, _ := ret[0].(*sts.AssumeRoleWithSAMLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithSAML indicates an expected call of AssumeRoleWithSAML
func (mr *MockSTSClientMockRecorder) AssumeRoleWithSAML(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithSAML", reflect.TypeOf((*MockSTSClient)(nil).AssumeRoleWithSAML), arg0)
}

// AssumeRoleWithSAMLRequest mocks base method
func (m *MockSTSClient) AssumeRoleWithSAMLRequest(arg0 *sts.AssumeRoleWithSAMLInput) (*request.Request, *sts.AssumeRoleWithSAMLOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleWithSAMLRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.AssumeRoleWithSAMLOutput)
	return ret0, ret1
}

// AssumeRoleWithSAMLRequest indicates an expected call of AssumeRoleWithSAMLRequest
func (mr *MockSTSClientMockRecorder) AssumeRoleWithSAMLRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithSAMLRequest", reflect.TypeOf((*MockSTSClient)(nil).AssumeRoleWithSAMLRequest), arg0)
}

// AssumeRoleWithSAMLWithContext mocks base method
func (m *MockSTSClient) AssumeRoleWithSAMLWithContext(arg0 context.Context, arg1 *sts.AssumeRoleWithSAMLInput, arg2 ...request.Option) (*sts.AssumeRoleWithSAMLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssumeRoleWithSAMLWithContext", varargs...)
	ret0, _ := ret[0].(*sts.AssumeRoleWithSAMLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithSAMLWithContext indicates an expected call of AssumeRoleWithSAMLWithContext
func (mr *MockSTSClientMockRecorder) AssumeRoleWithSAMLWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithSAMLWithContext", reflect.TypeOf((*MockSTSClient)(nil).AssumeRoleWithSAMLWithContext), varargs...)
}

// AssumeRoleWithWebIdentity mocks base method
func (m *MockSTSClient) AssumeRoleWithWebIdentity(arg0 *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleWithWebIdentity", arg0)
	ret0, _ := ret[0].(*sts.AssumeRoleWithWebIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithWebIdentity indicates an expected call of AssumeRoleWithWebIdentity
func (mr *MockSTSClientMockRecorder) AssumeRoleWithWebIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithWebIdentity", reflect.TypeOf((*MockSTSClient)(nil).AssumeRoleWithWebIdentity), arg0)
}

// AssumeRoleWithWebIdentityRequest mocks base method
func (m *MockSTSClient) AssumeRoleWithWebIdentityRequest(arg0 *sts.AssumeRoleWithWebIdentityInput) (*request.Request, *sts.AssumeRoleWithWebIdentityOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleWithWebIdentityRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.AssumeRoleWithWebIdentityOutput)
	return ret0, ret1
}

// AssumeRoleWithWebIdentityRequest indicates an expected call of AssumeRoleWithWebIdentityRequest
func (mr *MockSTSClientMockRecorder) AssumeRoleWithWebIdentityRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithWebIdentityRequest", reflect.TypeOf((*MockSTSClient)(nil).AssumeRoleWithWebIdentityRequest), arg0)
}

// AssumeRoleWithWebIdentityWithContext mocks base method
func (m *MockSTSClient) AssumeRoleWithWebIdentityWithContext(arg0 context.Context, arg1 *sts.AssumeRoleWithWebIdentityInput, arg2 ...request.Option) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssumeRoleWithWebIdentityWithContext", varargs...)
	ret0, _ := ret[0].(*sts.AssumeRoleWithWebIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithWebIdentityWithContext indicates an expected call of AssumeRoleWithWebIdentityWithContext
func (mr *MockSTSClientMockRecorder) AssumeRoleWithWebIdentityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithWebIdentityWithContext", reflect.TypeOf((*MockSTSClient)(nil).AssumeRoleWithWebIdentityWithContext), varargs...)
}

// DecodeAuthorizationMessage mocks base method
func (m *MockSTSClient) DecodeAuthorizationMessage(arg0 *sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeAuthorizationMessage", arg0)
	ret0, _ := ret[0].(*sts.DecodeAuthorizationMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeAuthorizationMessage indicates an expected call of DecodeAuthorizationMessage
func (mr *MockSTSClientMockRecorder) DecodeAuthorizationMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeAuthorizationMessage", reflect.TypeOf((*MockSTSClient)(nil).DecodeAuthorizationMessage), arg0)
}

// DecodeAuthorizationMessageRequest mocks base method
func (m *MockSTSClient) DecodeAuthorizationMessageRequest(arg0 *sts.DecodeAuthorizationMessageInput) (*request.Request, *sts.DecodeAuthorizationMessageOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeAuthorizationMessageRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.DecodeAuthorizationMessageOutput)
	return ret0, ret1
}

// DecodeAuthorizationMessageRequest indicates an expected call of DecodeAuthorizationMessageRequest
func (mr *MockSTSClientMockRecorder) DecodeAuthorizationMessageRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeAuthorizationMessageRequest", reflect.TypeOf((*MockSTSClient)(nil).DecodeAuthorizationMessageRequest), arg0)
}

// DecodeAuthorizationMessageWithContext mocks base method
func (m *MockSTSClient) DecodeAuthorizationMessageWithContext(arg0 context.Context, arg1 *sts.DecodeAuthorizationMessageInput, arg2 ...request.Option) (*sts.DecodeAuthorizationMessageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DecodeAuthorizationMessageWithContext", varargs...)
	ret0, _ := ret[0].(*sts.DecodeAuthorizationMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeAuthorizationMessageWithContext indicates an expected call of DecodeAuthorizationMessageWithContext
func (mr *MockSTSClientMockRecorder) DecodeAuthorizationMessageWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeAuthorizationMessageWithContext", reflect.TypeOf((*MockSTSClient)(nil).DecodeAuthorizationMessageWithContext), varargs...)
}

// GetAccessKeyInfo mocks base method
func (m *MockSTSClient) GetAccessKeyInfo(arg0 *sts.GetAccessKeyInfoInput) (*sts.GetAccessKeyInfoOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessKeyInfo", arg0)
	ret0, _ := ret[0].(*sts.GetAccessKeyInfoOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessKeyInfo indicates an expected call of GetAccessKeyInfo
func (mr *MockSTSClientMockRecorder) GetAccessKeyInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKeyInfo", reflect.TypeOf((*MockSTSClient)(nil).GetAccessKeyInfo), arg0)
}

// GetAccessKeyInfoRequest mocks base method
func (m *MockSTSClient) GetAccessKeyInfoRequest(arg0 *sts.GetAccessKeyInfoInput) (*request.Request, *sts.GetAccessKeyInfoOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessKeyInfoRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetAccessKeyInfoOutput)
	return ret0, ret1
}

// GetAccessKeyInfoRequest indicates an expected call of GetAccessKeyInfoRequest
func (mr *MockSTSClientMockRecorder) GetAccessKeyInfoRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKeyInfoRequest", reflect.TypeOf((*MockSTSClient)(nil).GetAccessKeyInfoRequest), arg0)
}

// GetAccessKeyInfoWithContext mocks base method
func (m *MockSTSClient) GetAccessKeyInfoWithContext(arg0 context.Context, arg1 *sts.GetAccessKeyInfoInput, arg2 ...request.Option) (*sts.GetAccessKeyInfoOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessKeyInfoWithContext", varargs...)
	ret0, _ := ret[0].(*sts.GetAccessKeyInfoOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessKeyInfoWithContext indicates an expected call of GetAccessKeyInfoWithContext
func (mr *MockSTSClientMockRecorder) GetAccessKeyInfoWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKeyInfoWithContext", reflect.TypeOf((*MockSTSClient)(nil).GetAccessKeyInfoWithContext), varargs...)
}

// GetCallerIdentity mocks base method
func (m *MockSTSClient) GetCallerIdentity(arg0 *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallerIdentity", arg0)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentity indicates an expected call of GetCallerIdentity
func (mr *MockSTSClientMockRecorder) GetCallerIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentity", reflect.TypeOf((*MockSTSClient)(nil).GetCallerIdentity), arg0)
}

// GetCallerIdentityRequest mocks base method
func (m *MockSTSClient) GetCallerIdentityRequest(arg0 *sts.GetCallerIdentityInput) (*request.Request, *sts.GetCallerIdentityOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallerIdentityRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetCallerIdentityOutput)
	return ret0, ret1
}

// GetCallerIdentityRequest indicates an expected call of GetCallerIdentityRequest
func (mr *MockSTSClientMockRecorder) GetCallerIdentityRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentityRequest", reflect.TypeOf((*MockSTSClient)(nil).GetCallerIdentityRequest), arg0)
}

// GetCallerIdentityWithContext mocks base method
func (m *MockSTSClient) GetCallerIdentityWithContext(arg0 context.Context, arg1 *sts.GetCallerIdentityInput, arg2 ...request.Option) (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCallerIdentityWithContext", varargs...)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentityWithContext indicates an expected call of GetCallerIdentityWithContext
func (mr *MockSTSClientMockRecorder) GetCallerIdentityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentityWithContext", reflect.TypeOf((*MockSTSClient)(nil).GetCallerIdentityWithContext), varargs...)
}

// GetFederationToken mocks base method
func (m *MockSTSClient) GetFederationToken(arg0 *sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFederationToken", arg0)
	ret0, _ := ret[0].(*sts.GetFederationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederationToken indicates an expected call of GetFederationToken
func (mr *MockSTSClientMockRecorder) GetFederationToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederationToken", reflect.TypeOf((*MockSTSClient)(nil).GetFederationToken), arg0)
}

// GetFederationTokenRequest mocks base method
func (m *MockSTSClient) GetFederationTokenRequest(arg0 *sts.GetFederationTokenInput) (*request.Request, *sts.GetFederationTokenOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFederationTokenRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetFederationTokenOutput)
	return ret0, ret1
}

// GetFederationTokenRequest indicates an expected call of GetFederationTokenRequest
func (mr *MockSTSClientMockRecorder) GetFederationTokenRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederationTokenRequest", reflect.TypeOf((*MockSTSClient)(nil).GetFederationTokenRequest), arg0)
}

// GetFederationTokenWithContext mocks base method
func (m *MockSTSClient) GetFederationTokenWithContext(arg0 context.Context, arg1 *sts.GetFederationTokenInput, arg2 ...request.Option) (*sts.GetFederationTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFederationTokenWithContext", varargs...)
	ret0, _ := ret[0].(*sts.GetFederationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederationTokenWithContext indicates an expected call of GetFederationTokenWithContext
func (mr *MockSTSClientMockRecorder) GetFederationTokenWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederationTokenWithContext", reflect.TypeOf((*MockSTSClient)(nil).GetFederationTokenWithContext), varargs...)
}

// GetSessionToken mocks base method
func (m *MockSTSClient) GetSessionToken(arg0 *sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionToken", arg0)
	ret0, _ := ret[0].(*sts.GetSessionTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionToken indicates an expected call of GetSessionToken
func (mr *MockSTSClientMockRecorder) GetSessionToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionToken", reflect.TypeOf((*MockSTSClient)(nil).GetSessionToken), arg0)
}

// GetSessionTokenRequest mocks base method
func (m *MockSTSClient) GetSessionTokenRequest(arg0 *sts.GetSessionTokenInput) (*request.Request, *sts.GetSessionTokenOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionTokenRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetSessionTokenOutput)
	return ret0, ret1
}

// GetSessionTokenRequest indicates an expected call of GetSessionTokenRequest
func (mr *MockSTSClientMockRecorder) GetSessionTokenRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionTokenRequest", reflect.TypeOf((*MockSTSClient)(nil).GetSessionTokenRequest), arg0)
}

// GetSessionTokenWithContext mocks base method
func (m *MockSTSClient) GetSessionTokenWithContext(arg0 context.Context, arg1 *sts.GetSessionTokenInput, arg2 ...request.Option) (*sts.GetSessionTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSessionTokenWithContext", varargs...)
	ret0, _ := ret[0].(*sts.GetSessionTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionTokenWithContext indicates an expected call of GetSessionTokenWithContext
func (mr *MockSTSClientMockRecorder) GetSessionTokenWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionTokenWithContext", reflect.TypeOf((*MockSTSClient)(nil).GetSessionTokenWithContext), varargs...)
}
