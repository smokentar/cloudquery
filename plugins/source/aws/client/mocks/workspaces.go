// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: WorkspacesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	workspaces "github.com/aws/aws-sdk-go-v2/service/workspaces"
	gomock "github.com/golang/mock/gomock"
)

// MockWorkspacesClient is a mock of WorkspacesClient interface.
type MockWorkspacesClient struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspacesClientMockRecorder
}

// MockWorkspacesClientMockRecorder is the mock recorder for MockWorkspacesClient.
type MockWorkspacesClientMockRecorder struct {
	mock *MockWorkspacesClient
}

// NewMockWorkspacesClient creates a new mock instance.
func NewMockWorkspacesClient(ctrl *gomock.Controller) *MockWorkspacesClient {
	mock := &MockWorkspacesClient{ctrl: ctrl}
	mock.recorder = &MockWorkspacesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspacesClient) EXPECT() *MockWorkspacesClientMockRecorder {
	return m.recorder
}

// DescribeAccount mocks base method.
func (m *MockWorkspacesClient) DescribeAccount(arg0 context.Context, arg1 *workspaces.DescribeAccountInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccount", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccount indicates an expected call of DescribeAccount.
func (mr *MockWorkspacesClientMockRecorder) DescribeAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccount", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeAccount), varargs...)
}

// DescribeAccountModifications mocks base method.
func (m *MockWorkspacesClient) DescribeAccountModifications(arg0 context.Context, arg1 *workspaces.DescribeAccountModificationsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeAccountModificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccountModifications", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeAccountModificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccountModifications indicates an expected call of DescribeAccountModifications.
func (mr *MockWorkspacesClientMockRecorder) DescribeAccountModifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccountModifications", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeAccountModifications), varargs...)
}

