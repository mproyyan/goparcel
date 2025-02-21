package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/gateway/middlewares"
	"github.com/mproyyan/goparcel/internal/gateway/responses"
)

type ShipmentService struct {
	router fiber.Router
	client genproto.ShipmentServiceClient
}

func NewShipmentService(router fiber.Router, client genproto.ShipmentServiceClient) ShipmentService {
	return ShipmentService{
		router: router,
		client: client,
	}
}

func (s ShipmentService) createShipment(c *fiber.Ctx) error {
	// Parse request body
	var request genproto.CreateShipmentRequest
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

func (s ShipmentService) getUnroutedShipments(c *fiber.Ctx) error {
	// Get location id from param
	request := genproto.GetUnroutedShipmentRequest{LocationId: c.Params("locationId")}

	// Call GetUnroutedShipment rpc
	shipments, err := s.client.GetUnroutedShipment(c.Context(), &request)
	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	// Send response
	return c.Status(fiber.StatusOK).JSON(shipments)
}

func (s ShipmentService) Bootstrap() {
	shipment := s.router.Group("/shipments")

	shipment.Post("/", middlewares.AuthMiddleware(), s.createShipment)
	shipment.Get("/:locationId/unrouted", middlewares.AuthMiddleware(), s.getUnroutedShipments)
}
