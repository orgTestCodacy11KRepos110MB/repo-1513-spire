// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/pkg/server/plugin/nodeattestor/aws (interfaces: Client)

// Package mock_aws is a generated GoMock package.
package mock_aws

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	iam "github.com/aws/aws-sdk-go/service/iam"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// DescribeInstancesWithContext mocks base method
func (m *MockClient) DescribeInstancesWithContext(arg0 context.Context, arg1 *ec2.DescribeInstancesInput, arg2 ...request.Option) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstancesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstancesWithContext indicates an expected call of DescribeInstancesWithContext
func (mr *MockClientMockRecorder) DescribeInstancesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstancesWithContext", reflect.TypeOf((*MockClient)(nil).DescribeInstancesWithContext), varargs...)
}

// GetInstanceProfileWithContext mocks base method
func (m *MockClient) GetInstanceProfileWithContext(arg0 context.Context, arg1 *iam.GetInstanceProfileInput, arg2 ...request.Option) (*iam.GetInstanceProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInstanceProfileWithContext", varargs...)
	ret0, _ := ret[0].(*iam.GetInstanceProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceProfileWithContext indicates an expected call of GetInstanceProfileWithContext
func (mr *MockClientMockRecorder) GetInstanceProfileWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceProfileWithContext", reflect.TypeOf((*MockClient)(nil).GetInstanceProfileWithContext), varargs...)
}