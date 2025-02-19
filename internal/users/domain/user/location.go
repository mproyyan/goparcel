package user

type Location struct {
	ID          string
	Name        string
	Type        string
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
