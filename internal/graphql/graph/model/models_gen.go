// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Courier struct {
	ID       string    `json:"id"`
	UserID   string    `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Status   string    `json:"status"`
	Location *Location `json:"location"`
}

type CreateLocationInput struct {
	Name          string       `json:"name"`
	Type          LocationType `json:"type"`
	WarehouseID   *string      `json:"warehouse_id,omitempty"`
	ZipCode       string       `json:"zip_code"`
	Latitude      *float64     `json:"latitude,omitempty"`
	Longitude     *float64     `json:"longitude,omitempty"`
	StreetAddress *string      `json:"street_address,omitempty"`
}

type CreateShipmentInput struct {
	Origin    string       `json:"origin"`
	Sender    *EntityInput `json:"sender"`
	Recipient *EntityInput `json:"recipient"`
	Items     []*ItemInput `json:"items"`
}

type EntityInput struct {
	Name          string `json:"name"`
	PhoneNumber   string `json:"phone_number"`
	ZipCode       string `json:"zip_code"`
	StreetAddress string `json:"street_address"`
}

type Item struct {
	Name   string `json:"name"`
	Amount int32  `json:"amount"`
	Weight int32  `json:"weight"`
	Volume *int32 `json:"volume,omitempty"`
}

type ItemInput struct {
	Name   string       `json:"name"`
	Amount int32        `json:"amount"`
	Weight int32        `json:"weight"`
	Volume *VolumeInput `json:"volume,omitempty"`
}

type ItineraryLog struct {
	ActivityType string    `json:"activity_type"`
	Timestamp    time.Time `json:"timestamp"`
	Location     *Location `json:"location,omitempty"`
}

type Location struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Type      string           `json:"type"`
	Warehouse *Location        `json:"warehouse,omitempty"`
	Address   *LocationAddress `json:"address"`
}

type LocationAddress struct {
	Province      *string  `json:"province,omitempty"`
	City          *string  `json:"city,omitempty"`
	District      *string  `json:"district,omitempty"`
	Subdistrict   *string  `json:"subdistrict,omitempty"`
	ZipCode       *string  `json:"zip_code,omitempty"`
	Latitude      *float64 `json:"latitude,omitempty"`
	Longitude     *float64 `json:"longitude,omitempty"`
	StreetAddress *string  `json:"street_address,omitempty"`
}

type Mutation struct {
}

type PartyDetail struct {
	Name        *string `json:"name,omitempty"`
	Contact     *string `json:"contact,omitempty"`
	Province    *string `json:"province,omitempty"`
	City        *string `json:"city,omitempty"`
	District    *string `json:"district,omitempty"`
	Subdistrict *string `json:"subdistrict,omitempty"`
	Address     *string `json:"address,omitempty"`
	ZipCode     *string `json:"zip_code,omitempty"`
}

type Query struct {
}

type Region struct {
	Province    string  `json:"province"`
	City        string  `json:"city"`
	District    string  `json:"district"`
	Subdistrict *string `json:"subdistrict,omitempty"`
	ZipCode     string  `json:"zip_code"`
}

type RegisterAsCarrierInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	LocationID string `json:"location_id"`
}

type RegisterAsCourierInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	LocationID string `json:"location_id"`
}

type RegisterAsOperatorInput struct {
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Password   string       `json:"password"`
	LocationID string       `json:"location_id"`
	Type       OperatorType `json:"type"`
}

type RequestTransitInput struct {
	ShipmentID  string `json:"shipment_id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	CourierID   string `json:"courier_id"`
}

type Shipment struct {
	ID              string          `json:"id"`
	AirwayBill      string          `json:"airway_bill"`
	TransportStatus string          `json:"transport_status"`
	RoutingStatus   string          `json:"routing_status"`
	Items           []*Item         `json:"items"`
	SenderDetail    *PartyDetail    `json:"sender_detail"`
	RecipientDetail *PartyDetail    `json:"recipient_detail"`
	Origin          *Location       `json:"origin,omitempty"`
	Destination     *Location       `json:"destination,omitempty"`
	ItineraryLogs   []*ItineraryLog `json:"itinerary_logs"`
}

type User struct {
	ID      string      `json:"id"`
	ModelID string      `json:"model_id"`
	Type    string      `json:"type"`
	Entity  *UserEntity `json:"entity,omitempty"`
}

type UserEntity struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Location *Location `json:"location,omitempty"`
}

type VolumeInput struct {
	Length *int32 `json:"length,omitempty"`
	Width  *int32 `json:"width,omitempty"`
	Height *int32 `json:"height,omitempty"`
}

type LocationType string

const (
	LocationTypeWarehouse LocationType = "warehouse"
	LocationTypeDepot     LocationType = "depot"
)

var AllLocationType = []LocationType{
	LocationTypeWarehouse,
	LocationTypeDepot,
}

func (e LocationType) IsValid() bool {
	switch e {
	case LocationTypeWarehouse, LocationTypeDepot:
		return true
	}
	return false
}

func (e LocationType) String() string {
	return string(e)
}

func (e *LocationType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LocationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LocationType", str)
	}
	return nil
}

func (e LocationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OperatorType string

const (
	OperatorTypeDepotOperator     OperatorType = "depot_operator"
	OperatorTypeWarehouseOperator OperatorType = "warehouse_operator"
)

var AllOperatorType = []OperatorType{
	OperatorTypeDepotOperator,
	OperatorTypeWarehouseOperator,
}

func (e OperatorType) IsValid() bool {
	switch e {
	case OperatorTypeDepotOperator, OperatorTypeWarehouseOperator:
		return true
	}
	return false
}

func (e OperatorType) String() string {
	return string(e)
}

func (e *OperatorType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OperatorType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OperatorType", str)
	}
	return nil
}

func (e OperatorType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
