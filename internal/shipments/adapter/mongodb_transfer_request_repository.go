package adapter

import (
	"context"
	"errors"
	"time"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Models
type TransferRequestModel struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	RequestType string              `bson:"request_type"`
	ShipmentID  primitive.ObjectID  `bson:"shipment_id"`
	Origin      Origin              `bson:"origin"`
	Destination Destination         `bson:"destination"`
	CourierID   *primitive.ObjectID `bson:"courier_id,omitempty"`
	CargoID     *primitive.ObjectID `bson:"cargo_id,omitempty"`
	Status      string              `bson:"status"`
	CreatedAt   time.Time           `bson:"created_at"`
}

type Origin struct {
	Location    primitive.ObjectID `bson:"location"`
	RequestedBy primitive.ObjectID `bson:"requested_by"`
}

type Destination struct {
	Location        *primitive.ObjectID `bson:"location,omitempty"`
	AcceptedBy      *primitive.ObjectID `bson:"accepted_by,omitempty"`
	RecipientDetail *RecipientDetail    `bson:"recipient_detail,omitempty"`
}

type RecipientDetail struct {
	Name        string `bson:"name"`
	Contact     string `bson:"contact"`
	Province    string `bson:"province"`
	City        string `bson:"city"`
	District    string `bson:"district"`
	Subdistrict string `bson:"subdistrict"`
	Address     string `bson:"address"`
	ZipCode     string `bson:"zip_code"`
}

// Implementation

type TransferRequestRepository struct {
	collection *mongo.Collection
}

func NewTransferRequestRepository(db *mongo.Database) *TransferRequestRepository {
	return &TransferRequestRepository{collection: db.Collection("transfer_requests")}
}

func (t *TransferRequestRepository) LatestPendingTransferRequest(ctx context.Context, shipmentId string) (*domain.TransferRequest, bool, error) {
	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return nil, false, status.Error(codes.InvalidArgument, "invalid shipment ID format")
	}

	filter := bson.M{
		"shipment_id": shipmentObjId,
		"status":      domain.StatusPending.String(),
	}

	logrus.WithField("filter", filter).Debug("Querying latest pending transfer request of shipment")

	var transferRequest TransferRequestModel
	opts := options.FindOne().SetSort(bson.M{"created_at": -1})
	err = t.collection.FindOne(ctx, filter, opts).Decode(&transferRequest)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, false, nil
		}
		return nil, false, cuserr.MongoError(err)
	}

	response := transferRequestModelToDomain(transferRequest)
	return &response, true, nil
}

func (t *TransferRequestRepository) CreateTransitRequest(ctx context.Context, shipmentId, origin, destination, courierId, requestedBy string) (string, error) {
	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return "", status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	originObjId, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return "", status.Error(codes.InvalidArgument, "origin is not valid object id")
	}

	destinationObjId, err := primitive.ObjectIDFromHex(destination)
	if err != nil {
		return "", status.Error(codes.InvalidArgument, "destination is not valid object id")
	}

	courierObjId, err := primitive.ObjectIDFromHex(courierId)
	if err != nil {
		return "", status.Error(codes.InvalidArgument, "courier is not valid object id")
	}

	requestedByObjId, err := primitive.ObjectIDFromHex(requestedBy)
	if err != nil {
		return "", status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	transitRequest := TransferRequestModel{
		ID:          primitive.NewObjectID(),
		RequestType: domain.RequestTypeTransit.String(),
		ShipmentID:  shipmentObjId,
		Origin: Origin{
			Location:    originObjId,
			RequestedBy: requestedByObjId,
		},
		Destination: Destination{
			Location: &destinationObjId,
		},
		CourierID: &courierObjId,
		Status:    domain.StatusPending.String(),
		CreatedAt: time.Now(),
	}

	result, err := t.collection.InsertOne(ctx, transitRequest)
	if err != nil {
		return "", cuserr.MongoError(err)
	}

	insertedID, _ := result.InsertedID.(primitive.ObjectID)
	return insertedID.Hex(), nil
}

func (t *TransferRequestRepository) IncomingShipments(ctx context.Context, locationId string) ([]*domain.TransferRequest, error) {
	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	filter := bson.M{
		"status":               domain.StatusPending.String(),
		"destination.location": locationObjId,
	}

	logrus.WithField("filter", filter).Debug("Querying incoming shipments in a location")

	var shipments []*TransferRequestModel
	cursor, err := t.collection.Find(ctx, filter)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &shipments); err != nil {
		return nil, cuserr.Decorate(err, "failed to decode query result to shipments")
	}

	return transferRequestModelsToDomain(shipments), nil
}

