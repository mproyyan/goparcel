// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go
//
// Generated by this command:
//
//	mockgen -source=./repository.go -destination=../../mock/mock_carrier.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	carrier "github.com/mproyyan/goparcel/internal/users/domain/carrier"
	gomock "go.uber.org/mock/gomock"
)

// MockCarrierRepository is a mock of CarrierRepository interface.
type MockCarrierRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCarrierRepositoryMockRecorder
	isgomock struct{}
}

// MockCarrierRepositoryMockRecorder is the mock recorder for MockCarrierRepository.
type MockCarrierRepositoryMockRecorder struct {
	mock *MockCarrierRepository
}

// NewMockCarrierRepository creates a new mock instance.
func NewMockCarrierRepository(ctrl *gomock.Controller) *MockCarrierRepository {
	mock := &MockCarrierRepository{ctrl: ctrl}
	mock.recorder = &MockCarrierRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarrierRepository) EXPECT() *MockCarrierRepositoryMockRecorder {
	return m.recorder
}

// CreateCarrier mocks base method.
func (m *MockCarrierRepository) CreateCarrier(ctx context.Context, carrier carrier.Carrier) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCarrier", ctx, carrier)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCarrier indicates an expected call of CreateCarrier.
func (mr *MockCarrierRepositoryMockRecorder) CreateCarrier(ctx, carrier any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCarrier", reflect.TypeOf((*MockCarrierRepository)(nil).CreateCarrier), ctx, carrier)
}
