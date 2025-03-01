//go:generate mockgen -source shipment_service.go -destination ../mock/mock_external_service.go -package mock

package app

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShipmentService struct {
	transaction               db.TransactionManager
	shipmentRepository        domain.ShipmentRepository
	transferRequestRepository domain.TransferRequestRepository
	locationService           LocationService
	cargoService              CargoService
}

func NewShipmentService(
	transaction db.TransactionManager,
	shipmentRepository domain.ShipmentRepository,
	transferRequestRepository domain.TransferRequestRepository,
	locationService LocationService,
	cargoService CargoService,
) ShipmentService {
	return ShipmentService{
		transaction:               transaction,
		shipmentRepository:        shipmentRepository,
		locationService:           locationService,
		transferRequestRepository: transferRequestRepository,
		cargoService:              cargoService,
	}
}

type LocationService interface {
	ResolveAddress(ctx context.Context, zipcode string) (*domain.Address, error)
}

type CargoService interface {
	UnloadShipment(ctx context.Context, cargoId, shipmentId string) error
}

func (s ShipmentService) CreateShipment(ctx context.Context, origin string, sender, recipient domain.Entity, items []domain.Item) error {
	originObjId, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return status.Error(codes.InvalidArgument, "origin is not valid object id")
	}

	// Retrieve the sender's detailed address
	senderDetailAddress, err := s.locationService.ResolveAddress(ctx, sender.Address.ZipCode)
	if err != nil {
		return cuserr.Decorate(err, "Failed to resolve sender's address")
	}

	// Fill sender address
	sender.Address.Province = senderDetailAddress.Province
	sender.Address.City = senderDetailAddress.City
	sender.Address.District = senderDetailAddress.District
	sender.Address.Subdistrict = senderDetailAddress.Subdistrict

	// Retrieve the recipient's detailed address
	recipientDetailAddress, err := s.locationService.ResolveAddress(ctx, recipient.Address.ZipCode)
	if err != nil {
		return cuserr.Decorate(err, "Failed to resolve recipient's address")
	}

	// Fill recipient address
	recipient.Address.Province = recipientDetailAddress.Province
	recipient.Address.City = recipientDetailAddress.City
	recipient.Address.District = recipientDetailAddress.District
	recipient.Address.Subdistrict = recipientDetailAddress.Subdistrict

	// Start transaction
	err = s.transaction.Execute(ctx, func(ctx context.Context) error {
		// Create shipment
		// TODO: should i validate the items?
		shipmentID, err := s.shipmentRepository.CreateShipment(ctx, origin, sender, recipient, items)
		if err != nil {
			return cuserr.Decorate(err, "Failed to create shipment")
		}

		// Log itinerary with with status receive and location where they inputted
		shipmentObjId, _ := primitive.ObjectIDFromHex(shipmentID)
		err = s.shipmentRepository.LogItinerary(ctx, []primitive.ObjectID{shipmentObjId}, originObjId, domain.Receive)
		if err != nil {
			return cuserr.Decorate(err, "Failed to create itinerary log")
		}

		return nil
	})

	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

func (s ShipmentService) UnroutedShipments(ctx context.Context, locationID string) ([]*domain.Shipment, error) {
	locationObjID, err := primitive.ObjectIDFromHex(locationID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	// Get unrouted shipments
	shipments, err := s.shipmentRepository.RetrieveShipmentsFromLocations(ctx, locationObjID, domain.NotRouted)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to retrieve shipments from location")
	}

	return shipments, nil
}

func (s ShipmentService) RoutedShipments(ctx context.Context, locationId string) ([]*domain.Shipment, error) {
	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	// Get routed shipments
	shipments, err := s.shipmentRepository.RetrieveShipmentsFromLocations(ctx, locationObjId, domain.Routed)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to retrieve shipments from location")
	}

	return shipments, nil
}

func (s ShipmentService) RequestTransit(ctx context.Context, shipmentId, origin, destination, courierId, requestedBy string) error {
	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	originObjId, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return status.Error(codes.InvalidArgument, "origin is not valid object id")
	}

	destinationObjId, err := primitive.ObjectIDFromHex(destination)
	if err != nil {
		return status.Error(codes.InvalidArgument, "destination is not valid object id")
	}

	courierObjId, err := primitive.ObjectIDFromHex(courierId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "courier is not valid object id")
	}

	requestedByObjId, err := primitive.ObjectIDFromHex(requestedBy)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	_, found, err := s.transferRequestRepository.LatestPendingTransferRequest(ctx, shipmentObjId)
	if err != nil {
		return cuserr.Decorate(err, "failed to get latest pending transfer request")
	}

	if found {
		return status.Error(codes.InvalidArgument, "failed to create transit request because this shipment still have pending transfer request")
	}

	_, err = s.transferRequestRepository.CreateTransitRequest(ctx, shipmentObjId, originObjId, destinationObjId, courierObjId, requestedByObjId)
	if err != nil {
		return cuserr.Decorate(err, "failed to create transit request")
	}

	return nil
}

func (s ShipmentService) IncomingShipments(ctx context.Context, locationId string) ([]*domain.TransferRequest, error) {
	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	shipments, err := s.transferRequestRepository.IncomingShipments(ctx, locationObjId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get incoming shipments from repository")
	}

	return shipments, nil
}

func (s ShipmentService) GetShipments(ctx context.Context, ids []string) ([]*domain.Shipment, error) {
	var objIds []primitive.ObjectID
	for _, id := range ids {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "id is not valid object id")
		}

		objIds = append(objIds, objId)
	}

	shipments, err := s.shipmentRepository.GetShipments(ctx, objIds)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get shipments from repository")
	}

	return shipments, nil
}