// DescribeClientBranding mocks base method.
func (m *MockWorkspacesClient) DescribeClientBranding(arg0 context.Context, arg1 *workspaces.DescribeClientBrandingInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeClientBrandingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClientBranding", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeClientBrandingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeClientBranding indicates an expected call of DescribeClientBranding.
func (mr *MockWorkspacesClientMockRecorder) DescribeClientBranding(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClientBranding", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeClientBranding), varargs...)
}

// DescribeClientProperties mocks base method.
func (m *MockWorkspacesClient) DescribeClientProperties(arg0 context.Context, arg1 *workspaces.DescribeClientPropertiesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeClientPropertiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClientProperties", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeClientPropertiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeClientProperties indicates an expected call of DescribeClientProperties.
func (mr *MockWorkspacesClientMockRecorder) DescribeClientProperties(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClientProperties", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeClientProperties), varargs...)
}

// DescribeConnectClientAddIns mocks base method.
func (m *MockWorkspacesClient) DescribeConnectClientAddIns(arg0 context.Context, arg1 *workspaces.DescribeConnectClientAddInsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeConnectClientAddInsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnectClientAddIns", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeConnectClientAddInsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnectClientAddIns indicates an expected call of DescribeConnectClientAddIns.
func (mr *MockWorkspacesClientMockRecorder) DescribeConnectClientAddIns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnectClientAddIns", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeConnectClientAddIns), varargs...)
}

// DescribeConnectionAliasPermissions mocks base method.
func (m *MockWorkspacesClient) DescribeConnectionAliasPermissions(arg0 context.Context, arg1 *workspaces.DescribeConnectionAliasPermissionsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeConnectionAliasPermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnectionAliasPermissions", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeConnectionAliasPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnectionAliasPermissions indicates an expected call of DescribeConnectionAliasPermissions.
func (mr *MockWorkspacesClientMockRecorder) DescribeConnectionAliasPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnectionAliasPermissions", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeConnectionAliasPermissions), varargs...)
}

// DescribeConnectionAliases mocks base method.
func (m *MockWorkspacesClient) DescribeConnectionAliases(arg0 context.Context, arg1 *workspaces.DescribeConnectionAliasesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeConnectionAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnectionAliases", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeConnectionAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnectionAliases indicates an expected call of DescribeConnectionAliases.
func (mr *MockWorkspacesClientMockRecorder) DescribeConnectionAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnectionAliases", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeConnectionAliases), varargs...)
}

// DescribeIpGroups mocks base method.
func (m *MockWorkspacesClient) DescribeIpGroups(arg0 context.Context, arg1 *workspaces.DescribeIpGroupsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeIpGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeIpGroups", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeIpGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeIpGroups indicates an expected call of DescribeIpGroups.
func (mr *MockWorkspacesClientMockRecorder) DescribeIpGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeIpGroups", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeIpGroups), varargs...)
}

// DescribeTags mocks base method.
func (m *MockWorkspacesClient) DescribeTags(arg0 context.Context, arg1 *workspaces.DescribeTagsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTags", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTags indicates an expected call of DescribeTags.
func (mr *MockWorkspacesClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTags", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeTags), varargs...)
}

// DescribeWorkspaceBundles mocks base method.
func (m *MockWorkspacesClient) DescribeWorkspaceBundles(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceBundlesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceBundlesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkspaceBundles", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceBundlesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkspaceBundles indicates an expected call of DescribeWorkspaceBundles.
func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkspaceBundles", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceBundles), varargs...)
}

// DescribeWorkspaceDirectories mocks base method.
func (m *MockWorkspacesClient) DescribeWorkspaceDirectories(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceDirectoriesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceDirectoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkspaceDirectories", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceDirectoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkspaceDirectories indicates an expected call of DescribeWorkspaceDirectories.
func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceDirectories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkspaceDirectories", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceDirectories), varargs...)
}

// DescribeWorkspaceImagePermissions mocks base method.
func (m *MockWorkspacesClient) DescribeWorkspaceImagePermissions(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceImagePermissionsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkspaceImagePermissions", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceImagePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkspaceImagePermissions indicates an expected call of DescribeWorkspaceImagePermissions.
func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceImagePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkspaceImagePermissions", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceImagePermissions), varargs...)
}

// DescribeWorkspaceImages mocks base method.
func (m *MockWorkspacesClient) DescribeWorkspaceImages(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceImagesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkspaceImages", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkspaceImages indicates an expected call of DescribeWorkspaceImages.
func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkspaceImages", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceImages), varargs...)
}

// DescribeWorkspaceSnapshots mocks base method.
func (m *MockWorkspacesClient) DescribeWorkspaceSnapshots(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceSnapshotsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkspaceSnapshots", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkspaceSnapshots indicates an expected call of DescribeWorkspaceSnapshots.
func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkspaceSnapshots", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceSnapshots), varargs...)
}

// DescribeWorkspaces mocks base method.
func (m *MockWorkspacesClient) DescribeWorkspaces(arg0 context.Context, arg1 *workspaces.DescribeWorkspacesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkspaces", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkspaces indicates an expected call of DescribeWorkspaces.
func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkspaces", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaces), varargs...)
}

// DescribeWorkspacesConnectionStatus mocks base method.
func (m *MockWorkspacesClient) DescribeWorkspacesConnectionStatus(arg0 context.Context, arg1 *workspaces.DescribeWorkspacesConnectionStatusInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkspacesConnectionStatus", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspacesConnectionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkspacesConnectionStatus indicates an expected call of DescribeWorkspacesConnectionStatus.
func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspacesConnectionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkspacesConnectionStatus", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspacesConnectionStatus), varargs...)
}

// ListAvailableManagementCidrRanges mocks base method.
func (m *MockWorkspacesClient) ListAvailableManagementCidrRanges(arg0 context.Context, arg1 *workspaces.ListAvailableManagementCidrRangesInput, arg2 ...func(*workspaces.Options)) (*workspaces.ListAvailableManagementCidrRangesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAvailableManagementCidrRanges", varargs...)
	ret0, _ := ret[0].(*workspaces.ListAvailableManagementCidrRangesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailableManagementCidrRanges indicates an expected call of ListAvailableManagementCidrRanges.
func (mr *MockWorkspacesClientMockRecorder) ListAvailableManagementCidrRanges(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableManagementCidrRanges", reflect.TypeOf((*MockWorkspacesClient)(nil).ListAvailableManagementCidrRanges), varargs...)
}
