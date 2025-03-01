//go:generate mockgen -source repository.go -destination ../mock/mock_shipment_repository.go -package mock

package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShipmentRepository interface {
	GetShipment(ctx context.Context, id primitive.ObjectID) (*Shipment, error)
	GetShipments(ctx context.Context, ids []primitive.ObjectID) ([]*Shipment, error)
	CreateShipment(ctx context.Context, origin string, sender, recipient Entity, items []Item) (string, error)
	LogItinerary(ctx context.Context, shipmentID []primitive.ObjectID, locationIds primitive.ObjectID, activityType ActivityType) error
	RetrieveShipmentsFromLocations(ctx context.Context, locationsID string, routingStatus RoutingStatus) ([]*Shipment, error)
	UpdateTransportStatus(ctx context.Context, shipmentId []primitive.ObjectID, status TransportStatus) error
	AddShipmentDestination(ctx context.Context, shipmentId, locationId primitive.ObjectID) error
}

type TransferRequestRepository interface {
	LatestPendingTransferRequest(ctx context.Context, shipmentId primitive.ObjectID) (*TransferRequest, bool, error)
	CreateTransitRequest(ctx context.Context, shipmentId, origin, destination, courierId, requestedBy primitive.ObjectID) (string, error)
	IncomingShipments(ctx context.Context, locationId primitive.ObjectID) ([]*TransferRequest, error)
	CompleteTransferRequest(ctx context.Context, requestId, acceptedBy primitive.ObjectID) error
	ShipPackage(ctx context.Context, shipmentId, cargoId, origin, destination, requestedBy primitive.ObjectID) error
}
