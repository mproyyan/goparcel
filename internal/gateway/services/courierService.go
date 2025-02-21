package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/gateway/middlewares"
	"github.com/mproyyan/goparcel/internal/gateway/responses"
)

type CourierService struct {
	router fiber.Router
	client genproto.CourierServiceClient
}

func NewCourierService(router fiber.Router, client genproto.CourierServiceClient) CourierService {
	return CourierService{
		router: router,
		client: client,
	}
}

func (co CourierService) getAvailableCouriers(c *fiber.Ctx) error {
	// Get location from parameter
	request := genproto.GetAvailableCourierRequest{LocationId: c.Params("locationId")}

	// Call GetAvailableCouriers rps
	couriers, err := co.client.GetAvailableCouriers(c.Context(), &request)
	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	// Send reponse
	return c.Status(fiber.StatusOK).JSON(couriers)
}

func (co CourierService) Bootstrap() {
	courier := co.router.Group("/couriers")

	courier.Get("/available/:locationId", middlewares.AuthMiddleware(), co.getAvailableCouriers)
}
