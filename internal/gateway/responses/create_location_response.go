package responses

type Address struct {
	Province      string  `json:"province"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	Subdistrict   string  `json:"subdistrict,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	StreetAddress string  `json:"street_address,omitempty"`
	ZipCode       string  `json:"zip_code"`
}

type LocationResponse struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	WarehouseID string  `json:"warehouse_id,omitempty"`
	Address     Address `json:"address"`
}
