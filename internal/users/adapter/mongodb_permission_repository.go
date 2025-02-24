package adapter

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PermissionRepository struct {
	collection *mongo.Collection
}

func NewPermissionRepository(db *mongo.Database) *PermissionRepository {
	return &PermissionRepository{collection: db.Collection("permissions")}
}

func (p *PermissionRepository) FindPermission(ctx context.Context, id primitive.ObjectID) (*user.Permission, error) {
	// Find permission by id
	var permission user.Permission
	err := p.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&permission)
	if err != nil {
		return nil, cuserr.MongoError(err)
	}

	return &permission, nil
}
