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
	cargoRepository domain.CargoRepository
}

func NewCargoService(cargoRepo domain.CargoRepository) CargoService {
	return CargoService{cargoRepository: cargoRepo}
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
