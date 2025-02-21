package domain

import "time"

type ActivityType int

// Daftar konstanta ActivityType menggunakan iota
const (
	Unknown ActivityType = iota
	Receive
	Transit
	Load
	Arrive
	Unload
)

func (a ActivityType) String() string {
	switch a {
	case Receive:
		return "receive"
	case Transit:
		return "transit"
	case Load:
		return "load"
	case Arrive:
		return "arrive"
	case Unload:
		return "unload"
	default:
		return "unknown"
	}
}

func StringToActivityType(s string) ActivityType {
	switch s {
	case "receive":
		return Receive
	case "transit":
		return Transit
	case "load":
		return Load
	case "arrive":
		return Arrive
	case "unload":
		return Unload
	default:
		return Unknown
	}
}

type ItineraryLog struct {
	ActivityType string
	Timestamp    time.Time
	Location     Location
}
