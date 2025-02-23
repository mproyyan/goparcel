package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/graphql/graph/generated"
	"github.com/mproyyan/goparcel/internal/graphql/graph/model"
)

// Location is the resolver for the location field.
func (r *itineraryLogResolver) Location(ctx context.Context, obj *model.ItineraryLog) (*model.Location, error) {
	return r.locationLoder.Load(ctx, obj.Location.ID)
}

// CreateShipment is the resolver for the CreateShipment field.
func (r *mutationResolver) CreateShipment(ctx context.Context, input *model.CreateShipmentInput) (string, error) {
	_, err := r.shipmentService.CreateShipment(ctx, &genproto.CreateShipmentRequest{
		Origin: input.Origin,
		Sender: &genproto.Entity{
			Name:          input.Sender.Name,
			PhoneNumber:   input.Sender.PhoneNumber,
			ZipCode:       input.Sender.ZipCode,
			StreetAddress: input.Sender.StreetAddress,
		},
		Recipient: &genproto.Entity{
			Name:          input.Recipient.Name,
			PhoneNumber:   input.Recipient.PhoneNumber,
			ZipCode:       input.Recipient.ZipCode,
			StreetAddress: input.Recipient.StreetAddress,
		},
		Package: itemInputToProtoRequest(input.Items),
	})

	if err != nil {
		return "", err
	}

	return "successfully made the shipment", nil
}

// GetUnroutedShipments is the resolver for the GetUnroutedShipments field.
func (r *queryResolver) GetUnroutedShipments(ctx context.Context, locationID string) ([]*model.Shipment, error) {
	shipment, err := r.shipmentService.GetUnroutedShipment(ctx, &genproto.GetUnroutedShipmentRequest{LocationId: locationID})
	if err != nil {
		return nil, err
	}

	return shipmentsToGraphResponse(shipment.Shipment), nil
}

// Origin is the resolver for the origin field.
func (r *shipmentResolver) Origin(ctx context.Context, obj *model.Shipment) (*model.Location, error) {
	return r.locationLoder.Load(ctx, obj.Origin.ID)
}

// Destination is the resolver for the destination field.
func (r *shipmentResolver) Destination(ctx context.Context, obj *model.Shipment) (*model.Location, error) {
	return r.locationLoder.Load(ctx, obj.Destination.ID)
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
