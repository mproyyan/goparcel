package domain

import "context"

type LocationRepository interface {
	FindLocation(ctx context.Context, locationID string) (*Location, error)
	CreateLocation(ctx context.Context, location Location) (string, error)
}
