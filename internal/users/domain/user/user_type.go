package user

type UserType struct {
	ID           string
	Name         string
	Description  string
	PermissionID string
	Permissions  []string
}
