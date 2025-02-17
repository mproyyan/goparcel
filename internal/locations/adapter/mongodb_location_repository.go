package adapter

import (
	"context"
	"errors"

	"github.com/mproyyan/goparcel/internal/common/db"
	"github.com/mproyyan/goparcel/internal/locations/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationRepository struct {
	collection *mongo.Collection
}

func NewLocationRepository(db *mongo.Database) *LocationRepository {
	return &LocationRepository{
		collection: db.Collection("locations"),
	}
}

func (l *LocationRepository) FindLocation(ctx context.Context, locationID string) (*domain.Location, error) {
	// Convert string ObjectId to literal ObjectId
	objID, err := primitive.ObjectIDFromHex(locationID)
	if err != nil {
		return nil, errors.New("invalid location ID format")
	}

	// Execute query
	var model LocationModel
	err = l.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("location not found")
		}
		return nil, err
	}

	// Convert model to domain
	location := locationModelToDomain(model)
	return &location, nil
}

func (l *LocationRepository) CreateLocation(ctx context.Context, location domain.Location) (string, error) {
	model, err := domainToLocationModel(location)
	if err != nil {
		return "", err
	}

	// Insert model
	result, err := l.collection.InsertOne(ctx, model)
	if err != nil {
		return "", err
	}

	// Ambil ID hasil insert
	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID.Hex(), nil
}

// Models
type LocationModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Type        string             `bson:"type"`
	WarehouseID primitive.ObjectID `bson:"warehouse_id,omitempty"`
	Address     Address            `bson:"address"`
}

type Address struct {
	ZipCode       string  `bson:"zip_code"`
	Province      string  `bson:"province"`
	City          string  `bson:"city"`
	District      string  `bson:"district"`
	Subdistrict   string  `bson:"subdistrict"`
	Latitude      float64 `bson:"latitude"`
	Longitude     float64 `bson:"longitude"`
	StreetAddress string  `bson:"street_address"`
}

func locationModelToDomain(model LocationModel) domain.Location {
	return domain.Location{
		ID:          model.ID.Hex(),
		Name:        model.Name,
		Type:        domain.LocationTypeFromString(model.Type),
		WarehouseID: model.WarehouseID.Hex(),
		Address: domain.Address{
			Province:      model.Address.Province,
			City:          model.Address.City,
			District:      model.Address.District,
			Subdistrict:   model.Address.Subdistrict,
			Latitude:      model.Address.Latitude,
			Longitude:     model.Address.Longitude,
			StreetAddress: model.Address.StreetAddress,
			ZipCode:       model.Address.ZipCode,
		},
	}
}

// Helper function to convert domain to model
func domainToLocationModel(location domain.Location) (*LocationModel, error) {
	// Convert string ObjectId to literal ObjectId
	locationID, err := db.ConvertToObjectId(location.ID)
	if err != nil {
		return nil, err
	}

	warehouseID, err := db.ConvertToObjectId(location.WarehouseID)
	if err != nil {
		return nil, err
	}

	return &LocationModel{
		ID:          locationID,
		Name:        location.Name,
		Type:        location.Type.String(),
		WarehouseID: warehouseID,
		Address: Address{
			ZipCode:       location.Address.ZipCode,
			Province:      location.Address.Province,
			City:          location.Address.City,
			District:      location.Address.District,
			Subdistrict:   location.Address.Subdistrict,
			Latitude:      location.Address.Latitude,
			Longitude:     location.Address.Longitude,
			StreetAddress: location.Address.StreetAddress,
		},
	}, nil
}
