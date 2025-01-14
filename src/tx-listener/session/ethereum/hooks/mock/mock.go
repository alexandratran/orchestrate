// Code generated by MockGen. DO NOT EDIT.
// Source: hook.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
	entities "github.com/consensys/orchestrate/src/entities"
	dynamic "github.com/consensys/orchestrate/src/tx-listener/dynamic"
	reflect "reflect"
)

// MockHook is a mock of Hook interface
type MockHook struct {
	ctrl     *gomock.Controller
	recorder *MockHookMockRecorder
}

// MockHookMockRecorder is the mock recorder for MockHook
type MockHookMockRecorder struct {
	mock *MockHook
}

// NewMockHook creates a new mock instance
func NewMockHook(ctrl *gomock.Controller) *MockHook {
	mock := &MockHook{ctrl: ctrl}
	mock.recorder = &MockHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHook) EXPECT() *MockHookMockRecorder {
	return m.recorder
}

// AfterNewBlock mocks base method
func (m *MockHook) AfterNewBlock(ctx context.Context, chain *dynamic.Chain, block *types.Block, jobs []*entities.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterNewBlock", ctx, chain, block, jobs)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterNewBlock indicates an expected call of AfterNewBlock
func (mr *MockHookMockRecorder) AfterNewBlock(ctx, chain, block, jobs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterNewBlock", reflect.TypeOf((*MockHook)(nil).AfterNewBlock), ctx, chain, block, jobs)
}
