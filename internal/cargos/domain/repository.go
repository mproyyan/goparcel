package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CargoRepository interface {
	GetCargos(ctx context.Context, ids []primitive.ObjectID) ([]*Cargo, error)
	GetCargo(ctx context.Context, id primitive.ObjectID) (*Cargo, error)
	FindMatchingCargos(ctx context.Context, origin, destination primitive.ObjectID) ([]*Cargo, error)
	LoadShipment(ctx context.Context, cargoId, shipmentId primitive.ObjectID) error
	MarkArrival(ctx context.Context, cargoId, locationId primitive.ObjectID) error
	UnloadShipment(ctx context.Context, cargoId, shipmentId primitive.ObjectID) error
	CreateCargo(ctx context.Context, cargo Cargo) (primitive.ObjectID, error)
}

type CarrierRepository interface {
	GetCarrier(ctx context.Context, id primitive.ObjectID) (*Carrier, error)
}
