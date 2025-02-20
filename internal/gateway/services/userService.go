package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mproyyan/goparcel/internal/common/genproto/users"
	"github.com/mproyyan/goparcel/internal/gateway/middlewares"
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
	var request users.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewInvalidRequestBodyErrorResponse(err))
	}

	// Call user service grpc
	response, err := u.client.Login(c.Context(), &request)
	if err != nil {
		code, response := responses.NewErrorResponse(err)
		return c.Status(code).JSON(response)
	}

	// Send response
	return c.Status(fiber.StatusOK).JSON(response)
}

func (u UserService) registerAsOperator(c *fiber.Ctx) error {
	// Parse request body
	var request users.RegisterAsOperatorRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewInvalidRequestBodyErrorResponse(err))
	}

	// Call user service
	_, err := u.client.RegisterAsOperator(c.Context(), &request)
	if err != nil {
		code, response := responses.NewErrorResponse(err)
		return c.Status(code).JSON(response)
	}

	// Send response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "registration as an operator has been successful",
	})
}

func (u UserService) registerAsCarrier(c *fiber.Ctx) error {
	// Parse request body
	var request users.RegisterAsCarrierRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewInvalidRequestBodyErrorResponse(err))
	}

	// Call user service grpc
	_, err := u.client.RegisterAsCarrier(c.Context(), &request)
	if err != nil {
		code, response := responses.NewErrorResponse(err)
		return c.Status(code).JSON(response)
	}

	// Send response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "registration as a carrier has been successful",
	})
}

func (u UserService) registerAsCourier(c *fiber.Ctx) error {
	// Parse request body
	var request users.RegisterAsCourierRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewInvalidRequestBodyErrorResponse(err))
	}

	// Call user service grpc
	_, err := u.client.RegisterAsCourier(c.Context(), &request)
	if err != nil {
		code, response := responses.NewErrorResponse(err)
		return c.Status(code).JSON(response)
	}

	// Send response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "registration as a courier has been successful",
	})
}

func (u UserService) userLocation(c *fiber.Ctx) error {
	// Parse request body
	var request users.GetUserLocationRequest
	if err := c.BodyParser(&request); err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	// Call GetUserLocation RPC
	location, err := u.client.GetUserLocation(c.Context(), &request)
	if err != nil {
		code, errResponse := responses.NewErrorResponse(err)
		return c.Status(code).JSON(errResponse)
	}

	// Send response
	return c.Status(fiber.StatusOK).JSON(location)
}

func (u UserService) Bootstrap() {
	user := u.router.Group("/users")

	// Auths
	user.Post("/login", u.login)
	user.Post("/register-as-operator", u.registerAsOperator)
	user.Post("/register-as-carrier", u.registerAsCarrier)
	user.Post("/register-as-courier", u.registerAsCourier)

	user.Post("/location", middlewares.AuthMiddleware(), u.userLocation)
}
