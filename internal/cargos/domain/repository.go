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
	AssignCarrier(ctx context.Context, cargoId primitive.ObjectID, carrierIds []primitive.ObjectID) error
	AssignRoute(ctx context.Context, cargoId primitive.ObjectID, itinerary []Itinerary) error
	GetUnroutedCargos(ctx context.Context, locationId primitive.ObjectID) ([]*Cargo, error)
	FindCargosWithoutCarrier(ctx context.Context, locationId primitive.ObjectID) ([]*Cargo, error)
	ResetCompletedCargo(ctx context.Context, cargoId primitive.ObjectID) error
}

type CarrierRepository interface {
	GetCarrier(ctx context.Context, id primitive.ObjectID) (*Carrier, error)
	GetIdleCarriers(ctx context.Context, locationId primitive.ObjectID) ([]*Carrier, error)
	ClearAssignedCargo(ctx context.Context, carrierIds []primitive.ObjectID) error
}
