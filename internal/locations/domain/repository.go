//go:generate mockgen -source=./repository.go -destination=../mock/mock_repository.go -package=mock

package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationRepository interface {
	FindLocation(ctx context.Context, locationID string) (*Location, error)
	CreateLocation(ctx context.Context, location Location) (string, error)
	FindTransitPlaces(ctx context.Context, locationID primitive.ObjectID) ([]Location, error)
}
