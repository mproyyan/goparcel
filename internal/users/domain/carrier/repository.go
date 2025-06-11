//go:generate mockgen -source=./repository.go -destination=../../mock/mock_carrier.go -package=mock

package carrier

import (
	"context"
)

type CarrierRepository interface {
	CreateCarrier(ctx context.Context, carrier Carrier) (string, error)
	GetCarriers(ctx context.Context, ids []string) ([]*Carrier, error)
}
