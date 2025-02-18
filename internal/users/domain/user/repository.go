package user

import "context"

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user User) (string, error)
	CheckEmailAvailability(ctx context.Context, email string) (bool, error)
	FetchUserEntity(ctx context.Context, userID, entity string) (*UserEntity, error)
}

type UserTypeRepository interface {
	FindUserType(ctx context.Context, userType string) (*UserType, error)
}

type CacheRepository interface {
	CacheUserPermissions(ctx context.Context, userID string, permissions Permissions) error
}
