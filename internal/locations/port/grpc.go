package port

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/locations/app"
	"github.com/mproyyan/goparcel/internal/locations/domain"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	service app.LocationService
	genproto.UnimplementedLocationServiceServer
}

func NewGrpcServer(service app.LocationService) GrpcServer {
	return GrpcServer{
		service: service,
	}
}

func (g GrpcServer) GetLocation(ctx context.Context, request *genproto.GetLocationRequest) (*genproto.Location, error) {
	location, err := g.service.GetLocation(ctx, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "service GetLocation failed")
	}

	return &genproto.Location{
		Id:          location.ID,
		Name:        location.Name,
		Type:        location.Type.String(),
		WarehouseId: location.WarehouseID,
		Address: &genproto.Address{
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

func (g GrpcServer) CreateLocation(ctx context.Context, request *genproto.CreateLocationRequest) (*emptypb.Empty, error) {
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

func (g GrpcServer) GetRegion(ctx context.Context, request *genproto.GetRegionRequest) (*genproto.Region, error) {
	region, err := g.service.GetRegion(ctx, request.Zipcode)
	if err != nil {
		return nil, cuserr.Decorate(err, "GetRegion failed")
	}

	return &genproto.Region{
		Province:    region.Province,
		City:        region.City,
		District:    region.District,
		Subdistrict: region.Subdistrict,
		ZipCode:     region.ZipCode,
	}, nil
}

func (g GrpcServer) GetTransitPlaces(ctx context.Context, request *genproto.GetTransitPlacesRequest) (*genproto.LocationResponse, error) {
	locations, err := g.service.TransitPlaces(ctx, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get transit places using location service")
	}

	locationsResponse := locationToProtoResponse(locations)
	return &genproto.LocationResponse{Locations: locationsResponse}, nil
}

func locationToProtoResponse(locations []domain.Location) []*genproto.Location {
	var protoLocations []*genproto.Location

	for _, loc := range locations {
		protoLocations = append(protoLocations, &genproto.Location{
			Id:          loc.ID,
			Name:        loc.Name,
			Type:        loc.Type.String(),
			WarehouseId: loc.WarehouseID,
			Address: &genproto.Address{
				ZipCode:       loc.Address.ZipCode,
				Province:      loc.Address.Province,
				City:          loc.Address.City,
				District:      loc.Address.District,
				Subdistrict:   loc.Address.Subdistrict,
				Latitude:      loc.Address.Latitude,
				Longitude:     loc.Address.Longitude,
				StreetAddress: loc.Address.StreetAddress,
			},
		})
	}

	return protoLocations
}
