package app

import (
	"context"

	"github.com/mproyyan/goparcel/internal/cargos/domain"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		"cargo_id": cargoId.Hex(),
	}).Info("Cargo created successfully")

	return nil
}

func (s CargoService) GetCargos(ctx context.Context, ids []string) ([]*domain.Cargo, error) {
	var objIds []primitive.ObjectID
	for _, id := range ids {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "id is not valid object id")
		}

		objIds = append(objIds, objId)
	}

	cargos, err := s.cargoRepository.GetCargos(ctx, objIds)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get cargos from repository")
	}

	return cargos, nil
}

func (c CargoService) FindMatchingCargos(ctx context.Context, origin, destination string) ([]*domain.Cargo, error) {
	originObjId, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "origin is not valid object id")
	}

	destinationObjId, err := primitive.ObjectIDFromHex(destination)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "destination is not valid object id")
	}

	cargos, err := c.cargoRepository.FindMatchingCargos(ctx, originObjId, destinationObjId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to find matching cargos from repository")
	}

	return cargos, nil
}

func (c CargoService) LoadShipment(ctx context.Context, carrierId, locationId, shipmentId string) error {
	carrierObjId, err := primitive.ObjectIDFromHex(carrierId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "carrier id is not valid object id")
	}

	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "location id is not valid object id")
	}

	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment id is not valid object id")
	}

	// Get carrier info
	carrier, err := c.carrierRepository.GetCarrier(ctx, carrierObjId)
	if err != nil {
		return cuserr.Decorate(err, "failed to get carrier")
	}

	cargoObjId, err := primitive.ObjectIDFromHex(carrier.CargoID)
	if err != nil {
		return status.Error(codes.InvalidArgument, "cargo id is not valid object id")
	}

	err = c.cargoRepository.LoadShipment(ctx, cargoObjId, shipmentObjId)
	if err != nil {
		return cuserr.Decorate(err, "failed to load shipment using repository")
	}

	err = c.shipmentService.AddItineraryHistory(ctx, []string{shipmentId}, locationObjId.Hex(), "load")
	if err != nil {
		return cuserr.Decorate(err, "shipment service failed to add itinerary hostory")
	}

	return nil
}

func (c CargoService) MarkArrival(ctx context.Context, cargoId, locationId string) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "cargo id is not valid object id")
	}

	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "location id is not valid object id")
	}

	// Get cargo
	cargo, err := c.cargoRepository.GetCargo(ctx, cargoObjId)
	if err != nil {
		return cuserr.Decorate(err, "failed to get cargo from repository")
	}

	// If cargo has shipment, add itineries for that shipments with current location
	if cargo.HasShipments() {
		err = c.shipmentService.AddItineraryHistory(ctx, cargo.Shipments, locationObjId.Hex(), "arrive")
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
	err = c.cargoRepository.MarkArrival(ctx, cargoObjId, locationObjId)
	if err != nil {
		return cuserr.Decorate(err, "failed to mark arrival with current location")
	}

	return nil
}

func (c CargoService) UnloadShipment(ctx context.Context, cargoId, shipmentId string) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "cargo id is not valid object id")
	}

	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment id is not valid object id")
	}

	err = c.cargoRepository.UnloadShipment(ctx, cargoObjId, shipmentObjId)
	if err != nil {
		return cuserr.Decorate(err, "failed to unload shipment using repository")
	}

	return nil
}

func (c CargoService) AssignCarrier(ctx context.Context, cargoId string, carrierIds []string) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "cargo id is not valid object id")
	}

	var carrierObjIds []primitive.ObjectID
	for _, carrierId := range carrierIds {
		objId, err := primitive.ObjectIDFromHex(carrierId)
		if err != nil {
			return status.Error(codes.InvalidArgument, "carrier id is not valid object id")
		}
		carrierObjIds = append(carrierObjIds, objId)
	}

	err = c.cargoRepository.AssignCarrier(ctx, cargoObjId, carrierObjIds)
	if err != nil {
		return cuserr.Decorate(err, "failed to assign carriers to cargo")
	}

	return nil
}

func (c CargoService) AssignRoute(ctx context.Context, cargoId string, itinerary []domain.Itinerary) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "cargo id is not valid object id")
	}

	err = c.cargoRepository.AssignRoute(ctx, cargoObjId, itinerary)
	if err != nil {
		return cuserr.Decorate(err, "failed to assign route to cargo")
	}

	return nil
}

func (c CargoService) GetUnroutedCargos(ctx context.Context, locationId string) ([]*domain.Cargo, error) {
	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location id is not valid object id")
	}

	cargos, err := c.cargoRepository.GetUnroutedCargos(ctx, locationObjId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get unrouted cargos from repository")
	}

	return cargos, nil
}

func (c CargoService) FindCargosWithoutCarrier(ctx context.Context, locationId string) ([]*domain.Cargo, error) {
	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location id is not valid object id")
	}

	cargos, err := c.cargoRepository.FindCargosWithoutCarrier(ctx, locationObjId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to find cargos without carrier from repository")
	}

	return cargos, nil
}

func (c CargoService) GetIdleCarriers(ctx context.Context, locationId string) ([]*domain.Carrier, error) {
	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location id is not valid object id")
	}

	carriers, err := c.carrierRepository.GetIdleCarriers(ctx, locationObjId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get idle carriers from repository")
	}

	return carriers, nil
}
