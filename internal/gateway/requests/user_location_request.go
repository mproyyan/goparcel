package requests

type UserLocationRequest struct {
	UserID string `json:"user_id"`
	Entity string `json:"entity"`
}
