package errors

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInvalidCredentials = status.Error(codes.Unauthenticated, "invalid credentials")
var ErrUserNotFound = status.Error(codes.NotFound, "user not found")
var ErrInvalidOperatorType = status.Error(codes.InvalidArgument, "invalid operator type, must be depot_operator or warehouse_operator")
var ErrInvalidHexString = status.Error(codes.InvalidArgument, "invalid hex string")

func MongoError(err error) error {
	if err == nil {
		return nil
	}

	// Document not found
	if errors.Is(err, mongo.ErrNoDocuments) {
		return status.Errorf(codes.NotFound, "document not found")
	}

	return status.Errorf(codes.Internal, "unexpected database error: %v", err)
}
