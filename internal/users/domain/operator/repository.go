package operator

import "context"

type OperatorRepository interface {
	CreateOperator(ctx context.Context, operator Operator) (string, error)
}
