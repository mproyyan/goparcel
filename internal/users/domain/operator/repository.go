//go:generate mockgen -source=./repository.go -destination=../../mock/mock_operator.go -package=mock

package operator

import (
	"context"
)

type OperatorRepository interface {
	CreateOperator(ctx context.Context, operator Operator) (string, error)
	GetOperators(ctx context.Context, ids []string) ([]*Operator, error)
}
