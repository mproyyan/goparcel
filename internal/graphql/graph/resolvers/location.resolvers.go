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

// Warehouse is the resolver for the warehouse field.
func (r *locationResolver) Warehouse(ctx context.Context, obj *model.Location) (*model.Location, error) {
	if obj.Warehouse != nil && obj.Warehouse.ID == "" {
		return nil, nil
	}

	return r.locationLoader.Load(ctx, obj.Warehouse.ID)
}

// CreateLocation is the resolver for the CreateLocation field.
func (r *mutationResolver) CreateLocation(ctx context.Context, input *model.CreateLocationInput) (string, error) {
	_, err := r.locationService.CreateLocation(ctx, &genproto.CreateLocationRequest{
		Name:          input.Name,
		Type:          input.Type.String(),
		WarehouseId:   safePointer(input.WarehouseID, ""),
		ZipCode:       input.ZipCode,
		Latitude:      safePointer(input.Latitude, 0),
		Longitude:     safePointer(input.Longitude, 0),
		StreetAddress: safePointer(input.StreetAddress, ""),
	})

	if err != nil {
		return "", err
	}

	return "successfully made the location", nil
}

// GetLocation is the resolver for the GetLocation field.
func (r *queryResolver) GetLocation(ctx context.Context, id string) (*model.Location, error) {
	location, err := r.locationService.GetLocation(ctx, &genproto.GetLocationRequest{LocationId: id})
	if err != nil {
		return nil, err
	}

	return locationToGraphResponse(location), nil
}

// GetRegion is the resolver for the GetRegion field.
func (r *queryResolver) GetRegion(ctx context.Context, zipCode string) (*model.Region, error) {
	region, err := r.locationService.GetRegion(ctx, &genproto.GetRegionRequest{Zipcode: zipCode})
	if err != nil {
		return nil, err
	}

	return regionToGraphResponse(region), nil
}

// GetTransitPlaces is the resolver for the GetTransitPlaces field.
func (r *queryResolver) GetTransitPlaces(ctx context.Context, id string) ([]*model.Location, error) {
	result, err := r.locationService.GetTransitPlaces(ctx, &genproto.GetTransitPlacesRequest{LocationId: id})
	if err != nil {
		return nil, err
	}

	return locationsToGraphResponse(result.Locations), nil
}

// SearchLocations is the resolver for the SearchLocations field.
func (r *queryResolver) SearchLocations(ctx context.Context, keyword string) ([]*model.Location, error) {
	result, err := r.locationService.SearchLocations(ctx, &genproto.SearchLocationRequest{Keyword: keyword})
	if err != nil {
		return nil, err
	}

	return locationsToGraphResponse(result.Locations), nil
}

// Location returns generated.LocationResolver implementation.
func (r *Resolver) Location() generated.LocationResolver { return &locationResolver{r} }

type locationResolver struct{ *Resolver }
