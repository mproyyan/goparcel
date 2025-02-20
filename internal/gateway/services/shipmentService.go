package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mproyyan/goparcel/internal/common/genproto/shipments"
	"github.com/mproyyan/goparcel/internal/gateway/requests"
	"github.com/mproyyan/goparcel/internal/gateway/responses"
)

type ShipmentService struct {
	router fiber.Router
	client shipments.ShipmentServiceClient
}

func NewShipmentService(router fiber.Router, client shipments.ShipmentServiceClient) ShipmentService {
	return ShipmentService{
		router: router,
		client: client,
	}
}

func (s ShipmentService) createShipment(c *fiber.Ctx) error {
	// Parse request body
	var request requests.CreateShipmentRequest
	if err := c.BodyParser(&request); err != nil {
		errResponse := responses.NewInvalidRequestBodyErrorResponse(err)
		return c.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	sender := entityToProtoRequest(request.Sender)
	recipient := entityToProtoRequest(request.Recipient)
	packages := packageToProtoRequest(request.Packages)

	// Call CreateShipment rpc
	_, err := s.client.CreateShipment(c.Context(), &shipments.CreateShipmentRequest{
		Origin:    request.Origin,
		Sender:    sender,
		Recipient: recipient,
		Package:   packages,
	})

	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully made the shipment"})
}

func (s ShipmentService) Bootstrap() {
	shipment := s.router.Group("/shipments")

	shipment.Post("/", s.createShipment)
}

func packageToProtoRequest(packages []requests.Package) []*shipments.Package {
	var protoPackages []*shipments.Package

	for _, p := range packages {
		protoVolume := &shipments.Volume{
			Length: p.Volume.Length,
			Width:  p.Volume.Width,
			Height: p.Volume.Height,
		}

		protoPackages = append(protoPackages, &shipments.Package{
			Name:   p.Name,
			Amount: p.Amount,
			Weight: p.Weight,
			Volume: protoVolume,
		})
	}

	return protoPackages
}

func entityToProtoRequest(entity requests.Entity) *shipments.Entity {
	return &shipments.Entity{
		Name:          entity.Name,
		PhoneNumber:   entity.PhoneNumber,
		ZipCode:       entity.ZipCode,
		StreetAddress: entity.StreetAddress,
	}
}
