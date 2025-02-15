package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	"github.com/mproyyan/goparcel/internal/users/domain/courier"
	cuserr "github.com/mproyyan/goparcel/internal/users/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		return "", err
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
	ID         primitive.ObjectID
	UserID     primitive.ObjectID
	Name       string
	Email      string
	Status     string
	LocationID primitive.ObjectID
}

// Helper function to convert domain to model
func domainToCourierModel(courier courier.Courier) (*CourierModel, error) {
	// Convert string ObjectId to literal ObjectId
	courierID, err := db.ConvertToObjectId(courier.ID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	userID, err := db.ConvertToObjectId(courier.UserID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	locationID, err := db.ConvertToObjectId(courier.LocationID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	return &CourierModel{
		ID:         courierID,
		UserID:     userID,
		Name:       courier.Name,
		Email:      courier.Email,
		Status:     courier.Status,
		LocationID: locationID,
	}, nil
}
