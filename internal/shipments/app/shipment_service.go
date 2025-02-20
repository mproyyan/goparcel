package app

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
)

type ShipmentService struct {
	transaction        db.TransactionManager
	shipmentRepository domain.ShipmentRepository
	locationService    LocationService
}

type LocationService interface {
	ResolveAddress(ctx context.Context, zipcode string) (*domain.Address, error)
}

func (s ShipmentService) CreateShipment(ctx context.Context, origin string, sender, recipient domain.Entity, items []domain.Item) error {
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
		err = s.shipmentRepository.LogItinerary(ctx, shipmentID, origin, domain.Receive)
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
