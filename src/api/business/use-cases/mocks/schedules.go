// Code generated by MockGen. DO NOT EDIT.
// Source: schedules.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	multitenancy "github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	entities "github.com/consensys/orchestrate/src/entities"
	usecases "github.com/consensys/orchestrate/src/api/business/use-cases"
	store "github.com/consensys/orchestrate/src/api/store"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockScheduleUseCases is a mock of ScheduleUseCases interface
type MockScheduleUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleUseCasesMockRecorder
}

// MockScheduleUseCasesMockRecorder is the mock recorder for MockScheduleUseCases
type MockScheduleUseCasesMockRecorder struct {
	mock *MockScheduleUseCases
}

// NewMockScheduleUseCases creates a new mock instance
func NewMockScheduleUseCases(ctrl *gomock.Controller) *MockScheduleUseCases {
	mock := &MockScheduleUseCases{ctrl: ctrl}
	mock.recorder = &MockScheduleUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScheduleUseCases) EXPECT() *MockScheduleUseCasesMockRecorder {
	return m.recorder
}

// CreateSchedule mocks base method
func (m *MockScheduleUseCases) CreateSchedule() usecases.CreateScheduleUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchedule")
	ret0, _ := ret[0].(usecases.CreateScheduleUseCase)
	return ret0
}

// CreateSchedule indicates an expected call of CreateSchedule
func (mr *MockScheduleUseCasesMockRecorder) CreateSchedule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchedule", reflect.TypeOf((*MockScheduleUseCases)(nil).CreateSchedule))
}

// GetSchedule mocks base method
func (m *MockScheduleUseCases) GetSchedule() usecases.GetScheduleUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule")
	ret0, _ := ret[0].(usecases.GetScheduleUseCase)
	return ret0
}

// GetSchedule indicates an expected call of GetSchedule
func (mr *MockScheduleUseCasesMockRecorder) GetSchedule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockScheduleUseCases)(nil).GetSchedule))
}

// SearchSchedules mocks base method
func (m *MockScheduleUseCases) SearchSchedules() usecases.SearchSchedulesUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchSchedules")
	ret0, _ := ret[0].(usecases.SearchSchedulesUseCase)
	return ret0
}

// SearchSchedules indicates an expected call of SearchSchedules
func (mr *MockScheduleUseCasesMockRecorder) SearchSchedules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchSchedules", reflect.TypeOf((*MockScheduleUseCases)(nil).SearchSchedules))
}

// MockCreateScheduleUseCase is a mock of CreateScheduleUseCase interface
type MockCreateScheduleUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCreateScheduleUseCaseMockRecorder
}

// MockCreateScheduleUseCaseMockRecorder is the mock recorder for MockCreateScheduleUseCase
type MockCreateScheduleUseCaseMockRecorder struct {
	mock *MockCreateScheduleUseCase
}

// NewMockCreateScheduleUseCase creates a new mock instance
func NewMockCreateScheduleUseCase(ctrl *gomock.Controller) *MockCreateScheduleUseCase {
	mock := &MockCreateScheduleUseCase{ctrl: ctrl}
	mock.recorder = &MockCreateScheduleUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreateScheduleUseCase) EXPECT() *MockCreateScheduleUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockCreateScheduleUseCase) Execute(ctx context.Context, schedule *entities.Schedule, userInfo *multitenancy.UserInfo) (*entities.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, schedule, userInfo)
	ret0, _ := ret[0].(*entities.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockCreateScheduleUseCaseMockRecorder) Execute(ctx, schedule, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCreateScheduleUseCase)(nil).Execute), ctx, schedule, userInfo)
}

// WithDBTransaction mocks base method
func (m *MockCreateScheduleUseCase) WithDBTransaction(dbtx store.Tx) usecases.CreateScheduleUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDBTransaction", dbtx)
	ret0, _ := ret[0].(usecases.CreateScheduleUseCase)
	return ret0
}

// WithDBTransaction indicates an expected call of WithDBTransaction
func (mr *MockCreateScheduleUseCaseMockRecorder) WithDBTransaction(dbtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDBTransaction", reflect.TypeOf((*MockCreateScheduleUseCase)(nil).WithDBTransaction), dbtx)
}

// MockGetScheduleUseCase is a mock of GetScheduleUseCase interface
type MockGetScheduleUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetScheduleUseCaseMockRecorder
}

// MockGetScheduleUseCaseMockRecorder is the mock recorder for MockGetScheduleUseCase
type MockGetScheduleUseCaseMockRecorder struct {
	mock *MockGetScheduleUseCase
}

// NewMockGetScheduleUseCase creates a new mock instance
func NewMockGetScheduleUseCase(ctrl *gomock.Controller) *MockGetScheduleUseCase {
	mock := &MockGetScheduleUseCase{ctrl: ctrl}
	mock.recorder = &MockGetScheduleUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetScheduleUseCase) EXPECT() *MockGetScheduleUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetScheduleUseCase) Execute(ctx context.Context, scheduleUUID string, userInfo *multitenancy.UserInfo) (*entities.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, scheduleUUID, userInfo)
	ret0, _ := ret[0].(*entities.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockGetScheduleUseCaseMockRecorder) Execute(ctx, scheduleUUID, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetScheduleUseCase)(nil).Execute), ctx, scheduleUUID, userInfo)
}

// MockSearchSchedulesUseCase is a mock of SearchSchedulesUseCase interface
type MockSearchSchedulesUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSearchSchedulesUseCaseMockRecorder
}

// MockSearchSchedulesUseCaseMockRecorder is the mock recorder for MockSearchSchedulesUseCase
type MockSearchSchedulesUseCaseMockRecorder struct {
	mock *MockSearchSchedulesUseCase
}

// NewMockSearchSchedulesUseCase creates a new mock instance
func NewMockSearchSchedulesUseCase(ctrl *gomock.Controller) *MockSearchSchedulesUseCase {
	mock := &MockSearchSchedulesUseCase{ctrl: ctrl}
	mock.recorder = &MockSearchSchedulesUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearchSchedulesUseCase) EXPECT() *MockSearchSchedulesUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSearchSchedulesUseCase) Execute(ctx context.Context, userInfo *multitenancy.UserInfo) ([]*entities.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, userInfo)
	ret0, _ := ret[0].([]*entities.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockSearchSchedulesUseCaseMockRecorder) Execute(ctx, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSearchSchedulesUseCase)(nil).Execute), ctx, userInfo)
}