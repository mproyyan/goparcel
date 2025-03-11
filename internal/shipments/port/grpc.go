package port

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/auth"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/mproyyan/goparcel/internal/shipments/app"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	logrus.WithFields(logrus.Fields{
		"origin":    request.Origin,
		"sender":    request.Sender.Name,
		"recipient": request.Recipient.Name,
	}).Info("Shipment created")

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) GetUnroutedShipment(ctx context.Context, request *genproto.GetUnroutedShipmentRequest) (*genproto.ShipmentResponse, error) {
	shipments, err := g.service.UnroutedShipments(ctx, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get unrouted shipments")
	}

	return shipmentsToProtoResponse(shipments), nil
}

func (g GrpcServer) GetRoutedShipments(ctx context.Context, request *genproto.GetRoutedShipmentsRequest) (*genproto.ShipmentResponse, error) {
	shipments, err := g.service.RoutedShipments(ctx, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get routed shipments")
	}

	return shipmentsToProtoResponse(shipments), nil
}

func (g GrpcServer) RequestTransit(ctx context.Context, request *genproto.RequestTransitRequest) (*emptypb.Empty, error) {
	authUser, err := auth.RetrieveAuthUser(ctx)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to retrieve user")
	}

	err = g.service.RequestTransit(ctx, request.ShipmentId, request.Origin, request.Destination, request.CourierId, authUser.UserID)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to request transit")
	}

	logrus.WithFields(logrus.Fields{
		"user_id":     authUser.UserID,
		"type":        "transit",
		"shipment_id": request.ShipmentId,
		"origin":      request.Origin,
		"destination": request.Destination,
		"courier_id":  request.CourierId,
	}).Info("Requesting transfer")

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) IncomingShipments(ctx context.Context, request *genproto.IncomingShipmentRequest) (*genproto.TransferRequestResponse, error) {
	transferRequests, err := g.service.IncomingShipments(ctx, request.LocationId)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get incoming shipments from shipment service")
	}

	response := transferRequestsToProtoResponse(transferRequests)
	return &genproto.TransferRequestResponse{TransferRequests: response}, nil
}

func (g GrpcServer) GetShipments(ctx context.Context, request *genproto.GetShipmentsRequest) (*genproto.ShipmentResponse, error) {
	shipments, err := g.service.GetShipments(ctx, request.Ids)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to get shipments from shipment service")
	}

	return shipmentsToProtoResponse(shipments), nil
}

func (g GrpcServer) ScanArrivingShipment(ctx context.Context, request *genproto.ScanArrivingShipmentRequest) (*emptypb.Empty, error) {
	authUser, err := auth.RetrieveAuthUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	err = g.service.ScanArrivingShipment(ctx, request.LocationId, request.ShipmentId, authUser.UserID)
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"user_id":     authUser.UserID,
		"shipment_id": request.ShipmentId,
		"location_id": request.LocationId,
	}).Info("Shipment scanned at location")

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) ShipPackage(ctx context.Context, request *genproto.ShipPackageRequest) (*emptypb.Empty, error) {
	authUser, err := auth.RetrieveAuthUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	err = g.service.ShipPackage(ctx, request.ShipmentId, request.CargoId, request.Origin, request.Destination, authUser.UserID)
	if err != nil {
		return nil, cuserr.Decorate(err, "ship package failed")
	}

	logrus.WithFields(logrus.Fields{
		"user_id":     authUser.UserID,
		"type":        "shipment",
		"shipment_id": request.ShipmentId,
		"cargo_id":    request.CargoId,
		"origin":      request.Origin,
		"destination": request.Destination,
	}).Info("Requesting transer")

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) AddItineraryHistory(ctx context.Context, request *genproto.AddItineraryHistoryRequest) (*emptypb.Empty, error) {
	activityType := domain.StringToActivityType(request.Activity)
	if activityType == domain.Unknown {
		return nil, status.Error(codes.InvalidArgument, "invalid activity type")
	}

	err := g.service.RecordItinerary(ctx, request.ShipmentIds, request.LocationId, activityType)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to record itinerary")
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) DeliverPackage(ctx context.Context, request *genproto.DeliverPackageRequest) (*emptypb.Empty, error) {
	authUser, err := auth.RetrieveAuthUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	err = g.service.DeliverPackage(ctx, request.Origin, request.ShipmentId, request.CourierId, authUser.UserID)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to deliver package")
	}

	logrus.WithFields(logrus.Fields{
		"user_id":     authUser.UserID,
		"type":        "delivery",
		"shipment_id": request.ShipmentId,
		"origin":      request.Origin,
		"courier_id":  request.CourierId,
	}).Info("Requesting transfer")

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) CompleteShipment(ctx context.Context, request *genproto.CompleteShipmentRequest) (*emptypb.Empty, error) {
	authUser, err := auth.RetrieveAuthUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	err = g.service.CompleteShipment(ctx, request.ShipmentId, authUser.UserID)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to complete shipment")
	}

	return &emptypb.Empty{}, nil
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
func shipmentsToProtoResponse(domainShipments []*domain.Shipment) *genproto.ShipmentResponse {
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
			Origin:          s.Origin,
			Destination:     s.Destination,
			ItineraryLogs:   itineraryToProtoResponse(s.ItineraryLogs),
			CreatedAt:       timestamppb.New(s.CreatedAt),
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

func itineraryToProtoResponse(logs []domain.ItineraryLog) []*genproto.ItineraryLog {
	var protoLogs []*genproto.ItineraryLog
	for _, log := range logs {
		protoLogs = append(protoLogs, &genproto.ItineraryLog{
			ActivityType: log.ActivityType.String(),
			Timestamp:    timestamppb.New(log.Timestamp),
			LocationId:   log.Location,
		})
	}
	return protoLogs
}

func transferRequestToProtoResponse(req *domain.TransferRequest) *genproto.TransferRequest {
	if req == nil {
		return nil
	}

	return &genproto.TransferRequest{
		Id:          req.ID,
		RequestType: req.RequestType.String(),
		ShipmentId:  req.ShipmentID,
		Origin: &genproto.TransferRequestOrigin{
			Location:    req.Origin.Location,
			RequestedBy: req.Origin.RequestedBy,
		},
		Destinaion: &genproto.TransferRequestDestination{
			Location:        req.Destination.Location,
			AcceptedBy:      req.Destination.AcceptedBy,
			RecipientDetail: entityToProtoResponse(req.Destination.RecipientDetail),
		},
		CourierId: req.CourierID,
		CargoId:   req.CargoID,
		Status:    req.Status.String(),
		CreatedAt: timestamppb.New(req.CreatedAt),
	}
}

func transferRequestsToProtoResponse(reqs []*domain.TransferRequest) []*genproto.TransferRequest {
	var protoRequests []*genproto.TransferRequest
	for _, req := range reqs {
		protoRequests = append(protoRequests, transferRequestToProtoResponse(req))
	}
	return protoRequests
}
