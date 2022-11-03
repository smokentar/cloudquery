// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: KinesisClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	kinesis "github.com/aws/aws-sdk-go-v2/service/kinesis"
	gomock "github.com/golang/mock/gomock"
)

// MockKinesisClient is a mock of KinesisClient interface.
type MockKinesisClient struct {
	ctrl     *gomock.Controller
	recorder *MockKinesisClientMockRecorder
}

// MockKinesisClientMockRecorder is the mock recorder for MockKinesisClient.
type MockKinesisClientMockRecorder struct {
	mock *MockKinesisClient
}

// NewMockKinesisClient creates a new mock instance.
func NewMockKinesisClient(ctrl *gomock.Controller) *MockKinesisClient {
	mock := &MockKinesisClient{ctrl: ctrl}
	mock.recorder = &MockKinesisClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKinesisClient) EXPECT() *MockKinesisClientMockRecorder {
	return m.recorder
}

// DescribeLimits mocks base method.
func (m *MockKinesisClient) DescribeLimits(arg0 context.Context, arg1 *kinesis.DescribeLimitsInput, arg2 ...func(*kinesis.Options)) (*kinesis.DescribeLimitsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLimits", varargs...)
	ret0, _ := ret[0].(*kinesis.DescribeLimitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLimits indicates an expected call of DescribeLimits.
func (mr *MockKinesisClientMockRecorder) DescribeLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLimits", reflect.TypeOf((*MockKinesisClient)(nil).DescribeLimits), varargs...)
}

// DescribeStream mocks base method.
func (m *MockKinesisClient) DescribeStream(arg0 context.Context, arg1 *kinesis.DescribeStreamInput, arg2 ...func(*kinesis.Options)) (*kinesis.DescribeStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStream", varargs...)
	ret0, _ := ret[0].(*kinesis.DescribeStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStream indicates an expected call of DescribeStream.
func (mr *MockKinesisClientMockRecorder) DescribeStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStream", reflect.TypeOf((*MockKinesisClient)(nil).DescribeStream), varargs...)
}

// DescribeStreamConsumer mocks base method.
func (m *MockKinesisClient) DescribeStreamConsumer(arg0 context.Context, arg1 *kinesis.DescribeStreamConsumerInput, arg2 ...func(*kinesis.Options)) (*kinesis.DescribeStreamConsumerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStreamConsumer", varargs...)
	ret0, _ := ret[0].(*kinesis.DescribeStreamConsumerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStreamConsumer indicates an expected call of DescribeStreamConsumer.
func (mr *MockKinesisClientMockRecorder) DescribeStreamConsumer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStreamConsumer", reflect.TypeOf((*MockKinesisClient)(nil).DescribeStreamConsumer), varargs...)
}

// DescribeStreamSummary mocks base method.
func (m *MockKinesisClient) DescribeStreamSummary(arg0 context.Context, arg1 *kinesis.DescribeStreamSummaryInput, arg2 ...func(*kinesis.Options)) (*kinesis.DescribeStreamSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStreamSummary", varargs...)
	ret0, _ := ret[0].(*kinesis.DescribeStreamSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStreamSummary indicates an expected call of DescribeStreamSummary.
func (mr *MockKinesisClientMockRecorder) DescribeStreamSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStreamSummary", reflect.TypeOf((*MockKinesisClient)(nil).DescribeStreamSummary), varargs...)
}

// GetRecords mocks base method.
func (m *MockKinesisClient) GetRecords(arg0 context.Context, arg1 *kinesis.GetRecordsInput, arg2 ...func(*kinesis.Options)) (*kinesis.GetRecordsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecords", varargs...)
	ret0, _ := ret[0].(*kinesis.GetRecordsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecords indicates an expected call of GetRecords.
func (mr *MockKinesisClientMockRecorder) GetRecords(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecords", reflect.TypeOf((*MockKinesisClient)(nil).GetRecords), varargs...)
}

// GetShardIterator mocks base method.
func (m *MockKinesisClient) GetShardIterator(arg0 context.Context, arg1 *kinesis.GetShardIteratorInput, arg2 ...func(*kinesis.Options)) (*kinesis.GetShardIteratorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetShardIterator", varargs...)
	ret0, _ := ret[0].(*kinesis.GetShardIteratorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShardIterator indicates an expected call of GetShardIterator.
func (mr *MockKinesisClientMockRecorder) GetShardIterator(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardIterator", reflect.TypeOf((*MockKinesisClient)(nil).GetShardIterator), varargs...)
}

// ListShards mocks base method.
func (m *MockKinesisClient) ListShards(arg0 context.Context, arg1 *kinesis.ListShardsInput, arg2 ...func(*kinesis.Options)) (*kinesis.ListShardsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListShards", varargs...)
	ret0, _ := ret[0].(*kinesis.ListShardsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListShards indicates an expected call of ListShards.
func (mr *MockKinesisClientMockRecorder) ListShards(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShards", reflect.TypeOf((*MockKinesisClient)(nil).ListShards), varargs...)
}

// ListStreamConsumers mocks base method.
func (m *MockKinesisClient) ListStreamConsumers(arg0 context.Context, arg1 *kinesis.ListStreamConsumersInput, arg2 ...func(*kinesis.Options)) (*kinesis.ListStreamConsumersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStreamConsumers", varargs...)
	ret0, _ := ret[0].(*kinesis.ListStreamConsumersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStreamConsumers indicates an expected call of ListStreamConsumers.
func (mr *MockKinesisClientMockRecorder) ListStreamConsumers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStreamConsumers", reflect.TypeOf((*MockKinesisClient)(nil).ListStreamConsumers), varargs...)
}

// ListStreams mocks base method.
func (m *MockKinesisClient) ListStreams(arg0 context.Context, arg1 *kinesis.ListStreamsInput, arg2 ...func(*kinesis.Options)) (*kinesis.ListStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStreams", varargs...)
	ret0, _ := ret[0].(*kinesis.ListStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStreams indicates an expected call of ListStreams.
func (mr *MockKinesisClientMockRecorder) ListStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStreams", reflect.TypeOf((*MockKinesisClient)(nil).ListStreams), varargs...)
}

// ListTagsForStream mocks base method.
func (m *MockKinesisClient) ListTagsForStream(arg0 context.Context, arg1 *kinesis.ListTagsForStreamInput, arg2 ...func(*kinesis.Options)) (*kinesis.ListTagsForStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForStream", varargs...)
	ret0, _ := ret[0].(*kinesis.ListTagsForStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForStream indicates an expected call of ListTagsForStream.
func (mr *MockKinesisClientMockRecorder) ListTagsForStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForStream", reflect.TypeOf((*MockKinesisClient)(nil).ListTagsForStream), varargs...)
}
