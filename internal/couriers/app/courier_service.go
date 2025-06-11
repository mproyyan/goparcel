package app

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/couriers/domain"
)

type CourierService struct {
	courierRepository domain.CourierRepository
}

func NewCourierService(courierRepository domain.CourierRepository) CourierService {
	return CourierService{courierRepository}
}

func (c CourierService) AvailableCouriers(ctx context.Context, locationID string) ([]domain.Courier, error) {
	couriers, err := c.courierRepository.AvailableCouriers(ctx, locationID)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get available couriers")
	}

	return couriers, nil
}
