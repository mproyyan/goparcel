package adapter

import (
	"context"
	"time"

	"github.com/mproyyan/goparcel/internal/cargos/domain"
	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Models

type CargoModel struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty"`
	Name              string               `bson:"name"`
	Status            string               `bson:"status"`
	MaxCapacity       Capacity             `bson:"max_capacity"`
	CurrentLoad       Capacity             `bson:"current_load"`
	Carriers          []primitive.ObjectID `bson:"carriers"`
	Itineraries       []Itinerary          `bson:"itineraries"`
	Shipments         []primitive.ObjectID `bson:"shipments"`
	LastKnownLocation *primitive.ObjectID  `bson:"last_known_location,omitempty"`
}

type Capacity struct {
	Weight float64 `bson:"weight"`
	Volume float64 `bson:"volume"`
}

type Itinerary struct {
	Location             primitive.ObjectID `bson:"location"`
	EstimatedTimeArrival time.Time          `bson:"estimated_time_arrival"`
	ActualTimeArrival    *time.Time         `bson:"actual_time_arrival,omitempty"`
}

// Implementation

type CargoRepository struct {
	collection *mongo.Collection
}

func NewCargoRepository(db *mongo.Database) *CargoRepository {
	return &CargoRepository{collection: db.Collection("cargos")}
}

func (c *CargoRepository) CreateCargo(ctx context.Context, cargo domain.Cargo) (string, error) {
	cargoModel, err := domainToCargoModel(&cargo)
	if err != nil {
		return "", cuserr.Decorate(err, "failed to convert domain to cargo model")
	}

	result, err := c.collection.InsertOne(ctx, cargoModel)
	if err != nil {
		return "", cuserr.MongoError(err)
	}

	logrus.WithField("cargo_id", result.InsertedID).Info("Cargo created successfully")
	return result.InsertedID.(string), nil
}

func (c *CargoRepository) GetCargos(ctx context.Context, ids []string) ([]*domain.Cargo, error) {
	cargoObjIds := make([]primitive.ObjectID, 0, len(ids))
	for _, id := range ids {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid cargo ID format")
		}

		cargoObjIds = append(cargoObjIds, objId)
	}

	filter := bson.M{}

	// If ids not empty then fetch cargos based on the ids
	if len(ids) > 0 {
		filter["_id"] = bson.M{"$in": cargoObjIds}
	}

	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var cargos []*CargoModel
	if err := cursor.All(ctx, &cargos); err != nil {
		return nil, err
	}

	return cargoModelsToDomain(cargos), nil
}

func (c *CargoRepository) FindMatchingCargos(ctx context.Context, origin, destination string) ([]*domain.Cargo, error) {
	originObjId, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid origin ID format")
	}

	destinationObjId, err := primitive.ObjectIDFromHex(destination)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid destination ID format")
	}

	var cargos []*CargoModel

	// Define the filter to find cargos that have both origin and destination in their itineraries
	filter := bson.M{
		// Ensure that the cargo contains both origin and destination in the itineraries array
		"itineraries": bson.M{
			"$elemMatch": bson.M{
				"location": bson.M{"$in": []primitive.ObjectID{originObjId, destinationObjId}},
			},
		},

		// Ensure that the index of origin in the itineraries array is before the index of destination
		"$expr": bson.M{
			"$lt": []interface{}{
				bson.M{"$indexOfArray": []interface{}{"$itineraries.location", originObjId}},
				bson.M{"$indexOfArray": []interface{}{"$itineraries.location", destinationObjId}},
			},
		},

		// Exclude cargos that have already passed the origin location, but allow cargos at the origin location
		"$or": []bson.M{
			// Include cargos where last_known_location is not set (new cargo)
			{"last_known_location": nil},
			// Include cargos where last_known_location is at or before the origin in the itineraries array
			{"$expr": bson.M{
				"$lte": []interface{}{
					bson.M{"$indexOfArray": []interface{}{"$itineraries.location", "$last_known_location"}},
					bson.M{"$indexOfArray": []interface{}{"$itineraries.location", originObjId}},
				},
			}},
		},
	}

	// Execute the query to find matching cargos
	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode the result into cargos slice
	if err = cursor.All(ctx, &cargos); err != nil {
		return nil, err
	}

	return cargoModelsToDomain(cargos), nil
}

func (c *CargoRepository) LoadShipment(ctx context.Context, cargoId, shipmentId string) error {
	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid shipment ID format")
	}

	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid cargo ID format")
	}

	filter := bson.M{"_id": cargoObjId}
	update := bson.M{"$addToSet": bson.M{"shipments": shipmentObjId}}
	_, err = c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithFields(logrus.Fields{
		"cargo_id":    cargoObjId,
		"shipment_id": shipmentObjId,
	}).Info("Shipment loaded to cargo")

	return nil
}

