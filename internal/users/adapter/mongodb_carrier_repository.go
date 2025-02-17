package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/users/domain/carrier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return "", cuserr.Decorate(err, "failed to convert domain to carrier model")
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
		return nil, status.Error(codes.InvalidArgument, "id is not valid object id")
	}

	userID, err := db.ConvertToObjectId(carrier.UserID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	cargoID, err := db.ConvertToObjectId(carrier.CargoID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "cargo_id is not valid object id")
	}

	locationID, err := db.ConvertToObjectId(carrier.LastKnownLocationID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
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
