package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CargoRepository interface {
	GetCargos(ctx context.Context, ids []primitive.ObjectID) ([]*Cargo, error)
	FindMatchingCargos(ctx context.Context, origin, destination primitive.ObjectID) ([]*Cargo, error)
}

type CarrierRepository interface {
	GetCarrier(ctx context.Context, id primitive.ObjectID) (*Carrier, error)
}
