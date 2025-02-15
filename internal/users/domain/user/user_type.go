package user

import (
	"encoding/json"
)

type UserType struct {
	ID           string
	Name         string
	Description  string
	PermissionID string
	Permissions  Permissions
}

type Permissions []string

// Implement encoding.BinaryMarshaler
func (p Permissions) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}
