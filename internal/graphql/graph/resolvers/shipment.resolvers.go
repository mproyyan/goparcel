package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"

	"github.com/mproyyan/goparcel/internal/common/auth"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/graphql/graph/generated"
	"github.com/mproyyan/goparcel/internal/graphql/graph/middlewares"
	"github.com/mproyyan/goparcel/internal/graphql/graph/model"
)

// Location is the resolver for the location field.
func (r *destinationResolver) Location(ctx context.Context, obj *model.Destination) (*model.Location, error) {
	if obj.Location != nil && obj.Location.ID == "" {
		return nil, nil
	}

	return r.locationLoader.Load(ctx, obj.Location.ID)
}

// AcceptedBy is the resolver for the accepted_by field.
func (r *destinationResolver) AcceptedBy(ctx context.Context, obj *model.Destination) (*model.User, error) {
	if obj.AcceptedBy != nil && obj.AcceptedBy.ID == "" {
		return nil, nil
	}

	return r.userLoader.Load(ctx, obj.AcceptedBy.ID)
}

// Location is the resolver for the location field.
func (r *itineraryLogResolver) Location(ctx context.Context, obj *model.ItineraryLog) (*model.Location, error) {
	if obj.Location != nil && obj.Location.ID == "" {
		return nil, nil
	}

	return r.locationLoader.Load(ctx, obj.Location.ID)
}

