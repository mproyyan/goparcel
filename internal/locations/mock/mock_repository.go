// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go
//
// Generated by this command:
//
//	mockgen -source=./repository.go -destination=../mock/mock_repository.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/mproyyan/goparcel/internal/locations/domain"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	gomock "go.uber.org/mock/gomock"
)

// MockLocationRepository is a mock of LocationRepository interface.
type MockLocationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLocationRepositoryMockRecorder
	isgomock struct{}
}

// MockLocationRepositoryMockRecorder is the mock recorder for MockLocationRepository.
type MockLocationRepositoryMockRecorder struct {
	mock *MockLocationRepository
}

// NewMockLocationRepository creates a new mock instance.
func NewMockLocationRepository(ctrl *gomock.Controller) *MockLocationRepository {
	mock := &MockLocationRepository{ctrl: ctrl}
	mock.recorder = &MockLocationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocationRepository) EXPECT() *MockLocationRepositoryMockRecorder {
	return m.recorder
}

// CreateLocation mocks base method.
func (m *MockLocationRepository) CreateLocation(ctx context.Context, location domain.Location) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLocation", ctx, location)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLocation indicates an expected call of CreateLocation.
func (mr *MockLocationRepositoryMockRecorder) CreateLocation(ctx, location any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLocation", reflect.TypeOf((*MockLocationRepository)(nil).CreateLocation), ctx, location)
}

// FindLocation mocks base method.
func (m *MockLocationRepository) FindLocation(ctx context.Context, locationID string) (*domain.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLocation", ctx, locationID)
	ret0, _ := ret[0].(*domain.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLocation indicates an expected call of FindLocation.
func (mr *MockLocationRepositoryMockRecorder) FindLocation(ctx, locationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLocation", reflect.TypeOf((*MockLocationRepository)(nil).FindLocation), ctx, locationID)
}

// FindTransitPlaces mocks base method.
func (m *MockLocationRepository) FindTransitPlaces(ctx context.Context, locationID primitive.ObjectID) ([]domain.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTransitPlaces", ctx, locationID)
	ret0, _ := ret[0].([]domain.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTransitPlaces indicates an expected call of FindTransitPlaces.
func (mr *MockLocationRepositoryMockRecorder) FindTransitPlaces(ctx, locationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTransitPlaces", reflect.TypeOf((*MockLocationRepository)(nil).FindTransitPlaces), ctx, locationID)
}
