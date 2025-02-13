package user

type User struct {
	ID       string
	ModelID  string
	Email    string
	Password string
	Type     UserType
}
