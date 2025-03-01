// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/mproyyan/goparcel/internal/graphql/graph/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Cargo() CargoResolver
	Carrier() CarrierResolver
	Courier() CourierResolver
	Destination() DestinationResolver
	Itinerary() ItineraryResolver
	ItineraryLog() ItineraryLogResolver
	Location() LocationResolver
	Mutation() MutationResolver
	Origin() OriginResolver
	Query() QueryResolver
	Shipment() ShipmentResolver
	TransferRequest() TransferRequestResolver
	User() UserResolver
	UserEntity() UserEntityResolver
}

type DirectiveRoot struct {
	SkipAuth func(ctx context.Context, obj any, next graphql.Resolver) (res any, err error)
}

type ComplexityRoot struct {
	Address struct {
		City          func(childComplexity int) int
		District      func(childComplexity int) int
		Province      func(childComplexity int) int
		StreetAddress func(childComplexity int) int
		Subdistrict   func(childComplexity int) int
		ZipCode       func(childComplexity int) int
	}

	Capacity struct {
		Volume func(childComplexity int) int
		Weight func(childComplexity int) int
	}

	Cargo struct {
		Carriers          func(childComplexity int) int
		CurrentLoad       func(childComplexity int) int
		ID                func(childComplexity int) int
		Itineraries       func(childComplexity int) int
		LastKnownLocation func(childComplexity int) int
		MaxCapacity       func(childComplexity int) int
		Name              func(childComplexity int) int
		Shipments         func(childComplexity int) int
		Status            func(childComplexity int) int
	}

	Carrier struct {
		Email    func(childComplexity int) int
		ID       func(childComplexity int) int
		Location func(childComplexity int) int
		Name     func(childComplexity int) int
		Status   func(childComplexity int) int
		UserID   func(childComplexity int) int
	}

	Courier struct {
		Email    func(childComplexity int) int
		ID       func(childComplexity int) int
		Location func(childComplexity int) int
		Name     func(childComplexity int) int
		Status   func(childComplexity int) int
		UserID   func(childComplexity int) int
	}

	Destination struct {
		AcceptedBy      func(childComplexity int) int
		Location        func(childComplexity int) int
		RecipientDetail func(childComplexity int) int
	}

	Entity struct {
		Address func(childComplexity int) int
		Contact func(childComplexity int) int
		Name    func(childComplexity int) int
	}

	Item struct {
		Amount func(childComplexity int) int
		Name   func(childComplexity int) int
		Volume func(childComplexity int) int
		Weight func(childComplexity int) int
	}

	Itinerary struct {
		ActualTimeArrival    func(childComplexity int) int
		EstimatedTimeArrival func(childComplexity int) int
		Location             func(childComplexity int) int
	}

	ItineraryLog struct {
		ActivityType func(childComplexity int) int
		Location     func(childComplexity int) int
		Timestamp    func(childComplexity int) int
	}

	Location struct {
		Address   func(childComplexity int) int
		ID        func(childComplexity int) int
		Name      func(childComplexity int) int
		Type      func(childComplexity int) int
		Warehouse func(childComplexity int) int
	}

	LocationAddress struct {
		City          func(childComplexity int) int
		District      func(childComplexity int) int
		Latitude      func(childComplexity int) int
		Longitude     func(childComplexity int) int
		Province      func(childComplexity int) int
		StreetAddress func(childComplexity int) int
		Subdistrict   func(childComplexity int) int
		ZipCode       func(childComplexity int) int
	}

	Mutation struct {
		CreateLocation       func(childComplexity int, input *model.CreateLocationInput) int
		CreateShipment       func(childComplexity int, input *model.CreateShipmentInput) int
		LoadShipment         func(childComplexity int, shipmentID string, locationID string) int
		Login                func(childComplexity int, email string, password string) int
		MarkArrival          func(childComplexity int, cargoID string, locationID string) int
		RegisterAsCarrier    func(childComplexity int, input model.RegisterAsCarrierInput) int
		RegisterAsCourier    func(childComplexity int, input model.RegisterAsCourierInput) int
		RegisterAsOperator   func(childComplexity int, input model.RegisterAsOperatorInput) int
		RequestTransit       func(childComplexity int, input *model.RequestTransitInput) int
		ScanArrivingShipment func(childComplexity int, locationID string, shipmentID string) int
		ShipPackage          func(childComplexity int, input model.ShipPackageInput) int
	}

	Origin struct {
		Location    func(childComplexity int) int
		RequestedBy func(childComplexity int) int
	}

	PartyDetail struct {
		Address     func(childComplexity int) int
		City        func(childComplexity int) int
		Contact     func(childComplexity int) int
		District    func(childComplexity int) int
		Name        func(childComplexity int) int
		Province    func(childComplexity int) int
		Subdistrict func(childComplexity int) int
		ZipCode     func(childComplexity int) int
	}

	Query struct {
		GetAvailableCouriers func(childComplexity int, locationID string) int
		GetLocation          func(childComplexity int, id string) int
		GetMatchingCargos    func(childComplexity int, origin string, destination string) int
		GetRegion            func(childComplexity int, zipCode string) int
		GetRoutedShipments   func(childComplexity int, locationID string) int
		GetTransitPlaces     func(childComplexity int, id string) int
		GetUnroutedShipments func(childComplexity int, locationID string) int
		GetUser              func(childComplexity int, id string) int
		IncomingShipments    func(childComplexity int, locationID string) int
		SearchLocations      func(childComplexity int, keyword string) int
	}

	Region struct {
		City        func(childComplexity int) int
		District    func(childComplexity int) int
		Province    func(childComplexity int) int
		Subdistrict func(childComplexity int) int
		ZipCode     func(childComplexity int) int
	}

	Shipment struct {
		AirwayBill      func(childComplexity int) int
		CreatedAt       func(childComplexity int) int
		Destination     func(childComplexity int) int
		ID              func(childComplexity int) int
		Items           func(childComplexity int) int
		ItineraryLogs   func(childComplexity int) int
		Origin          func(childComplexity int) int
		RecipientDetail func(childComplexity int) int
		RoutingStatus   func(childComplexity int) int
		SenderDetail    func(childComplexity int) int
		TransportStatus func(childComplexity int) int
	}

	TransferRequest struct {
		Cargo       func(childComplexity int) int
		Courier     func(childComplexity int) int
		CreatedAt   func(childComplexity int) int
		Destination func(childComplexity int) int
		ID          func(childComplexity int) int
		Origin      func(childComplexity int) int
		RequestType func(childComplexity int) int
		Shipment    func(childComplexity int) int
		Status      func(childComplexity int) int
	}

	User struct {
		Entity  func(childComplexity int) int
		ID      func(childComplexity int) int
		ModelID func(childComplexity int) int
		Type    func(childComplexity int) int
	}

	UserEntity struct {
		Email    func(childComplexity int) int
		Location func(childComplexity int) int
		Name     func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]any) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Address.city":
		if e.complexity.Address.City == nil {
			break
		}

		return e.complexity.Address.City(childComplexity), true

	case "Address.district":
		if e.complexity.Address.District == nil {
			break
		}

		return e.complexity.Address.District(childComplexity), true

	case "Address.province":
		if e.complexity.Address.Province == nil {
			break
		}

		return e.complexity.Address.Province(childComplexity), true

	case "Address.street_address":
		if e.complexity.Address.StreetAddress == nil {
			break
		}

		return e.complexity.Address.StreetAddress(childComplexity), true

	case "Address.subdistrict":
		if e.complexity.Address.Subdistrict == nil {
			break
		}

		return e.complexity.Address.Subdistrict(childComplexity), true

	case "Address.zip_code":
		if e.complexity.Address.ZipCode == nil {
			break
		}

		return e.complexity.Address.ZipCode(childComplexity), true

	case "Capacity.volume":
		if e.complexity.Capacity.Volume == nil {
			break
		}

		return e.complexity.Capacity.Volume(childComplexity), true

	case "Capacity.weight":
		if e.complexity.Capacity.Weight == nil {
			break
		}

		return e.complexity.Capacity.Weight(childComplexity), true

	case "Cargo.carriers":
		if e.complexity.Cargo.Carriers == nil {
			break
		}

		return e.complexity.Cargo.Carriers(childComplexity), true

	case "Cargo.currentLoad":
		if e.complexity.Cargo.CurrentLoad == nil {
			break
		}

		return e.complexity.Cargo.CurrentLoad(childComplexity), true

	case "Cargo.id":
		if e.complexity.Cargo.ID == nil {
			break
		}

		return e.complexity.Cargo.ID(childComplexity), true

	case "Cargo.itineraries":
		if e.complexity.Cargo.Itineraries == nil {
			break
		}

		return e.complexity.Cargo.Itineraries(childComplexity), true

	case "Cargo.last_known_location":
		if e.complexity.Cargo.LastKnownLocation == nil {
			break
		}

		return e.complexity.Cargo.LastKnownLocation(childComplexity), true

	case "Cargo.maxCapacity":
		if e.complexity.Cargo.MaxCapacity == nil {
			break
		}

		return e.complexity.Cargo.MaxCapacity(childComplexity), true

	case "Cargo.name":
		if e.complexity.Cargo.Name == nil {
			break
		}

		return e.complexity.Cargo.Name(childComplexity), true

	case "Cargo.shipments":
		if e.complexity.Cargo.Shipments == nil {
			break
		}

		return e.complexity.Cargo.Shipments(childComplexity), true

	case "Cargo.status":
		if e.complexity.Cargo.Status == nil {
			break
		}

		return e.complexity.Cargo.Status(childComplexity), true

	case "Carrier.email":
		if e.complexity.Carrier.Email == nil {
			break
		}

		return e.complexity.Carrier.Email(childComplexity), true

	case "Carrier.id":
		if e.complexity.Carrier.ID == nil {
			break
		}

		return e.complexity.Carrier.ID(childComplexity), true

	case "Carrier.location":
		if e.complexity.Carrier.Location == nil {
			break
		}

		return e.complexity.Carrier.Location(childComplexity), true

	case "Carrier.name":
		if e.complexity.Carrier.Name == nil {
			break
		}

		return e.complexity.Carrier.Name(childComplexity), true

	case "Carrier.status":
		if e.complexity.Carrier.Status == nil {
			break
		}

		return e.complexity.Carrier.Status(childComplexity), true

	case "Carrier.user_id":
		if e.complexity.Carrier.UserID == nil {
			break
		}

		return e.complexity.Carrier.UserID(childComplexity), true

	case "Courier.email":
		if e.complexity.Courier.Email == nil {
			break
		}

		return e.complexity.Courier.Email(childComplexity), true

	case "Courier.id":
		if e.complexity.Courier.ID == nil {
			break
		}

		return e.complexity.Courier.ID(childComplexity), true

	case "Courier.location":
		if e.complexity.Courier.Location == nil {
			break
		}

		return e.complexity.Courier.Location(childComplexity), true

	case "Courier.name":
		if e.complexity.Courier.Name == nil {
			break
		}

		return e.complexity.Courier.Name(childComplexity), true

	case "Courier.status":
		if e.complexity.Courier.Status == nil {
			break
		}

		return e.complexity.Courier.Status(childComplexity), true

	case "Courier.user_id":
		if e.complexity.Courier.UserID == nil {
			break
		}

		return e.complexity.Courier.UserID(childComplexity), true

	case "Destination.accepted_by":
		if e.complexity.Destination.AcceptedBy == nil {
			break
		}

		return e.complexity.Destination.AcceptedBy(childComplexity), true

	case "Destination.location":
		if e.complexity.Destination.Location == nil {
			break
		}

		return e.complexity.Destination.Location(childComplexity), true

	case "Destination.recipient_detail":
		if e.complexity.Destination.RecipientDetail == nil {
			break
		}

		return e.complexity.Destination.RecipientDetail(childComplexity), true

	case "Entity.address":
		if e.complexity.Entity.Address == nil {
			break
		}

		return e.complexity.Entity.Address(childComplexity), true

	case "Entity.contact":
		if e.complexity.Entity.Contact == nil {
			break
		}

		return e.complexity.Entity.Contact(childComplexity), true

	case "Entity.name":
		if e.complexity.Entity.Name == nil {
			break
		}

		return e.complexity.Entity.Name(childComplexity), true

	case "Item.amount":
		if e.complexity.Item.Amount == nil {
			break
		}

		return e.complexity.Item.Amount(childComplexity), true

	case "Item.name":
		if e.complexity.Item.Name == nil {
			break
		}

		return e.complexity.Item.Name(childComplexity), true

	case "Item.volume":
		if e.complexity.Item.Volume == nil {
			break
		}

		return e.complexity.Item.Volume(childComplexity), true

	case "Item.weight":
		if e.complexity.Item.Weight == nil {
			break
		}

		return e.complexity.Item.Weight(childComplexity), true

	case "Itinerary.actual_time_arrival":
		if e.complexity.Itinerary.ActualTimeArrival == nil {
			break
		}

		return e.complexity.Itinerary.ActualTimeArrival(childComplexity), true

	case "Itinerary.estimated_time_arrival":
		if e.complexity.Itinerary.EstimatedTimeArrival == nil {
			break
		}

		return e.complexity.Itinerary.EstimatedTimeArrival(childComplexity), true

	case "Itinerary.location":
		if e.complexity.Itinerary.Location == nil {
			break
		}

		return e.complexity.Itinerary.Location(childComplexity), true

	case "ItineraryLog.activity_type":
		if e.complexity.ItineraryLog.ActivityType == nil {
			break
		}

		return e.complexity.ItineraryLog.ActivityType(childComplexity), true

	case "ItineraryLog.location":
		if e.complexity.ItineraryLog.Location == nil {
			break
		}

		return e.complexity.ItineraryLog.Location(childComplexity), true

	case "ItineraryLog.timestamp":
		if e.complexity.ItineraryLog.Timestamp == nil {
			break
		}

		return e.complexity.ItineraryLog.Timestamp(childComplexity), true

	case "Location.address":
		if e.complexity.Location.Address == nil {
			break
		}

		return e.complexity.Location.Address(childComplexity), true

	case "Location.id":
		if e.complexity.Location.ID == nil {
			break
		}

		return e.complexity.Location.ID(childComplexity), true

	case "Location.name":
		if e.complexity.Location.Name == nil {
			break
		}

		return e.complexity.Location.Name(childComplexity), true

	case "Location.type":
		if e.complexity.Location.Type == nil {
			break
		}

		return e.complexity.Location.Type(childComplexity), true

	case "Location.warehouse":
		if e.complexity.Location.Warehouse == nil {
			break
		}

		return e.complexity.Location.Warehouse(childComplexity), true

	case "LocationAddress.city":
		if e.complexity.LocationAddress.City == nil {
			break
		}

		return e.complexity.LocationAddress.City(childComplexity), true

	case "LocationAddress.district":
		if e.complexity.LocationAddress.District == nil {
			break
		}

		return e.complexity.LocationAddress.District(childComplexity), true

	case "LocationAddress.latitude":
		if e.complexity.LocationAddress.Latitude == nil {
			break
		}

		return e.complexity.LocationAddress.Latitude(childComplexity), true

	case "LocationAddress.longitude":
		if e.complexity.LocationAddress.Longitude == nil {
			break
		}

		return e.complexity.LocationAddress.Longitude(childComplexity), true

	case "LocationAddress.province":
		if e.complexity.LocationAddress.Province == nil {
			break
		}

		return e.complexity.LocationAddress.Province(childComplexity), true

	case "LocationAddress.street_address":
		if e.complexity.LocationAddress.StreetAddress == nil {
			break
		}

		return e.complexity.LocationAddress.StreetAddress(childComplexity), true

	case "LocationAddress.subdistrict":
		if e.complexity.LocationAddress.Subdistrict == nil {
			break
		}

		return e.complexity.LocationAddress.Subdistrict(childComplexity), true

	case "LocationAddress.zip_code":
		if e.complexity.LocationAddress.ZipCode == nil {
			break
		}

		return e.complexity.LocationAddress.ZipCode(childComplexity), true

	case "Mutation.CreateLocation":
		if e.complexity.Mutation.CreateLocation == nil {
			break
		}

		args, err := ec.field_Mutation_CreateLocation_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateLocation(childComplexity, args["input"].(*model.CreateLocationInput)), true

	case "Mutation.CreateShipment":
		if e.complexity.Mutation.CreateShipment == nil {
			break
		}

		args, err := ec.field_Mutation_CreateShipment_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateShipment(childComplexity, args["input"].(*model.CreateShipmentInput)), true

	case "Mutation.LoadShipment":
		if e.complexity.Mutation.LoadShipment == nil {
			break
		}

		args, err := ec.field_Mutation_LoadShipment_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.LoadShipment(childComplexity, args["shipment_id"].(string), args["location_id"].(string)), true

	case "Mutation.Login":
		if e.complexity.Mutation.Login == nil {
			break
		}

		args, err := ec.field_Mutation_Login_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.Login(childComplexity, args["email"].(string), args["password"].(string)), true

	case "Mutation.MarkArrival":
		if e.complexity.Mutation.MarkArrival == nil {
			break
		}

		args, err := ec.field_Mutation_MarkArrival_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.MarkArrival(childComplexity, args["cargo_id"].(string), args["location_id"].(string)), true

	case "Mutation.RegisterAsCarrier":
		if e.complexity.Mutation.RegisterAsCarrier == nil {
			break
		}

		args, err := ec.field_Mutation_RegisterAsCarrier_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RegisterAsCarrier(childComplexity, args["input"].(model.RegisterAsCarrierInput)), true

	case "Mutation.RegisterAsCourier":
		if e.complexity.Mutation.RegisterAsCourier == nil {
			break
		}

		args, err := ec.field_Mutation_RegisterAsCourier_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RegisterAsCourier(childComplexity, args["input"].(model.RegisterAsCourierInput)), true

	case "Mutation.RegisterAsOperator":
		if e.complexity.Mutation.RegisterAsOperator == nil {
			break
		}

		args, err := ec.field_Mutation_RegisterAsOperator_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RegisterAsOperator(childComplexity, args["input"].(model.RegisterAsOperatorInput)), true

	case "Mutation.RequestTransit":
		if e.complexity.Mutation.RequestTransit == nil {
			break
		}

		args, err := ec.field_Mutation_RequestTransit_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RequestTransit(childComplexity, args["input"].(*model.RequestTransitInput)), true

	case "Mutation.ScanArrivingShipment":
		if e.complexity.Mutation.ScanArrivingShipment == nil {
			break
		}

		args, err := ec.field_Mutation_ScanArrivingShipment_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.ScanArrivingShipment(childComplexity, args["location_id"].(string), args["shipment_id"].(string)), true

	case "Mutation.ShipPackage":
		if e.complexity.Mutation.ShipPackage == nil {
			break
		}

		args, err := ec.field_Mutation_ShipPackage_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.ShipPackage(childComplexity, args["input"].(model.ShipPackageInput)), true

	case "Origin.location":
		if e.complexity.Origin.Location == nil {
			break
		}

		return e.complexity.Origin.Location(childComplexity), true

	case "Origin.requested_by":
		if e.complexity.Origin.RequestedBy == nil {
			break
		}

		return e.complexity.Origin.RequestedBy(childComplexity), true

	case "PartyDetail.address":
		if e.complexity.PartyDetail.Address == nil {
			break
		}

		return e.complexity.PartyDetail.Address(childComplexity), true

	case "PartyDetail.city":
		if e.complexity.PartyDetail.City == nil {
			break
		}

		return e.complexity.PartyDetail.City(childComplexity), true

	case "PartyDetail.contact":
		if e.complexity.PartyDetail.Contact == nil {
			break
		}

		return e.complexity.PartyDetail.Contact(childComplexity), true

	case "PartyDetail.district":
		if e.complexity.PartyDetail.District == nil {
			break
		}

		return e.complexity.PartyDetail.District(childComplexity), true

	case "PartyDetail.name":
		if e.complexity.PartyDetail.Name == nil {
			break
		}

		return e.complexity.PartyDetail.Name(childComplexity), true

	case "PartyDetail.province":
		if e.complexity.PartyDetail.Province == nil {
			break
		}

		return e.complexity.PartyDetail.Province(childComplexity), true

	case "PartyDetail.subdistrict":
		if e.complexity.PartyDetail.Subdistrict == nil {
			break
		}

		return e.complexity.PartyDetail.Subdistrict(childComplexity), true

	case "PartyDetail.zip_code":
		if e.complexity.PartyDetail.ZipCode == nil {
			break
		}

		return e.complexity.PartyDetail.ZipCode(childComplexity), true

	case "Query.GetAvailableCouriers":
		if e.complexity.Query.GetAvailableCouriers == nil {
			break
		}

		args, err := ec.field_Query_GetAvailableCouriers_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetAvailableCouriers(childComplexity, args["location_id"].(string)), true

	case "Query.GetLocation":
		if e.complexity.Query.GetLocation == nil {
			break
		}

		args, err := ec.field_Query_GetLocation_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetLocation(childComplexity, args["id"].(string)), true

	case "Query.GetMatchingCargos":
		if e.complexity.Query.GetMatchingCargos == nil {
			break
		}

		args, err := ec.field_Query_GetMatchingCargos_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetMatchingCargos(childComplexity, args["origin"].(string), args["destination"].(string)), true

	case "Query.GetRegion":
		if e.complexity.Query.GetRegion == nil {
			break
		}

		args, err := ec.field_Query_GetRegion_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetRegion(childComplexity, args["zip_code"].(string)), true

	case "Query.GetRoutedShipments":
		if e.complexity.Query.GetRoutedShipments == nil {
			break
		}

		args, err := ec.field_Query_GetRoutedShipments_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetRoutedShipments(childComplexity, args["location_id"].(string)), true

	case "Query.GetTransitPlaces":
		if e.complexity.Query.GetTransitPlaces == nil {
			break
		}

		args, err := ec.field_Query_GetTransitPlaces_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetTransitPlaces(childComplexity, args["id"].(string)), true

	case "Query.GetUnroutedShipments":
		if e.complexity.Query.GetUnroutedShipments == nil {
			break
		}

		args, err := ec.field_Query_GetUnroutedShipments_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetUnroutedShipments(childComplexity, args["location_id"].(string)), true

	case "Query.GetUser":
		if e.complexity.Query.GetUser == nil {
			break
		}

		args, err := ec.field_Query_GetUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetUser(childComplexity, args["id"].(string)), true

	case "Query.IncomingShipments":
		if e.complexity.Query.IncomingShipments == nil {
			break
		}

		args, err := ec.field_Query_IncomingShipments_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.IncomingShipments(childComplexity, args["location_id"].(string)), true

	case "Query.SearchLocations":
		if e.complexity.Query.SearchLocations == nil {
			break
		}

		args, err := ec.field_Query_SearchLocations_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.SearchLocations(childComplexity, args["keyword"].(string)), true

	case "Region.city":
		if e.complexity.Region.City == nil {
			break
		}

		return e.complexity.Region.City(childComplexity), true

	case "Region.district":
		if e.complexity.Region.District == nil {
			break
		}

		return e.complexity.Region.District(childComplexity), true

	case "Region.province":
		if e.complexity.Region.Province == nil {
			break
		}

		return e.complexity.Region.Province(childComplexity), true

	case "Region.subdistrict":
		if e.complexity.Region.Subdistrict == nil {
			break
		}

		return e.complexity.Region.Subdistrict(childComplexity), true

	case "Region.zip_code":
		if e.complexity.Region.ZipCode == nil {
			break
		}

		return e.complexity.Region.ZipCode(childComplexity), true

	case "Shipment.airway_bill":
		if e.complexity.Shipment.AirwayBill == nil {
			break
		}

		return e.complexity.Shipment.AirwayBill(childComplexity), true

	case "Shipment.created_at":
		if e.complexity.Shipment.CreatedAt == nil {
			break
		}

		return e.complexity.Shipment.CreatedAt(childComplexity), true

	case "Shipment.destination":
		if e.complexity.Shipment.Destination == nil {
			break
		}

		return e.complexity.Shipment.Destination(childComplexity), true

	case "Shipment.id":
		if e.complexity.Shipment.ID == nil {
			break
		}

		return e.complexity.Shipment.ID(childComplexity), true

	case "Shipment.items":
		if e.complexity.Shipment.Items == nil {
			break
		}

		return e.complexity.Shipment.Items(childComplexity), true

	case "Shipment.itinerary_logs":
		if e.complexity.Shipment.ItineraryLogs == nil {
			break
		}

		return e.complexity.Shipment.ItineraryLogs(childComplexity), true

	case "Shipment.origin":
		if e.complexity.Shipment.Origin == nil {
			break
		}

		return e.complexity.Shipment.Origin(childComplexity), true

	case "Shipment.recipient_detail":
		if e.complexity.Shipment.RecipientDetail == nil {
			break
		}

		return e.complexity.Shipment.RecipientDetail(childComplexity), true

	case "Shipment.routing_status":
		if e.complexity.Shipment.RoutingStatus == nil {
			break
		}

		return e.complexity.Shipment.RoutingStatus(childComplexity), true

	case "Shipment.sender_detail":
		if e.complexity.Shipment.SenderDetail == nil {
			break
		}

		return e.complexity.Shipment.SenderDetail(childComplexity), true

	case "Shipment.transport_status":
		if e.complexity.Shipment.TransportStatus == nil {
			break
		}

		return e.complexity.Shipment.TransportStatus(childComplexity), true

	case "TransferRequest.cargo":
		if e.complexity.TransferRequest.Cargo == nil {
			break
		}

		return e.complexity.TransferRequest.Cargo(childComplexity), true

	case "TransferRequest.courier":
		if e.complexity.TransferRequest.Courier == nil {
			break
		}

		return e.complexity.TransferRequest.Courier(childComplexity), true

	case "TransferRequest.created_at":
		if e.complexity.TransferRequest.CreatedAt == nil {
			break
		}

		return e.complexity.TransferRequest.CreatedAt(childComplexity), true

	case "TransferRequest.destination":
		if e.complexity.TransferRequest.Destination == nil {
			break
		}

		return e.complexity.TransferRequest.Destination(childComplexity), true

	case "TransferRequest.id":
		if e.complexity.TransferRequest.ID == nil {
			break
		}

		return e.complexity.TransferRequest.ID(childComplexity), true

	case "TransferRequest.origin":
		if e.complexity.TransferRequest.Origin == nil {
			break
		}

		return e.complexity.TransferRequest.Origin(childComplexity), true

	case "TransferRequest.request_type":
		if e.complexity.TransferRequest.RequestType == nil {
			break
		}

		return e.complexity.TransferRequest.RequestType(childComplexity), true

	case "TransferRequest.shipment":
		if e.complexity.TransferRequest.Shipment == nil {
			break
		}

		return e.complexity.TransferRequest.Shipment(childComplexity), true

	case "TransferRequest.status":
		if e.complexity.TransferRequest.Status == nil {
			break
		}

		return e.complexity.TransferRequest.Status(childComplexity), true

	case "User.entity":
		if e.complexity.User.Entity == nil {
			break
		}

		return e.complexity.User.Entity(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.model_id":
		if e.complexity.User.ModelID == nil {
			break
		}

		return e.complexity.User.ModelID(childComplexity), true

	case "User.type":
		if e.complexity.User.Type == nil {
			break
		}

		return e.complexity.User.Type(childComplexity), true

	case "UserEntity.email":
		if e.complexity.UserEntity.Email == nil {
			break
		}

		return e.complexity.UserEntity.Email(childComplexity), true

	case "UserEntity.location":
		if e.complexity.UserEntity.Location == nil {
			break
		}

		return e.complexity.UserEntity.Location(childComplexity), true

	case "UserEntity.name":
		if e.complexity.UserEntity.Name == nil {
			break
		}

		return e.complexity.UserEntity.Name(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	opCtx := graphql.GetOperationContext(ctx)
	ec := executionContext{opCtx, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCreateLocationInput,
		ec.unmarshalInputCreateShipmentInput,
		ec.unmarshalInputEntityInput,
		ec.unmarshalInputItemInput,
		ec.unmarshalInputRegisterAsCarrierInput,
		ec.unmarshalInputRegisterAsCourierInput,
		ec.unmarshalInputRegisterAsOperatorInput,
		ec.unmarshalInputRequestTransitInput,
		ec.unmarshalInputShipPackageInput,
		ec.unmarshalInputVolumeInput,
	)
	first := true

	switch opCtx.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, opCtx.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, opCtx.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../schemas/cargo.graphqls", Input: `type Cargo {
  id: ID!
  name: String!
  status: String
  maxCapacity: Capacity!
  currentLoad: Capacity!
  carriers: [Carrier!]!
  itineraries: [Itinerary!]!
  shipments: [Shipment!]!
  last_known_location: Location
}

type Capacity {
  weight: Float!
  volume: Float!
}

type Itinerary {
  location: Location!
  estimated_time_arrival: Time!
  actual_time_arrival: Time
}

type Carrier {
  id: ID!
  user_id: ID!
  name: String!
  email: String!
  status: String
  location: Location
}

extend type Query {
  GetMatchingCargos(origin: ID! destination: ID!): [Cargo!]!
}

extend type Mutation {
  LoadShipment(shipment_id: ID! location_id: ID!): String!
  MarkArrival(cargo_id: ID! location_id: ID!): String!
}`, BuiltIn: false},
	{Name: "../schemas/courier.graphqls", Input: `type Courier {
    id: ID!
    user_id: ID!
    name: String!
    email: String!
    status: String!
    location: Location!
}

extend type Query {
    GetAvailableCouriers(location_id: ID!): [Courier!]!
}`, BuiltIn: false},
	{Name: "../schemas/location.graphqls", Input: `type LocationAddress {
    province: String
    city: String
    district: String
    subdistrict: String
    zip_code: String
    latitude: Float
    longitude: Float
    street_address: String
}

type Region {
    province: String!
    city: String!
    district: String!
    subdistrict: String
    zip_code: String!
}

type Location {
    id: String!
    name: String!
    type: String!
    warehouse: Location
    address: LocationAddress!
}

# Enums

enum LocationType {
    warehouse
    depot
}

# Inputs

input CreateLocationInput {
    name: String!
    type: LocationType!
    warehouse_id: ID
    zip_code: String!
    latitude: Float
    longitude: Float
    street_address: String
}

extend type Query {
    GetLocation(id: ID!): Location
    GetRegion(zip_code: String!): Region
    GetTransitPlaces(id: ID!): [Location!]!
    SearchLocations(keyword: String!): [Location!]!
}

extend type Mutation {
    CreateLocation(input: CreateLocationInput): String!
}
`, BuiltIn: false},
	{Name: "../schemas/shipment.graphqls", Input: `scalar Time

type Item {
    name: String!
    amount: Int!
    weight: Int!
    volume: Int
}

type PartyDetail {
    name: String
    contact: String
    province: String
    city: String
    district: String
    subdistrict: String
    address: String
    zip_code: String
}

type ItineraryLog {
    activity_type: String!
    timestamp: Time!
    location: Location
}

type Shipment {
    id: String!
    airway_bill: String!
    transport_status: String!
    routing_status: String!
    items: [Item!]!
    sender_detail: PartyDetail!
    recipient_detail: PartyDetail!
    origin: Location
    destination: Location
    itinerary_logs: [ItineraryLog]!
    created_at: Time!
}

type TransferRequest {
	id: ID!
	request_type: String!
	shipment: Shipment!
	origin: Origin!
	destination: Destination!
	courier: Courier
	cargo: Cargo
	status: String!
	created_at: Time!
}

type Origin {
	location: Location
	requested_by: User
}

type Destination {
	location: Location
	accepted_by: User
	recipient_detail: Entity
}

type Entity {
	name: String
	contact: String
	address: Address
}

type Address {
	province: String
	city: String
	district: String
	subdistrict: String
	street_address: String
	zip_code: String
}

# Input

input CreateShipmentInput {
    origin: ID!
    sender: EntityInput!
    recipient: EntityInput!
    items: [ItemInput!]!
}

input EntityInput {
    name: String!
    phone_number: String!
    zip_code: String!
    street_address: String!
}

input ItemInput {
    name: String!
    amount: Int!
    weight: Int!
    volume: VolumeInput
}

input VolumeInput {
    length: Int
    width: Int
    height: Int
}

input RequestTransitInput {
    shipment_id: ID!
    origin: ID!
    destination: ID!
    courier_id: ID!
}

input ShipPackageInput {
    shipment_id: ID!
    origin: ID!
    destination: ID!
    cargo_id: ID!
}

type Query {
    GetUnroutedShipments(location_id: String!): [Shipment]!
    GetRoutedShipments(location_id: String!): [Shipment]!
    IncomingShipments(location_id: ID!): [TransferRequest!]!
}

type Mutation {
    CreateShipment(input: CreateShipmentInput): String!
    RequestTransit(input: RequestTransitInput): String!
    ScanArrivingShipment(location_id: ID! shipment_id: ID!): String!
    ShipPackage(input: ShipPackageInput!): String!
}
`, BuiltIn: false},
	{Name: "../schemas/user.graphqls", Input: `directive @skipAuth on FIELD_DEFINITION

type User {
    id: ID!
    model_id: ID!
    type: String!
    entity: UserEntity
}

type UserEntity {
    name: String!
    email: String!
    location: Location
}

# Enums

enum OperatorType {
    depot_operator
    warehouse_operator
}

# Inputs

input RegisterAsOperatorInput {
    name: String!
    email: String!
    password: String!
    location_id: String!
    type: OperatorType!
}

input RegisterAsCourierInput {
    name: String!
    email: String!
    password: String!
    location_id: String!
}

input RegisterAsCarrierInput {
    name: String!
    email: String!
    password: String!
    location_id: String!
}

extend type Query {
    GetUser(id: ID!): User
}

extend type Mutation {
    Login(email: String!, password: String!): String!
    RegisterAsOperator(input: RegisterAsOperatorInput!): String!
    RegisterAsCourier(input: RegisterAsCourierInput!): String!
    RegisterAsCarrier(input: RegisterAsCarrierInput!): String!
}`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
