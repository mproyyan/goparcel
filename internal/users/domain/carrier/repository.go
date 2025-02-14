package carrier

import "context"

type CarrierRepository interface {
	CreateCarrier(ctx context.Context, carrier Carrier) (string, error)
}
