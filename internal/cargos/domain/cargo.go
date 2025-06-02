package domain

import "time"

type Cargo struct {
	ID                string
	Name              string
	Status            string
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

func (c Cargo) HasShipments() bool {
	return len(c.Shipments) > 0
}
