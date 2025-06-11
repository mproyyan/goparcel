//go:generate mockgen -source repository.go -destination ../mock/mock_shipment_repository.go -package mock

package domain

import (
	"context"
)

type ShipmentRepository interface {
	GetShipment(ctx context.Context, id string) (*Shipment, error)
	GetShipments(ctx context.Context, ids []string) ([]*Shipment, error)
	CreateShipment(ctx context.Context, origin string, sender, recipient Entity, items []Item, awb string) (string, error)
	LogItinerary(ctx context.Context, shipmentID []string, locationIds string, activityType ActivityType) error
	RetrieveShipmentsFromLocations(ctx context.Context, locationsID string, routingStatus RoutingStatus) ([]*Shipment, error)
	UpdateTransportStatus(ctx context.Context, shipmentId []string, status TransportStatus) error
	AddShipmentDestination(ctx context.Context, shipmentId, locationId string) error
	TrackPackage(ctx context.Context, awb string) ([]*ItineraryLog, error)
}

type TransferRequestRepository interface {
	LatestPendingTransferRequest(ctx context.Context, shipmentId string) (*TransferRequest, bool, error)
	CreateTransitRequest(ctx context.Context, shipmentId, origin, destination, courierId, requestedBy string) (string, error)
	IncomingShipments(ctx context.Context, locationId string) ([]*TransferRequest, error)
	CompleteTransferRequest(ctx context.Context, requestId, acceptedBy string) error
	RequestShipPackage(ctx context.Context, shipmentId, cargoId, origin, destination, requestedBy string) error
	RequestPackageDelivery(ctx context.Context, origin, shipmentId, courierId, requestedBy string, recipient Entity) error
}
