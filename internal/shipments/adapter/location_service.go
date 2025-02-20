package adapter

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto/locations"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
)

type LocationService struct {
	client locations.LocationServiceClient
}

func NewLocationService(client locations.LocationServiceClient) *LocationService {
	return &LocationService{
		client: client,
	}
}

func (l *LocationService) ResolveAddress(ctx context.Context, zipcode string) (*domain.Address, error) {
	region, err := l.client.GetRegion(ctx, &locations.GetRegionRequest{Zipcode: zipcode})
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to call GetRegion")
	}

	return &domain.Address{
		Province:    region.Province,
		City:        region.City,
		District:    region.District,
		Subdistrict: region.Subdistrict,
		ZipCode:     region.ZipCode,
	}, nil
}
