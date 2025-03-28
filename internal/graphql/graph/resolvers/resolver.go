package resolvers

import (
	"context"
	"fmt"
	"log"
	"slices"
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
	locationLoader  *dataloadgen.Loader[string, *model.Location]
	entityLoader    *dataloadgen.Loader[string, *model.UserEntity]
	shipmentLoader  *dataloadgen.Loader[string, *model.Shipment]
	userLoader      *dataloadgen.Loader[string, *model.User]
	courierLoader   *dataloadgen.Loader[string, *model.Courier]
	cargoLoader     *dataloadgen.Loader[string, *model.Cargo]
	shipmentsLoader *dataloadgen.Loader[string, []*model.Shipment]
	carriersLoader  *dataloadgen.Loader[string, []*model.Carrier]

	// GRPC Clients
	locationService genproto.LocationServiceClient
	shipmentService genproto.ShipmentServiceClient
	courierService  genproto.CourierServiceClient
	userService     genproto.UserServiceClient
	cargoService    genproto.CargoServiceClient
}

func NewResolver(
	locationService genproto.LocationServiceClient,
	shipmentService genproto.ShipmentServiceClient,
	courierService genproto.CourierServiceClient,
	userService genproto.UserServiceClient,
	cargoService genproto.CargoServiceClient,
) *Resolver {
	resolver := &Resolver{
		locationService: locationService,
		shipmentService: shipmentService,
		courierService:  courierService,
		userService:     userService,
		cargoService:    cargoService,
	}

	// Initiate dataloader
	resolver.locationLoader = dataloadgen.NewLoader(resolver.loadLocations, dataloadgen.WithWait(time.Millisecond*5))
	resolver.entityLoader = dataloadgen.NewLoader(resolver.loadEntity, dataloadgen.WithWait(time.Millisecond*5))
	resolver.shipmentLoader = dataloadgen.NewLoader(resolver.loadShipment, dataloadgen.WithWait(time.Millisecond*5))
	resolver.userLoader = dataloadgen.NewLoader(resolver.loadUser, dataloadgen.WithWait(time.Millisecond*5))
	resolver.courierLoader = dataloadgen.NewLoader(resolver.loadCourier, dataloadgen.WithWait(time.Millisecond*5))
	resolver.shipmentsLoader = dataloadgen.NewLoader(resolver.loadShipments, dataloadgen.WithWait(time.Millisecond*5))
	resolver.carriersLoader = dataloadgen.NewLoader(resolver.loadCarriers, dataloadgen.WithWait(time.Millisecond*5))
	resolver.cargoLoader = dataloadgen.NewLoader(resolver.loadCargo, dataloadgen.WithWait(time.Millisecond*5))

	return resolver
}

