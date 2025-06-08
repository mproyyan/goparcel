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
	CarrierIdle
	CarrierActive
)

func (c CarrierStatus) String() string {
	switch c {
	case CarrierIdle:
		return "idle"
	case CarrierActive:
		return "active"
	default:
		return "unknown"
	}
}

func StringToCarrierStatus(s string) CarrierStatus {
	switch s {
	case "idle":
		return CarrierIdle
	case "active":
		return CarrierActive
	default:
		return UnknownCarrierStatus
	}
}
