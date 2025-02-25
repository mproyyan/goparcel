package middlewares

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func SkipAuthDirectiveHandler(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	// Just return without authentication
	return next(ctx)
}
