// Code generated by MockGen. DO NOT EDIT.
// Source: chains.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	multitenancy "github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	entities "github.com/consensys/orchestrate/src/entities"
	usecases "github.com/consensys/orchestrate/src/api/business/use-cases"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockChainUseCases is a mock of ChainUseCases interface
type MockChainUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockChainUseCasesMockRecorder
}

// MockChainUseCasesMockRecorder is the mock recorder for MockChainUseCases
type MockChainUseCasesMockRecorder struct {
	mock *MockChainUseCases
}

// NewMockChainUseCases creates a new mock instance
func NewMockChainUseCases(ctrl *gomock.Controller) *MockChainUseCases {
	mock := &MockChainUseCases{ctrl: ctrl}
	mock.recorder = &MockChainUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainUseCases) EXPECT() *MockChainUseCasesMockRecorder {
	return m.recorder
}

// RegisterChain mocks base method
func (m *MockChainUseCases) RegisterChain() usecases.RegisterChainUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterChain")
	ret0, _ := ret[0].(usecases.RegisterChainUseCase)
	return ret0
}

// RegisterChain indicates an expected call of RegisterChain
func (mr *MockChainUseCasesMockRecorder) RegisterChain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterChain", reflect.TypeOf((*MockChainUseCases)(nil).RegisterChain))
}

// GetChain mocks base method
func (m *MockChainUseCases) GetChain() usecases.GetChainUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChain")
	ret0, _ := ret[0].(usecases.GetChainUseCase)
	return ret0
}

// GetChain indicates an expected call of GetChain
func (mr *MockChainUseCasesMockRecorder) GetChain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChain", reflect.TypeOf((*MockChainUseCases)(nil).GetChain))
}

// SearchChains mocks base method
func (m *MockChainUseCases) SearchChains() usecases.SearchChainsUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchChains")
	ret0, _ := ret[0].(usecases.SearchChainsUseCase)
	return ret0
}

// SearchChains indicates an expected call of SearchChains
func (mr *MockChainUseCasesMockRecorder) SearchChains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchChains", reflect.TypeOf((*MockChainUseCases)(nil).SearchChains))
}

// UpdateChain mocks base method
func (m *MockChainUseCases) UpdateChain() usecases.UpdateChainUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChain")
	ret0, _ := ret[0].(usecases.UpdateChainUseCase)
	return ret0
}

// UpdateChain indicates an expected call of UpdateChain
func (mr *MockChainUseCasesMockRecorder) UpdateChain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChain", reflect.TypeOf((*MockChainUseCases)(nil).UpdateChain))
}

// DeleteChain mocks base method
func (m *MockChainUseCases) DeleteChain() usecases.DeleteChainUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChain")
	ret0, _ := ret[0].(usecases.DeleteChainUseCase)
	return ret0
}

// DeleteChain indicates an expected call of DeleteChain
func (mr *MockChainUseCasesMockRecorder) DeleteChain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChain", reflect.TypeOf((*MockChainUseCases)(nil).DeleteChain))
}

// MockRegisterChainUseCase is a mock of RegisterChainUseCase interface
type MockRegisterChainUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterChainUseCaseMockRecorder
}

// MockRegisterChainUseCaseMockRecorder is the mock recorder for MockRegisterChainUseCase
type MockRegisterChainUseCaseMockRecorder struct {
	mock *MockRegisterChainUseCase
}

// NewMockRegisterChainUseCase creates a new mock instance
func NewMockRegisterChainUseCase(ctrl *gomock.Controller) *MockRegisterChainUseCase {
	mock := &MockRegisterChainUseCase{ctrl: ctrl}
	mock.recorder = &MockRegisterChainUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegisterChainUseCase) EXPECT() *MockRegisterChainUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockRegisterChainUseCase) Execute(ctx context.Context, chain *entities.Chain, fromLatest bool, userInfo *multitenancy.UserInfo) (*entities.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, chain, fromLatest, userInfo)
	ret0, _ := ret[0].(*entities.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockRegisterChainUseCaseMockRecorder) Execute(ctx, chain, fromLatest, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockRegisterChainUseCase)(nil).Execute), ctx, chain, fromLatest, userInfo)
}

