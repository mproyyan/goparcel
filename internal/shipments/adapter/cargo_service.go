package adapter

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto"
)

type CargoService struct {
	client genproto.CargoServiceClient
}

func NewCargoService(client genproto.CargoServiceClient) *CargoService {
	return &CargoService{client: client}
}

func (c *CargoService) UnloadShipment(ctx context.Context, cargoId string, shipmentId string) error {
	_, err := c.client.UnloadShipment(ctx, &genproto.UnloadShipmentRequest{
		CargoId:    cargoId,
		ShipmentId: shipmentId,
	})

	if err != nil {
		return cuserr.Decorate(err, "cargo service failed to unload shipment")
	}

	return nil
}
