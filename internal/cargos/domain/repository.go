package domain

import (
	"context"
)

type CargoRepository interface {
	GetCargos(ctx context.Context, ids []string) ([]*Cargo, error)
	GetCargo(ctx context.Context, id string) (*Cargo, error)
	FindMatchingCargos(ctx context.Context, origin, destination string) ([]*Cargo, error)
	LoadShipment(ctx context.Context, cargoId, shipmentId string) error
	MarkArrival(ctx context.Context, cargoId, locationId string) error
	UnloadShipment(ctx context.Context, cargoId, shipmentId string) error
	CreateCargo(ctx context.Context, cargo Cargo) (string, error)
	AssignCarrier(ctx context.Context, cargoId string, carrierIds []string) error
	AssignRoute(ctx context.Context, cargoId string, itinerary []Itinerary) error
	GetUnroutedCargos(ctx context.Context, locationId string) ([]*Cargo, error)
	FindCargosWithoutCarrier(ctx context.Context, locationId string) ([]*Cargo, error)
	ResetCompletedCargo(ctx context.Context, cargoId string) error
	UpdateCargoStatus(ctx context.Context, cargoId string, status CargoStatus) error
}

type CarrierRepository interface {
	GetCarrier(ctx context.Context, id string) (*Carrier, error)
	GetIdleCarriers(ctx context.Context, locationId string) ([]*Carrier, error)
	AssignCargo(ctx context.Context, carrierIds []string, cargoId string) error
	ClearAssignedCargo(ctx context.Context, carrierIds []string) error
	UpdateCarrierStatus(ctx context.Context, carrierIds []string, status CarrierStatus) error
	UpdateLocation(ctx context.Context, carrierIds []string, locationId string) error
}
