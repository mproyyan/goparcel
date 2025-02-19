package middlewares

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/gateway/responses"
	"github.com/redis/go-redis/v9"
)

type Permissions []string

func (p Permissions) Empty() bool {
	return len(p) > 0
}

func PermissionMiddleware(client *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get auth user
		user := c.Locals("user")
		authUser, ok := user.(AuthUser)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Error:   fiber.ErrUnauthorized.Message,
				Message: "User not authenticated",
			})
		}

		// Get permission from cache
		permissionString, err := client.Get(c.Context(), authUser.UserID).Result()
		if err == redis.Nil {
			// TODO: Call rpc to get user permission
			var permissions Permissions

			if permissions.Empty() {
				return c.Status(fiber.StatusForbidden).JSON(responses.ErrorResponse{
					Error:   fiber.ErrForbidden.Message,
					Message: "Missing permission",
				})
			}
		} else if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(responses.ErrorResponse{
				Error:   fiber.ErrForbidden.Message,
				Message: cuserr.Decorate(err, "failed to get permission from cache").Error(),
			})
		}

		// Parse permission
		permissions, err := parsePermission(permissionString)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(responses.ErrorResponse{
				Error:   fiber.ErrForbidden.Message,
				Message: cuserr.Decorate(err, "failed to parse permission").Error(),
			})
		}

		if permissions.Empty() {
			return c.Status(fiber.StatusForbidden).JSON(responses.ErrorResponse{
				Error:   fiber.ErrForbidden.Message,
				Message: cuserr.Decorate(err, "Missing permission").Error(),
			})
		}

		// Get required permissions
		key := fmt.Sprintf("%s:%s", c.Method(), c.Path())
		requiredPermissions, exists := endpointPermissions[key]
		if !exists {
			return c.Status(fiber.StatusForbidden).JSON(responses.ErrorResponse{
				Error:   fiber.ErrForbidden.Message,
				Message: "Access denied",
			})
		}

		if !hasAllPermissions(permissions, requiredPermissions) {
			return c.Status(fiber.StatusForbidden).JSON(responses.ErrorResponse{
				Error:   fiber.ErrForbidden.Message,
				Message: "Insufficient permissions",
			})
		}

		return c.Next()
	}
}

func parsePermission(permissionString string) (Permissions, error) {
	var permission Permissions
	if permissionString == "" {
		return permission, nil
	}

	err := json.Unmarshal([]byte(permissionString), &permission)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to unmarshal permission")
	}

	return permission, nil
}

func hasAllPermissions(userPermissions Permissions, requiredPermissions Permissions) bool {
	permissionSet := make(map[string]bool)
	for _, perm := range userPermissions {
		permissionSet[perm] = true
	}

	for _, requiredPerm := range requiredPermissions {
		if !permissionSet[requiredPerm] {
			return false
		}
	}

	return true
}

var endpointPermissions = map[string]Permissions{
	"POST:/api/shipments": {"shipment.create_shipment"},
}
