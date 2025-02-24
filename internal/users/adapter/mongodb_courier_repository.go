package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/users/domain/courier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CourierRepository struct {
	collection *mongo.Collection
}

func NewCourierRepository(db *mongo.Database) *CourierRepository {
	return &CourierRepository{
		collection: db.Collection("couriers"),
	}
}

func (c *CourierRepository) CreateCourier(ctx context.Context, courier courier.Courier) (string, error) {
	// Prepare data to insert
	courierModel, err := domainToCourierModel(courier)
	if err != nil {
		return "", cuserr.Decorate(err, "failed to convert domain to courier model")
	}

	// Create new carrier
	result, err := c.collection.InsertOne(ctx, courierModel)
	if err != nil {
		return "", cuserr.MongoError(err)
	}

	// Return inserted id
	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID.Hex(), nil
}

func (c *CourierRepository) GetCouriers(ctx context.Context, ids []primitive.ObjectID) ([]*courier.Courier, error) {
	filter := bson.M{}

	// If ids not empty then fetch courier based on the ids
	if len(ids) > 0 {
		filter["_id"] = bson.M{"$in": ids}
	}

	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var courierModel []CourierModel
	if err := cursor.All(ctx, &courierModel); err != nil {
		return nil, err
	}

	return couriersModelsTodomain(courierModel), nil
}

// Models
type CourierModel struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty"`
	UserID     *primitive.ObjectID `bson:"user_id,omitempty"`
	Name       string              `bson:"name"`
	Email      string              `bson:"email"`
	Status     string              `bson:"status"`
	LocationID *primitive.ObjectID `bson:"location_id,omitempty"`
}

// Helper function to convert domain to model
func domainToCourierModel(courier courier.Courier) (*CourierModel, error) {
	// Convert string ObjectId to literal ObjectId
	courierID, err := db.ConvertToObjectId(courier.ID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not valid object id")
	}

	userID, err := db.ConvertToObjectId(courier.UserID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_id is not valid object id")
	}

	locationID, err := db.ConvertToObjectId(courier.LocationID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "location_id is not valid object id")
	}

	return &CourierModel{
		ID:         courierID,
		UserID:     &userID,
		Name:       courier.Name,
		Email:      courier.Email,
		Status:     courier.Status.String(), // default not_available
		LocationID: &locationID,
	}, nil
}

func courierModelToDomain(model CourierModel) *courier.Courier {
	return &courier.Courier{
		ID:         model.ID.Hex(),
		UserID:     model.UserID.Hex(),
		Name:       model.Name,
		Email:      model.Email,
		Status:     courier.StringToCourierStatus(model.Status),
		LocationID: model.LocationID.Hex(),
	}
}

func couriersModelsTodomain(models []CourierModel) []*courier.Courier {
	var couriers []*courier.Courier
	for _, model := range models {
		op := courierModelToDomain(model)
		couriers = append(couriers, op)
	}

	return couriers
}
