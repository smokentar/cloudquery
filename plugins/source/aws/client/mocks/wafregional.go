// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: WafregionalClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	wafregional "github.com/aws/aws-sdk-go-v2/service/wafregional"
	gomock "github.com/golang/mock/gomock"
)

// MockWafregionalClient is a mock of WafregionalClient interface.
type MockWafregionalClient struct {
	ctrl     *gomock.Controller
	recorder *MockWafregionalClientMockRecorder
}

// MockWafregionalClientMockRecorder is the mock recorder for MockWafregionalClient.
type MockWafregionalClientMockRecorder struct {
	mock *MockWafregionalClient
}

// NewMockWafregionalClient creates a new mock instance.
func NewMockWafregionalClient(ctrl *gomock.Controller) *MockWafregionalClient {
	mock := &MockWafregionalClient{ctrl: ctrl}
	mock.recorder = &MockWafregionalClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWafregionalClient) EXPECT() *MockWafregionalClientMockRecorder {
	return m.recorder
}

// GetByteMatchSet mocks base method.
func (m *MockWafregionalClient) GetByteMatchSet(arg0 context.Context, arg1 *wafregional.GetByteMatchSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetByteMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByteMatchSet", varargs...)
	ret0, _ := ret[0].(*wafregional.GetByteMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByteMatchSet indicates an expected call of GetByteMatchSet.
func (mr *MockWafregionalClientMockRecorder) GetByteMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByteMatchSet", reflect.TypeOf((*MockWafregionalClient)(nil).GetByteMatchSet), varargs...)
}

// GetChangeToken mocks base method.
func (m *MockWafregionalClient) GetChangeToken(arg0 context.Context, arg1 *wafregional.GetChangeTokenInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetChangeTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChangeToken", varargs...)
	ret0, _ := ret[0].(*wafregional.GetChangeTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChangeToken indicates an expected call of GetChangeToken.
func (mr *MockWafregionalClientMockRecorder) GetChangeToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChangeToken", reflect.TypeOf((*MockWafregionalClient)(nil).GetChangeToken), varargs...)
}

// GetChangeTokenStatus mocks base method.
func (m *MockWafregionalClient) GetChangeTokenStatus(arg0 context.Context, arg1 *wafregional.GetChangeTokenStatusInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetChangeTokenStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChangeTokenStatus", varargs...)
	ret0, _ := ret[0].(*wafregional.GetChangeTokenStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChangeTokenStatus indicates an expected call of GetChangeTokenStatus.
func (mr *MockWafregionalClientMockRecorder) GetChangeTokenStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChangeTokenStatus", reflect.TypeOf((*MockWafregionalClient)(nil).GetChangeTokenStatus), varargs...)
}

// GetGeoMatchSet mocks base method.
func (m *MockWafregionalClient) GetGeoMatchSet(arg0 context.Context, arg1 *wafregional.GetGeoMatchSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetGeoMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGeoMatchSet", varargs...)
	ret0, _ := ret[0].(*wafregional.GetGeoMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGeoMatchSet indicates an expected call of GetGeoMatchSet.
func (mr *MockWafregionalClientMockRecorder) GetGeoMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeoMatchSet", reflect.TypeOf((*MockWafregionalClient)(nil).GetGeoMatchSet), varargs...)
}

// GetIPSet mocks base method.
func (m *MockWafregionalClient) GetIPSet(arg0 context.Context, arg1 *wafregional.GetIPSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetIPSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIPSet", varargs...)
	ret0, _ := ret[0].(*wafregional.GetIPSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPSet indicates an expected call of GetIPSet.
func (mr *MockWafregionalClientMockRecorder) GetIPSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPSet", reflect.TypeOf((*MockWafregionalClient)(nil).GetIPSet), varargs...)
}

// GetLoggingConfiguration mocks base method.
func (m *MockWafregionalClient) GetLoggingConfiguration(arg0 context.Context, arg1 *wafregional.GetLoggingConfigurationInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetLoggingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLoggingConfiguration", varargs...)
	ret0, _ := ret[0].(*wafregional.GetLoggingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoggingConfiguration indicates an expected call of GetLoggingConfiguration.
func (mr *MockWafregionalClientMockRecorder) GetLoggingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoggingConfiguration", reflect.TypeOf((*MockWafregionalClient)(nil).GetLoggingConfiguration), varargs...)
}

// GetPermissionPolicy mocks base method.
func (m *MockWafregionalClient) GetPermissionPolicy(arg0 context.Context, arg1 *wafregional.GetPermissionPolicyInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetPermissionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPermissionPolicy", varargs...)
	ret0, _ := ret[0].(*wafregional.GetPermissionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermissionPolicy indicates an expected call of GetPermissionPolicy.
func (mr *MockWafregionalClientMockRecorder) GetPermissionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissionPolicy", reflect.TypeOf((*MockWafregionalClient)(nil).GetPermissionPolicy), varargs...)
}

// GetRateBasedRule mocks base method.
func (m *MockWafregionalClient) GetRateBasedRule(arg0 context.Context, arg1 *wafregional.GetRateBasedRuleInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRateBasedRule", varargs...)
	ret0, _ := ret[0].(*wafregional.GetRateBasedRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRateBasedRule indicates an expected call of GetRateBasedRule.
func (mr *MockWafregionalClientMockRecorder) GetRateBasedRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateBasedRule", reflect.TypeOf((*MockWafregionalClient)(nil).GetRateBasedRule), varargs...)
}

// GetRateBasedRuleManagedKeys mocks base method.
func (m *MockWafregionalClient) GetRateBasedRuleManagedKeys(arg0 context.Context, arg1 *wafregional.GetRateBasedRuleManagedKeysInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleManagedKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRateBasedRuleManagedKeys", varargs...)
	ret0, _ := ret[0].(*wafregional.GetRateBasedRuleManagedKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRateBasedRuleManagedKeys indicates an expected call of GetRateBasedRuleManagedKeys.
func (mr *MockWafregionalClientMockRecorder) GetRateBasedRuleManagedKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateBasedRuleManagedKeys", reflect.TypeOf((*MockWafregionalClient)(nil).GetRateBasedRuleManagedKeys), varargs...)
}

// GetRegexMatchSet mocks base method.
func (m *MockWafregionalClient) GetRegexMatchSet(arg0 context.Context, arg1 *wafregional.GetRegexMatchSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRegexMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegexMatchSet", varargs...)
	ret0, _ := ret[0].(*wafregional.GetRegexMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegexMatchSet indicates an expected call of GetRegexMatchSet.
func (mr *MockWafregionalClientMockRecorder) GetRegexMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegexMatchSet", reflect.TypeOf((*MockWafregionalClient)(nil).GetRegexMatchSet), varargs...)
}

// GetRegexPatternSet mocks base method.
func (m *MockWafregionalClient) GetRegexPatternSet(arg0 context.Context, arg1 *wafregional.GetRegexPatternSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRegexPatternSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegexPatternSet", varargs...)
	ret0, _ := ret[0].(*wafregional.GetRegexPatternSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegexPatternSet indicates an expected call of GetRegexPatternSet.
func (mr *MockWafregionalClientMockRecorder) GetRegexPatternSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegexPatternSet", reflect.TypeOf((*MockWafregionalClient)(nil).GetRegexPatternSet), varargs...)
}

// GetRule mocks base method.
func (m *MockWafregionalClient) GetRule(arg0 context.Context, arg1 *wafregional.GetRuleInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRule", varargs...)
	ret0, _ := ret[0].(*wafregional.GetRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRule indicates an expected call of GetRule.
func (mr *MockWafregionalClientMockRecorder) GetRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRule", reflect.TypeOf((*MockWafregionalClient)(nil).GetRule), varargs...)
}

// GetRuleGroup mocks base method.
func (m *MockWafregionalClient) GetRuleGroup(arg0 context.Context, arg1 *wafregional.GetRuleGroupInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRuleGroup", varargs...)
	ret0, _ := ret[0].(*wafregional.GetRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRuleGroup indicates an expected call of GetRuleGroup.
func (mr *MockWafregionalClientMockRecorder) GetRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuleGroup", reflect.TypeOf((*MockWafregionalClient)(nil).GetRuleGroup), varargs...)
}

// GetSampledRequests mocks base method.
func (m *MockWafregionalClient) GetSampledRequests(arg0 context.Context, arg1 *wafregional.GetSampledRequestsInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetSampledRequestsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSampledRequests", varargs...)
	ret0, _ := ret[0].(*wafregional.GetSampledRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSampledRequests indicates an expected call of GetSampledRequests.
func (mr *MockWafregionalClientMockRecorder) GetSampledRequests(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSampledRequests", reflect.TypeOf((*MockWafregionalClient)(nil).GetSampledRequests), varargs...)
}

// GetSizeConstraintSet mocks base method.
func (m *MockWafregionalClient) GetSizeConstraintSet(arg0 context.Context, arg1 *wafregional.GetSizeConstraintSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetSizeConstraintSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSizeConstraintSet", varargs...)
	ret0, _ := ret[0].(*wafregional.GetSizeConstraintSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSizeConstraintSet indicates an expected call of GetSizeConstraintSet.
func (mr *MockWafregionalClientMockRecorder) GetSizeConstraintSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSizeConstraintSet", reflect.TypeOf((*MockWafregionalClient)(nil).GetSizeConstraintSet), varargs...)
}

// GetSqlInjectionMatchSet mocks base method.
func (m *MockWafregionalClient) GetSqlInjectionMatchSet(arg0 context.Context, arg1 *wafregional.GetSqlInjectionMatchSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetSqlInjectionMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSqlInjectionMatchSet", varargs...)
	ret0, _ := ret[0].(*wafregional.GetSqlInjectionMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSqlInjectionMatchSet indicates an expected call of GetSqlInjectionMatchSet.
func (mr *MockWafregionalClientMockRecorder) GetSqlInjectionMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSqlInjectionMatchSet", reflect.TypeOf((*MockWafregionalClient)(nil).GetSqlInjectionMatchSet), varargs...)
}

// GetWebACL mocks base method.
func (m *MockWafregionalClient) GetWebACL(arg0 context.Context, arg1 *wafregional.GetWebACLInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWebACL", varargs...)
	ret0, _ := ret[0].(*wafregional.GetWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebACL indicates an expected call of GetWebACL.
func (mr *MockWafregionalClientMockRecorder) GetWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebACL", reflect.TypeOf((*MockWafregionalClient)(nil).GetWebACL), varargs...)
}

// GetWebACLForResource mocks base method.
func (m *MockWafregionalClient) GetWebACLForResource(arg0 context.Context, arg1 *wafregional.GetWebACLForResourceInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetWebACLForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWebACLForResource", varargs...)
	ret0, _ := ret[0].(*wafregional.GetWebACLForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebACLForResource indicates an expected call of GetWebACLForResource.
func (mr *MockWafregionalClientMockRecorder) GetWebACLForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebACLForResource", reflect.TypeOf((*MockWafregionalClient)(nil).GetWebACLForResource), varargs...)
}

// GetXssMatchSet mocks base method.
func (m *MockWafregionalClient) GetXssMatchSet(arg0 context.Context, arg1 *wafregional.GetXssMatchSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetXssMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetXssMatchSet", varargs...)
	ret0, _ := ret[0].(*wafregional.GetXssMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetXssMatchSet indicates an expected call of GetXssMatchSet.
func (mr *MockWafregionalClientMockRecorder) GetXssMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetXssMatchSet", reflect.TypeOf((*MockWafregionalClient)(nil).GetXssMatchSet), varargs...)
}

// ListActivatedRulesInRuleGroup mocks base method.
func (m *MockWafregionalClient) ListActivatedRulesInRuleGroup(arg0 context.Context, arg1 *wafregional.ListActivatedRulesInRuleGroupInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListActivatedRulesInRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListActivatedRulesInRuleGroup", varargs...)
	ret0, _ := ret[0].(*wafregional.ListActivatedRulesInRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActivatedRulesInRuleGroup indicates an expected call of ListActivatedRulesInRuleGroup.
func (mr *MockWafregionalClientMockRecorder) ListActivatedRulesInRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActivatedRulesInRuleGroup", reflect.TypeOf((*MockWafregionalClient)(nil).ListActivatedRulesInRuleGroup), varargs...)
}

// ListByteMatchSets mocks base method.
func (m *MockWafregionalClient) ListByteMatchSets(arg0 context.Context, arg1 *wafregional.ListByteMatchSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListByteMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByteMatchSets", varargs...)
	ret0, _ := ret[0].(*wafregional.ListByteMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByteMatchSets indicates an expected call of ListByteMatchSets.
func (mr *MockWafregionalClientMockRecorder) ListByteMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByteMatchSets", reflect.TypeOf((*MockWafregionalClient)(nil).ListByteMatchSets), varargs...)
}

// ListGeoMatchSets mocks base method.
func (m *MockWafregionalClient) ListGeoMatchSets(arg0 context.Context, arg1 *wafregional.ListGeoMatchSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListGeoMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGeoMatchSets", varargs...)
	ret0, _ := ret[0].(*wafregional.ListGeoMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGeoMatchSets indicates an expected call of ListGeoMatchSets.
func (mr *MockWafregionalClientMockRecorder) ListGeoMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGeoMatchSets", reflect.TypeOf((*MockWafregionalClient)(nil).ListGeoMatchSets), varargs...)
}

// ListIPSets mocks base method.
func (m *MockWafregionalClient) ListIPSets(arg0 context.Context, arg1 *wafregional.ListIPSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListIPSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIPSets", varargs...)
	ret0, _ := ret[0].(*wafregional.ListIPSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIPSets indicates an expected call of ListIPSets.
func (mr *MockWafregionalClientMockRecorder) ListIPSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIPSets", reflect.TypeOf((*MockWafregionalClient)(nil).ListIPSets), varargs...)
}

// ListLoggingConfigurations mocks base method.
func (m *MockWafregionalClient) ListLoggingConfigurations(arg0 context.Context, arg1 *wafregional.ListLoggingConfigurationsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListLoggingConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLoggingConfigurations", varargs...)
	ret0, _ := ret[0].(*wafregional.ListLoggingConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLoggingConfigurations indicates an expected call of ListLoggingConfigurations.
func (mr *MockWafregionalClientMockRecorder) ListLoggingConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLoggingConfigurations", reflect.TypeOf((*MockWafregionalClient)(nil).ListLoggingConfigurations), varargs...)
}

// ListRateBasedRules mocks base method.
func (m *MockWafregionalClient) ListRateBasedRules(arg0 context.Context, arg1 *wafregional.ListRateBasedRulesInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRateBasedRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRateBasedRules", varargs...)
	ret0, _ := ret[0].(*wafregional.ListRateBasedRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRateBasedRules indicates an expected call of ListRateBasedRules.
func (mr *MockWafregionalClientMockRecorder) ListRateBasedRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRateBasedRules", reflect.TypeOf((*MockWafregionalClient)(nil).ListRateBasedRules), varargs...)
}

// ListRegexMatchSets mocks base method.
func (m *MockWafregionalClient) ListRegexMatchSets(arg0 context.Context, arg1 *wafregional.ListRegexMatchSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRegexMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRegexMatchSets", varargs...)
	ret0, _ := ret[0].(*wafregional.ListRegexMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegexMatchSets indicates an expected call of ListRegexMatchSets.
func (mr *MockWafregionalClientMockRecorder) ListRegexMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegexMatchSets", reflect.TypeOf((*MockWafregionalClient)(nil).ListRegexMatchSets), varargs...)
}

// ListRegexPatternSets mocks base method.
func (m *MockWafregionalClient) ListRegexPatternSets(arg0 context.Context, arg1 *wafregional.ListRegexPatternSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRegexPatternSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRegexPatternSets", varargs...)
	ret0, _ := ret[0].(*wafregional.ListRegexPatternSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegexPatternSets indicates an expected call of ListRegexPatternSets.
func (mr *MockWafregionalClientMockRecorder) ListRegexPatternSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegexPatternSets", reflect.TypeOf((*MockWafregionalClient)(nil).ListRegexPatternSets), varargs...)
}

// ListResourcesForWebACL mocks base method.
func (m *MockWafregionalClient) ListResourcesForWebACL(arg0 context.Context, arg1 *wafregional.ListResourcesForWebACLInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListResourcesForWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourcesForWebACL", varargs...)
	ret0, _ := ret[0].(*wafregional.ListResourcesForWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourcesForWebACL indicates an expected call of ListResourcesForWebACL.
func (mr *MockWafregionalClientMockRecorder) ListResourcesForWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourcesForWebACL", reflect.TypeOf((*MockWafregionalClient)(nil).ListResourcesForWebACL), varargs...)
}

// ListRuleGroups mocks base method.
func (m *MockWafregionalClient) ListRuleGroups(arg0 context.Context, arg1 *wafregional.ListRuleGroupsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRuleGroups", varargs...)
	ret0, _ := ret[0].(*wafregional.ListRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRuleGroups indicates an expected call of ListRuleGroups.
func (mr *MockWafregionalClientMockRecorder) ListRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRuleGroups", reflect.TypeOf((*MockWafregionalClient)(nil).ListRuleGroups), varargs...)
}

// ListRules mocks base method.
func (m *MockWafregionalClient) ListRules(arg0 context.Context, arg1 *wafregional.ListRulesInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRules", varargs...)
	ret0, _ := ret[0].(*wafregional.ListRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRules indicates an expected call of ListRules.
func (mr *MockWafregionalClientMockRecorder) ListRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRules", reflect.TypeOf((*MockWafregionalClient)(nil).ListRules), varargs...)
}

// ListSizeConstraintSets mocks base method.
func (m *MockWafregionalClient) ListSizeConstraintSets(arg0 context.Context, arg1 *wafregional.ListSizeConstraintSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListSizeConstraintSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSizeConstraintSets", varargs...)
	ret0, _ := ret[0].(*wafregional.ListSizeConstraintSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSizeConstraintSets indicates an expected call of ListSizeConstraintSets.
func (mr *MockWafregionalClientMockRecorder) ListSizeConstraintSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSizeConstraintSets", reflect.TypeOf((*MockWafregionalClient)(nil).ListSizeConstraintSets), varargs...)
}

// ListSqlInjectionMatchSets mocks base method.
func (m *MockWafregionalClient) ListSqlInjectionMatchSets(arg0 context.Context, arg1 *wafregional.ListSqlInjectionMatchSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListSqlInjectionMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSqlInjectionMatchSets", varargs...)
	ret0, _ := ret[0].(*wafregional.ListSqlInjectionMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSqlInjectionMatchSets indicates an expected call of ListSqlInjectionMatchSets.
func (mr *MockWafregionalClientMockRecorder) ListSqlInjectionMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSqlInjectionMatchSets", reflect.TypeOf((*MockWafregionalClient)(nil).ListSqlInjectionMatchSets), varargs...)
}

// ListSubscribedRuleGroups mocks base method.
func (m *MockWafregionalClient) ListSubscribedRuleGroups(arg0 context.Context, arg1 *wafregional.ListSubscribedRuleGroupsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListSubscribedRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSubscribedRuleGroups", varargs...)
	ret0, _ := ret[0].(*wafregional.ListSubscribedRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubscribedRuleGroups indicates an expected call of ListSubscribedRuleGroups.
func (mr *MockWafregionalClientMockRecorder) ListSubscribedRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubscribedRuleGroups", reflect.TypeOf((*MockWafregionalClient)(nil).ListSubscribedRuleGroups), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockWafregionalClient) ListTagsForResource(arg0 context.Context, arg1 *wafregional.ListTagsForResourceInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*wafregional.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockWafregionalClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockWafregionalClient)(nil).ListTagsForResource), varargs...)
}

// ListWebACLs mocks base method.
func (m *MockWafregionalClient) ListWebACLs(arg0 context.Context, arg1 *wafregional.ListWebACLsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListWebACLsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWebACLs", varargs...)
	ret0, _ := ret[0].(*wafregional.ListWebACLsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWebACLs indicates an expected call of ListWebACLs.
func (mr *MockWafregionalClientMockRecorder) ListWebACLs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebACLs", reflect.TypeOf((*MockWafregionalClient)(nil).ListWebACLs), varargs...)
}

// ListXssMatchSets mocks base method.
func (m *MockWafregionalClient) ListXssMatchSets(arg0 context.Context, arg1 *wafregional.ListXssMatchSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListXssMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListXssMatchSets", varargs...)
	ret0, _ := ret[0].(*wafregional.ListXssMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListXssMatchSets indicates an expected call of ListXssMatchSets.
func (mr *MockWafregionalClientMockRecorder) ListXssMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListXssMatchSets", reflect.TypeOf((*MockWafregionalClient)(nil).ListXssMatchSets), varargs...)
}
