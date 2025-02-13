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

func (u UserService) Bootstrap() {
	user := u.router.Group("/users")

	user.Post("/login", u.login)
}
