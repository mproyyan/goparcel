package app

import (
	"context"
	"log"
	"time"

	"github.com/mproyyan/goparcel/internal/common/auth"
	"github.com/mproyyan/goparcel/internal/users/domain/carrier"
	"github.com/mproyyan/goparcel/internal/users/domain/courier"
	"github.com/mproyyan/goparcel/internal/users/domain/operator"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	cuserr "github.com/mproyyan/goparcel/internal/users/errors"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepository     user.UserRepository
	userTypeRepository user.UserTypeRepository
	operatorRepository operator.OperatorRepository
	carrierRepository  carrier.CarrierRepository
	courierRepository  courier.CourierRepository
	redisClient        *redis.Client
}

func NewUserService(
	userRepository user.UserRepository,
	userTypeRepository user.UserTypeRepository,
	operatorRepository operator.OperatorRepository,
	carrierRepository carrier.CarrierRepository,
	courierRepository courier.CourierRepository,
	redisClient *redis.Client,
) UserService {
	return UserService{
		userRepository:     userRepository,
		userTypeRepository: userTypeRepository,
		operatorRepository: operatorRepository,
		carrierRepository:  carrierRepository,
		courierRepository:  courierRepository,
		redisClient:        redisClient,
	}
}

func (u UserService) Login(ctx context.Context, email, password string) (string, error) {
	// Find user
	user, err := u.userRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	// Compare given password with password stored in db
	if authenticated := auth.CheckPassword(user.Password, password); !authenticated {
		return "", cuserr.ErrInvalidCredentials
	}

	// Save user permissions to redis
	err = u.redisClient.Set(ctx, user.ID, user.Type.Permissions, time.Hour).Err()
	if err != nil {
		log.Printf("failed to cache user permission: %v", err)
	}

	return auth.GenerateToken(user.ID, user.ModelID, time.Hour)
}

func (u UserService) RegisterAsOperator(ctx context.Context, name, email, password, location string, operatorType operator.OperatorType) error {
	// Email cannot be used by another user
	isAvailable, err := u.userRepository.CheckEmailAvailability(ctx, email)
	if err != nil {
		return cuserr.MongoError(err)
	}

	if !isAvailable {
		return cuserr.ErrEmailUsed
	}

	// Find user type
	userType, err := u.userTypeRepository.FindUserType(ctx, operatorType.String())
	if err != nil {
		return err
	}

	// Define ObjectId for user and operator
	userId := primitive.NewObjectID()
	operatorId := primitive.NewObjectID()

	// Encrypt password
	encryptedPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
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
		return err
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
		return err
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
		return cuserr.ErrEmailUsed
	}

	// Find carrier user type
	userType, err := u.userTypeRepository.FindUserType(ctx, "carrier")
	if err != nil {
		return err
	}

	// Define ObjectId for user and carrier
	userId := primitive.NewObjectID()
	carrierId := primitive.NewObjectID()

	// Encrypt password
	encryptedPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
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
		return err
	}

	// Create carrier
	_, err = u.carrierRepository.CreateCarrier(ctx, carrier.Carrier{
		ID:     carrierId.Hex(),
		UserID: userId.Hex(),
		Name:   name,
		Email:  email,
	})

	if err != nil {
		return err
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
		return cuserr.ErrEmailUsed
	}

	// Find carrier user type
	userType, err := u.userTypeRepository.FindUserType(ctx, "courier")
	if err != nil {
		return err
	}

	// Define ObjectId for user and carrier
	userId := primitive.NewObjectID()
	courierID := primitive.NewObjectID()

	// Encrypt password
	encryptedPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
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
		return err
	}

	// Create carrier
	_, err = u.courierRepository.CreateCourier(ctx, courier.Courier{
		ID:     courierID.Hex(),
		UserID: userId.Hex(),
		Name:   name,
		Email:  email,
	})

	if err != nil {
		return err
	}

	return nil
}
