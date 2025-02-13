package operator

import cuserr "github.com/mproyyan/goparcel/internal/users/errors"

type Operator struct {
	ID         string
	UserID     string
	Type       OperatorType
	Name       string
	Email      string
	LocationID string
}

type OperatorType int

const (
	DepotOperator OperatorType = iota + 1
	WarehouseOperator
)

func (o OperatorType) String() string {
	switch o {
	case DepotOperator:
		return "depot_operator"
	case WarehouseOperator:
		return "warehouse_operator"
	}

	return ""
}

func OperatorTypeFromString(stringType string) (OperatorType, error) {
	switch stringType {
	case "depot_operator":
		return DepotOperator, nil
	case "warehouse_operator":
		return WarehouseOperator, nil
	}

	return 0, cuserr.ErrInvalidOperatorType
}
