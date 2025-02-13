package errors

import "errors"

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidOperatorType = errors.New("invalid operator type, must be depot_operator or warehouse_operator")
