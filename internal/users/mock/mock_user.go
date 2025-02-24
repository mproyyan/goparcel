// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go
//
// Generated by this command:
//
//	mockgen -source=./repository.go -destination=../../mock/mock_user.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	user "github.com/mproyyan/goparcel/internal/users/domain/user"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CheckEmailAvailability mocks base method.
func (m *MockUserRepository) CheckEmailAvailability(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailAvailability", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmailAvailability indicates an expected call of CheckEmailAvailability.
func (mr *MockUserRepositoryMockRecorder) CheckEmailAvailability(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailAvailability", reflect.TypeOf((*MockUserRepository)(nil).CheckEmailAvailability), ctx, email)
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(ctx context.Context, user user.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), ctx, user)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(ctx context.Context, email string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", ctx, email)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), ctx, email)
}

// MockCacheRepository is a mock of CacheRepository interface.
type MockCacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCacheRepositoryMockRecorder
	isgomock struct{}
}

// MockCacheRepositoryMockRecorder is the mock recorder for MockCacheRepository.
type MockCacheRepositoryMockRecorder struct {
	mock *MockCacheRepository
}

// NewMockCacheRepository creates a new mock instance.
func NewMockCacheRepository(ctrl *gomock.Controller) *MockCacheRepository {
	mock := &MockCacheRepository{ctrl: ctrl}
	mock.recorder = &MockCacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheRepository) EXPECT() *MockCacheRepositoryMockRecorder {
	return m.recorder
}

// CacheUserPermissions mocks base method.
func (m *MockCacheRepository) CacheUserPermissions(ctx context.Context, userID string, permissions user.Permissions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CacheUserPermissions", ctx, userID, permissions)
	ret0, _ := ret[0].(error)
	return ret0
}

// CacheUserPermissions indicates an expected call of CacheUserPermissions.
func (mr *MockCacheRepositoryMockRecorder) CacheUserPermissions(ctx, userID, permissions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheUserPermissions", reflect.TypeOf((*MockCacheRepository)(nil).CacheUserPermissions), ctx, userID, permissions)
}
