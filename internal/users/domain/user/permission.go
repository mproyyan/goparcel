package user

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PermissionRepository interface {
	FindPermission(ctx context.Context, id primitive.ObjectID) (*Permission, error)
}

type UserManagementPermission struct {
	GetUserLocation bool `bson:"get_user_location"`
	GetUser         bool `bson:"get_user"`
}

type ShipmentPermission struct {
	CreateShipment        bool `bson:"create_shipment"`
	GetUnroutedShipments  bool `bson:"get_unrouted_shipments"`
	GetRoutedShipments    bool `bson:"get_routed_shipments"`
	GetIncomingShipments  bool `bson:"get_incoming_shipments"`
	ScanArrivingShipments bool `bson:"scan_arriving_shipments"`
	ShipPackage           bool `bson:"ship_package"`
	AddItineraryHistory   bool `bson:"add_itinerary_history"`
	SendPackage           bool `bson:"send_package"`
	CompleteShipment      bool `bson:"complete_shipment"`
}

type LocationPermission struct {
	GetTransitPlaces                  bool `bson:"get_transit_places"`
	GetRecommendedShippingDestination bool `bson:"get_recommended_shipping_destination"`
	GetLocation                       bool `bson:"get_location"`
}

type CourierPermission struct {
	GetAvailableCouriers bool `bson:"get_available_couriers"`
}

type CargoPermission struct {
	GetMatchingCargos   bool `bson:"get_matching_cargos"`
	RequestLoadShipment bool `bson:"request_load_shipment"`
	LoadShipment        bool `bson:"load_shipment"`
	UnloadShipment      bool `bson:"unload_shipment"`
	MarkArrival         bool `bson:"mark_arrival"`
}

type Permission struct {
	UserManagement UserManagementPermission `bson:"user_management"`
	Shipment       ShipmentPermission       `bson:"shipment"`
	Location       LocationPermission       `bson:"location"`
	Courier        CourierPermission        `bson:"courier"`
	Cargo          CargoPermission          `bson:"cargo"`
}

// Helper function to get all granted permissions as slice
func (p Permission) GrantedPermissions() Permissions {
	var permissions Permissions

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

type Permissions []string

// Implement encoding.BinaryMarshaler
func (p Permissions) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

// Helper function to create snake case for permission name
func toSnakeCase(s string) string {
	result := ""
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result += "_"
		}
		result += string(r)
	}

	return strings.ToLower(result)
}
