package domain

type Carrier struct {
	ID         string
	UserID     string
	Name       string
	Email      string
	CargoID    string
	Status     CarrierStatus
	LocationID string
}

type CarrierStatus int

const (
	UnknownCarrierStatus CarrierStatus = iota
	Idle
	Active
)

func (c CarrierStatus) String() string {
	switch c {
	case Idle:
		return "idle"
	case Active:
		return "active"
	default:
		return "unknown"
	}
}

func StringToCarrierStatus(s string) CarrierStatus {
	switch s {
	case "idle":
		return Idle
	case "active":
		return Active
	default:
		return UnknownCarrierStatus
	}
}
