package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"

	"github.com/mproyyan/goparcel/internal/graphql/graph/generated"
	"github.com/mproyyan/goparcel/internal/graphql/graph/model"
)

// Location is the resolver for the location field.
func (r *itineraryLogResolver) Location(ctx context.Context, obj *model.ItineraryLog) (*model.Location, error) {
	panic(fmt.Errorf("not implemented: Location - location"))
}

// CreateShipment is the resolver for the CreateShipment field.
func (r *mutationResolver) CreateShipment(ctx context.Context, input *model.CreateShipmentInput) (string, error) {
	panic(fmt.Errorf("not implemented: CreateShipment - CreateShipment"))
}

// GetUnroutedShipments is the resolver for the GetUnroutedShipments field.
func (r *queryResolver) GetUnroutedShipments(ctx context.Context, locationID string) ([]*model.Shipment, error) {
	panic(fmt.Errorf("not implemented: GetUnroutedShipments - GetUnroutedShipments"))
}

// Origin is the resolver for the origin field.
func (r *shipmentResolver) Origin(ctx context.Context, obj *model.Shipment) (*model.Location, error) {
	panic(fmt.Errorf("not implemented: Origin - origin"))
}

// Destination is the resolver for the destination field.
func (r *shipmentResolver) Destination(ctx context.Context, obj *model.Shipment) (*model.Location, error) {
	panic(fmt.Errorf("not implemented: Destination - destination"))
}

// ItineraryLog returns generated.ItineraryLogResolver implementation.
func (r *Resolver) ItineraryLog() generated.ItineraryLogResolver { return &itineraryLogResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Shipment returns generated.ShipmentResolver implementation.
func (r *Resolver) Shipment() generated.ShipmentResolver { return &shipmentResolver{r} }

type itineraryLogResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type shipmentResolver struct{ *Resolver }
