//go:generate mockgen -source=./repository.go -destination=../mock/mock_repository.go -package=mock

package domain

import (
	"context"
)

type LocationRepository interface {
	FindLocation(ctx context.Context, locationID string) (*Location, error)
	CreateLocation(ctx context.Context, location Location) (string, error)
	FindTransitPlaces(ctx context.Context, locationID string) ([]*Location, error)
	GetLocations(ctx context.Context, locationIds []string) ([]*Location, error)
	FindMatchingLocations(ctx context.Context, keyword string) ([]*Location, error)
}
