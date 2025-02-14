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
		return err
	}

	// Call user service grpc
	response, err := u.client.Login(c.Context(), &users.LoginRequest{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	})

	if err != nil {
		return err
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
		return err
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "registration as an operator has been successful",
	})
}

func (u UserService) registerAsCarrier(c *fiber.Ctx) error {
	// Parse request body
	var registerAsCarrierRequest requests.RegisterAsCarrierRequest
	if err := c.BodyParser(&registerAsCarrierRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	// Call user service grpc
	_, err := u.client.RegisterAsCarrier(c.Context(), &users.RegisterAsCarrierRequest{
		Name:     registerAsCarrierRequest.Name,
		Email:    registerAsCarrierRequest.Email,
		Password: registerAsCarrierRequest.Password,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "registration as a carrier has been successful",
	})
}

func (u UserService) registerAsCourier(c *fiber.Ctx) error {
	// Parse request body
	var registerAsCourierRequest requests.RegisterAsCourierRequest
	if err := c.BodyParser(&registerAsCourierRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	// Call user service grpc
	_, err := u.client.RegisterAsCourier(c.Context(), &users.RegisterAsCourierRequest{
		Name:     registerAsCourierRequest.Name,
		Email:    registerAsCourierRequest.Email,
		Password: registerAsCourierRequest.Password,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "registration as a courier has been successful",
	})
}

func (u UserService) Bootstrap() {
	user := u.router.Group("/users")

	user.Post("/login", u.login)
	user.Post("/register-as-operator", u.registerAsOperator)
	user.Post("/register-as-carrier", u.registerAsCarrier)
	user.Post("/register-as-courier", u.registerAsCourier)
}
