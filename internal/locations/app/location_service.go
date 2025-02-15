package app

import (
	"context"

	"github.com/mproyyan/goparcel/internal/locations/domain"
	cuserr "github.com/mproyyan/goparcel/internal/locations/errors"
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
			return cuserr.ErrDepotType
		}
	}

	// If location type is depot then cannot contain another warehouse id
	if location.IsWarehouse() {
		if location.WarehouseID != "" {
			return cuserr.ErrWarehouseType
		}
	}

	// Create new location
	_, err = l.locationRepository.CreateLocation(ctx, location)
	if err != nil {
		return err // TODO: handle mongo error
	}

	return nil
}

func (l LocationService) GetLocation(ctx context.Context, locationID string) (*domain.Location, error) {
	return l.locationRepository.FindLocation(ctx, locationID)
}