// CreateShipment is the resolver for the CreateShipment field.
func (r *mutationResolver) CreateShipment(ctx context.Context, input *model.CreateShipmentInput) (string, error) {
	response, err := r.shipmentService.CreateShipment(ctx, &genproto.CreateShipmentRequest{
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

	return response.AirwayBillNumber, nil
}

// RequestTransit is the resolver for the RequestTransit field.
func (r *mutationResolver) RequestTransit(ctx context.Context, input *model.RequestTransitInput) (string, error) {
	authUser, err := middlewares.GetAuthUser(ctx)
	if err != nil {
		return "", err
	}

	ctx = auth.SendAuthUser(ctx, authUser.UserID, authUser.ModelID)
	_, err = r.shipmentService.RequestTransit(ctx, &genproto.RequestTransitRequest{
		ShipmentId:  input.ShipmentID,
		Origin:      input.Origin,
		Destination: input.Destination,
		CourierId:   input.CourierID,
	})

	if err != nil {
		return "", err
	}

	return "Transit request created successfully", nil
}

// ScanArrivingShipment is the resolver for the ScanArrivingShipment field.
func (r *mutationResolver) ScanArrivingShipment(ctx context.Context, locationID string, shipmentID string) (string, error) {
	authUser, err := middlewares.GetAuthUser(ctx)
	if err != nil {
		return "", err
	}

	ctx = auth.SendAuthUser(ctx, authUser.UserID, authUser.ModelID)
	_, err = r.shipmentService.ScanArrivingShipment(ctx, &genproto.ScanArrivingShipmentRequest{
		ShipmentId: shipmentID,
		LocationId: locationID,
	})

	if err != nil {
		return "", err
	}

	return "The shipment has been scanned successfully", nil
}

// ShipPackage is the resolver for the ShipPackage field.
func (r *mutationResolver) ShipPackage(ctx context.Context, input model.ShipPackageInput) (string, error) {
	authUser, err := middlewares.GetAuthUser(ctx)
	if err != nil {
		return "", err
	}

	ctx = auth.SendAuthUser(ctx, authUser.UserID, authUser.ModelID)
	_, err = r.shipmentService.ShipPackage(ctx, &genproto.ShipPackageRequest{
		ShipmentId:  input.ShipmentID,
		CargoId:     input.CargoID,
		Origin:      input.Origin,
		Destination: input.Destination,
	})

	if err != nil {
		return "", err
	}

	return "successfully created a shipment request", nil
}

// DeliverPackage is the resolver for the DeliverPackage field.
func (r *mutationResolver) DeliverPackage(ctx context.Context, input model.DeliverPackageInput) (string, error) {
	authUser, err := middlewares.GetAuthUser(ctx)
	if err != nil {
		return "", err
	}

	ctx = auth.SendAuthUser(ctx, authUser.UserID, authUser.ModelID)
	_, err = r.shipmentService.DeliverPackage(ctx, &genproto.DeliverPackageRequest{
		Origin:     input.Origin,
		ShipmentId: input.ShipmentID,
		CourierId:  input.CourierID,
	})

	if err != nil {
		return "", err
	}

	return "successfully created a package delivery request", nil
}

// CompleteShipment is the resolver for the CompleteShipment field.
func (r *mutationResolver) CompleteShipment(ctx context.Context, shipmentID string) (string, error) {
	authUser, err := middlewares.GetAuthUser(ctx)
	if err != nil {
		return "", err
	}

	ctx = auth.SendAuthUser(ctx, authUser.UserID, authUser.ModelID)
	_, err = r.shipmentService.CompleteShipment(ctx, &genproto.CompleteShipmentRequest{ShipmentId: shipmentID})
	if err != nil {
		return "", err
	}

	return "successfully completed the shipment", nil
}

// Location is the resolver for the location field.
func (r *originResolver) Location(ctx context.Context, obj *model.Origin) (*model.Location, error) {
	if obj.Location != nil && obj.Location.ID == "" {
		return nil, nil
	}

	return r.locationLoader.Load(ctx, obj.Location.ID)
}

// RequestedBy is the resolver for the requested_by field.
func (r *originResolver) RequestedBy(ctx context.Context, obj *model.Origin) (*model.User, error) {
	if obj.RequestedBy != nil && obj.RequestedBy.ID == "" {
		return nil, nil
	}

	return r.userLoader.Load(ctx, obj.RequestedBy.ID)
}

// GetUnroutedShipments is the resolver for the GetUnroutedShipments field.
func (r *queryResolver) GetUnroutedShipments(ctx context.Context, locationID string) ([]*model.Shipment, error) {
	shipment, err := r.shipmentService.GetUnroutedShipment(ctx, &genproto.GetUnroutedShipmentRequest{LocationId: locationID})
	if err != nil {
		return nil, err
	}

	return shipmentsToGraphResponse(shipment.Shipment), nil
}

// GetRoutedShipments is the resolver for the GetRoutedShipments field.
func (r *queryResolver) GetRoutedShipments(ctx context.Context, locationID string) ([]*model.Shipment, error) {
	result, err := r.shipmentService.GetRoutedShipments(ctx, &genproto.GetRoutedShipmentsRequest{LocationId: locationID})
	if err != nil {
		return nil, err
	}

	return shipmentsToGraphResponse(result.Shipment), nil
}

// IncomingShipments is the resolver for the IncomingShipments field.
func (r *queryResolver) IncomingShipments(ctx context.Context, locationID string) ([]*model.TransferRequest, error) {
	result, err := r.shipmentService.IncomingShipments(ctx, &genproto.IncomingShipmentRequest{LocationId: locationID})
	if err != nil {
		return nil, err
	}

	return transferRequestsToGraphResponse(result.TransferRequests), nil
}

// TrackPackage is the resolver for the TrackPackage field.
func (r *queryResolver) TrackPackage(ctx context.Context, airwayBill string) ([]*model.ItineraryLog, error) {
	response, err := r.shipmentService.TrackPackage(ctx, &genproto.TrackPackageRequest{Awb: airwayBill})
	if err != nil {
		return nil, fmt.Errorf("failed to track package: %w", err)
	}

	return itineraryLogsToGraphResponse(response.Itineraries), nil
}

// Origin is the resolver for the origin field.
func (r *shipmentResolver) Origin(ctx context.Context, obj *model.Shipment) (*model.Location, error) {
	if obj.Origin != nil && obj.Origin.ID == "" {
		return nil, nil
	}

	return r.locationLoader.Load(ctx, obj.Origin.ID)
}

// Destination is the resolver for the destination field.
func (r *shipmentResolver) Destination(ctx context.Context, obj *model.Shipment) (*model.Location, error) {
	if obj.Destination != nil && obj.Destination.ID == "" {
		return nil, nil
	}

	return r.locationLoader.Load(ctx, obj.Destination.ID)
}

// Shipment is the resolver for the shipment field.
func (r *transferRequestResolver) Shipment(ctx context.Context, obj *model.TransferRequest) (*model.Shipment, error) {
	if obj.Shipment != nil && obj.Shipment.ID == "" {
		return nil, nil
	}

	return r.shipmentLoader.Load(ctx, obj.Shipment.ID)
}

// Courier is the resolver for the courier field.
func (r *transferRequestResolver) Courier(ctx context.Context, obj *model.TransferRequest) (*model.Courier, error) {
	if obj.Courier != nil && obj.Courier.ID == "" {
		return nil, nil
	}

	return r.courierLoader.Load(ctx, obj.Courier.ID)
}

// Cargo is the resolver for the cargo field.
func (r *transferRequestResolver) Cargo(ctx context.Context, obj *model.TransferRequest) (*model.Cargo, error) {
	if obj.Cargo == nil {
		return nil, nil
	}

	if obj.Cargo.ID == "" {
		return nil, nil
	}

	return r.cargoLoader.Load(ctx, obj.Cargo.ID)
}

// Destination returns generated.DestinationResolver implementation.
func (r *Resolver) Destination() generated.DestinationResolver { return &destinationResolver{r} }

// ItineraryLog returns generated.ItineraryLogResolver implementation.
func (r *Resolver) ItineraryLog() generated.ItineraryLogResolver { return &itineraryLogResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Origin returns generated.OriginResolver implementation.
func (r *Resolver) Origin() generated.OriginResolver { return &originResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Shipment returns generated.ShipmentResolver implementation.
func (r *Resolver) Shipment() generated.ShipmentResolver { return &shipmentResolver{r} }

// TransferRequest returns generated.TransferRequestResolver implementation.
func (r *Resolver) TransferRequest() generated.TransferRequestResolver {
	return &transferRequestResolver{r}
}

type destinationResolver struct{ *Resolver }
type itineraryLogResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type originResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type shipmentResolver struct{ *Resolver }
type transferRequestResolver struct{ *Resolver }
