package port

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/couriers/app"
	"github.com/mproyyan/goparcel/internal/couriers/domain"
)

type GrpcServer struct {
	service app.CourierService
	genproto.UnimplementedCourierServiceServer
}

func NewGrpcServer(service app.CourierService) GrpcServer {
	return GrpcServer{service: service}
}

func (g GrpcServer) GetAvailableCouriers(ctx context.Context, request *genproto.GetAvailableCourierRequest) (*genproto.CourierResponse, error) {
	couriers, err := g.service.AvailableCouriers(ctx, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "AvailableCouriers service failed")
	}

	return &genproto.CourierResponse{Couriers: couriersToProtoResponse(couriers)}, nil
}

func couriersToProtoResponse(couriers []domain.Courier) []*genproto.Courier {
	var protoCouriers []*genproto.Courier
	for _, c := range couriers {
		protoCouriers = append(protoCouriers, &genproto.Courier{
			Id:         c.ID,
			UserID:     c.UserID,
			Name:       c.Name,
			Email:      c.Email,
			Status:     c.Status.String(),
			LocationId: c.LocationID,
		})
	}
	return protoCouriers
}
