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

	return nil, status.Errorf(codes.NotFound, "user with email %s not found", email)
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

func (u *UserRepository) FetchUserEntity(ctx context.Context, userID string, entity user.UserEntityName) (*user.UserEntity, error) {
	// Validate entity
	if entity == user.Unknown {
		return nil, status.Error(codes.InvalidArgument, "invalid entity")
	}

	// Convert userID to ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	// Get collection name to lookup from current entity name
	collectionName := entity.CollectionName()

	// Define aggregation pipeline with lookup
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: objID}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: collectionName},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "user_id"},
			{Key: "as", Value: "entity"},
		}}},
		{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$entity"},
			{Key: "preserveNullAndEmptyArrays", Value: false},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: "$entity._id"},
			{Key: "user_id", Value: "$_id"},
			{Key: "name", Value: "$entity.name"},
			{Key: "email", Value: "$entity.email"},
			{Key: "location_id", Value: "$entity.location_id"},
		}}},
	}

	// Execute aggregation
	cursor, err := u.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}
	defer cursor.Close(ctx)

	// Parse result
	var result UserEntity
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to decode user entity: %v", err)
		}
	} else {
		return nil, status.Error(codes.NotFound, "user entity not found")
	}

	userEntityModel := userEntityModelToDomain(result)
	return &userEntityModel, nil
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

type UserEntity struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id"`
	Name       string             `bson:"name"`
	Email      string             `bson:"email"`
	LocationID primitive.ObjectID `bson:"location_id"`
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
	userID, err := db.ConvertToObjectId(user.ID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "permission_id is not valid object id")
	}

	modelID, err := db.ConvertToObjectId(user.ModelID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "model_id is not valid object id")
	}

	userTypeID, err := db.ConvertToObjectId(user.Type.ID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_type_id is not valid object id")
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

// Helper function to conver user entity model to user entity domain
func userEntityModelToDomain(userEntity UserEntity) user.UserEntity {
	return user.UserEntity{
		ID:         userEntity.ID.Hex(),
		UserID:     userEntity.UserID.Hex(),
		Name:       userEntity.Name,
		Email:      userEntity.Email,
		LocationID: userEntity.LocationID.Hex(),
	}
}
