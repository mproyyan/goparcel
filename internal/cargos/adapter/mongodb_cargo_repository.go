package adapter

import (
	"context"
	"time"

	"github.com/mproyyan/goparcel/internal/cargos/domain"
	"github.com/mproyyan/goparcel/internal/common/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	ActualTimeArrival    time.Time          `bson:"actual_time_arrival"`
}

// Implementation

type CargoRepository struct {
	collection *mongo.Collection
}

func NewCargoRepository(db *mongo.Database) *CargoRepository {
	return &CargoRepository{collection: db.Collection("cargos")}
}

func (c *CargoRepository) FindMatchingCargos(ctx context.Context, origin primitive.ObjectID, destination primitive.ObjectID) ([]*domain.Cargo, error) {
	var cargos []*CargoModel

	// Define the filter to find cargos that have both origin and destination in their itineraries
	filter := bson.M{
		// Ensure that the cargo contains both origin and destination in the itineraries array
		"itineraries": bson.M{
			"$elemMatch": bson.M{
				"location": bson.M{"$in": []primitive.ObjectID{origin, destination}},
			},
		},
		// Ensure that the index of origin in the itineraries array is before the index of destination
		"$expr": bson.M{
			"$lt": []interface{}{
				bson.M{"$indexOfArray": []interface{}{"$itineraries.location", origin}},
				bson.M{"$indexOfArray": []interface{}{"$itineraries.location", destination}},
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
					bson.M{"$indexOfArray": []interface{}{"$itineraries.location", origin}},
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

// Helper function

func cargoModelToDomain(model *CargoModel) *domain.Cargo {
	if model == nil {
		return nil
	}

	return &domain.Cargo{
		ID:                model.ID.Hex(),
		Name:              model.Name,
		Status:            model.Status,
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
