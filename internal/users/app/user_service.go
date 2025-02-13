package app

import (
	"context"
	"time"

	"github.com/mproyyan/goparcel/internal/common/auth"
	"github.com/mproyyan/goparcel/internal/users/domain/operator"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	cuserr "github.com/mproyyan/goparcel/internal/users/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	UserRepository     user.UserRepository
	UserTypeRepository user.UserTypeRepository
	OperatorRepository operator.OperatorRepository
}

func NewUserService(
	userRepository user.UserRepository,
	userTypeRepository user.UserTypeRepository,
	operatorRepository operator.OperatorRepository,
) UserService {
	return UserService{
		UserRepository:     userRepository,
		UserTypeRepository: userTypeRepository,
		OperatorRepository: operatorRepository,
	}
}

func (u UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.UserRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if authenticated := auth.CheckPassword(user.Password, password); !authenticated {
		return "", cuserr.ErrInvalidCredentials
	}

	return auth.GenerateToken(user.ID, user.ModelID, time.Hour)
}

func (u UserService) RegisterAsOperator(ctx context.Context, name, email, password, location string, operatorType operator.OperatorType) error {
	// Find user type
	userType, err := u.UserTypeRepository.FindUserType(ctx, operatorType.String())
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
	_, err = u.UserRepository.CreateUser(ctx, user.User{
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
	_, err = u.OperatorRepository.CreateOperator(ctx, operator.Operator{
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
