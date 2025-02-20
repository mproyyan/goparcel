package requests

type CreateShipmentRequest struct {
	Origin    string    `json:"origin"`
	Sender    Entity    `json:"sender"`
	Recipient Entity    `json:"recipient"`
	Packages  []Package `json:"packages"`
}

type Entity struct {
	Name          string `json:"name"`
	PhoneNumber   string `json:"phone_number"`
	ZipCode       string `json:"zip_code"`
	StreetAddress string `json:"street_address"`
}

type Package struct {
	Name   string `json:"name"`
	Amount int32  `json:"amount"`
	Weight int32  `json:"weight"`
	Volume Volume `json:"volume"`
}

type Volume struct {
	Length int32 `json:"length"`
	Width  int32 `json:"width"`
	Height int32 `json:"height"`
}
