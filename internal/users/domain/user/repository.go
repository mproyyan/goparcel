//go:generate mockgen -source=./repository.go -destination=../../mock/mock_user.go -package=mock

package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user User) (string, error)
	CheckEmailAvailability(ctx context.Context, email string) (bool, error)
	FetchUserEntity(ctx context.Context, userID string, entity UserEntityName) (*UserEntity, error)
}

type UserTypeRepository interface {
	FindUserType(ctx context.Context, userType string) (*UserType, error)
	FindUserTypeById(ctx context.Context, id primitive.ObjectID) (*UserType, error)
}

type CacheRepository interface {
	CacheUserPermissions(ctx context.Context, userID string, permissions Permissions) error
}
