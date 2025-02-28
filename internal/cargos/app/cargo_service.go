package app

import (
	"context"

	"github.com/mproyyan/goparcel/internal/cargos/domain"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
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
