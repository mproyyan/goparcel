package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/users/domain/courier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CourierRepository struct {
	collection *mongo.Collection
}

func NewCourierRepository(db *mongo.Database) *CourierRepository {
	return &CourierRepository{
		collection: db.Collection("couriers"),
	}
}

func (c *CourierRepository) CreateCourier(ctx context.Context, courier courier.Courier) (string, error) {
	// Prepare data to insert
	courierModel, err := domainToCourierModel(courier)
	if err != nil {
		return "", cuserr.Decorate(err, "failed to convert domain to courier model")
	}

	// Create new carrier
	result, err := c.collection.InsertOne(ctx, courierModel)
	if err != nil {
		return "", cuserr.MongoError(err)
	}

	// Return inserted id
	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID.Hex(), nil
}

// Models
type CourierModel struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty"`
	UserID     *primitive.ObjectID `bson:"user_id,omitempty"`
	Name       string              `bson:"name"`
	Email      string              `bson:"email"`
	Status     string              `bson:"status"`
	LocationID *primitive.ObjectID `bson:"location_id,omitempty"`
}

// Helper function to convert domain to model
func domainToCourierModel(courier courier.Courier) (*CourierModel, error) {
	// Convert string ObjectId to literal ObjectId
	courierID, err := db.ConvertToObjectId(courier.ID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not valid object id")
	}

	userID, err := db.ConvertToObjectId(courier.UserID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	locationID, err := db.ConvertToObjectId(courier.LocationID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	return &CourierModel{
		ID:         courierID,
		UserID:     &userID,
		Name:       courier.Name,
		Email:      courier.Email,
		Status:     courier.Status.String(), // default not_available
		LocationID: &locationID,
	}, nil
}
