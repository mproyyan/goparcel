package app

import (
	"context"
	"errors"
	"testing"

	"github.com/mproyyan/goparcel/internal/common/auth"
	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/users/domain/carrier"
	"github.com/mproyyan/goparcel/internal/users/domain/courier"
	"github.com/mproyyan/goparcel/internal/users/domain/operator"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	"github.com/mproyyan/goparcel/internal/users/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	cacheRepo := mock.NewMockCacheRepository(ctrl)
	userService := NewUserService(nil, userRepo, nil, nil, nil, nil, cacheRepo, nil)
	ctx := context.Background()

	tests := []struct {
		name          string
		user          *user.User
		rawPassword   string
		setupMock     func(*user.User)
		expectedError error
	}{
		{
			name:        "Login success",
			user:        newDataTestUser(),
			rawPassword: "password",
			setupMock: func(u *user.User) {
				userRepo.EXPECT().
					FindUserByEmail(ctx, u.Email).
					Return(u, nil)

				cacheRepo.EXPECT().
					CacheUserPermissions(ctx, u.ID, u.Type.Permissions).
					Return(nil)
			},
		},
		{
			name: "Cannot find user by email",
			user: newDataTestUser(),
			setupMock: func(u *user.User) {
				userRepo.EXPECT().
					FindUserByEmail(ctx, u.Email).
					Return(nil, errors.New("user not found"))
			},
			expectedError: cuserr.Decorate(errors.New("user not found"), "failed to find user by email"),
		},
		{
			name:        "Authentication failed",
			user:        newDataTestUser(),
			rawPassword: "wrongpassword",
			setupMock: func(u *user.User) {
				userRepo.EXPECT().
					FindUserByEmail(ctx, u.Email).
					Return(u, nil)
			},
			expectedError: status.Error(codes.Unauthenticated, "invalid credentials"),
		},
		{
			name:        "Failed to cache user permission",
			user:        newDataTestUser(),
			rawPassword: "password",
			setupMock: func(u *user.User) {
				userRepo.EXPECT().
					FindUserByEmail(ctx, u.Email).
					Return(u, nil)

				cacheRepo.EXPECT().
					CacheUserPermissions(ctx, u.ID, u.Type.Permissions).
					Return(errors.New("failed to cache user permission"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setupMock(test.user)
			token, err := userService.Login(ctx, test.user.Email, test.rawPassword)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, test.expectedError, err.Error())
			} else {
				claim, err := auth.Authenticate(token)
				assert.NoError(t, err)
				assert.Equal(t, claim.UserID, test.user.ID)
				assert.Equal(t, claim.ModelID, test.user.ModelID)
			}
		})
	}
}

