// Code generated by MockGen. DO NOT EDIT.
// Source: common/remote/rpc/rpc_client.go

// Package remote is a generated GoMock package.
package rpc

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRpcClient is a mock of IRpcClient interface.
type MockIRpcClient struct {
	ctrl     *gomock.Controller
	recorder *MockIRpcClientMockRecorder
}

// MockIRpcClientMockRecorder is the mock recorder for MockIRpcClient.
type MockIRpcClientMockRecorder struct {
	mock *MockIRpcClient
}

// NewMockIRpcClient creates a new mock instance.
func NewMockIRpcClient(ctrl *gomock.Controller) *MockIRpcClient {
	mock := &MockIRpcClient{ctrl: ctrl}
	mock.recorder = &MockIRpcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRpcClient) EXPECT() *MockIRpcClientMockRecorder {
	return m.recorder
}

// GetRpcClient mocks base method.
func (m *MockIRpcClient) GetRpcClient() *RpcClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRpcClient")
	ret0, _ := ret[0].(*RpcClient)
	return ret0
}

// GetRpcClient indicates an expected call of GetRpcClient.
func (mr *MockIRpcClientMockRecorder) GetRpcClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRpcClient", reflect.TypeOf((*MockIRpcClient)(nil).GetRpcClient))
}

// connectToServer mocks base method.
func (m *MockIRpcClient) connectToServer(serverInfo ServerInfo) (IConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "connectToServer", serverInfo)
	ret0, _ := ret[0].(IConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// connectToServer indicates an expected call of connectToServer.
func (mr *MockIRpcClientMockRecorder) connectToServer(serverInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "connectToServer", reflect.TypeOf((*MockIRpcClient)(nil).connectToServer), serverInfo)
}

// getConnectionType mocks base method.
func (m *MockIRpcClient) getConnectionType() ConnectionType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getConnectionType")
	ret0, _ := ret[0].(ConnectionType)
	return ret0
}

// getConnectionType indicates an expected call of getConnectionType.
func (mr *MockIRpcClientMockRecorder) getConnectionType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getConnectionType", reflect.TypeOf((*MockIRpcClient)(nil).getConnectionType))
}

// putAllLabels mocks base method.
func (m *MockIRpcClient) putAllLabels(labels map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "putAllLabels", labels)
}

// putAllLabels indicates an expected call of putAllLabels.
func (mr *MockIRpcClientMockRecorder) putAllLabels(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "putAllLabels", reflect.TypeOf((*MockIRpcClient)(nil).putAllLabels), labels)
}

// rpcPortOffset mocks base method.
func (m *MockIRpcClient) rpcPortOffset() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "rpcPortOffset")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// rpcPortOffset indicates an expected call of rpcPortOffset.
func (mr *MockIRpcClientMockRecorder) rpcPortOffset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "rpcPortOffset", reflect.TypeOf((*MockIRpcClient)(nil).rpcPortOffset))
}
