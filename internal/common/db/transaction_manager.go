package db

import "context"

// TransactionManager abstracts the execution of operations within transactions.
// The Execute function executes the fn callback with the transaction context, so
// If an error occurs, the transaction will be canceled (rollback).
type TransactionManager interface {
	Execute(ctx context.Context, fn func(ctx context.Context) error) error
}