func (c *CargoRepository) GetCargo(ctx context.Context, id string) (*domain.Cargo, error) {
	cargoObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid cargo ID format")
	}

	var cargo *CargoModel
	filter := bson.M{"_id": cargoObjId}
	err = c.collection.FindOne(ctx, filter).Decode(&cargo)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	return cargoModelToDomain(cargo), nil
}

func (c *CargoRepository) MarkArrival(ctx context.Context, cargoId, locationId string) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid cargo ID format")
	}

	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid location ID format")
	}

	filter := bson.M{
		"_id": cargoObjId,
		"itineraries": bson.M{
			"$elemMatch": bson.M{
				"location":            locationObjId,
				"actual_time_arrival": bson.M{"$exists": false}, // Ensure we only update itineraries that haven't been marked
			},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"last_known_location":               locationObjId,
			"itineraries.$.actual_time_arrival": time.Now(), // Only update the first matching itinerary
		},
	}

	_, err = c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithFields(logrus.Fields{
		"cargo_id":    cargoId,
		"location_id": locationId,
	}).Info("Cargo last location updated")

	return nil
}

func (c *CargoRepository) UnloadShipment(ctx context.Context, cargoId, shipmentId string) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid cargo ID format")
	}

	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid shipment ID format")
	}

	filter := bson.M{"_id": cargoObjId}
	update := bson.M{
		"$pull": bson.M{"shipments": shipmentObjId},
	}

	_, err = c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithFields(logrus.Fields{
		"cargo_id":    cargoObjId,
		"shipment_id": shipmentObjId,
	}).Info("Shipment unloaded from cargo")

	return nil
}

func (c *CargoRepository) AssignCarrier(ctx context.Context, cargoId string, carrierIds []string) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "cargo id is not valid object id")
	}

	var carrierObjIds []primitive.ObjectID
	for _, carrierId := range carrierIds {
		objId, err := primitive.ObjectIDFromHex(carrierId)
		if err != nil {
			return status.Error(codes.InvalidArgument, "carrier id is not valid object id")
		}

		carrierObjIds = append(carrierObjIds, objId)
	}

	filter := bson.M{"_id": cargoObjId}
	update := bson.M{
		"$set": bson.M{
			"carriers": carrierObjIds,
		},
	}

	_, err = c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithFields(logrus.Fields{
		"cargo_id":    cargoObjId,
		"carrier_ids": carrierObjIds,
	}).Info("Carriers replaced for cargo")

	return nil
}

func (c *CargoRepository) AssignRoute(ctx context.Context, cargoId string, itinerary []domain.Itinerary) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "cargo id is not valid object id")
	}

	filter := bson.M{"_id": cargoObjId}
	update := bson.M{
		"$set": bson.M{
			"itineraries": domainToItinerariesModel(itinerary),
		},
	}

	_, err = c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithFields(logrus.Fields{
		"cargo_id": cargoObjId,
	}).Info("Route assigned to cargo")

	return nil
}

func (c *CargoRepository) GetUnroutedCargos(ctx context.Context, locationId string) ([]*domain.Cargo, error) {
	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location id is not valid object id")
	}

	filter := bson.M{
		"last_known_location": locationObjId,
		"$or": []bson.M{
			{"itineraries": bson.M{"$size": 0}},
			{"itineraries": bson.M{"$exists": false}},
			{"itineraries": nil},
		},
	}

	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	defer cursor.Close(ctx)

	var cargos []*CargoModel
	if err := cursor.All(ctx, &cargos); err != nil {
		return nil, cuserr.MongoError(err)
	}

	return cargoModelsToDomain(cargos), nil
}

func (c *CargoRepository) FindCargosWithoutCarrier(ctx context.Context, locationId string) ([]*domain.Cargo, error) {
	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location id is not valid object id")
	}

	filter := bson.M{
		"last_known_location": locationObjId,
		"$or": []bson.M{
			{"carriers": bson.M{"$size": 0}},       // Cargos with empty carriers array
			{"carriers": bson.M{"$exists": false}}, // Cargos with no carriers field
			{"carriers": nil},                      // Cargos with carriers explicitly set to nil
		},
	}

	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	defer cursor.Close(ctx)

	var cargos []*CargoModel
	if err := cursor.All(ctx, &cargos); err != nil {
		return nil, cuserr.MongoError(err)
	}

	return cargoModelsToDomain(cargos), nil
}

