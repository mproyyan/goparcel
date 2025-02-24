//go:generate mockgen -source=./repository.go -destination=../../mock/mock_carrier.go -package=mock

package carrier

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CarrierRepository interface {
	CreateCarrier(ctx context.Context, carrier Carrier) (string, error)
	GetCarriers(ctx context.Context, ids []primitive.ObjectID) ([]*Carrier, error)
}
