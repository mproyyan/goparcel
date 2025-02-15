package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrDepotType = status.Error(codes.InvalidArgument, "depot type location must contain warehouse id")
var ErrWarehouseType = status.Error(codes.InvalidArgument, "warehouse type location cannot contain another warehouse id")
