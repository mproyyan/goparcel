package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/cargos/domain"
	"github.com/mproyyan/goparcel/internal/common/db"
	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Models
type CarrierModel struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty"`
	UserID     primitive.ObjectID  `bson:"user_id,omitempty"`
	Name       string              `bson:"name"`
	Email      string              `bson:"email"`
	CargoID    *primitive.ObjectID `bson:"cargo_id,omitempty"`
	Status     string              `bson:"status"`
	LocationID *primitive.ObjectID `bson:"location_id,omitempty"`
}

// Implementation

type CarrierRepository struct {
	collection *mongo.Collection
}

func NewCarrierRepository(db *mongo.Database) *CarrierRepository {
	return &CarrierRepository{collection: db.Collection("carriers")}
}

func (c *CarrierRepository) GetCarrier(ctx context.Context, id primitive.ObjectID) (*domain.Carrier, error) {
	var carrier *CarrierModel
	filter := bson.M{"_id": id}
	err := c.collection.FindOne(ctx, filter).Decode(&carrier)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	return carrierModelToDomain(carrier), nil
}

func (c *CarrierRepository) GetIdleCarriers(ctx context.Context, locationId primitive.ObjectID) ([]*domain.Carrier, error) {
	var carriers []*CarrierModel
	filter := bson.M{
		"location_id": locationId,
		"status":      domain.Idle.String(),
	}

	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var carrier CarrierModel
		if err := cursor.Decode(&carrier); err != nil {
			return nil, cuserr.MongoError(err)
		}
		carriers = append(carriers, &carrier)
	}

	if err := cursor.Err(); err != nil {
		return nil, cuserr.MongoError(err)
	}

	return carriersModelsTodomain(carriers), nil
}

func (c *CarrierRepository) ClearAssignedCargo(ctx context.Context, carrierIds []primitive.ObjectID) error {
	if len(carrierIds) == 0 {
		return nil // No carriers to clear
	}

	filter := bson.M{"_id": bson.M{"$in": carrierIds}}
	update := bson.M{"$set": bson.M{"cargo_id": nil}}

	_, err := c.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return cuserr.MongoError(err)
	}

	return nil
}

// Helper functions
func carrierModelToDomain(model *CarrierModel) *domain.Carrier {
	return &domain.Carrier{
		ID:         model.ID.Hex(),
		UserID:     model.UserID.Hex(),
		Name:       model.Name,
		Email:      model.Email,
		CargoID:    db.ObjectIdToString(model.CargoID),
		Status:     domain.StringToCarrierStatus(model.Status),
		LocationID: db.ObjectIdToString(model.LocationID),
	}
}

func carriersModelsTodomain(models []*CarrierModel) []*domain.Carrier {
	var carriers []*domain.Carrier
	for _, model := range models {
		op := carrierModelToDomain(model)
		carriers = append(carriers, op)
	}

	return carriers
}
