package port

import (
	"context"
	"time"

	"github.com/mproyyan/goparcel/internal/cargos/app"
	"github.com/mproyyan/goparcel/internal/cargos/domain"
	"github.com/mproyyan/goparcel/internal/common/auth"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	service app.CargoService
	genproto.UnimplementedCargoServiceServer
}

func NewGrpcServer(service app.CargoService) GrpcServer {
	return GrpcServer{service: service}
}

func (g GrpcServer) GetMatchingCargos(ctx context.Context, request *genproto.GetMatchingCargosRequest) (*genproto.CargoResponse, error) {
	cargos, err := g.service.FindMatchingCargos(ctx, request.Origin, request.Destination)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to find matching cargos")
	}

	return &genproto.CargoResponse{Cargos: cargosToProtoResponse(cargos)}, nil
}

func (g GrpcServer) GetCargos(ctx context.Context, request *genproto.GetCargosRequest) (*genproto.CargoResponse, error) {
	cargos, err := g.service.GetCargos(ctx, request.Ids)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get cargos")
	}

	return &genproto.CargoResponse{Cargos: cargosToProtoResponse(cargos)}, nil
}

func (g GrpcServer) CreateCargo(ctx context.Context, request *genproto.CreateCargoRequest) (*emptypb.Empty, error) {
	err := g.service.CreateCargo(ctx, domain.Cargo{
		Name:   request.Name,
		Status: domain.CargoIdle,
		MaxCapacity: domain.Capacity{
			Weight: request.MaxCapacity.Weight,
			Volume: request.MaxCapacity.Volume,
		},
		LastKnownLocation: request.Origin,
	})

	if err != nil {
		return nil, cuserr.Decorate(err, "failed to create cargo")
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) LoadShipment(ctx context.Context, request *genproto.LoadShipmentRequest) (*emptypb.Empty, error) {
	authUser, err := auth.RetrieveAuthUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	err = g.service.LoadShipment(ctx, authUser.ModelID, request.LocationId, request.ShipmentId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to load shipment")
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) MarkArrival(ctx context.Context, request *genproto.MarkArrivalRequest) (*emptypb.Empty, error) {
	err := g.service.MarkArrival(ctx, request.CargoId, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to mark arrival")
	}

	return &emptypb.Empty{}, err
}

func (g GrpcServer) UnloadShipment(ctx context.Context, request *genproto.UnloadShipmentRequest) (*emptypb.Empty, error) {
	err := g.service.UnloadShipment(ctx, request.CargoId, request.ShipmentId)
	if err != nil {
		return nil, cuserr.Decorate(err, "cargo service failed to unload shipment")
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) AssignCarrier(ctx context.Context, request *genproto.AssignCarrierRequest) (*emptypb.Empty, error) {
	err := g.service.AssignCarrier(ctx, request.CargoId, request.CarrierIds)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to assign carrier")
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) AssignRoute(ctx context.Context, request *genproto.AssignRouteRequest) (*emptypb.Empty, error) {
	itineraries := convertProtoToItineraries(request.Itineraries)
	err := g.service.AssignRoute(ctx, request.CargoId, itineraries)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to assign route")
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) GetUnroutedCargos(ctx context.Context, request *genproto.GetUnroutedCargosRequest) (*genproto.CargoResponse, error) {
	cargos, err := g.service.GetUnroutedCargos(ctx, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get unrouted cargos")
	}

	return &genproto.CargoResponse{Cargos: cargosToProtoResponse(cargos)}, nil
}

func (g GrpcServer) FindCargosWithoutCarrier(ctx context.Context, request *genproto.FindCargosWithoutCarrierRequest) (*genproto.CargoResponse, error) {
	cargos, err := g.service.FindCargosWithoutCarrier(ctx, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to find cargos without carrier")
	}

	return &genproto.CargoResponse{Cargos: cargosToProtoResponse(cargos)}, nil
}

func (g GrpcServer) GetIdleCarriers(ctx context.Context, request *genproto.GetIdleCarriersRequest) (*genproto.CarrierResponse, error) {
	carriers, err := g.service.GetIdleCarriers(ctx, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get idle carriers")
	}

	var protoCarriers []*genproto.Carrier
	for _, carrier := range carriers {
		protoCarriers = append(protoCarriers, &genproto.Carrier{
			Id:         carrier.ID,
			UserId:     carrier.UserID,
			Name:       carrier.Name,
			Email:      carrier.Email,
			CargoId:    carrier.CargoID,
			Status:     carrier.Status.String(),
			LocationId: carrier.LocationID,
		})
	}

	return &genproto.CarrierResponse{Carriers: protoCarriers}, nil
}

func cargoToProtoResponse(cargo *domain.Cargo) *genproto.Cargo {
	if cargo == nil {
		return nil
	}

	return &genproto.Cargo{
		Id:                cargo.ID,
		Name:              cargo.Name,
		Status:            cargo.Status.String(),
		MaxCapacity:       convertCapacityToProto(cargo.MaxCapacity),
		CurrentLoad:       convertCapacityToProto(cargo.CurrentLoad),
		Carriers:          cargo.Carriers,
		Itineraries:       convertItinerariesToProto(cargo.Itineraries),
		Shipments:         cargo.Shipments,
		LastKnownLocation: cargo.LastKnownLocation,
	}
}

func cargosToProtoResponse(cargos []*domain.Cargo) []*genproto.Cargo {
	var result []*genproto.Cargo
	for _, cargo := range cargos {
		result = append(result, cargoToProtoResponse(cargo))
	}
	return result
}

func convertCapacityToProto(capacity domain.Capacity) *genproto.Capacity {
	return &genproto.Capacity{
		Weight: capacity.Weight,
		Volume: capacity.Volume,
	}
}

func convertItinerariesToProto(itineraries []domain.Itinerary) []*genproto.Itinerary {
	var result []*genproto.Itinerary
	for _, itinerary := range itineraries {
		var actualTime *timestamppb.Timestamp
		if itinerary.ActualTimeArrival != nil {
			actualTime = timestamppb.New(*itinerary.ActualTimeArrival)
		}

		result = append(result, &genproto.Itinerary{
			Location:             itinerary.Location,
			EstimatedTimeArrival: timestamppb.New(itinerary.EstimatedTimeArrival),
			ActualTimeArrival:    actualTime,
		})
	}
	return result
}

func convertProtoToItineraries(itineraries []*genproto.Itinerary) []domain.Itinerary {
	var result []domain.Itinerary
	for _, itinerary := range itineraries {
		var actualTime *time.Time
		if itinerary.ActualTimeArrival != nil {
			t := itinerary.ActualTimeArrival.AsTime()
			actualTime = &t
		}

		result = append(result, domain.Itinerary{
			Location:             itinerary.Location,
			EstimatedTimeArrival: itinerary.EstimatedTimeArrival.AsTime(),
			ActualTimeArrival:    actualTime,
		})
	}

	return result
}
