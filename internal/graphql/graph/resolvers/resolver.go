package resolvers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/graphql/graph/model"
	"github.com/vikstrous/dataloadgen"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// Loaders
	locationLoder *dataloadgen.Loader[string, *model.Location]

	// GRPC Clients
	locationService genproto.LocationServiceClient
	shipmentService genproto.ShipmentServiceClient
	courierService  genproto.CourierServiceClient
}

func NewResolver(
	locationService genproto.LocationServiceClient,
	shipmentService genproto.ShipmentServiceClient,
	courierService genproto.CourierServiceClient,
) *Resolver {
	resolver := &Resolver{
		locationService: locationService,
		shipmentService: shipmentService,
		courierService:  courierService,
	}

	// Initiate dataloader
	resolver.locationLoder = dataloadgen.NewLoader(resolver.loadLocations, dataloadgen.WithWait(time.Millisecond*5))

	return resolver
}

func (r *Resolver) loadLocations(ctx context.Context, keys []string) ([]*model.Location, []error) {
	if len(keys) <= 0 {
		return nil, nil
	}

	log.Println("Fetching locations from RPC service for keys:", keys)

	resp, err := r.locationService.GetLocations(ctx, &genproto.GetLocationsRequest{LocationIds: keys})
	if err != nil {
		log.Println("Error fetching locations:", err)
		return nil, []error{err}
	}

	locationMap := make(map[string]*model.Location)
	for _, loc := range resp.Locations {
		locationMap[loc.Id] = &model.Location{
			ID:        loc.Id,
			Name:      loc.Name,
			Type:      loc.Type,
			Warehouse: &model.Location{ID: loc.WarehouseId},
			Address: &model.LocationAddress{
				Province:      &loc.Address.Province,
				City:          &loc.Address.City,
				District:      &loc.Address.District,
				Subdistrict:   &loc.Address.Subdistrict,
				ZipCode:       &loc.Address.ZipCode,
				Latitude:      &loc.Address.Latitude,
				Longitude:     &loc.Address.Longitude,
				StreetAddress: &loc.Address.StreetAddress,
			},
		}
	}

	results := make([]*model.Location, len(keys))
	errors := make([]error, len(keys))
	for i, key := range keys {
		if location, found := locationMap[key]; found {
			results[i] = location
		} else {
			errors[i] = fmt.Errorf("location not found for ID: %s", key)
		}
	}

	return results, errors
}
