package operator

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
	OperatorUnknown OperatorType = iota
	DepotOperator
	WarehouseOperator
)

func (o OperatorType) String() string {
	switch o {
	case DepotOperator:
		return "depot_operator"
	case WarehouseOperator:
		return "warehouse_operator"
	}

	return "unknown"
}

func OperatorTypeFromString(stringType string) OperatorType {
	switch stringType {
	case "depot_operator":
		return DepotOperator
	case "warehouse_operator":
		return WarehouseOperator
	}

	return OperatorUnknown
}
