package adapter

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShipmentRepository struct {
	collection *mongo.Collection
}

func NewShipementRepository(db *mongo.Database) *ShipmentRepository {
	return &ShipmentRepository{
		collection: db.Collection("shipments"),
	}
}

func (s *ShipmentRepository) CreateShipment(ctx context.Context, origin string, sender domain.Entity, recipient domain.Entity, items []domain.Item) (string, error) {
	// Convert origin to object id
	locationID, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return "", status.Error(codes.Internal, "origin is not valid object id")
	}

	shipment := ShipmentModel{
		AirwayBill:      generateAWB(12),
		TransportStatus: domain.InPort.String(),
		RoutingStatus:   domain.NotRouted.String(),
		Items:           domainToItemModel(items),
		SenderDetail: EntityDetail{
			Name:        sender.Name,
			Contact:     sender.Contact,
			Province:    sender.Address.Province,
			City:        sender.Address.City,
			District:    sender.Address.District,
			Subdistrict: sender.Address.Subdistrict,
			Address:     sender.Address.StreetAddress,
			ZipCode:     sender.Address.ZipCode,
		},
		RecipientDetail: EntityDetail{
			Name:        recipient.Name,
			Contact:     recipient.Contact,
			Province:    recipient.Address.Province,
			City:        recipient.Address.City,
			District:    recipient.Address.District,
			Subdistrict: recipient.Address.Subdistrict,
			Address:     recipient.Address.StreetAddress,
			ZipCode:     recipient.Address.ZipCode,
		},
		Origin:        locationID,
		ItineraryLogs: []ItineraryLog{},
	}

	result, err := s.collection.InsertOne(ctx, shipment)
	if err != nil {
		return "", cuserr.MongoError(err)
	}

	id := result.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (s *ShipmentRepository) LogItinerary(ctx context.Context, shipmentID string, locationID string, activityType domain.ActivityType) error {
	// Convert string object id to literal object id
	shipmentObjID, err := primitive.ObjectIDFromHex(shipmentID)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	locationObjID, err := primitive.ObjectIDFromHex(locationID)
	if err != nil {
		return status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	// Validate activity type
	if activityType == domain.Unknown {
		return status.Error(codes.InvalidArgument, "activity type is not valid")
	}

	// Define itinerary
	logEntry := ItineraryLog{
		ActivityType: activityType.String(),
		Timestamp:    time.Now(),
		Location:     locationObjID,
	}

	// Push itinerary to itinerary logs
	update := bson.M{"$push": bson.M{"itinerary_logs": logEntry}}
	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": shipmentObjID}, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

func (s *ShipmentRepository) RetrieveShipmentsFromLocations(ctx context.Context, locationsID string, routingStatus domain.RoutingStatus) ([]domain.Shipment, error) {
	// Conver location id to object id
	locationObjID, err := primitive.ObjectIDFromHex(locationsID)
	if err != nil {
		return []domain.Shipment{}, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	// Create pipeline
	pipeline := mongo.Pipeline{
		// Step 1: Filter shipments where the last itinerary log's location matches the given locationObjID
		{
			{Key: "$match", Value: bson.M{
				"routing_status": routingStatus.String(),
				"$expr": bson.M{
					"$eq": bson.A{
						bson.M{"$arrayElemAt": bson.A{"$itinerary_logs.location", -1}}, // Get last location
						locationObjID, // Compare with the given location ID
					},
				},
			}},
		},

		// Step 2: Lookup origin location details from the "locations" collection
		{
			{Key: "$lookup", Value: bson.M{
				"from":         "locations",
				"localField":   "origin",
				"foreignField": "_id",
				"as":           "origin_location",
			}},
		},
		{{Key: "$unwind", Value: bson.M{"path": "$origin_location", "preserveNullAndEmptyArrays": true}}},

		// Step 3: Lookup destination location details from the "locations" collection
		{
			{Key: "$lookup", Value: bson.M{
				"from":         "locations",
				"localField":   "destination",
				"foreignField": "_id",
				"as":           "destination_location",
			}},
		},
		{{Key: "$unwind", Value: bson.M{"path": "$destination_location", "preserveNullAndEmptyArrays": true}}},

		// Step 4: Lookup all location details for each itinerary log's location
		{
			{Key: "$lookup", Value: bson.M{
				"from":         "locations",
				"localField":   "itinerary_logs.location",
				"foreignField": "_id",
				"as":           "location_details",
			}},
		},

		// Step 5: Merge location_details into itinerary_logs
		{
			{Key: "$set", Value: bson.M{
				"itinerary_logs": bson.M{
					"$map": bson.M{
						"input": "$itinerary_logs",
						"as":    "log",
						"in": bson.M{
							"activity_type": "$$log.activity_type",
							"timestamp":     "$$log.timestamp",
							"location":      "$$log.location",
							"location_detail": bson.M{
								"$arrayElemAt": bson.A{
									bson.M{
										"$filter": bson.M{
											"input": "$location_details",
											"as":    "detail",
											"cond":  bson.M{"$eq": bson.A{"$$detail._id", "$$log.location"}},
										},
									},
									0, // Get the first matching location detail
								},
							},
						},
					},
				},
			}},
		},

		// Step 6: Project the final fields to return
		{
			{Key: "$project", Value: bson.M{
				"_id":                  1,
				"airway_bill":          1,
				"transport_status":     1,
				"routing_status":       1,
				"items":                1,
				"sender_detail":        1,
				"recipient_detail":     1,
				"origin":               1,
				"destination":          1,
				"itinerary_logs":       1, // Now contains merged location_detail
				"origin_location":      1,
				"destination_location": 1,
			}},
		},
	}

	cursor, err := s.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}
	defer cursor.Close(ctx)

	var shipments []ShipmentModel
	if err := cursor.All(ctx, &shipments); err != nil {
		return nil, cuserr.MongoError(err)
	}

	return shipmentModelToDomain(shipments), nil
}

// generateAWB generates a unique Airway Bill (AWB) number
func generateAWB(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	randomNumber := r.Int63n(10_000_000_000)

	awbNumber := fmt.Sprintf("%s%0*d", "GP", length-2, randomNumber)
	return awbNumber
}

func shipmentModelToDomain(models []ShipmentModel) []domain.Shipment {
	var shipments []domain.Shipment

	for _, model := range models {
		shipments = append(shipments, domain.Shipment{
			ID:              model.ID.Hex(),
			AirwayBill:      model.AirwayBill,
			TransportStatus: domain.StringToTransportStatus(model.TransportStatus),
			RoutingStatus:   domain.StringToRoutingStatus(model.RoutingStatus),
			Items:           itemModelToDomain(model.Items),
			Sender:          entityModelToDomain(model.SenderDetail),
			Recipient:       entityModelToDomain(model.RecipientDetail),
			Origin:          locationModelToDomain(model.OriginLocation),
			Destination:     locationModelToDomain(model.DestinationLocation),
			ItineraryLogs:   itineraryModelToDomain(model.ItineraryLogs),
		})
	}

	return shipments
}

func itemModelToDomain(items []Item) []domain.Item {
	var domainItems []domain.Item
	for _, item := range items {
		domainItems = append(domainItems, domain.Item{
			Name:   item.Name,
			Amount: item.Amount,
			Weight: item.Weight,
			Volume: item.Volume,
		})
	}
	return domainItems
}

func entityModelToDomain(detail EntityDetail) domain.Entity {
	return domain.Entity{
		Name:    detail.Name,
		Contact: detail.Contact,
		Address: domain.Address{
			Province:      detail.Province,
			City:          detail.City,
			District:      detail.District,
			Subdistrict:   detail.Subdistrict,
			StreetAddress: detail.Address,
			ZipCode:       detail.ZipCode,
		},
	}
}

func locationModelToDomain(location *Location) domain.Location {
	if location == nil {
		return domain.Location{}
	}
	return domain.Location{
		ID:   location.ID.Hex(),
		Name: location.Name,
		Type: location.Type,
	}
}

func itineraryModelToDomain(logs []ItineraryLog) []domain.ItineraryLog {
	var domainLogs []domain.ItineraryLog
	for _, log := range logs {
		domainLogs = append(domainLogs, domain.ItineraryLog{
			ActivityType: domain.StringToActivityType(log.ActivityType),
			Timestamp:    log.Timestamp,
			Location:     locationModelToDomain(log.LocationDetail),
		})
	}
	return domainLogs
}
