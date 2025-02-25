package domain

import "time"

type RequestType int

const (
	RequestTypeUnknown RequestType = iota
	RequestTypeTransit
	RequestTypeShipment
	RequestTypeDelivery
)

func ParseRequestType(value string) RequestType {
	switch value {
	case "unknown":
		return RequestTypeUnknown
	case "transit":
		return RequestTypeTransit
	case "shipment":
		return RequestTypeShipment
	case "delivery":
		return RequestTypeDelivery
	default:
		return RequestTypeUnknown
	}
}

func (r RequestType) String() string {
	switch r {
	case RequestTypeUnknown:
		return "unknown"
	case RequestTypeTransit:
		return "transit"
	case RequestTypeShipment:
		return "shipment"
	case RequestTypeDelivery:
		return "delivery"
	default:
		return "unknown"
	}
}

type Status int

const (
	StatusPending Status = iota
	StatusCompleted
)

func ParseStatus(value string) Status {
	switch value {
	case "pending":
		return StatusPending
	case "completed":
		return StatusCompleted
	default:
		return StatusPending
	}
}

func (s Status) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusCompleted:
		return "completed"
	default:
		return "pending"
	}
}

type TransferRequest struct {
	ID          string
	RequestType RequestType
	ShipmentID  string
	Origin      Origin
	Destination Destination
	CourierID   string
	CargoID     string
	Status      Status
	CreatedAt   time.Time
}

type Origin struct {
	Location    string
	RequestedBy string
}

type Destination struct {
	Location        string
	AcceptedBy      string
	RecipientDetail Entity
}

func (t TransferRequest) IsPending() bool {
	return t.Status == StatusPending
}
