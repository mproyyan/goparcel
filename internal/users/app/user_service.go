package app

import (
	"context"
	"errors"
	"time"

	"github.com/mproyyan/goparcel/internal/common/auth"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
)

type UserService struct {
	UserRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) UserService {
	return UserService{
		UserRepository: userRepository,
	}
}

func (u UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.UserRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if authenticated := auth.CheckPassword(user.Password, password); !authenticated {
		return "", errors.New("invalid credentials")
	}

	return auth.GenerateToken(user.ID, user.ModelID, time.Hour)
}
