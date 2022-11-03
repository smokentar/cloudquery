// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: TransferClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	transfer "github.com/aws/aws-sdk-go-v2/service/transfer"
	gomock "github.com/golang/mock/gomock"
)

// MockTransferClient is a mock of TransferClient interface.
type MockTransferClient struct {
	ctrl     *gomock.Controller
	recorder *MockTransferClientMockRecorder
}

// MockTransferClientMockRecorder is the mock recorder for MockTransferClient.
type MockTransferClientMockRecorder struct {
	mock *MockTransferClient
}

// NewMockTransferClient creates a new mock instance.
func NewMockTransferClient(ctrl *gomock.Controller) *MockTransferClient {
	mock := &MockTransferClient{ctrl: ctrl}
	mock.recorder = &MockTransferClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferClient) EXPECT() *MockTransferClientMockRecorder {
	return m.recorder
}

// DescribeAccess mocks base method.
func (m *MockTransferClient) DescribeAccess(arg0 context.Context, arg1 *transfer.DescribeAccessInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeAccessOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccess", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeAccessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccess indicates an expected call of DescribeAccess.
func (mr *MockTransferClientMockRecorder) DescribeAccess(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccess", reflect.TypeOf((*MockTransferClient)(nil).DescribeAccess), varargs...)
}

// DescribeAgreement mocks base method.
func (m *MockTransferClient) DescribeAgreement(arg0 context.Context, arg1 *transfer.DescribeAgreementInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeAgreementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAgreement", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeAgreementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAgreement indicates an expected call of DescribeAgreement.
func (mr *MockTransferClientMockRecorder) DescribeAgreement(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAgreement", reflect.TypeOf((*MockTransferClient)(nil).DescribeAgreement), varargs...)
}

// DescribeCertificate mocks base method.
func (m *MockTransferClient) DescribeCertificate(arg0 context.Context, arg1 *transfer.DescribeCertificateInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCertificate", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCertificate indicates an expected call of DescribeCertificate.
func (mr *MockTransferClientMockRecorder) DescribeCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCertificate", reflect.TypeOf((*MockTransferClient)(nil).DescribeCertificate), varargs...)
}

// DescribeConnector mocks base method.
func (m *MockTransferClient) DescribeConnector(arg0 context.Context, arg1 *transfer.DescribeConnectorInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeConnectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnector", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeConnectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnector indicates an expected call of DescribeConnector.
func (mr *MockTransferClientMockRecorder) DescribeConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnector", reflect.TypeOf((*MockTransferClient)(nil).DescribeConnector), varargs...)
}

// DescribeExecution mocks base method.
func (m *MockTransferClient) DescribeExecution(arg0 context.Context, arg1 *transfer.DescribeExecutionInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeExecution", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeExecution indicates an expected call of DescribeExecution.
func (mr *MockTransferClientMockRecorder) DescribeExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeExecution", reflect.TypeOf((*MockTransferClient)(nil).DescribeExecution), varargs...)
}

// DescribeHostKey mocks base method.
func (m *MockTransferClient) DescribeHostKey(arg0 context.Context, arg1 *transfer.DescribeHostKeyInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeHostKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeHostKey", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeHostKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeHostKey indicates an expected call of DescribeHostKey.
func (mr *MockTransferClientMockRecorder) DescribeHostKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHostKey", reflect.TypeOf((*MockTransferClient)(nil).DescribeHostKey), varargs...)
}

// DescribeProfile mocks base method.
func (m *MockTransferClient) DescribeProfile(arg0 context.Context, arg1 *transfer.DescribeProfileInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeProfile", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeProfile indicates an expected call of DescribeProfile.
func (mr *MockTransferClientMockRecorder) DescribeProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProfile", reflect.TypeOf((*MockTransferClient)(nil).DescribeProfile), varargs...)
}

// DescribeSecurityPolicy mocks base method.
func (m *MockTransferClient) DescribeSecurityPolicy(arg0 context.Context, arg1 *transfer.DescribeSecurityPolicyInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeSecurityPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSecurityPolicy", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeSecurityPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecurityPolicy indicates an expected call of DescribeSecurityPolicy.
func (mr *MockTransferClientMockRecorder) DescribeSecurityPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityPolicy", reflect.TypeOf((*MockTransferClient)(nil).DescribeSecurityPolicy), varargs...)
}

// DescribeServer mocks base method.
func (m *MockTransferClient) DescribeServer(arg0 context.Context, arg1 *transfer.DescribeServerInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeServerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServer", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServer indicates an expected call of DescribeServer.
func (mr *MockTransferClientMockRecorder) DescribeServer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServer", reflect.TypeOf((*MockTransferClient)(nil).DescribeServer), varargs...)
}

// DescribeUser mocks base method.
func (m *MockTransferClient) DescribeUser(arg0 context.Context, arg1 *transfer.DescribeUserInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUser", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUser indicates an expected call of DescribeUser.
func (mr *MockTransferClientMockRecorder) DescribeUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUser", reflect.TypeOf((*MockTransferClient)(nil).DescribeUser), varargs...)
}

// DescribeWorkflow mocks base method.
func (m *MockTransferClient) DescribeWorkflow(arg0 context.Context, arg1 *transfer.DescribeWorkflowInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeWorkflowOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkflow", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeWorkflowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkflow indicates an expected call of DescribeWorkflow.
func (mr *MockTransferClientMockRecorder) DescribeWorkflow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkflow", reflect.TypeOf((*MockTransferClient)(nil).DescribeWorkflow), varargs...)
}

// ListAccesses mocks base method.
func (m *MockTransferClient) ListAccesses(arg0 context.Context, arg1 *transfer.ListAccessesInput, arg2 ...func(*transfer.Options)) (*transfer.ListAccessesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccesses", varargs...)
	ret0, _ := ret[0].(*transfer.ListAccessesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccesses indicates an expected call of ListAccesses.
func (mr *MockTransferClientMockRecorder) ListAccesses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccesses", reflect.TypeOf((*MockTransferClient)(nil).ListAccesses), varargs...)
}

// ListAgreements mocks base method.
func (m *MockTransferClient) ListAgreements(arg0 context.Context, arg1 *transfer.ListAgreementsInput, arg2 ...func(*transfer.Options)) (*transfer.ListAgreementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAgreements", varargs...)
	ret0, _ := ret[0].(*transfer.ListAgreementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAgreements indicates an expected call of ListAgreements.
func (mr *MockTransferClientMockRecorder) ListAgreements(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAgreements", reflect.TypeOf((*MockTransferClient)(nil).ListAgreements), varargs...)
}

// ListCertificates mocks base method.
func (m *MockTransferClient) ListCertificates(arg0 context.Context, arg1 *transfer.ListCertificatesInput, arg2 ...func(*transfer.Options)) (*transfer.ListCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCertificates", varargs...)
	ret0, _ := ret[0].(*transfer.ListCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCertificates indicates an expected call of ListCertificates.
func (mr *MockTransferClientMockRecorder) ListCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificates", reflect.TypeOf((*MockTransferClient)(nil).ListCertificates), varargs...)
}

// ListConnectors mocks base method.
func (m *MockTransferClient) ListConnectors(arg0 context.Context, arg1 *transfer.ListConnectorsInput, arg2 ...func(*transfer.Options)) (*transfer.ListConnectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConnectors", varargs...)
	ret0, _ := ret[0].(*transfer.ListConnectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnectors indicates an expected call of ListConnectors.
func (mr *MockTransferClientMockRecorder) ListConnectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnectors", reflect.TypeOf((*MockTransferClient)(nil).ListConnectors), varargs...)
}

// ListExecutions mocks base method.
func (m *MockTransferClient) ListExecutions(arg0 context.Context, arg1 *transfer.ListExecutionsInput, arg2 ...func(*transfer.Options)) (*transfer.ListExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListExecutions", varargs...)
	ret0, _ := ret[0].(*transfer.ListExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExecutions indicates an expected call of ListExecutions.
func (mr *MockTransferClientMockRecorder) ListExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExecutions", reflect.TypeOf((*MockTransferClient)(nil).ListExecutions), varargs...)
}

// ListHostKeys mocks base method.
func (m *MockTransferClient) ListHostKeys(arg0 context.Context, arg1 *transfer.ListHostKeysInput, arg2 ...func(*transfer.Options)) (*transfer.ListHostKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHostKeys", varargs...)
	ret0, _ := ret[0].(*transfer.ListHostKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostKeys indicates an expected call of ListHostKeys.
func (mr *MockTransferClientMockRecorder) ListHostKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostKeys", reflect.TypeOf((*MockTransferClient)(nil).ListHostKeys), varargs...)
}

// ListProfiles mocks base method.
func (m *MockTransferClient) ListProfiles(arg0 context.Context, arg1 *transfer.ListProfilesInput, arg2 ...func(*transfer.Options)) (*transfer.ListProfilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProfiles", varargs...)
	ret0, _ := ret[0].(*transfer.ListProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProfiles indicates an expected call of ListProfiles.
func (mr *MockTransferClientMockRecorder) ListProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProfiles", reflect.TypeOf((*MockTransferClient)(nil).ListProfiles), varargs...)
}

// ListSecurityPolicies mocks base method.
func (m *MockTransferClient) ListSecurityPolicies(arg0 context.Context, arg1 *transfer.ListSecurityPoliciesInput, arg2 ...func(*transfer.Options)) (*transfer.ListSecurityPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSecurityPolicies", varargs...)
	ret0, _ := ret[0].(*transfer.ListSecurityPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecurityPolicies indicates an expected call of ListSecurityPolicies.
func (mr *MockTransferClientMockRecorder) ListSecurityPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecurityPolicies", reflect.TypeOf((*MockTransferClient)(nil).ListSecurityPolicies), varargs...)
}

// ListServers mocks base method.
func (m *MockTransferClient) ListServers(arg0 context.Context, arg1 *transfer.ListServersInput, arg2 ...func(*transfer.Options)) (*transfer.ListServersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServers", varargs...)
	ret0, _ := ret[0].(*transfer.ListServersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServers indicates an expected call of ListServers.
func (mr *MockTransferClientMockRecorder) ListServers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServers", reflect.TypeOf((*MockTransferClient)(nil).ListServers), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockTransferClient) ListTagsForResource(arg0 context.Context, arg1 *transfer.ListTagsForResourceInput, arg2 ...func(*transfer.Options)) (*transfer.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*transfer.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockTransferClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockTransferClient)(nil).ListTagsForResource), varargs...)
}

// ListUsers mocks base method.
func (m *MockTransferClient) ListUsers(arg0 context.Context, arg1 *transfer.ListUsersInput, arg2 ...func(*transfer.Options)) (*transfer.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUsers", varargs...)
	ret0, _ := ret[0].(*transfer.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockTransferClientMockRecorder) ListUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockTransferClient)(nil).ListUsers), varargs...)
}

// ListWorkflows mocks base method.
func (m *MockTransferClient) ListWorkflows(arg0 context.Context, arg1 *transfer.ListWorkflowsInput, arg2 ...func(*transfer.Options)) (*transfer.ListWorkflowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWorkflows", varargs...)
	ret0, _ := ret[0].(*transfer.ListWorkflowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkflows indicates an expected call of ListWorkflows.
func (mr *MockTransferClientMockRecorder) ListWorkflows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflows", reflect.TypeOf((*MockTransferClient)(nil).ListWorkflows), varargs...)
}
