package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/users/domain/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OperatorRepository struct {
	collection *mongo.Collection
}

func NewOperatorRepository(db *mongo.Database) *OperatorRepository {
	return &OperatorRepository{
		collection: db.Collection("operators"),
	}
}

func (o *OperatorRepository) CreateOperator(ctx context.Context, operator operator.Operator) (string, error) {
	// Convert string ObjectId to literal ObjectId
	id, err := primitive.ObjectIDFromHex(operator.ID)
	if err != nil {
		return "", err
	}

	userID, err := primitive.ObjectIDFromHex(operator.UserID)
	if err != nil {
		return "", err
	}

	locationID, err := primitive.ObjectIDFromHex(operator.LocationID)
	if err != nil {
		return "", err
	}

	// Create new operator
	result, err := o.collection.InsertOne(ctx, bson.M{
		"_id":         id,
		"user_id":     userID,
		"type":        operator.Type.String(),
		"name":        operator.Name,
		"email":       operator.Email,
		"location_id": locationID,
	})

	if err != nil {
		return "", err
	}

	// Return inserted id
	return result.InsertedID.(string), nil
}
