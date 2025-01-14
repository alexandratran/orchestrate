// Code generated by MockGen. DO NOT EDIT.
// Source: use-cases.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	usecases "github.com/consensys/orchestrate/src/tx-sender/tx-sender/use-cases"
	reflect "reflect"
)

// MockUseCases is a mock of UseCases interface
type MockUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockUseCasesMockRecorder
}

// MockUseCasesMockRecorder is the mock recorder for MockUseCases
type MockUseCasesMockRecorder struct {
	mock *MockUseCases
}

// NewMockUseCases creates a new mock instance
func NewMockUseCases(ctrl *gomock.Controller) *MockUseCases {
	mock := &MockUseCases{ctrl: ctrl}
	mock.recorder = &MockUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCases) EXPECT() *MockUseCasesMockRecorder {
	return m.recorder
}

// SendETHRawTx mocks base method
func (m *MockUseCases) SendETHRawTx() usecases.SendETHRawTxUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendETHRawTx")
	ret0, _ := ret[0].(usecases.SendETHRawTxUseCase)
	return ret0
}

// SendETHRawTx indicates an expected call of SendETHRawTx
func (mr *MockUseCasesMockRecorder) SendETHRawTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendETHRawTx", reflect.TypeOf((*MockUseCases)(nil).SendETHRawTx))
}

// SendETHTx mocks base method
func (m *MockUseCases) SendETHTx() usecases.SendETHTxUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendETHTx")
	ret0, _ := ret[0].(usecases.SendETHTxUseCase)
	return ret0
}

// SendETHTx indicates an expected call of SendETHTx
func (mr *MockUseCasesMockRecorder) SendETHTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendETHTx", reflect.TypeOf((*MockUseCases)(nil).SendETHTx))
}

// SendEEAPrivateTx mocks base method
func (m *MockUseCases) SendEEAPrivateTx() usecases.SendEEAPrivateTxUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEEAPrivateTx")
	ret0, _ := ret[0].(usecases.SendEEAPrivateTxUseCase)
	return ret0
}

// SendEEAPrivateTx indicates an expected call of SendEEAPrivateTx
func (mr *MockUseCasesMockRecorder) SendEEAPrivateTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEEAPrivateTx", reflect.TypeOf((*MockUseCases)(nil).SendEEAPrivateTx))
}

// SendTesseraPrivateTx mocks base method
func (m *MockUseCases) SendTesseraPrivateTx() usecases.SendTesseraPrivateTxUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTesseraPrivateTx")
	ret0, _ := ret[0].(usecases.SendTesseraPrivateTxUseCase)
	return ret0
}

// SendTesseraPrivateTx indicates an expected call of SendTesseraPrivateTx
func (mr *MockUseCasesMockRecorder) SendTesseraPrivateTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTesseraPrivateTx", reflect.TypeOf((*MockUseCases)(nil).SendTesseraPrivateTx))
}

// SendTesseraMarkingTx mocks base method
func (m *MockUseCases) SendTesseraMarkingTx() usecases.SendTesseraMarkingTxUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTesseraMarkingTx")
	ret0, _ := ret[0].(usecases.SendTesseraMarkingTxUseCase)
	return ret0
}

// SendTesseraMarkingTx indicates an expected call of SendTesseraMarkingTx
func (mr *MockUseCasesMockRecorder) SendTesseraMarkingTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTesseraMarkingTx", reflect.TypeOf((*MockUseCases)(nil).SendTesseraMarkingTx))
}
