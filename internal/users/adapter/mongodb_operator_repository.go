package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/utils"
	"github.com/mproyyan/goparcel/internal/users/domain/operator"
	cuserr "github.com/mproyyan/goparcel/internal/users/errors"
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
	// Prepare data to insert
	operatorModel, err := domainToOperatorModel(operator)
	if err != nil {
		return "", err
	}

	// Create new operator
	result, err := o.collection.InsertOne(ctx, operatorModel)
	if err != nil {
		return "", cuserr.MongoError(err)
	}

	// Return inserted id
	insertedId := result.InsertedID.(primitive.ObjectID)
	return insertedId.Hex(), nil
}

// Models
type OperatorModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty"`
	Type       string             `bson:"type"`
	Name       string             `bson:"name"`
	Email      string             `bson:"email"`
	LocationID primitive.ObjectID `bson:"location_id,omitempty"`
}

// Helper function to convert domain to operator model
func domainToOperatorModel(operator operator.Operator) (*OperatorModel, error) {
	// Convert string ObjectId to literal ObjectId
	id, err := utils.ConvertToObjectId(operator.ID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	userID, err := utils.ConvertToObjectId(operator.UserID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	locationID, err := utils.ConvertToObjectId(operator.LocationID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	return &OperatorModel{
		ID:         id,
		UserID:     userID,
		Type:       operator.Type.String(),
		Name:       operator.Name,
		Email:      operator.Email,
		LocationID: locationID,
	}, nil
}