func (s ShipmentService) ScanArrivingShipment(ctx context.Context, locationId, shipmentId, userId string) error {
	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	request, found, err := s.transferRequestRepository.LatestPendingTransferRequest(ctx, shipmentObjId)
	if err != nil {
		return cuserr.Decorate(err, "failed to get latest pending transfer request")
	}

	if !found {
		return status.Error(codes.InvalidArgument, "shipment doesn't have pending transfer request")
	}

	// Check location, is location where the shipment get scanned same with destination location
	if request.Destination.Location != locationObjId.Hex() {
		return status.Error(codes.InvalidArgument, "the destination of the shipment does not match the current scanned shipment location")
	}

	err = s.transaction.Execute(ctx, func(ctx context.Context) error {
		// Check transfer request type, if 'shipment' then call UnloadShipment rpc
		if request.RequestType == domain.RequestTypeShipment {
			err = s.cargoService.UnloadShipment(ctx, request.CargoID, shipmentObjId.Hex())
			if err != nil {
				return cuserr.Decorate(err, "failed to unload shipment")
			}

			// Update transport status to 'in_port'
			s.shipmentRepository.UpdateTransportStatus(ctx, []primitive.ObjectID{shipmentObjId}, domain.InPort)
		}

		reqObjId, _ := primitive.ObjectIDFromHex(request.ID)
		err := s.transferRequestRepository.CompleteTransferRequest(ctx, reqObjId, userObjId)
		if err != nil {
			return cuserr.Decorate(err, "failed to complete pending transfer request")
		}

		// Check next activity type
		activityType := domain.NextActivityTypeBasedOnRequestType(request.RequestType)
		err = s.shipmentRepository.LogItinerary(ctx, []primitive.ObjectID{shipmentObjId}, locationObjId, activityType)
		if err != nil {
			return cuserr.Decorate(err, "failed to log itinerary")
		}

		return nil
	})

	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

func (s ShipmentService) ShipPackage(ctx context.Context, shipmentId, cargoId, origin, destnation, userId string) error {
	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	cargoObjId, err := primitive.ObjectIDFromHex(cargoId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "cargo_id is not valid object id")
	}

	originObjId, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return status.Error(codes.InvalidArgument, "origin is not valid object id")
	}

	destinationObjId, err := primitive.ObjectIDFromHex(destnation)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	// Cannot create another transfer request if shipment still have pending transfer request
	_, found, err := s.transferRequestRepository.LatestPendingTransferRequest(ctx, shipmentObjId)
	if err != nil {
		return cuserr.Decorate(err, "failed to find latest pending transfer request from repository")
	}

	if found {
		return status.Error(codes.InvalidArgument, "shipment still have pending transfer request")
	}

	// Start transaction
	s.transaction.Execute(ctx, func(ctx context.Context) error {
		// Create new transfer request with shipment type
		err := s.transferRequestRepository.RequestShipPackage(ctx, shipmentObjId, cargoObjId, originObjId, destinationObjId, userObjId)
		if err != nil {
			return cuserr.Decorate(err, "failed to create new transfer request with shipment type")
		}

		// Add shipment destination and update routing status
		err = s.shipmentRepository.AddShipmentDestination(ctx, shipmentObjId, destinationObjId)
		if err != nil {
			return cuserr.Decorate(err, "failed to update routing status to 'routed'")
		}

		return nil
	})

	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

func (s ShipmentService) RecordItinerary(ctx context.Context, shipmentIds []string, locationId string, activity domain.ActivityType) error {
	var objIds []primitive.ObjectID
	for _, id := range shipmentIds {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return status.Error(codes.InvalidArgument, "shipment id is not valid object id")
		}
		objIds = append(objIds, objId)
	}

	locationObjId, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "location id is not valid object id")
	}

	s.transaction.Execute(ctx, func(ctx context.Context) error {
		// if activity is load then update transport status
		if activity == domain.Load {
			err = s.shipmentRepository.UpdateTransportStatus(ctx, objIds, domain.OnBoardCargo)
			if err != nil {
				return cuserr.Decorate(err, "failed to update transport status")
			}
		}

		err = s.shipmentRepository.LogItinerary(ctx, objIds, locationObjId, activity)
		if err != nil {
			return cuserr.Decorate(err, "failed to create log itinerary")
		}

		return nil
	})

	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

func (s ShipmentService) DeliverPackage(ctx context.Context, origin, shipmentId, courierId, userId string) error {
	shipmentObjId, err := primitive.ObjectIDFromHex(shipmentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "shipment_id is not valid object id")
	}

	originObjId, err := primitive.ObjectIDFromHex(origin)
	if err != nil {
		return status.Error(codes.InvalidArgument, "origin is not valid object id")
	}

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	courierObjId, err := primitive.ObjectIDFromHex(courierId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "courier id is not valid object id")
	}

	shipment, err := s.shipmentRepository.GetShipment(ctx, shipmentObjId)
	if err != nil {
		return cuserr.Decorate(err, "failed to find shipment")
	}

	err = s.transaction.Execute(ctx, func(ctx context.Context) error {
		err = s.transferRequestRepository.RequestPackageDelivery(ctx, originObjId, shipmentObjId, courierObjId, userObjId, shipment.Recipient)
		if err != nil {
			return cuserr.Decorate(err, "failed to request package delivery")
		}

		err = s.shipmentRepository.UpdateTransportStatus(ctx, []primitive.ObjectID{shipmentObjId}, domain.OnBoardCourier)
		if err != nil {
			return cuserr.Decorate(err, "failed to update transport status")
		}

		return nil
	})

	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}
