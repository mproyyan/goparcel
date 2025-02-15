package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	"github.com/mproyyan/goparcel/internal/users/domain/carrier"
	cuserr "github.com/mproyyan/goparcel/internal/users/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CarrierRepository struct {
	collection *mongo.Collection
}

func NewCarrierRepository(db *mongo.Database) *CarrierRepository {
	return &CarrierRepository{
		collection: db.Collection("carriers"),
	}
}

func (c *CarrierRepository) CreateCarrier(ctx context.Context, carrier carrier.Carrier) (string, error) {
	// Prepare data to insert
	carrierModel, err := domainToCarrierModel(carrier)
	if err != nil {
		return "", err
	}

	// Create new carrier
	result, err := c.collection.InsertOne(ctx, carrierModel)
	if err != nil {
		return "", cuserr.MongoError(err)
	}

	// Return inserted id
	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID.Hex(), nil
}

// Models
type CarrierModel struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	UserID              primitive.ObjectID `bson:"user_id,omitempty"`
	Name                string             `bson:"name"`
	Email               string             `bson:"email"`
	CargoID             primitive.ObjectID `bson:"cargo_id,omitempty"`
	Status              string             `bson:"status"`
	LastKnownLocationID primitive.ObjectID `bson:"last_known_location_id,omitempty"`
}

// Helper function to convert domain to model
func domainToCarrierModel(carrier carrier.Carrier) (*CarrierModel, error) {
	// Convert string ObjectId to literal ObjectId
	carrierID, err := db.ConvertToObjectId(carrier.ID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	userID, err := db.ConvertToObjectId(carrier.UserID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	cargoID, err := db.ConvertToObjectId(carrier.CargoID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	locationID, err := db.ConvertToObjectId(carrier.LastKnownLocationID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	return &CarrierModel{
		ID:                  carrierID,
		UserID:              userID,
		Name:                carrier.Name,
		Email:               carrier.Email,
		CargoID:             cargoID,
		Status:              carrier.Name,
		LastKnownLocationID: locationID,
	}, nil
}
