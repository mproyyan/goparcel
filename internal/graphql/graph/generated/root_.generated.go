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
	Courier() CourierResolver
	ItineraryLog() ItineraryLogResolver
	Location() LocationResolver
	Mutation() MutationResolver
	Query() QueryResolver
	Shipment() ShipmentResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Courier struct {
		Email    func(childComplexity int) int
		ID       func(childComplexity int) int
		Location func(childComplexity int) int
		Name     func(childComplexity int) int
		Status   func(childComplexity int) int
		UserID   func(childComplexity int) int
	}

	Item struct {
		Amount func(childComplexity int) int
		Name   func(childComplexity int) int
		Volume func(childComplexity int) int
		Weight func(childComplexity int) int
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
		CreateLocation func(childComplexity int, input *model.CreateLocationInput) int
		CreateShipment func(childComplexity int, input *model.CreateShipmentInput) int
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
		GetRegion            func(childComplexity int, zipCode string) int
		GetTransitPlaces     func(childComplexity int, id string) int
		GetUnroutedShipments func(childComplexity int, locationID string) int
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

	case "Query.GetRegion":
		if e.complexity.Query.GetRegion == nil {
			break
		}

		args, err := ec.field_Query_GetRegion_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetRegion(childComplexity, args["zip_code"].(string)), true

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

type Query {
    GetUnroutedShipments(location_id: String!): [Shipment]!
}

type Mutation {
    CreateShipment(input: CreateShipmentInput): String!
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
