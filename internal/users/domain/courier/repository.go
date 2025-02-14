package courier

import "context"

type CourierRepository interface {
	CreateCourier(ctx context.Context, courier Courier) (string, error)
}
