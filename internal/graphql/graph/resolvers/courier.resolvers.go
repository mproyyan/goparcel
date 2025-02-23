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
func (r *courierResolver) Location(ctx context.Context, obj *model.Courier) (*model.Location, error) {
	return r.locationLoder.Load(ctx, obj.Location.ID)
}

// GetAvailableCouriers is the resolver for the GetAvailableCouriers field.
func (r *queryResolver) GetAvailableCouriers(ctx context.Context, locationID string) ([]*model.Courier, error) {
	couriers, err := r.courierService.GetAvailableCouriers(ctx, &genproto.GetAvailableCourierRequest{LocationId: locationID})
	if err != nil {
		return nil, err
	}

	return couriersToGraphResponse(couriers.Couriers), nil
}

// Courier returns generated.CourierResolver implementation.
func (r *Resolver) Courier() generated.CourierResolver { return &courierResolver{r} }

type courierResolver struct{ *Resolver }