func TestRegisterAsOperator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTypeRepo := mock.NewMockUserTypeRepository(ctrl)
	mockOperatorRepo := mock.NewMockOperatorRepository(ctrl)
	mockTransaction := db.NewMockTransactionManager(ctrl)
	userService := NewUserService(mockTransaction, mockUserRepo, mockUserTypeRepo, mockOperatorRepo, nil, nil, nil, nil)
	ctx := context.Background()

	tests := []struct {
		name          string
		user          *user.User
		operator      *operator.Operator
		setupMock     func(user.User, operator.Operator)
		expectedError error
	}{
		// Register as operator success
		{
			name:     "Register as operator success",
			user:     newDataTestUser(),
			operator: newDataTestOperator(),
			setupMock: func(u user.User, o operator.Operator) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, o.Type.String()).Return(&user.UserType{ID: u.Type.ID}, nil)
				mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return("", nil)
				mockOperatorRepo.EXPECT().CreateOperator(ctx, gomock.Any()).Return("", nil)
				mockTransaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
		},
		// Failed to check email availability
		{
			name:     "Failed to check email availability",
			user:     newDataTestUser(),
			operator: newDataTestOperator(),
			setupMock: func(u user.User, o operator.Operator) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(false, errors.New("failed to check email availability"))
			},
			expectedError: status.Errorf(codes.Internal, "unexpected database error: %v", errors.New("failed to check email availability").Error()),
		},
		// Email already in use
		{
			name:     "Email already in use",
			user:     newDataTestUser(),
			operator: newDataTestOperator(),
			setupMock: func(u user.User, o operator.Operator) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(false, nil)
			},
			expectedError: status.Error(codes.AlreadyExists, "email already in use"),
		},
		// Failed to find user type
		{
			name:     "Failed to find user type",
			user:     newDataTestUser(),
			operator: newDataTestOperator(),
			setupMock: func(u user.User, o operator.Operator) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, o.Type.String()).Return(nil, errors.New("user type not found"))
			},
			expectedError: cuserr.Decorate(errors.New("user type not found"), "failed to find user type"),
		},
		// Failed to create user
		{
			name:     "Failed to create user",
			user:     newDataTestUser(),
			operator: newDataTestOperator(),
			setupMock: func(u user.User, o operator.Operator) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, o.Type.String()).Return(&user.UserType{ID: u.Type.ID}, nil)
				mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return("", errors.New("internal error"))
				mockTransaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
			expectedError: status.Error(codes.Internal, "unexpected database error: failed to create new user, cause: internal error"),
		},
		// Failed to create operator
		{
			name:     "Failed to create operator",
			user:     newDataTestUser(),
			operator: newDataTestOperator(),
			setupMock: func(u user.User, o operator.Operator) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, o.Type.String()).Return(&user.UserType{ID: u.Type.ID}, nil)
				mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return("", nil)
				mockOperatorRepo.EXPECT().CreateOperator(ctx, gomock.Any()).Return("", errors.New("internal error"))
				mockTransaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
			expectedError: status.Error(codes.Internal, "unexpected database error: failed to create new operator, cause: internal error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setupMock(*test.user, *test.operator)
			err := userService.RegisterAsOperator(ctx, test.operator.Name, test.user.Email, "password", test.operator.LocationID, test.operator.Type)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, test.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRegisterAsCarrier(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTypeRepo := mock.NewMockUserTypeRepository(ctrl)
	mockCarrier := mock.NewMockCarrierRepository(ctrl)
	mockTransaction := db.NewMockTransactionManager(ctrl)
	userService := NewUserService(mockTransaction, mockUserRepo, mockUserTypeRepo, nil, mockCarrier, nil, nil, nil)
	ctx := context.Background()

	tests := []struct {
		name          string
		user          *user.User
		carrier       *carrier.Carrier
		setupMock     func(user.User, carrier.Carrier)
		expectedError error
	}{
		// Register as carrier success
		{
			name:    "Register as carrier success",
			user:    newDataTestUser(),
			carrier: newDataTestCarrier(),
			setupMock: func(u user.User, c carrier.Carrier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, "carrier").Return(&user.UserType{ID: u.Type.ID}, nil)
				mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return("", nil)
				mockCarrier.EXPECT().CreateCarrier(ctx, gomock.Any()).Return("", nil)
				mockTransaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
		},
		// Failed to check email availability
		{
			name:    "Failed to check email availability",
			user:    newDataTestUser(),
			carrier: newDataTestCarrier(),
			setupMock: func(u user.User, c carrier.Carrier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(false, errors.New("failed to check email availability"))
			},
			expectedError: status.Errorf(codes.Internal, "unexpected database error: %v", errors.New("failed to check email availability").Error()),
		},
		// Email already in use
		{
			name:    "Email already in use",
			user:    newDataTestUser(),
			carrier: newDataTestCarrier(),
			setupMock: func(u user.User, c carrier.Carrier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(false, nil)
			},
			expectedError: status.Error(codes.AlreadyExists, "email already in use"),
		},
		// Failed to find user type
		{
			name:    "Failed to find user type",
			user:    newDataTestUser(),
			carrier: newDataTestCarrier(),
			setupMock: func(u user.User, c carrier.Carrier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, "carrier").Return(nil, errors.New("user type not found"))
			},
			expectedError: cuserr.Decorate(errors.New("user type not found"), "failed to find user type"),
		},
		// Failed to create user
		{
			name:    "Failed to create user",
			user:    newDataTestUser(),
			carrier: newDataTestCarrier(),
			setupMock: func(u user.User, c carrier.Carrier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, "carrier").Return(&user.UserType{ID: u.Type.ID}, nil)
				mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return("", errors.New("internal error"))
				mockTransaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
			expectedError: status.Error(codes.Internal, "unexpected database error: failed to create new user, cause: internal error"),
		},
		// Failed to create carrier
		{
			name:    "Failed to create carrier",
			user:    newDataTestUser(),
			carrier: newDataTestCarrier(),
			setupMock: func(u user.User, c carrier.Carrier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, "carrier").Return(&user.UserType{ID: u.Type.ID}, nil)
				mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return("", nil)
				mockCarrier.EXPECT().CreateCarrier(ctx, gomock.Any()).Return("", errors.New("internal error"))
				mockTransaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
			expectedError: status.Error(codes.Internal, "unexpected database error: failed to create new carrier, cause: internal error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setupMock(*test.user, *test.carrier)
			err := userService.RegisterAsCarrier(ctx, test.carrier.Name, test.user.Email, test.user.Password)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, test.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRegisterAsCourier(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTypeRepo := mock.NewMockUserTypeRepository(ctrl)
	mockCourier := mock.NewMockCourierRepository(ctrl)
	mockTransaction := db.NewMockTransactionManager(ctrl)
	userService := NewUserService(mockTransaction, mockUserRepo, mockUserTypeRepo, nil, nil, mockCourier, nil, nil)
	ctx := context.Background()

	tests := []struct {
		name          string
		user          *user.User
		courier       *courier.Courier
		setupMock     func(user.User, courier.Courier)
		expectedError error
	}{
		// Register as carrier success
		{
			name:    "Register as courier success",
			user:    newDataTestUser(),
			courier: newDataTestCourier(),
			setupMock: func(u user.User, c courier.Courier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, "courier").Return(&user.UserType{ID: u.Type.ID}, nil)
				mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return("", nil)
				mockCourier.EXPECT().CreateCourier(ctx, gomock.Any()).Return("", nil)
				mockTransaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
		},
		// Failed to check email availability
		{
			name:    "Failed to check email availability",
			user:    newDataTestUser(),
			courier: newDataTestCourier(),
			setupMock: func(u user.User, c courier.Courier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(false, errors.New("failed to check email availability"))
			},
			expectedError: status.Errorf(codes.Internal, "unexpected database error: %v", errors.New("failed to check email availability").Error()),
		},
		// Email already in use
		{
			name:    "Email already in use",
			user:    newDataTestUser(),
			courier: newDataTestCourier(),
			setupMock: func(u user.User, c courier.Courier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(false, nil)
			},
			expectedError: status.Error(codes.AlreadyExists, "email already in use"),
		},
		// Failed to find user type
		{
			name:    "Failed to find user type",
			user:    newDataTestUser(),
			courier: newDataTestCourier(),
			setupMock: func(u user.User, c courier.Courier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, "courier").Return(nil, errors.New("user type not found"))
			},
			expectedError: cuserr.Decorate(errors.New("user type not found"), "failed to find user type"),
		},
		// Failed to create user
		{
			name:    "Failed to create user",
			user:    newDataTestUser(),
			courier: newDataTestCourier(),
			setupMock: func(u user.User, c courier.Courier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, "courier").Return(&user.UserType{ID: u.Type.ID}, nil)
				mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return("", errors.New("internal error"))
				mockTransaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
			expectedError: status.Error(codes.Internal, "unexpected database error: failed to create new user, cause: internal error"),
		},
		// Failed to create courier
		{
			name:    "Failed to create courier",
			user:    newDataTestUser(),
			courier: newDataTestCourier(),
			setupMock: func(u user.User, c courier.Courier) {
				mockUserRepo.EXPECT().CheckEmailAvailability(ctx, u.Email).Return(true, nil)
				mockUserTypeRepo.EXPECT().FindUserType(ctx, "courier").Return(&user.UserType{ID: u.Type.ID}, nil)
				mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return("", nil)
				mockCourier.EXPECT().CreateCourier(ctx, gomock.Any()).Return("", errors.New("internal error"))
				mockTransaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
			expectedError: status.Error(codes.Internal, "unexpected database error: failed to create new courier, cause: internal error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setupMock(*test.user, *test.courier)
			err := userService.RegisterAsCourier(ctx, test.courier.Name, test.user.Email, test.user.Password)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, test.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetUserLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	locationService := mock.NewMockLocationService(ctrl)
	userService := NewUserService(nil, userRepo, nil, nil, nil, nil, nil, locationService)
	ctx := context.Background()

	tests := []struct {
		name          string
		entity        user.UserEntityName
		userID        string
		setupMock     func(entity user.UserEntityName, userID string)
		expectedError error
	}{
		// Get user location success
		{
			name:   "Get user location success",
			entity: user.Operator,
			userID: "123",
			setupMock: func(e user.UserEntityName, uid string) {
				userRepo.EXPECT().FetchUserEntity(ctx, uid, e).Return(&user.UserEntity{UserID: uid, Entity: e, LocationID: "x"}, nil)
				locationService.EXPECT().GetLocation(ctx, "x").Return(&user.Location{ID: "x"}, nil)
			},
		},
		// Fetch user entity failed
		{
			name:   "Fetch user entity failed",
			entity: user.Operator,
			userID: "123",
			setupMock: func(e user.UserEntityName, uid string) {
				userRepo.EXPECT().FetchUserEntity(ctx, uid, e).Return(nil, errors.New("entity not found"))
			},
			expectedError: cuserr.Decorate(errors.New("entity not found"), "failed to fetch user entity"),
		},
		// Location not found
		{
			name:   "Location not found",
			entity: user.Operator,
			userID: "123",
			setupMock: func(e user.UserEntityName, uid string) {
				userRepo.EXPECT().FetchUserEntity(ctx, uid, e).Return(&user.UserEntity{UserID: uid, Entity: e, LocationID: "x"}, nil)
				locationService.EXPECT().GetLocation(ctx, "x").Return(nil, errors.New("location not found"))
			},
			expectedError: cuserr.Decorate(errors.New("location not found"), "failed to get location using location service"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setupMock(test.entity, test.userID)
			location, err := userService.UserLocation(ctx, test.userID, test.entity.String())

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, test.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, location)
			}
		})
	}
}

func newDataTestUser() *user.User {
	return &user.User{
		ID:       "xyz",
		ModelID:  "111",
		Email:    "user@test.com",
		Password: "$2a$10$W6T337haJUXIbJJsT46t1uBWFnNrJmQ9ZfkSVG3M7Dju59ZjtsTR.", // password
		Type: user.UserType{
			ID:          "zzz",
			Permissions: user.Permissions{"test.create_unit_test"},
		},
	}
}

func newDataTestOperator() *operator.Operator {
	return &operator.Operator{
		ID:         "111",
		UserID:     "xyz",
		Type:       operator.DepotOperator,
		LocationID: "x city",
		Name:       "Operator X",
	}
}

func newDataTestCarrier() *carrier.Carrier {
	return &carrier.Carrier{
		ID:     "222",
		UserID: "xyz",
		Name:   "Carrier X",
	}
}

func newDataTestCourier() *courier.Courier {
	return &courier.Courier{
		ID:     "333",
		UserID: "xyz",
		Name:   "Carrier X",
	}
}