func (t *TransferRequestRepository) CompleteTransferRequest(ctx context.Context, requestId, acceptedBy string) error {
	requestedByObjId, err := primitive.ObjectIDFromHex(requestId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "request_id is not valid object id")
	}

	acceptedByObjId, err := primitive.ObjectIDFromHex(acceptedBy)
	if err != nil {
		return status.Error(codes.InvalidArgument, "accepted_by is not valid object id")
	}

	filter := bson.M{"_id": requestedByObjId, "status": domain.StatusPending.String()}
	update := bson.M{"$set": bson.M{"status": domain.StatusCompleted.String(), "destination.accepted_by": acceptedByObjId}}

	logrus.WithFields(logrus.Fields{
		"filter": filter,
		"update": update,
	}).Debug("Completing transfer request")

	_, err = t.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	logrus.WithFields(logrus.Fields{
		"user_id":    acceptedByObjId,
		"status":     domain.StatusCompleted.String(),
		"request_id": requestedByObjId,
	}).Info("Transfer request status updated")

	return nil
}

func (t *TransferRequestRepository) RequestShipPackage(ctx context.Context, shipmentId, cargoId, origin, destination, requestedBy string) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "cargo_id is not valid object id")
	}

	originObjId, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return status.Error(codes.InvalidArgument, "origin is not valid object id")
	}

	requestedByObjId, err := primitive.ObjectIDFromHex(requestedBy)
	if err != nil {
		return status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	destinationObjId, err := primitive.ObjectIDFromHex(destination)
	if err != nil {
		return status.Error(codes.InvalidArgument, "destination is not valid object id")
	}

	shipmentIdObj, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	transitRequest := TransferRequestModel{
		RequestType: domain.RequestTypeShipment.String(),
		ShipmentID:  shipmentIdObj,
		Origin: Origin{
			Location:    originObjId,
			RequestedBy: requestedByObjId,
		},
		Destination: Destination{
			Location: &destinationObjId,
		},
		CargoID:   &cargoObjId,
		Status:    domain.StatusPending.String(),
		CreatedAt: time.Now(),
	}

	_, err = t.collection.InsertOne(ctx, transitRequest)
	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

func (t *TransferRequestRepository) RequestPackageDelivery(ctx context.Context, origin, shipmentId, courierId, requestedBy string, recipient domain.Entity) error {
	originObjId, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return status.Error(codes.InvalidArgument, "origin is not valid object id")
	}

	requestedByObjId, err := primitive.ObjectIDFromHex(requestedBy)
	if err != nil {
		return status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	courierObjId, err := primitive.ObjectIDFromHex(courierId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "courier id is not valid object id")
	}

	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	transitRequest := TransferRequestModel{
		RequestType: domain.RequestTypeDelivery.String(),
		ShipmentID:  shipmentObjId,
		Origin: Origin{
			Location:    originObjId,
			RequestedBy: requestedByObjId,
		},
		Destination: Destination{
			RecipientDetail: &RecipientDetail{
				Name:        recipient.Name,
				Contact:     recipient.Contact,
				Province:    recipient.Address.Province,
				City:        recipient.Address.City,
				District:    recipient.Address.District,
				Subdistrict: recipient.Address.Subdistrict,
				Address:     recipient.Address.StreetAddress,
				ZipCode:     recipient.Address.ZipCode,
			},
		},
		CourierID: &courierObjId,
		Status:    domain.StatusPending.String(),
		CreatedAt: time.Now(),
	}

	_, err = t.collection.InsertOne(ctx, transitRequest)
	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

// Helper function

func transferRequestModelToDomain(model TransferRequestModel) domain.TransferRequest {
	return domain.TransferRequest{
		ID:          model.ID.Hex(),
		RequestType: domain.ParseRequestType(model.RequestType),
		ShipmentID:  model.ShipmentID.Hex(),
		Origin: domain.Origin{
			Location:    model.Origin.Location.Hex(),
			RequestedBy: model.Origin.RequestedBy.Hex(),
		},
		Destination: domain.Destination{
			Location:        db.ObjectIdToString(model.Destination.Location),
			AcceptedBy:      db.ObjectIdToString(model.Destination.AcceptedBy),
			RecipientDetail: convertRecipientDetail(model.Destination.RecipientDetail),
		},
		CourierID: db.ObjectIdToString(model.CourierID),
		CargoID:   db.ObjectIdToString(model.CargoID),
		Status:    domain.ParseStatus(model.Status),
		CreatedAt: model.CreatedAt,
	}
}

func transferRequestModelsToDomain(models []*TransferRequestModel) []*domain.TransferRequest {
	var requests []*domain.TransferRequest
	for _, model := range models {
		req := transferRequestModelToDomain(*model)
		requests = append(requests, &req)
	}

	return requests
}

func convertRecipientDetail(detail *RecipientDetail) domain.Entity {
	if detail == nil {
		return domain.Entity{}
	}

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
