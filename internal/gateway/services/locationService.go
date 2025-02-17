package services

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mproyyan/goparcel/internal/common/genproto/locations"
	"github.com/mproyyan/goparcel/internal/gateway/requests"
	"github.com/mproyyan/goparcel/internal/gateway/responses"
)

type LocationService struct {
	router fiber.Router
	client locations.LocationServiceClient
}

func NewLocationService(router fiber.Router, client locations.LocationServiceClient) LocationService {
	return LocationService{
		router: router,
		client: client,
	}
}

func (l LocationService) createLocation(c *fiber.Ctx) error {
	// Parse request body
	var request requests.CreateLocationRequest
	if err := c.BodyParser(&request); err != nil {
		errResponse := responses.NewInvalidRequestBodyErrorResponse(err)
		return c.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	// Call CreateLocation RPC
	_, err := l.client.CreateLocation(c.Context(), &locations.CreateLocationRequest{
		Name:          request.Name,
		Type:          request.Type,
		WarehouseId:   request.WarehouseID,
		ZipCode:       strconv.Itoa(request.ZipCode),
		Latitude:      request.Latitude,
		Longitude:     request.Longitude,
		StreetAddress: request.StreetAddress,
	})

	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "location has been successfully created",
	})
}

func (l LocationService) Bootstrap() {
	location := l.router.Group("/locations")

	location.Post("/", l.createLocation)
}
