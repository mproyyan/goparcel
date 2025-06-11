//go:generate mockgen -source repository.go -destination ../mock/mock_repository.go -package mock

package domain

import (
	"context"
)

type CourierRepository interface {
	AvailableCouriers(ctx context.Context, locationID string) ([]Courier, error)
}
