package adapter

import (
	"context"

	"github.com/mproyyan/goparcel/internal/users/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserTypeRepository struct {
	collection *mongo.Collection
}

func NewUserTypeRepository(db *mongo.Database) *UserTypeRepository {
	return &UserTypeRepository{
		collection: db.Collection("user_types"),
	}
}

func (u *UserTypeRepository) FindUserType(ctx context.Context, userType string) (*user.UserType, error) {
	// Find user type by name
	var userTypeResult UserType
	err := u.collection.FindOne(ctx, bson.D{
		{Key: "name", Value: userType},
	}).Decode(&userTypeResult)

	if err != nil {
		return nil, err
	}

	return UserTypeModelToDomain(&userTypeResult), nil
}

// models
type UserType struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	PermissionID primitive.ObjectID `bson:"permission_id"`
	Permission   Permission         `bson:"permission"`
}

type Permission struct {
	UserManagement struct {
		GetUserLocation bool `bson:"get_user_location"`
		GetUser         bool `bson:"get_user"`
	} `bson:"user_management"`

	Shipment struct {
		CreateShipment        bool `bson:"create_shipment"`
		GetUnroutedShipments  bool `bson:"get_unrouted_shipments"`
		GetRoutedShipments    bool `bson:"get_routed_shipments"`
		GetIncomingShipments  bool `bson:"get_incoming_shipments"`
		ScanArrivingShipments bool `bson:"scan_arriving_shipments"`
		ShipPackage           bool `bson:"ship_package"`
		AddItineraryHistory   bool `bson:"add_itinerary_history"`
		SendPackage           bool `bson:"send_package"`
		CompleteShipment      bool `bson:"complete_shipment"`
	} `bson:"shipment"`

	Location struct {
		GetTransitPlaces                  bool `bson:"get_transit_places"`
		GetRecommendedShippingDestination bool `bson:"get_recommended_shipping_destination"`
		GetLocation                       bool `bson:"get_location"`
	} `bson:"location"`

	Courier struct {
		GetAvailableCouriers bool `bson:"get_available_couriers"`
	} `bson:"courier"`

	Cargo struct {
		GetMatchingCargos   bool `bson:"get_matching_cargos"`
		RequestLoadShipment bool `bson:"request_load_shipment"`
		LoadShipment        bool `bson:"load_shipment"`
		UnloadShipment      bool `bson:"unload_shipment"`
		MarkArrival         bool `bson:"mark_arrival"`
	} `bson:"cargo"`
}

func UserTypeModelToDomain(userTypeModel *UserType) *user.UserType {
	return &user.UserType{
		ID:           userTypeModel.ID.Hex(),
		Name:         userTypeModel.Name,
		Description:  userTypeModel.Description,
		PermissionID: userTypeModel.PermissionID.Hex(),
	}
}
