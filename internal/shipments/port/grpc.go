package port

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto/shipments"
	"github.com/mproyyan/goparcel/internal/shipments/app"
	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	service app.ShipmentService
	shipments.UnimplementedShipmentServiceServer
}

func NewGrpcServer(service app.ShipmentService) GrpcServer {
	return GrpcServer{service: service}
}

func (g GrpcServer) CreateShipment(ctx context.Context, request *shipments.CreateShipmentRequest) (*emptypb.Empty, error) {
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
	items := packageProtoToDomain(request.Package)

	// Call CreateShipment rpc
	err := g.service.CreateShipment(ctx, request.Origin, sender, recipient, items)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to call CreateShipment service")
	}

	return &emptypb.Empty{}, nil
}

func packageProtoToDomain(packages []*shipments.Package) []domain.Item {
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
