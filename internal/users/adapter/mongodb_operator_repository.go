package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/users/domain/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return "", cuserr.Decorate(err, "failed to convert domain to operator model")
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

func (o *OperatorRepository) GetOperators(ctx context.Context, ids []primitive.ObjectID) ([]*operator.Operator, error) {
	filter := bson.M{}

	// If ids not empty then fetch operator based on the ids
	if len(ids) > 0 {
		filter["_id"] = bson.M{"$in": ids}
	}

	cursor, err := o.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var opreatorModel []OperatorModel
	if err := cursor.All(ctx, &opreatorModel); err != nil {
		return nil, err
	}

	return operatorModelsTodomain(opreatorModel), nil
}

// Models
type OperatorModel struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty"`
	UserID     primitive.ObjectID  `bson:"user_id,omitempty"`
	Type       string              `bson:"type"`
	Name       string              `bson:"name"`
	Email      string              `bson:"email"`
	LocationID *primitive.ObjectID `bson:"location_id,omitempty"`
}

// Helper function to convert domain to operator model
func domainToOperatorModel(operator operator.Operator) (*OperatorModel, error) {
	// Convert string ObjectId to literal ObjectId
	id, err := db.ConvertToObjectId(operator.ID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not valid object id")
	}

	userID, err := db.ConvertToObjectId(operator.UserID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	locationID, err := db.ConvertToObjectId(operator.LocationID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	return &OperatorModel{
		ID:         id,
		UserID:     userID,
		Type:       operator.Type.String(),
		Name:       operator.Name,
		Email:      operator.Email,
		LocationID: &locationID,
	}, nil
}

func operatorModelToDomain(model OperatorModel) *operator.Operator {
	return &operator.Operator{
		ID:         model.ID.Hex(),
		UserID:     model.UserID.Hex(),
		Type:       operator.OperatorTypeFromString(model.Type),
		Name:       model.Name,
		Email:      model.Email,
		LocationID: model.LocationID.Hex(),
	}
}

func operatorModelsTodomain(models []OperatorModel) []*operator.Operator {
	var operators []*operator.Operator
	for _, model := range models {
		op := operatorModelToDomain(model)
		operators = append(operators, op)
	}

	return operators
}
