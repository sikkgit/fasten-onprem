// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	models "github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockDatabaseRepository is a mock of DatabaseRepository interface.
type MockDatabaseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRepositoryMockRecorder
}

// MockDatabaseRepositoryMockRecorder is the mock recorder for MockDatabaseRepository.
type MockDatabaseRepositoryMockRecorder struct {
	mock *MockDatabaseRepository
}

// NewMockDatabaseRepository creates a new mock instance.
func NewMockDatabaseRepository(ctrl *gomock.Controller) *MockDatabaseRepository {
	mock := &MockDatabaseRepository{ctrl: ctrl}
	mock.recorder = &MockDatabaseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRepository) EXPECT() *MockDatabaseRepositoryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDatabaseRepository) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDatabaseRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDatabaseRepository)(nil).Close))
}

// CreateSource mocks base method.
func (m *MockDatabaseRepository) CreateSource(arg0 context.Context, arg1 *models.Source) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSource", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSource indicates an expected call of CreateSource.
func (mr *MockDatabaseRepositoryMockRecorder) CreateSource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSource", reflect.TypeOf((*MockDatabaseRepository)(nil).CreateSource), arg0, arg1)
}

// GetCurrentUser mocks base method.
func (m *MockDatabaseRepository) GetCurrentUser() models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUser")
	ret0, _ := ret[0].(models.User)
	return ret0
}

// GetCurrentUser indicates an expected call of GetCurrentUser.
func (mr *MockDatabaseRepositoryMockRecorder) GetCurrentUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUser", reflect.TypeOf((*MockDatabaseRepository)(nil).GetCurrentUser))
}

// GetSources mocks base method.
func (m *MockDatabaseRepository) GetSources(arg0 context.Context) ([]models.Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSources", arg0)
	ret0, _ := ret[0].([]models.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSources indicates an expected call of GetSources.
func (mr *MockDatabaseRepositoryMockRecorder) GetSources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSources", reflect.TypeOf((*MockDatabaseRepository)(nil).GetSources), arg0)
}

// UpsertOrganziation mocks base method.
func (m *MockDatabaseRepository) UpsertOrganziation(arg0 context.Context, arg1 *models.Organization) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertOrganziation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertOrganziation indicates an expected call of UpsertOrganziation.
func (mr *MockDatabaseRepositoryMockRecorder) UpsertOrganziation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertOrganziation", reflect.TypeOf((*MockDatabaseRepository)(nil).UpsertOrganziation), arg0, arg1)
}

// UpsertProfile mocks base method.
func (m *MockDatabaseRepository) UpsertProfile(arg0 context.Context, arg1 *models.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertProfile indicates an expected call of UpsertProfile.
func (mr *MockDatabaseRepositoryMockRecorder) UpsertProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertProfile", reflect.TypeOf((*MockDatabaseRepository)(nil).UpsertProfile), arg0, arg1)
}