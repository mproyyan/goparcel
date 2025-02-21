package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/couriers/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourierRepository struct {
	collection *mongo.Collection
}

func NewCourierRepository(db *mongo.Database) *CourierRepository {
	return &CourierRepository{
		collection: db.Collection("couriers"),
	}
}

func (c *CourierRepository) AvailableCouriers(ctx context.Context, locationID primitive.ObjectID) ([]domain.Courier, error) {
	filter := bson.M{"location_id": locationID, "status": "available"}
	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var couriers []CourierModel
	if err := cursor.All(ctx, &couriers); err != nil {
		return nil, err
	}

	return couriersModelToDomain(couriers), nil
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

func couriersModelToDomain(models []CourierModel) []domain.Courier {
	couriers := make([]domain.Courier, len(models))
	for i, model := range models {
		var userID, locationID string

		if model.UserID != nil {
			userID = model.UserID.Hex()
		}

		if model.LocationID != nil {
			locationID = model.LocationID.Hex()
		}

		couriers[i] = domain.Courier{
			ID:         model.ID.Hex(),
			UserID:     userID,
			Name:       model.Name,
			Email:      model.Email,
			Status:     domain.StringToCourierStatus(model.Status),
			LocationID: locationID,
		}
	}
	return couriers
}
