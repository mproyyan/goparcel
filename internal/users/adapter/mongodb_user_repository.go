package adapter

import (
	"context"
	"fmt"

	"github.com/mproyyan/goparcel/internal/common/utils"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	cuserr "github.com/mproyyan/goparcel/internal/users/errors"

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
				{Key: "localField", Value: "user_type_id"},
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
		return nil, cuserr.MongoError(err)
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, cuserr.MongoError(err)
		}

		userDomain := userModelToDomain(user)
		return &userDomain, nil
	}

	fmt.Println("user not found")
	return nil, cuserr.ErrUserNotFound
}

func (u *UserRepository) CreateUser(ctx context.Context, user user.User) (string, error) {
	// Prepare data to insert
	userModel, err := domainToUserModel(user)
	if err != nil {
		return "", err
	}

	// Insert new user
	result, err := u.collection.InsertOne(ctx, userModel)
	if err != nil {
		return "", cuserr.MongoError(err)
	}

	// Return inserted id
	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID.Hex(), nil
}

func (u *UserRepository) CheckEmailAvailability(ctx context.Context, email string) (bool, error) {
	results, err := u.collection.Distinct(ctx, "email", bson.M{"email": email})
	if err != nil {
		return false, cuserr.MongoError(err)
	}

	// If result 0 means email available
	return len(results) == 0, nil
}

// User models
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ModelID    primitive.ObjectID `bson:"model_id,omitempty"`
	Email      string             `bson:"email"`
	Password   string             `bson:"password"`
	UserTypeID primitive.ObjectID `bson:"user_type_id"`

	// Relations
	UserType *UserType `bson:"user_type,omitempty"`
}

func (u *User) hasUserType() bool {
	return u.UserType != nil
}

// Helper function to convert user model to user domain
func userModelToDomain(userModel User) user.User {
	// Check if has relation to user type
	userType := user.UserType{ID: userModel.UserTypeID.Hex()}
	if userModel.hasUserType() {
		userType = userTypeModelToDomain(*userModel.UserType)
	}

	return user.User{
		ID:       userModel.ID.Hex(),
		ModelID:  userModel.ModelID.Hex(),
		Email:    userModel.Email,
		Password: userModel.Password,
		Type:     userType,
	}
}

// Helper function to convert domain to user model
func domainToUserModel(user user.User) (*User, error) {
	// Convert string ObjectId to literal ObjectId
	userID, err := utils.ConvertToObjectId(user.ID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	modelID, err := utils.ConvertToObjectId(user.ModelID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	userTypeID, err := utils.ConvertToObjectId(user.Type.ID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	// Return user model
	return &User{
		ID:         userID,
		ModelID:    modelID,
		Email:      user.Email,
		Password:   user.Password,
		UserTypeID: userTypeID,
	}, nil
}
