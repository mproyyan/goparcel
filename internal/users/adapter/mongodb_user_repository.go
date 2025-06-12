package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// Find user by email
	var user User
	err := u.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	// Return query result
	userResponse := userModelToDomain(user)
	return &userResponse, nil
}

func (u *UserRepository) CreateUser(ctx context.Context, user user.User) (string, error) {
	// Prepare data to insert
	userModel, err := domainToUserModel(user)
	if err != nil {
		return "", cuserr.Decorate(err, "failed to convert domain to user model")
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

func (u *UserRepository) GetUser(ctx context.Context, id string) (*user.User, error) {
	userObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user id is not valid object id")
	}

	// Find a single user by ID with a projection to only retrieve specific fields
	var usr User
	err = u.collection.FindOne(ctx, bson.M{"_id": userObjId}, options.FindOne().SetProjection(bson.M{
		"_id":      1,
		"model_id": 1,
		"entity":   1,
	})).Decode(&usr)

	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	userResponse := userModelToDomain(usr)
	return &userResponse, nil
}

func (u *UserRepository) GetUsers(ctx context.Context, ids []string) ([]*user.User, error) {
	userObjIds := make([]primitive.ObjectID, 0, len(ids))
	for _, id := range ids {
		userObjId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "user id is not valid object id")
		}

		userObjIds = append(userObjIds, userObjId)
	}

	var users []*User
	query := bson.M{}

	// If a list of IDs is provided, filter users by these IDs
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": userObjIds}
	}

	// Execute the find query with a projection to retrieve only specific fields
	cursor, err := u.collection.Find(ctx, query, options.Find().SetProjection(bson.M{
		"_id":      1,
		"model_id": 1,
		"entity":   1,
	}))

	if err != nil {
		return nil, cuserr.MongoError(err)
	}
	defer cursor.Close(ctx)

	// Decode the retrieved users into the users slice
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return userModelsToDomain(users), nil
}

// User models
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ModelID    primitive.ObjectID `bson:"model_id,omitempty"`
	Entity     string             `bson:"entity"`
	Email      string             `bson:"email"`
	Password   string             `bson:"password"`
	UserTypeID primitive.ObjectID `bson:"user_type_id"`
}

// Helper function to convert user model to user domain
func userModelToDomain(userModel User) user.User {
	return user.User{
		ID:         userModel.ID.Hex(),
		ModelID:    userModel.ModelID.Hex(),
		Entity:     user.StringToUserEntityName(userModel.Entity),
		Email:      userModel.Email,
		Password:   userModel.Password,
		UserTypeID: userModel.UserTypeID.Hex(),
	}
}

func userModelsToDomain(models []*User) []*user.User {
	var users []*user.User
	for _, model := range models {
		user := userModelToDomain(*model)
		users = append(users, &user)
	}

	return users
}

// Helper function to convert domain to user model
func domainToUserModel(user user.User) (*User, error) {
	// Convert string ObjectId to literal ObjectId
	userID, err := db.ConvertToObjectId(user.ID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	modelID, err := db.ConvertToObjectId(user.ModelID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "model_id is not valid object id")
	}

	userTypeID, err := db.ConvertToObjectId(user.UserTypeID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_type_id is not valid object id")
	}

	// Return user model
	return &User{
		ID:         userID,
		ModelID:    modelID,
		Entity:     user.Entity.String(),
		Email:      user.Email,
		Password:   user.Password,
		UserTypeID: userTypeID,
	}, nil
}
