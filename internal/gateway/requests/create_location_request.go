package requests

type CreateLocationRequest struct {
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	WarehouseID   string  `json:"warehouse_id"`
	ZipCode       int     `json:"zip_code"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	StreetAddress string  `json:"street_address"`
}
