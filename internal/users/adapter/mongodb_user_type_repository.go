package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserTypeRepository struct {
	collection *mongo.Collection
}

func NewUserTypeRepository(db *mongo.Database) *UserTypeRepository {
	return &UserTypeRepository{
		collection: db.Collection("user_types"),
	}
}

func (u *UserTypeRepository) FindUserType(ctx context.Context, userType string) (*user.UserType, error) {
	// Find user type by name
	var userTypeResult UserType
	logrus.WithField("user_type", userType).Info("Finding user type")
	err := u.collection.FindOne(ctx, bson.M{"name": userType}).Decode(&userTypeResult)
	if err != nil {
		logrus.WithError(err).Error("Failed to find user type")
		return nil, cuserr.MongoError(err)
	}

	userTypeModel := userTypeModelToDomain(userTypeResult)
	return &userTypeModel, nil
}

func (u *UserTypeRepository) FindUserTypeById(ctx context.Context, id string) (*user.UserType, error) {
	userTypeObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user type id is not valid object id")
	}

	// Find user type by id
	var userTypeResult UserType
	err = u.collection.FindOne(ctx, bson.M{"_id": userTypeObjId}).Decode(&userTypeResult)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	userTypeModel := userTypeModelToDomain(userTypeResult)
	return &userTypeModel, nil
}

// models
type UserType struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	PermissionID primitive.ObjectID `bson:"permission_id"`
}

// Helper function to convert user type model to domain
func userTypeModelToDomain(userTypeModel UserType) user.UserType {
	return user.UserType{
		ID:           userTypeModel.ID.Hex(),
		Name:         userTypeModel.Name,
		Description:  userTypeModel.Description,
		PermissionID: userTypeModel.PermissionID.Hex(),
	}
}

// Helper function to convert domain to user type model
func domainToUserTypeModel(userType user.UserType) (*UserType, error) {
	userTypeID, err := db.ConvertToObjectId(userType.ID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_type_id is not valid object id")
	}

	permissionID, err := db.ConvertToObjectId(userType.PermissionID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "permission_id is not valid object id")
	}

	return &UserType{
		ID:           userTypeID,
		Name:         userType.Name,
		Description:  userType.Description,
		PermissionID: permissionID,
	}, nil
}
