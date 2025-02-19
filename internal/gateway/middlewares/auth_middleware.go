package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mproyyan/goparcel/internal/common/auth"
	"github.com/mproyyan/goparcel/internal/gateway/responses"
)

type AuthUser struct {
	UserID  string
	ModelID string
}

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Error:   fiber.ErrUnauthorized.Message,
				Message: "Missing token",
			})
		}

		// Get token from header
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Error:   fiber.ErrUnauthorized.Message,
				Message: "Invalid token format",
			})
		}

		// Validate token
		tokenString := tokenParts[1]
		claims, err := auth.Authenticate(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Error:   fiber.ErrUnauthorized.Message,
				Message: "Invalid token",
			})
		}

		// Build auth user and pass to the next handler
		c.Locals("user", AuthUser{UserID: claims.UserID, ModelID: claims.ModelID})
		return c.Next()
	}
}
