// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package gomailer is a generated GoMock package.
package gomailer

import (
	context "context"
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

// Send mocks base method
func (m *MockClient) Send(msg *Message) error {
	ret := m.ctrl.Call(m, "Send", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockClientMockRecorder) Send(msg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockClient)(nil).Send), msg)
}

// SendContext mocks base method
func (m *MockClient) SendContext(ctx context.Context, msg *Message) error {
	ret := m.ctrl.Call(m, "SendContext", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendContext indicates an expected call of SendContext
func (mr *MockClientMockRecorder) SendContext(ctx, msg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendContext", reflect.TypeOf((*MockClient)(nil).SendContext), ctx, msg)
}

// Close mocks base method
func (m *MockClient) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}
