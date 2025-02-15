package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoTransactionManager mengimplementasikan interface TransactionManager untuk MongoDB.
type MongoTransactionManager struct {
	client *mongo.Client
}

func NewMongoTransactionManager(client *mongo.Client) *MongoTransactionManager {
	return &MongoTransactionManager{client: client}
}

func (m *MongoTransactionManager) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	session, err := m.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sc mongo.SessionContext) (interface{}, error) {
		return nil, fn(sc)
	}

	_, err = session.WithTransaction(ctx, callback)
	return err
}
