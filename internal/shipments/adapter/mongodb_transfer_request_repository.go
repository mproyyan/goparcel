package adapter

import (
	"context"
	"errors"
	"time"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (t *TransferRequestRepository) LatestPendingTransferRequest(ctx context.Context, shipmentId primitive.ObjectID) (*domain.TransferRequest, bool, error) {
	filter := bson.M{
		"shipment_id": shipmentId,
		"status":      domain.StatusPending.String(),
	}

	var transferRequest TransferRequestModel
	opts := options.FindOne().SetSort(bson.M{"created_at": -1})
	err := t.collection.FindOne(ctx, filter, opts).Decode(&transferRequest)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, false, nil
		}
		return nil, false, cuserr.MongoError(err)
	}

	response := transferRequestModelToDomain(transferRequest)
	return &response, true, nil
}

func (t *TransferRequestRepository) CreateTransitRequest(ctx context.Context, shipmentId, origin, destination, courierId, requestedBy primitive.ObjectID) (string, error) {
	transitRequest := TransferRequestModel{
		ID:          primitive.NewObjectID(),
		RequestType: domain.RequestTypeTransit.String(),
		ShipmentID:  shipmentId,
		Origin: Origin{
			Location:    origin,
			RequestedBy: requestedBy,
		},
		Destination: Destination{
			Location: &destination,
		},
		CourierID: &courierId,
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

func (t *TransferRequestRepository) IncomingShipments(ctx context.Context, locationId primitive.ObjectID) ([]*domain.TransferRequest, error) {
	filter := bson.M{
		"status":               domain.StatusPending.String(),
		"destination.location": locationId,
	}

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

func (t *TransferRequestRepository) CompleteTransferRequest(ctx context.Context, requestId primitive.ObjectID) error {
	filter := bson.M{"_id": requestId, "status": domain.StatusPending.String()}
	update := bson.M{"status": domain.StatusCompleted.String()}

	_, err := t.collection.UpdateOne(ctx, filter, update)
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