func (c *CargoRepository) ResetCompletedCargo(ctx context.Context, cargoId string) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid cargo ID format")
	}

	filter := bson.M{"_id": cargoObjId}
	update := bson.M{
		"$set": bson.M{
			"status":      domain.CargoIdle.String(),
			"carriers":    []primitive.ObjectID{},
			"itineraries": []Itinerary{},
		},
	}

	_, err = c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithField("cargo_id", cargoId).Info("Cargo reset to new state")
	return nil
}

func (c *CargoRepository) UpdateCargoStatus(ctx context.Context, cargoId string, cargoStatus domain.CargoStatus) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid cargo ID format")
	}

	filter := bson.M{"_id": cargoObjId}
	update := bson.M{"$set": bson.M{"status": cargoStatus.String()}}

	_, err = c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithFields(logrus.Fields{
		"cargo_id": cargoId,
		"status":   cargoStatus.String(),
	}).Info("Cargo status updated")

	return nil
}

// Helper function

func domainToCargoModel(cargo *domain.Cargo) (*CargoModel, error) {
	if cargo == nil {
		return nil, status.Error(codes.InvalidArgument, "cargo cannot be nil")
	}

	id, err := db.ConvertToObjectId(cargo.ID)
	if err != nil {
		return nil, cuserr.Decorate(err, "invalid cargo ID format")
	}

	lastKnownLocation, err := db.ConvertToObjectId(cargo.LastKnownLocation)
	if err != nil {
		return nil, cuserr.Decorate(err, "invalid last known location format")
	}

	return &CargoModel{
		ID:                id,
		Name:              cargo.Name,
		Status:            cargo.Status.String(),
		MaxCapacity:       Capacity{Weight: cargo.MaxCapacity.Weight, Volume: cargo.MaxCapacity.Volume},
		CurrentLoad:       Capacity{Weight: cargo.CurrentLoad.Weight, Volume: cargo.CurrentLoad.Volume},
		Carriers:          convertStringSliceToObjectIDSlice(cargo.Carriers),
		Itineraries:       domainToItinerariesModel(cargo.Itineraries),
		Shipments:         convertStringSliceToObjectIDSlice(cargo.Shipments),
		LastKnownLocation: &lastKnownLocation,
	}, nil
}

func cargoModelToDomain(model *CargoModel) *domain.Cargo {
	if model == nil {
		return nil
	}

	return &domain.Cargo{
		ID:                model.ID.Hex(),
		Name:              model.Name,
		Status:            domain.StringToCargoStatus(model.Status),
		MaxCapacity:       domain.Capacity{Weight: model.MaxCapacity.Weight, Volume: model.MaxCapacity.Volume},
		CurrentLoad:       domain.Capacity{Weight: model.CurrentLoad.Weight, Volume: model.CurrentLoad.Volume},
		Carriers:          convertObjectIDSliceToStringSlice(model.Carriers),
		Itineraries:       convertItineraries(model.Itineraries),
		Shipments:         convertObjectIDSliceToStringSlice(model.Shipments),
		LastKnownLocation: db.ObjectIdToString(model.LastKnownLocation),
	}
}

func cargoModelsToDomain(models []*CargoModel) []*domain.Cargo {
	var result []*domain.Cargo
	for _, model := range models {
		result = append(result, cargoModelToDomain(model))
	}
	return result
}

func convertObjectIDSliceToStringSlice(ids []primitive.ObjectID) []string {
	var result []string
	for _, id := range ids {
		result = append(result, id.Hex())
	}
	return result
}

func convertStringSliceToObjectIDSlice(ids []string) []primitive.ObjectID {
	if len(ids) == 0 {
		return []primitive.ObjectID{}
	}

	var result []primitive.ObjectID
	for _, id := range ids {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logrus.WithError(err).Errorf("Invalid ObjectID: %s", id)
			continue
		}

		result = append(result, objId)
	}

	return result
}

func convertItineraries(itineraries []Itinerary) []domain.Itinerary {
	var result []domain.Itinerary
	for _, itinerary := range itineraries {
		result = append(result, domain.Itinerary{
			Location:             itinerary.Location.Hex(),
			EstimatedTimeArrival: itinerary.EstimatedTimeArrival,
			ActualTimeArrival:    itinerary.ActualTimeArrival,
		})
	}
	return result
}

func domainToItineraryModel(itinerary domain.Itinerary) Itinerary {
	location, _ := db.ConvertToObjectId(itinerary.Location)

	return Itinerary{
		Location:             location,
		EstimatedTimeArrival: itinerary.EstimatedTimeArrival,
		ActualTimeArrival:    itinerary.ActualTimeArrival,
	}
}

func domainToItinerariesModel(itineraries []domain.Itinerary) []Itinerary {
	var result []Itinerary
	for _, itinerary := range itineraries {
		result = append(result, domainToItineraryModel(itinerary))
	}

	return result
}
