package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/genproto"
)

type ShipmentService struct {
	client genproto.ShipmentServiceClient
}

func NewShipmentService(client genproto.ShipmentServiceClient) *ShipmentService {
	return &ShipmentService{client: client}
}

func (s *ShipmentService) AddItineraryHistory(ctx context.Context, shipmentId string, locationId string, activityType string) error {
	_, err := s.client.AddItineraryHistory(ctx, &genproto.AddItineraryHistoryRequest{
		ShipmentIds: []string{shipmentId},
		LocationId:  locationId,
		Activity:    activityType,
	})

	return err
}
