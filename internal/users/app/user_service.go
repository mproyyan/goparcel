package app

import (
	"context"
	"log"
	"time"

	"github.com/mproyyan/goparcel/internal/common/auth"
	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/users/domain/carrier"
	"github.com/mproyyan/goparcel/internal/users/domain/courier"
	"github.com/mproyyan/goparcel/internal/users/domain/operator"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	transaction        db.TransactionManager
	userRepository     user.UserRepository
	userTypeRepository user.UserTypeRepository
	operatorRepository operator.OperatorRepository
	carrierRepository  carrier.CarrierRepository
	courierRepository  courier.CourierRepository
	cacheRepository    user.CacheRepository
	locationService    LocationService
}

func NewUserService(
	transaction db.TransactionManager,
	userRepository user.UserRepository,
	userTypeRepository user.UserTypeRepository,
	operatorRepository operator.OperatorRepository,
	carrierRepository carrier.CarrierRepository,
	courierRepository courier.CourierRepository,
	cacheRepository user.CacheRepository,
	locationService LocationService,
) UserService {
	return UserService{
		transaction:        transaction,
		userRepository:     userRepository,
		userTypeRepository: userTypeRepository,
		operatorRepository: operatorRepository,
		carrierRepository:  carrierRepository,
		courierRepository:  courierRepository,
		cacheRepository:    cacheRepository,
		locationService:    locationService,
	}
}

type LocationService interface {
	GetLocation(ctx context.Context, locationID string) (*user.Location, error)
}

func (u UserService) Login(ctx context.Context, email, password string) (string, error) {
	// Find user
	user, err := u.userRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return "", cuserr.Decorate(err, "failed to find user by email")
	}

	// Compare given password with password stored in db
	if authenticated := auth.CheckPassword(user.Password, password); !authenticated {
		return "", status.Error(codes.Unauthenticated, "invalid credentials")
	}

	// Cache user permission
	err = u.cacheRepository.CacheUserPermissions(ctx, user.ID, user.Type.Permissions)
	if err != nil {
		log.Printf("failed to cache user permission: %v", err)
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.ModelID, time.Hour)
	if err != nil {
		return "", status.Error(codes.Internal, "failed to generate jwt token")
	}

	return token, nil
}

