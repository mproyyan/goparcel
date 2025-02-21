package adapter

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
)

type LocationService struct {
	client genproto.LocationServiceClient
}

func NewLocationService(client genproto.LocationServiceClient) *LocationService {
	return &LocationService{
		client: client,
	}
}

func (l *LocationService) GetLocation(ctx context.Context, locationID string) (*user.Location, error) {
	location, err := l.client.GetLocation(ctx, &genproto.GetLocationRequest{
		LocationId: locationID,
	})

	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get location using grpc client location service")
	}

	return &user.Location{
		ID:          location.Id,
		Name:        location.Name,
		Type:        location.Type,
		WarehouseID: location.WarehouseId,
		Address: user.Address{
			ZipCode:       location.Address.ZipCode,
			Province:      location.Address.Province,
			City:          location.Address.City,
			District:      location.Address.District,
			Subdistrict:   location.Address.Subdistrict,
			Latitude:      location.Address.Latitude,
			Longitude:     location.Address.Longitude,
			StreetAddress: location.Address.StreetAddress,
		},
	}, nil
}
