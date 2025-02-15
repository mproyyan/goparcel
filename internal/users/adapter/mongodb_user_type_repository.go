package adapter

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mproyyan/goparcel/internal/common/utils"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	cuserr "github.com/mproyyan/goparcel/internal/users/errors"
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
		return nil, cuserr.MongoError(err)
	}

	userTypeModel := userTypeModelToDomain(userTypeResult)
	return &userTypeModel, nil
}

// models
type UserType struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	PermissionID primitive.ObjectID `bson:"permission_id"`

	// Relations
	Permission *Permission `bson:"permission,omitempty"`
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

func (u *UserType) hasPermission() bool {
	return u.Permission != nil
}

// Helper function to get all granted permissions as slice
func (p Permission) grantedPermissions() []string {
	var permissions []string

	// Use reflect to read each field in a struct
	v := reflect.ValueOf(p)
	for i := 0; i < v.NumField(); i++ {
		category := v.Type().Field(i).Name // Category name (user_management, shipment, etc)
		subStruct := v.Field(i)

		// Iterate each field in sub-struct
		for j := 0; j < subStruct.NumField(); j++ {
			field := subStruct.Type().Field(j).Name // Permission field (get_user, create_shipment, etc)
			value := subStruct.Field(j).Bool()      // Permission value

			if value {
				// Only append granted permissionS
				permissions = append(permissions, fmt.Sprintf("%s.%s", toSnakeCase(category), toSnakeCase(field)))
			}
		}
	}

	return permissions
}

// Helper function to convert user type model to domain
func userTypeModelToDomain(userTypeModel UserType) user.UserType {
	var permissions user.Permissions
	if userTypeModel.hasPermission() {
		permissions = userTypeModel.Permission.grantedPermissions()
	}

	return user.UserType{
		ID:           userTypeModel.ID.Hex(),
		Name:         userTypeModel.Name,
		Description:  userTypeModel.Description,
		PermissionID: userTypeModel.PermissionID.Hex(),
		Permissions:  permissions,
	}
}

// Helper function to convert domain to user type model
func domainToUserTypeModel(userType user.UserType) (*UserType, error) {
	userTypeID, err := utils.ConvertToObjectId(userType.ID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	permissionID, err := utils.ConvertToObjectId(userType.PermissionID)
	if err != nil {
		return nil, cuserr.ErrInvalidHexString
	}

	return &UserType{
		ID:           userTypeID,
		Name:         userType.Name,
		Description:  userType.Description,
		PermissionID: permissionID,
	}, nil
}