func (u UserService) RegisterAsOperator(ctx context.Context, name, email, password, location string, operatorType operator.OperatorType) error {
	// Email cannot be used by another user
	isAvailable, err := u.userRepository.CheckEmailAvailability(ctx, email)
	if err != nil {
		return cuserr.MongoError(err)
	}

	if !isAvailable {
		return status.Error(codes.AlreadyExists, "email already in use")
	}

	// Find user type
	userType, err := u.userTypeRepository.FindUserType(ctx, operatorType.String())
	if err != nil {
		return cuserr.Decorate(err, "failed to find user type")
	}

	// Run transaction
	// Create user and model (operator)
	err = u.transaction.Execute(ctx, func(ctx context.Context) error {
		// Define ObjectId for user and operator
		userId := primitive.NewObjectID()
		operatorId := primitive.NewObjectID()

		// Encrypt password
		encryptedPassword, err := auth.HashPassword(password)
		if err != nil {
			return cuserr.Decorate(err, "failed to hash password")
		}

		// Create user
		_, err = u.userRepository.CreateUser(ctx, user.User{
			ID:       userId.Hex(),
			ModelID:  operatorId.Hex(),
			Email:    email,
			Password: encryptedPassword,
			Type:     user.UserType{ID: userType.ID},
		})

		if err != nil {
			return cuserr.Decorate(err, "failed to create new user")
		}

		// Create operator
		_, err = u.operatorRepository.CreateOperator(ctx, operator.Operator{
			ID:         operatorId.Hex(),
			UserID:     userId.Hex(),
			Type:       operatorType,
			Name:       name,
			Email:      email,
			LocationID: location,
		})

		if err != nil {
			return cuserr.Decorate(err, "failed to create new operator")
		}

		return nil
	})

	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

func (u UserService) RegisterAsCarrier(ctx context.Context, name, email, password string) error {
	// Email cannot be used by another user
	isAvailable, err := u.userRepository.CheckEmailAvailability(ctx, email)
	if err != nil {
		return cuserr.MongoError(err)
	}

	if !isAvailable {
		return status.Error(codes.AlreadyExists, "email already in use")
	}

	// Find carrier user type
	userType, err := u.userTypeRepository.FindUserType(ctx, "carrier")
	if err != nil {
		return cuserr.Decorate(err, "failed to find user type")
	}

	// Run transaction
	// Create user and model (carrier)
	err = u.transaction.Execute(ctx, func(ctx context.Context) error {
		// Define ObjectId for user and carrier
		userId := primitive.NewObjectID()
		carrierId := primitive.NewObjectID()

		// Encrypt password
		encryptedPassword, err := auth.HashPassword(password)
		if err != nil {
			return cuserr.Decorate(err, "failed to hash password")
		}

		// Create user
		_, err = u.userRepository.CreateUser(ctx, user.User{
			ID:       userId.Hex(),
			ModelID:  carrierId.Hex(),
			Email:    email,
			Password: encryptedPassword,
			Type:     user.UserType{ID: userType.ID},
		})

		if err != nil {
			return cuserr.Decorate(err, "failed to create new user")
		}

		// Create carrier
		_, err = u.carrierRepository.CreateCarrier(ctx, carrier.Carrier{
			ID:     carrierId.Hex(),
			UserID: userId.Hex(),
			Name:   name,
			Email:  email,
		})

		if err != nil {
			return cuserr.Decorate(err, "failed to create new carrier")
		}

		return nil
	})

	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

func (u UserService) RegisterAsCourier(ctx context.Context, name, email, password string) error {
	// Email cannot be used by another user
	isAvailable, err := u.userRepository.CheckEmailAvailability(ctx, email)
	if err != nil {
		return cuserr.MongoError(err)
	}

	if !isAvailable {
		return status.Error(codes.AlreadyExists, "email already in use")
	}

	// Find carrier user type
	userType, err := u.userTypeRepository.FindUserType(ctx, "courier")
	if err != nil {
		return cuserr.Decorate(err, "failed to find user type")
	}

	// Run transaction
	// Create user and model (carrier)
	err = u.transaction.Execute(ctx, func(ctx context.Context) error {
		// Define ObjectId for user and carrier
		userId := primitive.NewObjectID()
		courierID := primitive.NewObjectID()

		// Encrypt password
		encryptedPassword, err := auth.HashPassword(password)
		if err != nil {
			return cuserr.Decorate(err, "failed to hash password")
		}

		// Create user
		_, err = u.userRepository.CreateUser(ctx, user.User{
			ID:       userId.Hex(),
			ModelID:  courierID.Hex(),
			Email:    email,
			Password: encryptedPassword,
			Type:     user.UserType{ID: userType.ID},
		})

		if err != nil {
			return cuserr.Decorate(err, "failed to create new user")
		}

		// Create carrier
		_, err = u.courierRepository.CreateCourier(ctx, courier.Courier{
			ID:     courierID.Hex(),
			UserID: userId.Hex(),
			Name:   name,
			Email:  email,
		})

		if err != nil {
			return cuserr.Decorate(err, "failed to create new courier")
		}

		return nil
	})

	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

func (u UserService) UserLocation(ctx context.Context, userID, entity string) (*user.Location, error) {
	entityName := user.StringToUserEntityName(entity)
	userEntity, err := u.userRepository.FetchUserEntity(ctx, userID, entityName)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to fetch user entity")
	}

	location, err := u.locationService.GetLocation(ctx, userEntity.LocationID)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get location using location service")
	}

	return location, nil
}
