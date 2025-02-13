package adapter

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mproyyan/goparcel/internal/users/domain/user"
	"github.com/mproyyan/goparcel/internal/users/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (u *UserRepository) FindUserByEmail(ctx context.Context, email string) (*user.User, error) {
	// Lookup to user_types then to permissions
	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.M{"email": email}},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "user_types"},
				{Key: "localField", Value: "user_type"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "user_type"},
			}},
		},
		{
			{Key: "$unwind", Value: "$user_type"},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "permissions"},
				{Key: "localField", Value: "user_type.permission_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "user_type.permission"},
			}},
		},
		{
			{Key: "$unwind", Value: "$user_type.permission"},
		},
		{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "model_id", Value: 1},
				{Key: "email", Value: 1},
				{Key: "password", Value: 1},
				{Key: "user_type._id", Value: 1},
				{Key: "user_type.name", Value: 1},
				{Key: "user_type.description", Value: 1},
				{Key: "user_type.permission", Value: 1},
			}},
		},
	}

	// Execute query
	cursor, err := u.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		return modelToDomain(user), nil
	}

	return nil, errors.ErrUserNotFound
}

func (u *UserRepository) CreateUser(ctx context.Context, user user.User) (string, error) {
	// Insert new user
	result, err := u.collection.InsertOne(ctx, bson.M{
		"_id":          user.ID,
		"model_id":     user.ModelID,
		"email":        user.Email,
		"password":     user.Password,
		"user_type_id": user.Type.ID,
	})

	if err != nil {
		return "", err
	}

	// Return inserted id
	return result.InsertedID.(string), nil
}

// User models
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ModelID  primitive.ObjectID `bson:"model_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	UserType UserType           `bson:"user_type"`
}

// Helper function to get all granted permissions
func (p Permission) GrantedPermissions() []string {
	var permissions []string

	// Use reflect to read each field in a struct
	v := reflect.ValueOf(p)
	for i := 0; i < v.NumField(); i++ {
		category := v.Type().Field(i).Name // Category name (user_management, shipment, etc)
		subStruct := v.Field(i)

		// Iterate each field in sub-struct
		for j := 0; j < subStruct.NumField(); j++ {
			field := subStruct.Type().Field(j).Name // Permission field (get_user, create_shipment, etc)
			value := subStruct.Field(j).Bool()      // Permission value

			if value {
				// Only append granted permission
				permissions = append(permissions, fmt.Sprintf("%s.%s", toSnakeCase(category), toSnakeCase(field)))
			}
		}
	}

	return permissions
}

// Helper function to convert user model to user domain
func modelToDomain(userModel User) *user.User {
	return &user.User{
		ID:       userModel.ID.Hex(),
		ModelID:  userModel.ModelID.Hex(),
		Email:    userModel.Email,
		Password: userModel.Password,
		Type: user.UserType{
			ID:          userModel.UserType.ID.Hex(),
			Name:        userModel.UserType.Name,
			Description: userModel.UserType.Description,
			Permissions: userModel.UserType.Permission.GrantedPermissions(),
		},
	}
}

// Helper function to create snake case for permission name
func toSnakeCase(s string) string {
	result := ""
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result += "_"
		}
		result += string(r)
	}
	return result
}
