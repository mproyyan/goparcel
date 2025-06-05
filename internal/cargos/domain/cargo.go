package domain

import "time"

type Cargo struct {
	ID                string
	Name              string
	Status            CargoStatus
	MaxCapacity       Capacity
	CurrentLoad       Capacity
	Carriers          []string
	Itineraries       []Itinerary
	Shipments         []string
	LastKnownLocation string
}

type Capacity struct {
	Weight float64
	Volume float64
}

type Itinerary struct {
	Location             string
	EstimatedTimeArrival time.Time
	ActualTimeArrival    *time.Time
}

type CargoStatus int

const (
	UnknownCargoStatus CargoStatus = iota
	CargoIdle
	CargoActive
)

func (c CargoStatus) String() string {
	switch c {
	case CargoIdle:
		return "idle"
	case CargoActive:
		return "active"
	default:
		return "unknown"
	}
}

func StringToCargoStatus(s string) CargoStatus {
	switch s {
	case "idle":
		return CargoIdle
	case "active":
		return CargoActive
	default:
		return UnknownCargoStatus
	}
}

func (c Cargo) HasShipments() bool {
	return len(c.Shipments) > 0
}
