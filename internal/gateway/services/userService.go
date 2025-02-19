package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mproyyan/goparcel/internal/common/genproto/users"
	"github.com/mproyyan/goparcel/internal/gateway/requests"
	"github.com/mproyyan/goparcel/internal/gateway/responses"
)

type UserService struct {
	router fiber.Router
	client users.UserServiceClient
}

func NewUserService(router fiber.Router, client users.UserServiceClient) UserService {
	return UserService{
		router: router,
		client: client,
	}
}

func (u UserService) login(c *fiber.Ctx) error {
	// Parse request body
	var loginRequest = requests.LoginRequest{}
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewInvalidRequestBodyErrorResponse(err))
	}

	// Call user service grpc
	response, err := u.client.Login(c.Context(), &users.LoginRequest{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	})

	if err != nil {
		code, response := responses.NewErrorResponse(err)
		return c.Status(code).JSON(response)
	}

	// send response
	c.JSON(responses.LoginResponse{
		Token: response.Token,
	})

	return nil
}

func (u UserService) registerAsOperator(c *fiber.Ctx) error {
	// Parse request body
	var registerAsOperatorRequest requests.RegisterAsOperatorRequest
	if err := c.BodyParser(&registerAsOperatorRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewInvalidRequestBodyErrorResponse(err))
	}

	// Call user service
	_, err := u.client.RegisterAsOperator(c.Context(), &users.RegisterAsOperatorRequest{
		Name:     registerAsOperatorRequest.Name,
		Email:    registerAsOperatorRequest.Email,
		Password: registerAsOperatorRequest.Password,
		Location: registerAsOperatorRequest.LocationID,
		Type:     registerAsOperatorRequest.Type,
	})

	if err != nil {
		code, response := responses.NewErrorResponse(err)
		return c.Status(code).JSON(response)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "registration as an operator has been successful",
	})
}

func (u UserService) registerAsCarrier(c *fiber.Ctx) error {
	// Parse request body
	var registerAsCarrierRequest requests.RegisterAsCarrierRequest
	if err := c.BodyParser(&registerAsCarrierRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewInvalidRequestBodyErrorResponse(err))
	}

	// Call user service grpc
	_, err := u.client.RegisterAsCarrier(c.Context(), &users.RegisterAsCarrierRequest{
		Name:     registerAsCarrierRequest.Name,
		Email:    registerAsCarrierRequest.Email,
		Password: registerAsCarrierRequest.Password,
	})

	if err != nil {
		code, response := responses.NewErrorResponse(err)
		return c.Status(code).JSON(response)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "registration as a carrier has been successful",
	})
}

func (u UserService) registerAsCourier(c *fiber.Ctx) error {
	// Parse request body
	var registerAsCourierRequest requests.RegisterAsCourierRequest
	if err := c.BodyParser(&registerAsCourierRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewInvalidRequestBodyErrorResponse(err))
	}

	// Call user service grpc
	_, err := u.client.RegisterAsCourier(c.Context(), &users.RegisterAsCourierRequest{
		Name:     registerAsCourierRequest.Name,
		Email:    registerAsCourierRequest.Email,
		Password: registerAsCourierRequest.Password,
	})

	if err != nil {
		code, response := responses.NewErrorResponse(err)
		return c.Status(code).JSON(response)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "registration as a courier has been successful",
	})
}

func (u UserService) userLocation(c *fiber.Ctx) error {
	// Parse request body
	var request requests.UserLocationRequest
	if err := c.BodyParser(&request); err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	// Call GetUserLocation RPC
	location, err := u.client.GetUserLocation(c.Context(), &users.GetUserLocationRequest{
		UserId: request.UserID,
		Entity: request.Entity,
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

func (u UserService) Bootstrap() {
	user := u.router.Group("/users")

	// Auths
	user.Post("/login", u.login)
	user.Post("/register-as-operator", u.registerAsOperator)
	user.Post("/register-as-carrier", u.registerAsCarrier)
	user.Post("/register-as-courier", u.registerAsCourier)

	user.Post("/location", u.userLocation)
}
