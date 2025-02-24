package resolvers

import (
	"context"
	"fmt"
	"log"
	"strings"
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
	locationLoader *dataloadgen.Loader[string, *model.Location]
	entityLoader   *dataloadgen.Loader[string, *model.UserEntity]

	// GRPC Clients
	locationService genproto.LocationServiceClient
	shipmentService genproto.ShipmentServiceClient
	courierService  genproto.CourierServiceClient
	userService     genproto.UserServiceClient
}

func NewResolver(
	locationService genproto.LocationServiceClient,
	shipmentService genproto.ShipmentServiceClient,
	courierService genproto.CourierServiceClient,
	userService genproto.UserServiceClient,
) *Resolver {
	resolver := &Resolver{
		locationService: locationService,
		shipmentService: shipmentService,
		courierService:  courierService,
		userService:     userService,
	}

	// Initiate dataloader
	resolver.locationLoader = dataloadgen.NewLoader(resolver.loadLocations, dataloadgen.WithWait(time.Millisecond*5))
	resolver.entityLoader = dataloadgen.NewLoader(resolver.loadEntity, dataloadgen.WithWait(time.Millisecond*5))

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

func (r *Resolver) loadEntity(ctx context.Context, keys []string) ([]*model.UserEntity, []error) {
	if len(keys) <= 0 {
		return nil, nil
	}

	log.Println("Fetching entity from RPC service for keys:", keys)

	// 1. Map keys into a map[string][]string (entity -> user_ids)
	entityMap := make(map[string][]string)
	for _, key := range keys {
		parts := strings.Split(key, ":")
		if len(parts) != 2 {
			return nil, []error{fmt.Errorf("invalid key format: %s", key)}
		}

		entity, userID := parts[0], parts[1]
		entityMap[entity] = append(entityMap[entity], userID)
	}

	// 2. Prepare a map to store the results
	resultMap := make(map[string]*model.UserEntity)

	// TODO: use goroutine
	// 3. Call RPCs based on entity
	// Get operators
	if ids, found := entityMap["operator"]; found && len(ids) > 0 {
		result, err := r.userService.GetOperators(ctx, &genproto.GetOperatorsRequest{Ids: entityMap["operator"]})
		if err != nil {
			return nil, []error{err}
		}

		for _, operator := range result.Operators {
			key := fmt.Sprintf("%s:%s", "operator", operator.Id)
			resultMap[key] = &model.UserEntity{
				Name:     operator.Name,
				Email:    operator.Email,
				Location: &model.Location{ID: operator.LocationId},
			}
		}
	}

	// Get couriers
	if ids, found := entityMap["courier"]; found && len(ids) > 0 {
		result, err := r.userService.GetCouriers(ctx, &genproto.GetCouriersRequest{Ids: entityMap["couriers"]})
		if err != nil {
			return nil, []error{err}
		}

		for _, courier := range result.Couriers {
			key := fmt.Sprintf("%s:%s", "courier", courier.Id)
			resultMap[key] = &model.UserEntity{
				Name:     courier.Name,
				Email:    courier.Email,
				Location: &model.Location{ID: courier.LocationId},
			}
		}
	}

	// Get carriers
	if ids, found := entityMap["carrier"]; found && len(ids) > 0 {
		result, err := r.userService.GetCarriers(ctx, &genproto.GetCarriersRequest{Ids: entityMap["carrier"]})
		if err != nil {
			return nil, []error{err}
		}

		for _, carrier := range result.Carriers {
			key := fmt.Sprintf("%s:%s", "carrier", carrier.Id)
			resultMap[key] = &model.UserEntity{
				Name:     carrier.Name,
				Email:    carrier.Email,
				Location: &model.Location{ID: carrier.LocationId},
			}
		}
	}

	// 4. Sort results according to the order of keys
	result := make([]*model.UserEntity, len(keys))
	errors := make([]error, len(keys))
	for i, key := range keys {
		if entity, found := resultMap[key]; found {
			result[i] = entity
		} else {
			errors[i] = fmt.Errorf("entity not found for key %s", key)
		}
	}

	return result, errors
}
