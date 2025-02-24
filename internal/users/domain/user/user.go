package user

type User struct {
	ID         string
	ModelID    string
	Entity     string
	Email      string
	Password   string
	UserTypeID string
}

type UserEntityName int

const (
	Unknown = iota
	Operator
	Carrier
	Courier
)

func (u UserEntityName) String() string {
	switch u {
	case Operator:
		return "operator"
	case Carrier:
		return "carrier"
	case Courier:
		return "courier"
	}

	return ""
}

func StringToUserEntityName(entity string) UserEntityName {
	switch entity {
	case "operator":
		return Operator
	case "carrier":
		return Carrier
	case "courier":
		return Courier
	}

	return Unknown
}
