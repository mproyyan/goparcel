//go:generate mockgen -source=./repository.go -destination=../../mock/mock_operator.go -package=mock

package operator

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OperatorRepository interface {
	CreateOperator(ctx context.Context, operator Operator) (string, error)
	GetOperators(ctx context.Context, ids []primitive.ObjectID) ([]*Operator, error)
}
