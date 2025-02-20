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

// generateAWB generates a unique Airway Bill (AWB) number
func generateAWB(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	randomNumber := r.Int63n(10_000_000_000)

	awbNumber := fmt.Sprintf("%s%0*d", "GP", length-2, randomNumber)
	return awbNumber
}
