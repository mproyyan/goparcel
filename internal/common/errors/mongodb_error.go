package cuserr

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MongoError(err error) error {
	if err == nil {
		return nil
	}

	// Document not found
	if errors.Is(err, mongo.ErrNoDocuments) {
		return status.Errorf(codes.NotFound, "document not found: %v", err.Error())
	}

	// Handle MongoDB Command Errors (Transaction Error)
	var commandErr mongo.CommandError
	if errors.As(err, &commandErr) {
		if commandErr.HasErrorLabel("TransientTransactionError") {
			return status.Errorf(codes.Internal, "transaction aborted: %v", err)
		}
		if commandErr.HasErrorLabel("UnknownTransactionCommitResult") {
			return status.Errorf(codes.Internal, "transaction commit result unknown, possible rollback: %v", err)
		}
	}

	// Handle Write Errors during Transactions
	var writeErr mongo.WriteException
	if errors.As(err, &writeErr) {
		for _, we := range writeErr.WriteErrors {
			if we.Code == 112 { // WriteConflict
				return status.Errorf(codes.AlreadyExists, "transaction rollback due to write conflict: %v", err)
			}
		}
	}

	return status.Errorf(codes.Internal, "unexpected database error: %v", err)
}
