package port

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/shipments/app"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	service app.ShipmentService
	genproto.UnimplementedShipmentServiceServer
}

func NewGrpcServer(service app.ShipmentService) GrpcServer {
	return GrpcServer{service: service}
}

func (g GrpcServer) CreateShipment(ctx context.Context, request *genproto.CreateShipmentRequest) (*emptypb.Empty, error) {
	// Build sender detail
	sender := domain.Entity{
		Name:    request.Sender.Name,
		Contact: request.Sender.PhoneNumber,
		Address: domain.Address{
			ZipCode:       request.Sender.ZipCode,
			StreetAddress: request.Sender.ZipCode,
		},
	}

	// Build recipient detail
	recipient := domain.Entity{
		Name:    request.Recipient.Name,
		Contact: request.Recipient.PhoneNumber,
		Address: domain.Address{
			ZipCode:       request.Recipient.ZipCode,
			StreetAddress: request.Recipient.ZipCode,
		},
	}

	// Convert package from request to item domain
	items := protoRequestToItems(request.Package)

	// Call CreateShipment rpc
	err := g.service.CreateShipment(ctx, request.Origin, sender, recipient, items)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to call CreateShipment service")
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) GetUnroutedShipment(ctx context.Context, request *genproto.GetUnroutedShipmentRequest) (*genproto.ShipmentResponse, error) {
	shipments, err := g.service.UnroutedShipments(ctx, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get unrouted shipments")
	}

	return shipmentsToProtoResponse(shipments), nil
}

func protoRequestToItems(packages []*genproto.Package) []domain.Item {
	var items []domain.Item
	for _, pkg := range packages {
		volume := int32(0)
		if pkg.Volume != nil {
			volume = pkg.Volume.Length * pkg.Volume.Width * pkg.Volume.Height
		}

		item := domain.Item{
			Name:   pkg.Name,
			Amount: int(pkg.Amount),
			Weight: pkg.Weight,
			Volume: volume,
		}
		items = append(items, item)
	}
	return items
}

// shipmentsToProtoResponse converts a slice of domain.Shipment to *proto.ShipmentResponse
func shipmentsToProtoResponse(domainShipments []domain.Shipment) *genproto.ShipmentResponse {
	var protoShipments []*genproto.Shipment

	for _, s := range domainShipments {
		protoShipments = append(protoShipments, &genproto.Shipment{
			Id:              s.ID,
			AirwayBill:      s.AirwayBill,
			TransportStatus: s.TransportStatus.String(),
			RoutingStatus:   s.RoutingStatus.String(),
			Items:           itemsToProtoResponse(s.Items),
			Sender:          entityToProtoResponse(s.Sender),
			Recipient:       entityToProtoResponse(s.Recipient),
			Origin:          locationToProtoResponse(s.Origin),
			Destination:     locationToProtoResponse(s.Destination),
			ItineraryLogs:   itineraryToProtoResponse(s.ItineraryLogs),
		})
	}

	return &genproto.ShipmentResponse{
		Shipment: protoShipments,
	}
}

func itemsToProtoResponse(items []domain.Item) []*genproto.Item {
	var protoItems []*genproto.Item
	for _, item := range items {
		protoItems = append(protoItems, &genproto.Item{
			Name:   item.Name,
			Amount: int32(item.Amount),
			Weight: item.Weight,
			Volume: item.Volume,
		})
	}
	return protoItems
}

func entityToProtoResponse(e domain.Entity) *genproto.EntityDetail {
	return &genproto.EntityDetail{
		Name:    e.Name,
		Contact: e.Contact,
		Address: addressToProtoResponse(e.Address),
	}
}

func addressToProtoResponse(a domain.Address) *genproto.Address {
	return &genproto.Address{
		Province:      a.Province,
		City:          a.City,
		District:      a.District,
		Subdistrict:   a.Subdistrict,
		StreetAddress: a.StreetAddress,
		ZipCode:       a.ZipCode,
	}
}

func locationToProtoResponse(l domain.Location) *genproto.Location {
	return &genproto.Location{
		Id:   l.ID,
		Name: l.Name,
		Type: l.Type,
	}
}

func itineraryToProtoResponse(logs []domain.ItineraryLog) []*genproto.ItineraryLog {
	var protoLogs []*genproto.ItineraryLog
	for _, log := range logs {
		protoLogs = append(protoLogs, &genproto.ItineraryLog{
			ActivityType: log.ActivityType.String(),
			Timestamp:    timestamppb.New(log.Timestamp),
			Location:     locationToProtoResponse(log.Location),
		})
	}
	return protoLogs
}
