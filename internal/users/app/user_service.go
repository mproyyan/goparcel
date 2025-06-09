//go:generate mockgen -source user_service.go -destination ../mock/mock_external_service.go -package mock

package app

import (
	"context"
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
}

func NewUserService(
	transaction db.TransactionManager,
	userRepository user.UserRepository,
	userTypeRepository user.UserTypeRepository,
	operatorRepository operator.OperatorRepository,
	carrierRepository carrier.CarrierRepository,
	courierRepository courier.CourierRepository,
) UserService {
	return UserService{
		transaction:        transaction,
		userRepository:     userRepository,
		userTypeRepository: userTypeRepository,
		operatorRepository: operatorRepository,
		carrierRepository:  carrierRepository,
		courierRepository:  courierRepository,
	}
}

func (u UserService) GetUser(ctx context.Context, id string) (*user.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user id is not valid object id")
	}

	user, err := u.userRepository.GetUser(ctx, objId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get user from repository")
	}

	return user, nil
}

func (u UserService) GetUsers(ctx context.Context, ids []string) ([]*user.User, error) {
	var objIds []primitive.ObjectID
	for _, id := range ids {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.Internal, "user id is not valid object id")
		}
		objIds = append(objIds, objId)
	}

	users, err := u.userRepository.GetUsers(ctx, objIds)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get users from repository")
	}

	return users, nil
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

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.ModelID, time.Hour*24*30*3)
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
			ID:         userId.Hex(),
			ModelID:    operatorId.Hex(),
			Entity:     user.Operator,
			Email:      email,
			Password:   encryptedPassword,
			UserTypeID: userType.ID,
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

func (u UserService) RegisterAsCarrier(ctx context.Context, name, email, password, location string) error {
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
			ID:         userId.Hex(),
			ModelID:    carrierId.Hex(),
			Entity:     user.Carrier,
			Email:      email,
			Password:   encryptedPassword,
			UserTypeID: userType.ID,
		})

		if err != nil {
			return cuserr.Decorate(err, "failed to create new user")
		}

		// Create carrier
		_, err = u.carrierRepository.CreateCarrier(ctx, carrier.Carrier{
			ID:         carrierId.Hex(),
			UserID:     userId.Hex(),
			Name:       name,
			Email:      email,
			LocationID: location,
			Status:     "idle",
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

func (u UserService) RegisterAsCourier(ctx context.Context, name, email, password, location string) error {
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
			ID:         userId.Hex(),
			ModelID:    courierID.Hex(),
			Entity:     user.Courier,
			Email:      email,
			Password:   encryptedPassword,
			UserTypeID: userType.ID,
		})

		if err != nil {
			return cuserr.Decorate(err, "failed to create new user")
		}

		// Create carrier
		_, err = u.courierRepository.CreateCourier(ctx, courier.Courier{
			ID:         courierID.Hex(),
			UserID:     userId.Hex(),
			Name:       name,
			Email:      email,
			Status:     courier.Available,
			LocationID: location,
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

func (u UserService) GetOperators(ctx context.Context, ids []string) ([]*operator.Operator, error) {
	var objIds []primitive.ObjectID
	for _, id := range ids {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.Internal, "operator id is not valid object id")
		}
		objIds = append(objIds, objId)
	}

	operators, err := u.operatorRepository.GetOperators(ctx, objIds)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get operators from repository")
	}

	return operators, nil
}

func (u UserService) GetCouriers(ctx context.Context, ids []string) ([]*courier.Courier, error) {
	var objIds []primitive.ObjectID
	for _, id := range ids {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.Internal, "operator id is not valid object id")
		}
		objIds = append(objIds, objId)
	}

	couriers, err := u.courierRepository.GetCouriers(ctx, objIds)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get couriers from repository")
	}

	return couriers, nil
}

func (u UserService) GetCarriers(ctx context.Context, ids []string) ([]*carrier.Carrier, error) {
	var objIds []primitive.ObjectID
	for _, id := range ids {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.Internal, "operator id is not valid object id")
		}
		objIds = append(objIds, objId)
	}

	carriers, err := u.carrierRepository.GetCarriers(ctx, objIds)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get carriers from repository")
	}

	return carriers, nil
}
