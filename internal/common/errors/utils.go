package cuserr

import (
	"fmt"
)

func Decorate(err error, message string, args ...any) error {
	m := fmt.Sprintf(message, args...)
	return fmt.Errorf("%s, cause: %w", m, err)
}
