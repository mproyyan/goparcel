//go:generate mockgen -source repository.go -destination ../mock/mock_shipment_repository.go -package mock

package domain

import "context"

type ShipmentRepository interface {
	CreateShipment(ctx context.Context, origin string, sender, recipient Entity, items []Item) (string, error)
	LogItinerary(ctx context.Context, shipmentID, locationID string, activityType ActivityType) error
	RetrieveShipmentsFromLocations(ctx context.Context, locationsID string, routingStatus RoutingStatus) ([]Shipment, error)
}
