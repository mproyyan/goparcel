package domain

type Item struct {
	Name   string
	Amount int
	Weight int32
	Volume int32
}

type Entity struct {
	Name    string
	Contact string
	Address Address
}

type Address struct {
	Province      string
	City          string
	District      string
	Subdistrict   string
	StreetAddress string
	ZipCode       string
}

type Shipment struct {
	ID              string
	AirwayBill      string
	TransportStatus TransportStatus
	RoutingStatus   RoutingStatus
	Items           []Item
	Sender          Entity
	Recipient       Entity
	Origin          string
	Destination     string
	ItineraryLogs   []ItineraryLog
}

type TransportStatus int

const (
	TransportUnknown TransportStatus = iota
	NotReceive
	InPort
	OnBoardCourier
	OnBoardCargo
	Claimed
)

func (t TransportStatus) String() string {
	switch t {
	case NotReceive:
		return "not_receive"
	case InPort:
		return "in_port"
	case OnBoardCourier:
		return "on_board_courier"
	case OnBoardCargo:
		return "on_board_cargo"
	case Claimed:
		return "claimed"
	default:
		return "unknown"
	}
}

func StringToTransportStatus(s string) TransportStatus {
	switch s {
	case "not_receive":
		return NotReceive
	case "in_port":
		return InPort
	case "on_board_courier":
		return OnBoardCourier
	case "on_board_cargo":
		return OnBoardCargo
	case "claimed":
		return Claimed
	default:
		return TransportUnknown
	}
}

type RoutingStatus int

const (
	RoutingUnknown RoutingStatus = iota
	NotRouted
	Routed
	MisRouted
)

func (r RoutingStatus) String() string {
	switch r {
	case NotRouted:
		return "not_routed"
	case Routed:
		return "routed"
	case MisRouted:
		return "mis_routed"
	default:
		return "unknown"
	}
}

func StringToRoutingStatus(s string) RoutingStatus {
	switch s {
	case "not_routed":
		return NotRouted
	case "routed":
		return Routed
	case "mis_routed":
		return MisRouted
	default:
		return RoutingUnknown
	}
}

func (s Shipment) LatestItinerary() ItineraryLog {
	index := len(s.ItineraryLogs) - 1
	return s.ItineraryLogs[index]
}
