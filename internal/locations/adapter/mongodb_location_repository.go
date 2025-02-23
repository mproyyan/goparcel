package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/locations/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	// Execute query
	var model LocationModel
	err = l.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&model)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	// Convert model to domain
	location := locationModelToDomain(model)
	return &location, nil
}

func (l *LocationRepository) CreateLocation(ctx context.Context, location domain.Location) (string, error) {
	model, err := domainToLocationModel(location)
	if err != nil {
		return "", cuserr.Decorate(err, "failed to convert domain to location model")
	}

	// Insert model
	result, err := l.collection.InsertOne(ctx, model)
	if err != nil {
		return "", cuserr.MongoError(err)
	}

	// Ambil ID hasil insert
	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID.Hex(), nil
}

func (l *LocationRepository) FindTransitPlaces(ctx context.Context, locationID primitive.ObjectID) ([]domain.Location, error) {
	// Fetch current location
	location, err := l.FindLocation(ctx, locationID.Hex())
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to find current location")
	}

	var transitPlaces []LocationModel

	// Fetch warehouse and depot that belongs to the warehouse except the current depot location
	if location.IsDepot() {
		// Cconvert warehouse_id to object id
		warehouseObjId, err := primitive.ObjectIDFromHex(location.WarehouseID)
		if err != nil {
			return []domain.Location{}, status.Error(codes.InvalidArgument, "warehouse_id is not valid object id")
		}

		// Create filter
		filter := bson.M{
			"$and": []bson.M{
				{"_id": bson.M{"$ne": locationID}}, // Exclude the current location
				{
					"$or": []bson.M{
						{"_id": warehouseObjId},          // Get warehouse that belong to this depot
						{"warehouse_id": warehouseObjId}, // Get depots in the same warehouse
					},
				},
			},
		}

		cursor, err := l.collection.Find(ctx, filter)
		if err != nil {
			return nil, cuserr.MongoError(err)
		}
		defer cursor.Close(ctx)

		if err := cursor.All(ctx, &transitPlaces); err != nil {
			return nil, cuserr.Decorate(err, "failed to decode warehouse transit places")
		}
	}

	// Fetch depot that belong to this warehouse
	if location.IsWarehouse() {
		filter := bson.M{
			"type":         domain.Depot,
			"warehouse_id": locationID, // Depots belonging to this warehouse
		}

		cursor, err := l.collection.Find(ctx, filter)
		if err != nil {
			return nil, cuserr.Decorate(err, "failed to find depot transit places")
		}
		defer cursor.Close(ctx)

		if err := cursor.All(ctx, &transitPlaces); err != nil {
			return nil, cuserr.Decorate(err, "failed to decode depot transit places")
		}
	}

	return locationsModelToDomain(transitPlaces), nil
}

func (l *LocationRepository) GetLocations(ctx context.Context, locationIds []primitive.ObjectID) ([]domain.Location, error) {
	filter := bson.M{}

	// If location ids not empty then fetch location based on the location ids
	if len(locationIds) > 0 {
		filter["_id"] = bson.M{"$in": locationIds}
	}

	cursor, err := l.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var locationModel []LocationModel
	if err := cursor.All(ctx, &locationModel); err != nil {
		return nil, err
	}

	return locationsModelToDomain(locationModel), nil
}

// Models
type LocationModel struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	Name        string              `bson:"name"`
	Type        string              `bson:"type"`
	WarehouseID *primitive.ObjectID `bson:"warehouse_id,omitempty"`
	Address     Address             `bson:"address"`
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
		WarehouseID: convertObjectIdToHex(model.WarehouseID),
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

func locationsModelToDomain(models []LocationModel) []domain.Location {
	var locations []domain.Location
	for _, model := range models {
		location := locationModelToDomain(model)
		locations = append(locations, location)
	}

	return locations
}

// Helper function to convert domain to model
func domainToLocationModel(location domain.Location) (*LocationModel, error) {
	// Convert string ObjectId to literal ObjectId
	locationID, err := db.ConvertToObjectId(location.ID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	warehouseID, err := db.ConvertToObjectId(location.WarehouseID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "warehouse_id is not valid object id")
	}

	return &LocationModel{
		ID:          locationID,
		Name:        location.Name,
		Type:        location.Type.String(),
		WarehouseID: &warehouseID,
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

func convertObjectIdToHex(objId *primitive.ObjectID) string {
	if objId == nil {
		return ""
	}

	return objId.Hex()
}
