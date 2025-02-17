package app

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/locations/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LocationService struct {
	regionService      RegionService
	locationRepository domain.LocationRepository
}

func NewLocationService(regionServie RegionService, locationRepository domain.LocationRepository) LocationService {
	return LocationService{
		regionService:      regionServie,
		locationRepository: locationRepository,
	}
}

type RegionService interface {
	GetRegion(ctx context.Context, zipcode string) (*domain.Region, error)
}

func (l LocationService) CreateLocation(ctx context.Context, location domain.Location) error {
	// Find region detail by zipcode
	region, err := l.regionService.GetRegion(ctx, location.Address.ZipCode)
	if err != nil {
		return err
	}

	// Fill location address with region data
	location.Address.Province = region.Province
	location.Address.City = region.City
	location.Address.District = region.District
	location.Address.Subdistrict = region.Subdistrict

	// If location type is depot then must have warehouse id
	if location.IsDepot() {
		if location.WarehouseID == "" {
			return status.Error(codes.InvalidArgument, "depot type location must contain warehouse id")
		}
	}

	// If location type is depot then cannot contain another warehouse id
	if location.IsWarehouse() {
		if location.WarehouseID != "" {
			return status.Error(codes.InvalidArgument, "warehouse type location cannot contain another warehouse id")
		}
	}

	// Create new location
	_, err = l.locationRepository.CreateLocation(ctx, location)
	if err != nil {
		return cuserr.Decorate(err, "repository CreateLocation failed")
	}

	return nil
}

func (l LocationService) GetLocation(ctx context.Context, locationID string) (*domain.Location, error) {
	location, err := l.locationRepository.FindLocation(ctx, locationID)
	if err != nil {
		return nil, cuserr.Decorate(err, "repository FindLocation failed")
	}

	return location, nil
}
