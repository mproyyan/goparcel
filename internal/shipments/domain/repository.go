package domain

import "context"

type ShipmentRepository interface {
	CreateShipment(ctx context.Context, origin string, sender, recipient Entity, items []Item) (string, error)
	LogItinerary(ctx context.Context, shipmentID, locationID string, activityType ActivityType) error
}
