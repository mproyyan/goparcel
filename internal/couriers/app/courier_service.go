package app

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/couriers/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CourierService struct {
	courierRepository domain.CourierRepository
}

func NewCourierService(courierRepository domain.CourierRepository) CourierService {
	return CourierService{courierRepository}
}

func (c CourierService) AvailableCouriers(ctx context.Context, locationID string) ([]domain.Courier, error) {
	// Convert string to object id
	locationObjId, err := primitive.ObjectIDFromHex(locationID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}
	couriers, err := c.courierRepository.AvailableCouriers(ctx, locationObjId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get available couriers")
	}

	return couriers, nil
}