func (r *Resolver) loadLocations(ctx context.Context, keys []string) ([]*model.Location, []error) {
	keys = slices.DeleteFunc(keys, func(s string) bool { return s == "" })
	if len(keys) <= 0 {
		result := make([]*model.Location, len(keys))
		return result, nil
	}

	log.Println("Fetching locations from RPC service for keys:", keys)

	resp, err := r.locationService.GetLocations(ctx, &genproto.GetLocationsRequest{LocationIds: keys})
	if err != nil {
		log.Println("Error fetching locations:", err)
		return nil, []error{err}
	}

	locationMap := make(map[string]*model.Location)
	for _, loc := range resp.Locations {
		locationMap[loc.Id] = locationToGraphResponse(loc)
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
	keys = slices.DeleteFunc(keys, func(s string) bool { return s == "" })
	if len(keys) <= 0 {
		result := make([]*model.UserEntity, len(keys))
		return result, nil
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

func (r *Resolver) loadShipment(ctx context.Context, keys []string) ([]*model.Shipment, []error) {
	keys = slices.DeleteFunc(keys, func(s string) bool { return s == "" })
	if len(keys) <= 0 {
		result := make([]*model.Shipment, len(keys))
		return result, nil
	}

	log.Println("Fetching shipment from RPC service for keys:", keys)

	resp, err := r.shipmentService.GetShipments(ctx, &genproto.GetShipmentsRequest{Ids: keys})
	if err != nil {
		log.Println("Error fetching shipments:", err)
		return nil, []error{err}
	}

	shipmentMap := make(map[string]*model.Shipment)
	for _, shipment := range resp.Shipment {
		shipmentMap[shipment.Id] = shipmentToGraphResponse(shipment)
	}

	results := make([]*model.Shipment, len(keys))
	errors := make([]error, len(keys))
	for i, key := range keys {
		if shipment, found := shipmentMap[key]; found {
			results[i] = shipment
		} else {
			errors[i] = fmt.Errorf("shipment not found for ID: %s", key)
		}
	}

	return results, errors
}

func (r *Resolver) loadUser(ctx context.Context, keys []string) ([]*model.User, []error) {
	keys = slices.DeleteFunc(keys, func(s string) bool { return s == "" })
	if len(keys) == 0 {
		result := make([]*model.User, len(keys))
		return result, nil
	}

	log.Println("Fetching users from RPC service for keys:", keys)

	resp, err := r.userService.GetUsers(ctx, &genproto.GetUsersRequest{Id: keys})
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil, []error{err}
	}

	userMap := make(map[string]*model.User)
	for _, user := range resp.Users {
		userMap[user.Id] = userToGraphResponse(user)
	}

	results := make([]*model.User, len(keys))
	errors := make([]error, len(keys))
	for i, key := range keys {
		if user, found := userMap[key]; found {
			results[i] = user
		} else {
			errors[i] = fmt.Errorf("user not found for ID: %s", key)
		}
	}

	return results, errors
}

func (r *Resolver) loadCourier(ctx context.Context, keys []string) ([]*model.Courier, []error) {
	if len(keys) <= 0 {
		return nil, nil
	}

	log.Println("Fetching couriers from RPC service for keys:", keys)

	resp, err := r.userService.GetCouriers(ctx, &genproto.GetCouriersRequest{Ids: keys})
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil, []error{err}
	}

	courierMap := make(map[string]*model.Courier)
	for _, courier := range resp.Couriers {
		courierMap[courier.Id] = courierToGraphResponse(courier)
	}

	results := make([]*model.Courier, len(keys))
	errors := make([]error, len(keys))
	for i, key := range keys {
		if courier, found := courierMap[key]; found {
			results[i] = courier
		} else {
			errors[i] = fmt.Errorf("courier not found for ID: %s", key)
		}
	}

	return results, errors
}

func (r *Resolver) loadShipments(ctx context.Context, keys []string) ([][]*model.Shipment, []error) {
	if len(keys) == 0 {
		return nil, nil
	}

	// Parse the keys into a map of Cargo ID -> Shipment IDs
	cargoShipmentMap := make(map[string][]string)
	var allShipmentIDs []string
	seen := make(map[string]struct{}) // To prevent duplicate shipment ID retrievals

	for _, key := range keys {
		// Example key: "cargo123:shipmentA,shipmentB,shipmentC"
		parts := strings.SplitN(key, ":", 2)
		if len(parts) != 2 {
			log.Println("Invalid key format:", key)
			continue
		}

		cargoID := parts[0]
		shipmentIDs := strings.Split(parts[1], ",") // Extract shipment IDs

		cargoShipmentMap[cargoID] = shipmentIDs

		// Collect unique shipment IDs
		for _, id := range shipmentIDs {
			if _, exists := seen[id]; !exists {
				seen[id] = struct{}{}
				allShipmentIDs = append(allShipmentIDs, id)
			}
		}
	}

	// If no valid shipment IDs, return empty results
	if len(allShipmentIDs) == 0 {
		return nil, nil
	}

	// Fetch shipments from the RPC service
	log.Println("Fetching shipments from RPC service for shipment IDs:", allShipmentIDs)
	resp, err := r.shipmentService.GetShipments(ctx, &genproto.GetShipmentsRequest{Ids: allShipmentIDs})
	if err != nil {
		log.Println("Error fetching shipments:", err)
		return nil, []error{err}
	}

	// Create a map of Shipment ID -> Shipment object
	shipmentMap := make(map[string]*model.Shipment)
	for _, shipment := range resp.Shipment {
		shipmentMap[shipment.Id] = shipmentToGraphResponse(shipment)
	}

	// Prepare the results in the same order as the keys
	results := make([][]*model.Shipment, len(keys))
	for i, key := range keys {
		parts := strings.SplitN(key, ":", 2)
		if len(parts) != 2 {
			continue
		}

		shipmentIDs := strings.Split(parts[1], ",")
		var shipments []*model.Shipment

		for _, shipmentID := range shipmentIDs {
			if shipment, found := shipmentMap[shipmentID]; found {
				shipments = append(shipments, shipment)
			}
		}

		results[i] = shipments // Assign shipments to the corresponding cargo key
	}

	return results, nil
}

func (r *Resolver) loadCarriers(ctx context.Context, keys []string) ([][]*model.Carrier, []error) {
	if len(keys) == 0 {
		return nil, nil
	}

	// Parse the keys into a map of Cargo ID -> Carrier IDs
	cargoCarrierMap := make(map[string][]string)
	var allCarrierIDs []string
	seen := make(map[string]struct{}) // To prevent duplicate carrier ID retrievals

	for _, key := range keys {
		// Example key: "cargo123:CarrierA,CarrierB,CarrierC"
		parts := strings.SplitN(key, ":", 2)
		if len(parts) != 2 {
			log.Println("Invalid key format:", key)
			continue
		}

		cargoID := parts[0]
		carrierIds := strings.Split(parts[1], ",") // Extract carrier IDs

		cargoCarrierMap[cargoID] = carrierIds

		// Collect unique carrier IDs
		for _, id := range carrierIds {
			if _, exists := seen[id]; !exists {
				seen[id] = struct{}{}
				allCarrierIDs = append(allCarrierIDs, id)
			}
		}
	}

	// If no valid carrier IDs, return empty results
	if len(allCarrierIDs) == 0 {
		return nil, nil
	}

	// Fetch carriers from the RPC service
	log.Println("Fetching carriers from RPC service for carrier IDs:", allCarrierIDs)
	resp, err := r.userService.GetCarriers(ctx, &genproto.GetCarriersRequest{Ids: allCarrierIDs})
	if err != nil {
		log.Println("Error fetching carriers:", err)
		return nil, []error{err}
	}

	// Create a map of Carrier ID -> carrier object
	carrierMap := make(map[string]*model.Carrier)
	for _, carrier := range resp.Carriers {
		carrierMap[carrier.Id] = carrierToGraphResponse(carrier)
	}

	// Prepare the results in the same order as the keys
	results := make([][]*model.Carrier, len(keys))
	for i, key := range keys {
		parts := strings.SplitN(key, ":", 2)
		if len(parts) != 2 {
			continue
		}

		carrierIds := strings.Split(parts[1], ",")
		var carriers []*model.Carrier

		for _, carrierId := range carrierIds {
			if carrier, found := carrierMap[carrierId]; found {
				carriers = append(carriers, carrier)
			}
		}

		results[i] = carriers // Assign carrier to the corresponding cargo key
	}

	return results, nil
}

func (r *Resolver) loadCargo(ctx context.Context, keys []string) ([]*model.Cargo, []error) {
	if len(keys) <= 0 {
		return nil, nil
	}

	log.Println("Fetching cargos from RPC service for keys:", keys)

	resp, err := r.cargoService.GetCargos(ctx, &genproto.GetCargosRequest{Ids: keys})
	if err != nil {
		log.Println("Error fetching cargos:", err)
		return nil, []error{err}
	}

	cargoMap := make(map[string]*model.Cargo)
	for _, cargo := range resp.Cargos {
		cargoMap[cargo.Id] = cargoToGraphResponse(cargo)
	}

	results := make([]*model.Cargo, len(keys))
	errors := make([]error, len(keys))
	for i, key := range keys {
		if cargo, found := cargoMap[key]; found {
			results[i] = cargo
		} else {
			errors[i] = fmt.Errorf("cargo not found for ID: %s", key)
		}
	}

	return results, errors
}
