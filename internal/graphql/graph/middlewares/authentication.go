package middlewares

import (
	"context"
	"errors"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mproyyan/goparcel/internal/common/auth"
)

type contextKey string

const AuthUserKey contextKey = "user"

type AuthUser struct {
	UserID  string
	ModelID string
}

// AuthMiddleware is a GraphQL middleware that checks for authentication unless @skipAuth is applied
func AuthMiddleware() graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		opCtx := graphql.GetOperationContext(ctx)

		// Operation list that doesn't need authentication
		skippedOperations := map[string]bool{
			"Login":              true,
			"RegisterAsOperator": true,
			"RegisterAsCourier":  true,
			"RegisterAsCarrier":  true,
		}

		// Check if operation name on the list
		if skippedOperations[opCtx.OperationName] {
			return next(ctx)
		}

		// Extract Authorization header
		authHeader := opCtx.Headers.Get("Authorization")
		if authHeader == "" {
			return func(ctx context.Context) *graphql.Response {
				return graphql.ErrorResponse(ctx, "Missing token")
			}
		}

		// Validate token format
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return func(ctx context.Context) *graphql.Response {
				return graphql.ErrorResponse(ctx, "Invalid token format")
			}
		}

		// Validate token
		claims, err := auth.Authenticate(tokenParts[1])
		if err != nil {
			return func(ctx context.Context) *graphql.Response {
				return graphql.ErrorResponse(ctx, "Invalid token")
			}
		}

		// Store user data in context
		ctx = context.WithValue(ctx, AuthUserKey, AuthUser{UserID: claims.UserID, ModelID: claims.ModelID})
		return next(ctx)
	}
}

// GetAuthUser retrieves authenticated user from context
func GetAuthUser(ctx context.Context) (*AuthUser, error) {
	user, ok := ctx.Value(AuthUserKey).(AuthUser)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	return &user, nil
}
