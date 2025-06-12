package user

import (
	"context"
)

type UserTypeRepository interface {
	FindUserType(ctx context.Context, userType string) (*UserType, error)
	FindUserTypeById(ctx context.Context, id string) (*UserType, error)
}

type UserType struct {
	ID          string
	Name        string
	Description string
}
