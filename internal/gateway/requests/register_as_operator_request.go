package requests

type RegisterAsOperatorRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	LocationID string `json:"location_id"`
	Type       string `json:"type"`
}
