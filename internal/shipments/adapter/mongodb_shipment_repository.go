package adapter

import (
	"context"
	"time"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"github.com/sirupsen/logrus"
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

func (s *ShipmentRepository) CreateShipment(ctx context.Context, origin string, sender domain.Entity, recipient domain.Entity, items []domain.Item, awb string) (string, error) {
	// Convert origin to object id
	locationID, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return "", status.Error(codes.Internal, "origin is not valid object id")
	}

	shipment := ShipmentModel{
		AirwayBill:      awb,
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
		Origin:        &locationID,
		ItineraryLogs: []ItineraryLog{},
		CreatedAt:     time.Now(),
	}

	result, err := s.collection.InsertOne(ctx, shipment)
	if err != nil {
		return "", cuserr.MongoError(err)
	}

	id := result.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (s *ShipmentRepository) LogItinerary(ctx context.Context, shipmentIds []primitive.ObjectID, locationID primitive.ObjectID, activityType domain.ActivityType) error {
	// Validate activity type
	if activityType == domain.Unknown {
		return status.Error(codes.InvalidArgument, "activity type is not valid")
	}

	// Define itinerary
	logEntry := ItineraryLog{
		ActivityType: activityType.String(),
		Timestamp:    time.Now(),
		Location:     &locationID,
	}

	// Push itinerary to itinerary logs
	filter := bson.M{"_id": bson.M{"$in": shipmentIds}}
	update := bson.M{"$push": bson.M{"itinerary_logs": logEntry}}
	_, err := s.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithFields(logrus.Fields{
		"shipment_ids":    shipmentIds,
		"total_shipments": len(shipmentIds),
		"location_id":     locationID,
	})

	return nil
}

func (s *ShipmentRepository) RetrieveShipmentsFromLocations(ctx context.Context, locationId primitive.ObjectID, routingStatus domain.RoutingStatus) ([]*domain.Shipment, error) {
	// Build query
	query := bson.M{
		"$expr": bson.M{
			"$eq": bson.A{
				// Retrieve the location of the last itinerary log entry
				bson.M{"$arrayElemAt": bson.A{"$itinerary_logs.location", -1}},
				// Compare it with the provided location ID
				locationId,
			},
		},
		// Filter by routing_status to ensure it matches the given parameter
		"routing_status": routingStatus.String(),
	}

	logrus.WithField("query", query).Debug("Retrieving shipments from location")

	if routingStatus == domain.Routed {
		query["$expr"] = bson.M{
			"$and": bson.A{
				bson.M{
					"$eq": bson.A{
						bson.M{"$arrayElemAt": bson.A{"$itinerary_logs.location", -1}},
						locationId,
					},
				},
				bson.M{
					"$in": bson.A{
						bson.M{"$arrayElemAt": bson.A{"$itinerary_logs.activity_type", -1}},
						bson.A{"transit", "unload"},
					},
				},
			},
		}

		logrus.WithField("query", query).Debug("Retrieving routed shipments from location")
	}

	cursor, err := s.collection.Find(ctx, query)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}
	defer cursor.Close(ctx)

	var shipments []*ShipmentModel
	if err := cursor.All(ctx, &shipments); err != nil {
		return nil, cuserr.MongoError(err)
	}

	return shipmentModelsToDomain(shipments), nil
}

func (s *ShipmentRepository) GetShipments(ctx context.Context, ids []primitive.ObjectID) ([]*domain.Shipment, error) {
	filter := bson.M{}

	// If ids not empty then fetch shipments based on the ids
	if len(ids) > 0 {
		filter["_id"] = bson.M{"$in": ids}
	}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var shipments []*ShipmentModel
	if err := cursor.All(ctx, &shipments); err != nil {
		return nil, err
	}

	return shipmentModelsToDomain(shipments), nil
}

func (s *ShipmentRepository) GetShipment(ctx context.Context, id primitive.ObjectID) (*domain.Shipment, error) {
	var shipment ShipmentModel
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&shipment)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	return shipmentModelToDomain(&shipment), nil
}

func (s *ShipmentRepository) UpdateTransportStatus(ctx context.Context, shipmentIds []primitive.ObjectID, status domain.TransportStatus) error {
	filter := bson.M{"_id": bson.M{"$in": shipmentIds}}
	update := bson.M{"$set": bson.M{"transport_status": status.String()}}

	logrus.WithFields(logrus.Fields{
		"filter": filter,
		"update": update,
	}).Debug("Updating transport status")

	_, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithFields(logrus.Fields{
		"shipment_ids":    shipmentIds,
		"total_shipments": len(shipmentIds),
		"status":          status.String(),
	}).Info("Shipment transport status updated")

	return nil
}

func (s *ShipmentRepository) AddShipmentDestination(ctx context.Context, shipmentId, locationId primitive.ObjectID) error {
	filter := bson.M{"_id": shipmentId}
	update := bson.M{"$set": bson.M{"destination": locationId, "routing_status": domain.Routed.String()}}

	logrus.WithFields(logrus.Fields{
		"filter": filter,
		"update": update,
	}).Debug("Routing shipment")

	_, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithFields(logrus.Fields{
		"shipment_id": shipmentId,
		"status":      domain.Routed.String(),
	}).Info("Shipment routing status updated")

	return nil
}

func (s *ShipmentRepository) TrackPackage(ctx context.Context, awb string) ([]*domain.ItineraryLog, error) {
	var shipment ShipmentModel
	err := s.collection.FindOne(ctx, bson.M{"airway_bill": awb}).Decode(&shipment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "shipment not found")
		}

		return nil, cuserr.MongoError(err)
	}

	if len(shipment.ItineraryLogs) == 0 {
		return nil, status.Error(codes.NotFound, "no itinerary logs found for this shipment")
	}

	logrus.WithField("airway_bill", awb).Info("Tracking package")

	var itineraries []*domain.ItineraryLog
	for _, log := range shipment.ItineraryLogs {
		itineraries = append(itineraries, &domain.ItineraryLog{
			ActivityType: domain.StringToActivityType(log.ActivityType),
			Timestamp:    log.Timestamp,
			Location:     convertObjIdToHex(log.Location),
		})
	}

	return itineraries, nil
}

func shipmentModelToDomain(model *ShipmentModel) *domain.Shipment {
	return &domain.Shipment{
		ID:              model.ID.Hex(),
		AirwayBill:      model.AirwayBill,
		TransportStatus: domain.StringToTransportStatus(model.TransportStatus),
		RoutingStatus:   domain.StringToRoutingStatus(model.RoutingStatus),
		Items:           itemModelToDomain(model.Items),
		Sender:          entityModelToDomain(model.SenderDetail),
		Recipient:       entityModelToDomain(model.RecipientDetail),
		Origin:          convertObjIdToHex(model.Origin),
		Destination:     convertObjIdToHex(model.Destination),
		ItineraryLogs:   itineraryModelToDomain(model.ItineraryLogs),
		CreatedAt:       model.CreatedAt,
	}
}

func shipmentModelsToDomain(models []*ShipmentModel) []*domain.Shipment {
	var shipments []*domain.Shipment
	for _, model := range models {
		shipments = append(shipments, shipmentModelToDomain(model))
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

func itineraryModelToDomain(logs []ItineraryLog) []domain.ItineraryLog {
	var domainLogs []domain.ItineraryLog
	for _, log := range logs {
		domainLogs = append(domainLogs, domain.ItineraryLog{
			ActivityType: domain.StringToActivityType(log.ActivityType),
			Timestamp:    log.Timestamp,
			Location:     convertObjIdToHex(log.Location),
		})
	}
	return domainLogs
}
