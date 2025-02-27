package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CargoRepository interface {
	FindMatchingCargos(ctx context.Context, origin, destination primitive.ObjectID) ([]*Cargo, error)
}
