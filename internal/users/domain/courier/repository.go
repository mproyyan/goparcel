//go:generate mockgen -source=./repository.go -destination=../../mock/mock_courier.go -package=mock

package courier

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourierRepository interface {
	CreateCourier(ctx context.Context, courier Courier) (string, error)
	GetCouriers(ctx context.Context, ids []primitive.ObjectID) ([]*Courier, error)
}
