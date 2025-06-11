//go:generate mockgen -source shipment_service.go -destination ../mock/mock_external_service.go -package mock

package app

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"github.com/sirupsen/logrus"
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

func (s ShipmentService) CreateShipment(ctx context.Context, origin string, sender, recipient domain.Entity, items []domain.Item) (string, error) {
	// Retrieve the sender's detailed address
	logrus.WithField("zip_code", sender.Address.ZipCode).Info("Call location service to resolve address")
	senderDetailAddress, err := s.locationService.ResolveAddress(ctx, sender.Address.ZipCode)
	if err != nil {
		return "", cuserr.Decorate(err, "Failed to resolve sender's address")
	}

	// Fill sender address
	sender.Address.Province = senderDetailAddress.Province
	sender.Address.City = senderDetailAddress.City
	sender.Address.District = senderDetailAddress.District
	sender.Address.Subdistrict = senderDetailAddress.Subdistrict

	// Retrieve the recipient's detailed address
	logrus.WithField("zip_code", sender.Address.ZipCode).Info("Call location service to resolve address")
	recipientDetailAddress, err := s.locationService.ResolveAddress(ctx, recipient.Address.ZipCode)
	if err != nil {
		return "", cuserr.Decorate(err, "Failed to resolve recipient's address")
	}

	// Fill recipient address
	recipient.Address.Province = recipientDetailAddress.Province
	recipient.Address.City = recipientDetailAddress.City
	recipient.Address.District = recipientDetailAddress.District
	recipient.Address.Subdistrict = recipientDetailAddress.Subdistrict

	awb := generateAWB(12) // Generate a unique Airway Bill number

	// Start transaction
	err = s.transaction.Execute(ctx, func(ctx context.Context) error {
		// Create shipment
		// TODO: should i validate the items?
		shipmentID, err := s.shipmentRepository.CreateShipment(ctx, origin, sender, recipient, items, awb)
		if err != nil {
			return cuserr.Decorate(err, "Failed to create shipment")
		}

		// Log itinerary with with status receive and location where they inputted
		err = s.shipmentRepository.LogItinerary(ctx, []string{shipmentID}, origin, domain.Receive)
		if err != nil {
			return cuserr.Decorate(err, "Failed to create itinerary log")
		}

		return nil
	})

	if err != nil {
		return "", cuserr.MongoError(err)
	}

	return awb, nil
}

func (s ShipmentService) UnroutedShipments(ctx context.Context, locationID string) ([]*domain.Shipment, error) {
	// Get unrouted shipments
	shipments, err := s.shipmentRepository.RetrieveShipmentsFromLocations(ctx, locationID, domain.NotRouted)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to retrieve shipments from location")
	}

	return shipments, nil
}

func (s ShipmentService) RoutedShipments(ctx context.Context, locationId string) ([]*domain.Shipment, error) {
	// Get routed shipments
	shipments, err := s.shipmentRepository.RetrieveShipmentsFromLocations(ctx, locationId, domain.Routed)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to retrieve shipments from location")
	}

	return shipments, nil
}

func (s ShipmentService) RequestTransit(ctx context.Context, shipmentId, origin, destination, courierId, requestedBy string) error {
	_, found, err := s.transferRequestRepository.LatestPendingTransferRequest(ctx, shipmentId)
	if err != nil {
		return cuserr.Decorate(err, "failed to get latest pending transfer request")
	}

	if found {
		return status.Error(codes.InvalidArgument, "failed to create transit request because this shipment still have pending transfer request")
	}

	_, err = s.transferRequestRepository.CreateTransitRequest(ctx, shipmentId, origin, destination, courierId, requestedBy)
	if err != nil {
		return cuserr.Decorate(err, "failed to create transit request")
	}

	return nil
}

func (s ShipmentService) IncomingShipments(ctx context.Context, locationId string) ([]*domain.TransferRequest, error) {
	shipments, err := s.transferRequestRepository.IncomingShipments(ctx, locationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get incoming shipments from repository")
	}

	return shipments, nil
}

func (s ShipmentService) GetShipments(ctx context.Context, ids []string) ([]*domain.Shipment, error) {
	shipments, err := s.shipmentRepository.GetShipments(ctx, ids)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get shipments from repository")
	}

	return shipments, nil
}

