package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserTypeRepository interface {
	FindUserType(ctx context.Context, userType string) (*UserType, error)
	FindUserTypeById(ctx context.Context, id primitive.ObjectID) (*UserType, error)
}

type UserType struct {
	ID           string
	Name         string
	Description  string
	PermissionID string
}
