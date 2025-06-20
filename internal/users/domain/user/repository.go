//go:generate mockgen -source=./repository.go -destination=../../mock/mock_user.go -package=mock

package user

import (
	"context"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user User) (string, error)
	CheckEmailAvailability(ctx context.Context, email string) (bool, error)
	GetUser(ctx context.Context, id string) (*User, error)
	GetUsers(ctx context.Context, ids []string) ([]*User, error)
}
