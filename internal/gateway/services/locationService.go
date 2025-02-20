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

func (l LocationService) findLocations(c *fiber.Ctx) error {
	// Get id from url parameter
	locationID := c.Params("id")

	// Call GetLocation RPC
	location, err := l.client.GetLocation(c.Context(), &locations.GetLocationRequest{
		LocationID: locationID,
	})

	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	// Build response
	locationResponse := responses.LocationResponse{
		Name:        location.Name,
		Type:        location.Type,
		WarehouseID: location.WarehouseId,
		Address: responses.Address{
			Province:      location.Address.Province,
			City:          location.Address.City,
			District:      location.Address.District,
			Subdistrict:   location.Address.Subdistrict,
			Latitude:      location.Address.Latitude,
			Longitude:     location.Address.Longitude,
			StreetAddress: location.Address.StreetAddress,
			ZipCode:       location.Address.ZipCode,
		},
	}

	return c.Status(fiber.StatusOK).JSON(locationResponse)
}

func (l LocationService) getRegion(c *fiber.Ctx) error {
	// Get zipcode from url parameter
	zipcode := c.Params("zipcode")

	region, err := l.client.GetRegion(c.Context(), &locations.GetRegionRequest{
		Zipcode: zipcode,
	})

	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	return c.Status(fiber.StatusOK).JSON(responses.RegionResponse{
		ZipCode:     region.ZipCode,
		Province:    region.Province,
		City:        region.City,
		District:    region.District,
		Subdistrict: region.Subdistrict,
	})
}

func (l LocationService) Bootstrap() {
	location := l.router.Group("/locations")

	location.Post("/", l.createLocation)
	location.Get("/:id", l.findLocations)
	location.Get("/region/:zipcode", l.getRegion)
}
