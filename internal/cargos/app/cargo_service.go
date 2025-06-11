package app

import (
	"context"

	"github.com/mproyyan/goparcel/internal/cargos/domain"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/sirupsen/logrus"
)

type CargoService struct {
	cargoRepository   domain.CargoRepository
	carrierRepository domain.CarrierRepository
	shipmentService   ShipmentService
}

type ShipmentService interface {
	AddItineraryHistory(ctx context.Context, shipmentIds []string, locationId, activityType string) error
}

func NewCargoService(
	cargoRepo domain.CargoRepository,
	carrierRepository domain.CarrierRepository,
	shipmentService ShipmentService,
) CargoService {
	return CargoService{
		cargoRepository:   cargoRepo,
		carrierRepository: carrierRepository,
		shipmentService:   shipmentService,
	}
}

func (s CargoService) CreateCargo(ctx context.Context, cargo domain.Cargo) error {
	// Create cargo in repository
	cargoId, err := s.cargoRepository.CreateCargo(ctx, cargo)
	if err != nil {
		return cuserr.Decorate(err, "failed to create cargo in repository")
	}

	logrus.WithFields(logrus.Fields{
		"cargo_id": cargoId,
	}).Info("Cargo created successfully")

	return nil
}

func (s CargoService) GetCargos(ctx context.Context, ids []string) ([]*domain.Cargo, error) {
	cargos, err := s.cargoRepository.GetCargos(ctx, ids)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get cargos from repository")
	}

	return cargos, nil
}

func (c CargoService) FindMatchingCargos(ctx context.Context, origin, destination string) ([]*domain.Cargo, error) {
	cargos, err := c.cargoRepository.FindMatchingCargos(ctx, origin, destination)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to find matching cargos from repository")
	}

	return cargos, nil
}

func (c CargoService) LoadShipment(ctx context.Context, carrierId, locationId, shipmentId string) error {
	// Get carrier info
	carrier, err := c.carrierRepository.GetCarrier(ctx, carrierId)
	if err != nil {
		return cuserr.Decorate(err, "failed to get carrier")
	}

	err = c.cargoRepository.LoadShipment(ctx, carrier.CargoID, shipmentId)
	if err != nil {
		return cuserr.Decorate(err, "failed to load shipment using repository")
	}

	err = c.shipmentService.AddItineraryHistory(ctx, []string{shipmentId}, locationId, "load")
	if err != nil {
		return cuserr.Decorate(err, "shipment service failed to add itinerary hostory")
	}

	return nil
}

func (c CargoService) MarkArrival(ctx context.Context, cargoId, locationId string) error {
	// Get cargo
	cargo, err := c.cargoRepository.GetCargo(ctx, cargoId)
	if err != nil {
		return cuserr.Decorate(err, "failed to get cargo from repository")
	}

	// If cargo has shipment, add itineries for that shipments with current location
	if cargo.HasShipments() {
		err = c.shipmentService.AddItineraryHistory(ctx, cargo.Shipments, locationId, "arrive")
		if err != nil {
			return cuserr.Decorate(err, "failed to add itinerary history")
		}

		logrus.WithFields(logrus.Fields{
			"shipments_ids":   cargo.Shipments,
			"total_shipments": len(cargo.Shipments),
			"location_id":     locationId,
		}).Info("Shipments arrive at location")
	}

	// Update last known location of cargo with current location
	err = c.cargoRepository.MarkArrival(ctx, cargoId, locationId)
	if err != nil {
		return cuserr.Decorate(err, "failed to mark arrival with current location")
	}

	err = c.carrierRepository.UpdateLocation(ctx, cargo.Carriers, locationId)
	if err != nil {
		return cuserr.Decorate(err, "failed to update carrier location")
	}

	cargo, _ = c.cargoRepository.GetCargo(ctx, cargoId)
	if cargo.AllItinerariesCompleted() {
		// Reset cargo if all itineraries are completed
		err = c.cargoRepository.ResetCompletedCargo(ctx, cargoId)
		if err != nil {
			return cuserr.Decorate(err, "failed to reset completed cargo")
		}

		// Reset carriers assigned to this cargo
		err = c.carrierRepository.ClearAssignedCargo(ctx, cargo.Carriers)
		if err != nil {
			return cuserr.Decorate(err, "failed to clear assigned cargo for carriers")
		}
	}

	return nil
}

func (c CargoService) UnloadShipment(ctx context.Context, cargoId, shipmentId string) error {
	err := c.cargoRepository.UnloadShipment(ctx, cargoId, shipmentId)
	if err != nil {
		return cuserr.Decorate(err, "failed to unload shipment using repository")
	}

	return nil
}

func (c CargoService) AssignCarrier(ctx context.Context, cargoId string, carrierIds []string) error {
	cargo, err := c.cargoRepository.GetCargo(ctx, cargoId)
	if err != nil {
		return cuserr.Decorate(err, "failed to get cargo from repository")
	}

	err = c.cargoRepository.AssignCarrier(ctx, cargo.ID, carrierIds)
	if err != nil {
		return cuserr.Decorate(err, "failed to assign carriers to cargo")
	}

	err = c.carrierRepository.AssignCargo(ctx, carrierIds, cargoId)
	if err != nil {
		return cuserr.Decorate(err, "failed to assign cargo to carriers")
	}

	err = c.carrierRepository.UpdateCarrierStatus(ctx, carrierIds, domain.CarrierActive)
	if err != nil {
		return cuserr.Decorate(err, "failed to update carrier status to active")
	}

	if cargo.HasItineraries() {
		err = c.cargoRepository.UpdateCargoStatus(ctx, cargo.ID, domain.CargoActive)
		if err != nil {
			return cuserr.Decorate(err, "failed to update cargo status to active")
		}
	}

	return nil
}

func (c CargoService) AssignRoute(ctx context.Context, cargoId string, itinerary []domain.Itinerary) error {
	cargo, err := c.cargoRepository.GetCargo(ctx, cargoId)
	if err != nil {
		return cuserr.Decorate(err, "failed to get cargo from repository")
	}

	err = c.cargoRepository.AssignRoute(ctx, cargo.ID, itinerary)
	if err != nil {
		return cuserr.Decorate(err, "failed to assign route to cargo")
	}

	if cargo.HasCarriers() {
		err = c.cargoRepository.UpdateCargoStatus(ctx, cargo.ID, domain.CargoActive)
		if err != nil {
			return cuserr.Decorate(err, "failed to update cargo status to active")
		}
	}

	return nil
}

func (c CargoService) GetUnroutedCargos(ctx context.Context, locationId string) ([]*domain.Cargo, error) {
	cargos, err := c.cargoRepository.GetUnroutedCargos(ctx, locationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get unrouted cargos from repository")
	}

	return cargos, nil
}

func (c CargoService) FindCargosWithoutCarrier(ctx context.Context, locationId string) ([]*domain.Cargo, error) {
	cargos, err := c.cargoRepository.FindCargosWithoutCarrier(ctx, locationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to find cargos without carrier from repository")
	}

	return cargos, nil
}

func (c CargoService) GetIdleCarriers(ctx context.Context, locationId string) ([]*domain.Carrier, error) {
	carriers, err := c.carrierRepository.GetIdleCarriers(ctx, locationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get idle carriers from repository")
	}

	return carriers, nil
}
