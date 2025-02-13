package user

import "context"

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*User, error)
}
