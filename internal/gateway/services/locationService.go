package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/gateway/responses"
)

type LocationService struct {
	router fiber.Router
	client genproto.LocationServiceClient
}

func NewLocationService(router fiber.Router, client genproto.LocationServiceClient) LocationService {
	return LocationService{
		router: router,
		client: client,
	}
}

func (l LocationService) createLocation(c *fiber.Ctx) error {
	// Parse request body
	var request genproto.CreateLocationRequest
	if err := c.BodyParser(&request); err != nil {
		errResponse := responses.NewInvalidRequestBodyErrorResponse(err)
		return c.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	// Call CreateLocation RPC
	_, err := l.client.CreateLocation(c.Context(), &request)
	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	// Send response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "location has been successfully created",
	})
}

func (l LocationService) findLocations(c *fiber.Ctx) error {
	// Get id from url parameter
	locationID := c.Params("id")

	// Call GetLocation RPC
	location, err := l.client.GetLocation(c.Context(), &genproto.GetLocationRequest{LocationId: locationID})
	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	// Send response
	return c.Status(fiber.StatusOK).JSON(location)
}

func (l LocationService) getRegion(c *fiber.Ctx) error {
	// Get zipcode from url parameter
	zipcode := c.Params("zipcode")

	// Call GetRegion rpc
	region, err := l.client.GetRegion(c.Context(), &genproto.GetRegionRequest{Zipcode: zipcode})
	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	// Send response
	return c.Status(fiber.StatusOK).JSON(region)
}

func (l LocationService) Bootstrap() {
	location := l.router.Group("/locations")

	location.Post("/", l.createLocation)
	location.Get("/:id", l.findLocations)
	location.Get("/region/:zipcode", l.getRegion)
}
