package domain

type Location struct {
	ID          string
	Name        string
	Type        LocationType
	WarehouseID string
	Address     Address
}

type Address struct {
	ZipCode       string
	Province      string
	City          string
	District      string
	Subdistrict   string
	Latitude      float64
	Longitude     float64
	StreetAddress string
}

type LocationType int

const (
	Depot LocationType = iota + 1
	Warehouse
)

func (l LocationType) String() string {
	switch l {
	case Depot:
		return "depot"
	case Warehouse:
		return "warehouse"
	}

	return ""
}

// Check is location a depot or not by the type
func (l Location) IsDepot() bool {
	return l.Type == Depot
}

// Check is location a depot or not by the type
func (l Location) IsWarehouse() bool {
	return l.Type == Warehouse
}

func (l Location) InvalidType() bool {
	return !l.IsDepot() && !l.IsWarehouse()
}

func LocationTypeFromString(locationType string) LocationType {
	switch locationType {
	case "depot":
		return Depot
	case "warehouse":
		return Warehouse
	}

	return 0
}
