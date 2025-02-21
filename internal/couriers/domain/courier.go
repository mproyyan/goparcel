package domain

type Courier struct {
	ID         string
	UserID     string
	Name       string
	Email      string
	Status     CourierStatus
	LocationID string
}

type CourierStatus int

const (
	NotAvailable CourierStatus = iota
	Available
)

func (c CourierStatus) String() string {
	switch c {
	case Available:
		return "available"
	default:
		return "not_available"
	}
}

func StringToCourierStatus(s string) CourierStatus {
	switch s {
	case "available":
		return Available
	default:
		return NotAvailable
	}
}
