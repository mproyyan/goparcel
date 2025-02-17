package port

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto/locations"
	"github.com/mproyyan/goparcel/internal/locations/app"
	"github.com/mproyyan/goparcel/internal/locations/domain"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	service app.LocationService
	locations.UnimplementedLocationServiceServer
}

func NewGrpcServer(service app.LocationService) GrpcServer {
	return GrpcServer{
		service: service,
	}
}

func (g GrpcServer) GetLocation(ctx context.Context, request *locations.GetLocationRequest) (*locations.Location, error) {
	location, err := g.service.GetLocation(ctx, request.LocationID)
	if err != nil {
		return nil, cuserr.Decorate(err, "service GetLocation failed")
	}

	return &locations.Location{
		Name:        location.Name,
		Type:        location.Type.String(),
		WarehouseId: location.WarehouseID,
		Address: &locations.Address{
			Province:      location.Address.Province,
			City:          location.Address.City,
			District:      location.Address.District,
			Subdistrict:   location.Address.Subdistrict,
			Latitude:      location.Address.Latitude,
			Longitude:     location.Address.Longitude,
			StreetAddress: location.Address.StreetAddress,
			ZipCode:       location.Address.ZipCode,
		},
	}, nil
}

func (g GrpcServer) CreateLocation(ctx context.Context, request *locations.CreateLocationRequest) (*emptypb.Empty, error) {
	err := g.service.CreateLocation(ctx, domain.Location{
		Name:        request.Name,
		Type:        domain.LocationTypeFromString(request.Type),
		WarehouseID: request.WarehouseId,
		Address: domain.Address{
			ZipCode:       request.ZipCode,
			Latitude:      request.Latitude,
			Longitude:     request.Longitude,
			StreetAddress: request.StreetAddress,
		},
	})

	if err != nil {
		return nil, cuserr.Decorate(err, "service CreateLocation failed")
	}

	return &emptypb.Empty{}, nil
}
