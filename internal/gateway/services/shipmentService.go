package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mproyyan/goparcel/internal/common/genproto/shipments"
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
	var request shipments.CreateShipmentRequest
	if err := c.BodyParser(&request); err != nil {
		errResponse := responses.NewInvalidRequestBodyErrorResponse(err)
		return c.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	// Call CreateShipment rpc
	_, err := s.client.CreateShipment(c.Context(), &request)
	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	// Send response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully made the shipment"})
}

func (s ShipmentService) Bootstrap() {
	shipment := s.router.Group("/shipments")

	shipment.Post("/", s.createShipment)
}
