//go:generate mockgen -source repository.go -destination ../mock/mock_repository.go -package mock

package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourierRepository interface {
	AvailableCouriers(ctx context.Context, locationID primitive.ObjectID) ([]Courier, error)
}