func (s ShipmentService) ScanArrivingShipment(ctx context.Context, locationId, shipmentId, userId string) error {
	request, found, err := s.transferRequestRepository.LatestPendingTransferRequest(ctx, shipmentId)
	if err != nil {
		return cuserr.Decorate(err, "failed to get latest pending transfer request")
	}

	if !found {
		return status.Error(codes.InvalidArgument, "shipment doesn't have pending transfer request")
	}

	// Check location, is location where the shipment get scanned same with destination location
	if request.Destination.Location != locationId {
		return status.Error(codes.InvalidArgument, "the destination of the shipment does not match the current scanned shipment location")
	}

	err = s.transaction.Execute(ctx, func(ctx context.Context) error {
		// Check transfer request type, if 'shipment' then call UnloadShipment rpc
		if request.RequestType == domain.RequestTypeShipment {
			err = s.cargoService.UnloadShipment(ctx, request.CargoID, shipmentId)
			if err != nil {
				return cuserr.Decorate(err, "failed to unload shipment")
			}

			// Update transport status to 'in_port'
			s.shipmentRepository.UpdateTransportStatus(ctx, []string{shipmentId}, domain.InPort)
		}

		err := s.transferRequestRepository.CompleteTransferRequest(ctx, request.ID, userId)
		if err != nil {
			return cuserr.Decorate(err, "failed to complete pending transfer request")
		}

		// Check next activity type
		activityType := domain.NextActivityTypeBasedOnRequestType(request.RequestType)
		err = s.shipmentRepository.LogItinerary(ctx, []string{shipmentId}, locationId, activityType)
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
	// Cannot create another transfer request if shipment still have pending transfer request
	_, found, err := s.transferRequestRepository.LatestPendingTransferRequest(ctx, shipmentId)
	if err != nil {
		return cuserr.Decorate(err, "failed to find latest pending transfer request from repository")
	}

	if found {
		return status.Error(codes.InvalidArgument, "shipment still have pending transfer request")
	}

	// Start transaction
	err = s.transaction.Execute(ctx, func(ctx context.Context) error {
		// Create new transfer request with shipment type
		err := s.transferRequestRepository.RequestShipPackage(ctx, shipmentId, cargoId, origin, destnation, userId)
		if err != nil {
			return cuserr.Decorate(err, "failed to create new transfer request with shipment type")
		}

		// Add shipment destination and update routing status
		err = s.shipmentRepository.AddShipmentDestination(ctx, shipmentId, destnation)
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
	err := s.transaction.Execute(ctx, func(ctx context.Context) error {
		// if activity is load then update transport status
		if activity == domain.Load {
			err := s.shipmentRepository.UpdateTransportStatus(ctx, shipmentIds, domain.OnBoardCargo)
			if err != nil {
				return cuserr.Decorate(err, "failed to update transport status")
			}
		}

		err := s.shipmentRepository.LogItinerary(ctx, shipmentIds, locationId, activity)
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
	shipment, err := s.shipmentRepository.GetShipment(ctx, shipmentId)
	if err != nil {
		return cuserr.Decorate(err, "failed to find shipment")
	}

	err = s.transaction.Execute(ctx, func(ctx context.Context) error {
		err = s.transferRequestRepository.RequestPackageDelivery(ctx, origin, shipmentId, courierId, userId, shipment.Recipient)
		if err != nil {
			return cuserr.Decorate(err, "failed to request package delivery")
		}

		err = s.shipmentRepository.UpdateTransportStatus(ctx, []string{shipmentId}, domain.OnBoardCourier)
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

func (s ShipmentService) CompleteShipment(ctx context.Context, shipmentId string, userId string) error {
	request, found, err := s.transferRequestRepository.LatestPendingTransferRequest(ctx, shipmentId)
	if err != nil {
		return cuserr.Decorate(err, "failed to find latest pending transfer request")
	}

	if !found {
		return status.Error(codes.NotFound, "shipment doesnt have pending transfer request")
	}

	err = s.transaction.Execute(ctx, func(ctx context.Context) error {
		err = s.transferRequestRepository.CompleteTransferRequest(ctx, request.ID, userId)
		if err != nil {
			return cuserr.Decorate(err, "failed to complete transfer request")
		}

		err = s.shipmentRepository.UpdateTransportStatus(ctx, []string{shipmentId}, domain.Claimed)
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

func (s ShipmentService) TrackPackage(ctx context.Context, awb string) ([]*domain.ItineraryLog, error) {
	if awb == "" {
		return nil, status.Error(codes.InvalidArgument, "awb cannot be empty")
	}

	logrus.WithField("awb", awb).Info("Tracking package with AWB")

	itineraryLogs, err := s.shipmentRepository.TrackPackage(ctx, awb)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to track package")
	}

	return itineraryLogs, nil
}

// generateAWB generates a unique Airway Bill (AWB) number
func generateAWB(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	randomNumber := r.Int63n(10_000_000_000)

	awbNumber := fmt.Sprintf("%s%0*d", "GP", length-2, randomNumber)
	return awbNumber
}
