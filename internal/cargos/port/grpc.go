package port

import (
	"context"

	"github.com/mproyyan/goparcel/internal/cargos/app"
	"github.com/mproyyan/goparcel/internal/cargos/domain"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto"
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

func cargoToProtoResponse(cargo *domain.Cargo) *genproto.Cargo {
	if cargo == nil {
		return nil
	}

	return &genproto.Cargo{
		Id:                cargo.ID,
		Name:              cargo.Name,
		Status:            cargo.Status,
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
		result = append(result, &genproto.Itinerary{
			Location:             itinerary.Location,
			EstimatedTimeArrival: timestamppb.New(itinerary.EstimatedTimeArrival),
			ActualTimeArrival:    timestamppb.New(itinerary.ActualTimeArrival),
		})
	}
	return result
}
