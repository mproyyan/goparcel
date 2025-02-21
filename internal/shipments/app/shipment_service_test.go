package app

import (
	"context"
	"errors"
	"testing"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"github.com/mproyyan/goparcel/internal/shipments/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateShipment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationService := mock.NewMockLocationService(ctrl)
	shipmentRepo := mock.NewMockShipmentRepository(ctrl)
	transaction := db.NewMockTransactionManager(ctrl)
	shipmentService := NewShipmentService(transaction, shipmentRepo, locationService)
	ctx := context.Background()

	tests := []struct {
		name          string
		shipment      *domain.Shipment
		setupMock     func(s domain.Shipment)
		expectedError error
	}{
		// Create shipment success
		{
			name:     "Create shipment success",
			shipment: newShipmentDataTest(),
			setupMock: func(s domain.Shipment) {
				locationService.EXPECT().ResolveAddress(ctx, s.Sender.Address.ZipCode).Return(&s.Sender.Address, nil)
				locationService.EXPECT().ResolveAddress(ctx, s.Recipient.Address.ZipCode).Return(&s.Recipient.Address, nil)
				shipmentRepo.EXPECT().CreateShipment(ctx, s.Origin.ID, s.Sender, s.Recipient, s.Items).Return("123", nil)
				shipmentRepo.EXPECT().LogItinerary(ctx, "123", s.Origin.ID, domain.Receive).Return(nil)
				transaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
		},
		// Failed to resolver sender address
		{
			name:     "Failed to resolver sender address",
			shipment: newShipmentDataTest(),
			setupMock: func(s domain.Shipment) {
				locationService.EXPECT().ResolveAddress(ctx, s.Sender.Address.ZipCode).Return(nil, errors.New("region not found"))
			},
			expectedError: cuserr.Decorate(errors.New("region not found"), "Failed to resolve sender's address"),
		},
		// Failed to resolve recipient address
		{
			name:     "Failed to resolve recipient address",
			shipment: newShipmentDataTest(),
			setupMock: func(s domain.Shipment) {
				locationService.EXPECT().ResolveAddress(ctx, s.Sender.Address.ZipCode).Return(&s.Sender.Address, nil)
				locationService.EXPECT().ResolveAddress(ctx, s.Recipient.Address.ZipCode).Return(nil, errors.New("region not found"))
			},
			expectedError: cuserr.Decorate(errors.New("region not found"), "Failed to resolve recipient's address"),
		},
		// Repository create shipment failed
		{
			name:     "Repository create shipment failed",
			shipment: newShipmentDataTest(),
			setupMock: func(s domain.Shipment) {
				locationService.EXPECT().ResolveAddress(ctx, s.Sender.Address.ZipCode).Return(&s.Sender.Address, nil)
				locationService.EXPECT().ResolveAddress(ctx, s.Recipient.Address.ZipCode).Return(&s.Recipient.Address, nil)
				shipmentRepo.EXPECT().CreateShipment(ctx, s.Origin.ID, s.Sender, s.Recipient, s.Items).Return("", errors.New("create shipment failed"))
				transaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
			expectedError: status.Error(codes.Internal, "unexpected database error: Failed to create shipment, cause: create shipment failed"),
		},
		// Failed to log itinerary
		{
			name:     "Failed to log itinerary",
			shipment: newShipmentDataTest(),
			setupMock: func(s domain.Shipment) {
				locationService.EXPECT().ResolveAddress(ctx, s.Sender.Address.ZipCode).Return(&s.Sender.Address, nil)
				locationService.EXPECT().ResolveAddress(ctx, s.Recipient.Address.ZipCode).Return(&s.Recipient.Address, nil)
				shipmentRepo.EXPECT().CreateShipment(ctx, s.Origin.ID, s.Sender, s.Recipient, s.Items).Return("123", nil)
				shipmentRepo.EXPECT().LogItinerary(ctx, "123", s.Origin.ID, domain.Receive).Return(errors.New("failed to log"))
				transaction.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				})
			},
			expectedError: cuserr.MongoError(cuserr.Decorate(errors.New("failed to log"), "Failed to create itinerary log")),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setupMock(*test.shipment)
			err := shipmentService.CreateShipment(ctx, test.shipment.Origin.ID, test.shipment.Sender, test.shipment.Recipient, test.shipment.Items)

			if err != nil {
				assert.Error(t, err)
				assert.EqualError(t, test.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUnroutedShipments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	shipmentRepo := mock.NewMockShipmentRepository(ctrl)
	shipmentService := NewShipmentService(nil, shipmentRepo, nil)
	ctx := context.Background()

	tests := []struct {
		name          string
		shipments     *domain.Shipment
		setupMock     func(s domain.Shipment)
		expectedError error
	}{
		// Get unrouted shipment success
		{
			name:      "Get unrouted shipments success",
			shipments: newShipmentDataTest(),
			setupMock: func(s domain.Shipment) {
				shipmentRepo.EXPECT().RetrieveShipmentsFromLocations(ctx, s.LatestItinerary().Location.ID, domain.NotRouted).Return([]domain.Shipment{s}, nil)
			},
		},
		// Get unrouted shipment failed
		{
			name:      "Get unrouted shipments failed",
			shipments: newShipmentDataTest(),
			setupMock: func(s domain.Shipment) {
				shipmentRepo.EXPECT().RetrieveShipmentsFromLocations(ctx, s.LatestItinerary().Location.ID, domain.NotRouted).Return([]domain.Shipment{}, errors.New("shipment not found"))
			},
			expectedError: cuserr.Decorate(errors.New("shipment not found"), "failed to retrieve shipments from location"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setupMock(*test.shipments)
			s, err := shipmentService.UnroutedShipments(ctx, test.shipments.LatestItinerary().Location.ID)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, test.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, []domain.Shipment{*test.shipments}, s)
			}
		})
	}
}

func newShipmentDataTest() *domain.Shipment {
	return &domain.Shipment{
		Origin: domain.Location{
			ID: "jakarta",
		},
		Sender: domain.Entity{
			Address: domain.Address{
				ZipCode: "11111",
			},
		},
		Recipient: domain.Entity{
			Address: domain.Address{
				ZipCode: "55555",
			},
		},
		ItineraryLogs: []domain.ItineraryLog{
			{
				ActivityType: domain.Receive,
				Location: domain.Location{
					ID: "jakarta",
				},
			},
		},
	}
}