// MockGetChainUseCase is a mock of GetChainUseCase interface
type MockGetChainUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetChainUseCaseMockRecorder
}

// MockGetChainUseCaseMockRecorder is the mock recorder for MockGetChainUseCase
type MockGetChainUseCaseMockRecorder struct {
	mock *MockGetChainUseCase
}

// NewMockGetChainUseCase creates a new mock instance
func NewMockGetChainUseCase(ctrl *gomock.Controller) *MockGetChainUseCase {
	mock := &MockGetChainUseCase{ctrl: ctrl}
	mock.recorder = &MockGetChainUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetChainUseCase) EXPECT() *MockGetChainUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetChainUseCase) Execute(ctx context.Context, uuid string, userInfo *multitenancy.UserInfo) (*entities.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, uuid, userInfo)
	ret0, _ := ret[0].(*entities.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockGetChainUseCaseMockRecorder) Execute(ctx, uuid, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetChainUseCase)(nil).Execute), ctx, uuid, userInfo)
}

// MockSearchChainsUseCase is a mock of SearchChainsUseCase interface
type MockSearchChainsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSearchChainsUseCaseMockRecorder
}

// MockSearchChainsUseCaseMockRecorder is the mock recorder for MockSearchChainsUseCase
type MockSearchChainsUseCaseMockRecorder struct {
	mock *MockSearchChainsUseCase
}

// NewMockSearchChainsUseCase creates a new mock instance
func NewMockSearchChainsUseCase(ctrl *gomock.Controller) *MockSearchChainsUseCase {
	mock := &MockSearchChainsUseCase{ctrl: ctrl}
	mock.recorder = &MockSearchChainsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearchChainsUseCase) EXPECT() *MockSearchChainsUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSearchChainsUseCase) Execute(ctx context.Context, filters *entities.ChainFilters, userInfo *multitenancy.UserInfo) ([]*entities.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, filters, userInfo)
	ret0, _ := ret[0].([]*entities.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockSearchChainsUseCaseMockRecorder) Execute(ctx, filters, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSearchChainsUseCase)(nil).Execute), ctx, filters, userInfo)
}

// MockUpdateChainUseCase is a mock of UpdateChainUseCase interface
type MockUpdateChainUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateChainUseCaseMockRecorder
}

// MockUpdateChainUseCaseMockRecorder is the mock recorder for MockUpdateChainUseCase
type MockUpdateChainUseCaseMockRecorder struct {
	mock *MockUpdateChainUseCase
}

// NewMockUpdateChainUseCase creates a new mock instance
func NewMockUpdateChainUseCase(ctrl *gomock.Controller) *MockUpdateChainUseCase {
	mock := &MockUpdateChainUseCase{ctrl: ctrl}
	mock.recorder = &MockUpdateChainUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpdateChainUseCase) EXPECT() *MockUpdateChainUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockUpdateChainUseCase) Execute(ctx context.Context, chain *entities.Chain, userInfo *multitenancy.UserInfo) (*entities.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, chain, userInfo)
	ret0, _ := ret[0].(*entities.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockUpdateChainUseCaseMockRecorder) Execute(ctx, chain, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUpdateChainUseCase)(nil).Execute), ctx, chain, userInfo)
}

// MockDeleteChainUseCase is a mock of DeleteChainUseCase interface
type MockDeleteChainUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteChainUseCaseMockRecorder
}

// MockDeleteChainUseCaseMockRecorder is the mock recorder for MockDeleteChainUseCase
type MockDeleteChainUseCaseMockRecorder struct {
	mock *MockDeleteChainUseCase
}

// NewMockDeleteChainUseCase creates a new mock instance
func NewMockDeleteChainUseCase(ctrl *gomock.Controller) *MockDeleteChainUseCase {
	mock := &MockDeleteChainUseCase{ctrl: ctrl}
	mock.recorder = &MockDeleteChainUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeleteChainUseCase) EXPECT() *MockDeleteChainUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockDeleteChainUseCase) Execute(ctx context.Context, uuid string, userInfo *multitenancy.UserInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, uuid, userInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockDeleteChainUseCaseMockRecorder) Execute(ctx, uuid, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockDeleteChainUseCase)(nil).Execute), ctx, uuid, userInfo)
}
