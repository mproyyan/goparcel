package adapter

import (
	"time"

	"github.com/mproyyan/goparcel/internal/shipments/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShipmentModel struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	AirwayBill      string              `bson:"airway_bill"`
	TransportStatus string              `bson:"transport_status"`
	RoutingStatus   string              `bson:"routing_status"`
	Items           []Item              `bson:"items"`
	SenderDetail    EntityDetail        `bson:"sender_detail"`
	RecipientDetail EntityDetail        `bson:"recipient_detail"`
	Origin          *primitive.ObjectID `bson:"origin,omitempty"`
	Destination     *primitive.ObjectID `bson:"destination,omitempty"`
	ItineraryLogs   []ItineraryLog      `bson:"itinerary_logs"`
}

type Item struct {
	Name   string `bson:"name"`
	Amount int    `bson:"amount"`
	Weight int32  `bson:"weight"`
	Volume int32  `bson:"volume"`
}

type EntityDetail struct {
	Name        string `bson:"name"`
	Contact     string `bson:"contact"`
	Province    string `bson:"province"`
	City        string `bson:"city"`
	District    string `bson:"district"`
	Subdistrict string `bson:"subdistrict"`
	Address     string `bson:"address"`
	ZipCode     string `bson:"zip_code"`
}

type ItineraryLog struct {
	ActivityType string              `bson:"activity_type"`
	Timestamp    time.Time           `bson:"timestamp"`
	Location     *primitive.ObjectID `bson:"location"`
}

type Location struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Type string             `bson:"type"`
}

func domainToItemModel(domainItems []domain.Item) []Item {
	var items []Item
	for _, d := range domainItems {
		items = append(items, Item{
			Name:   d.Name,
			Amount: d.Amount,
			Weight: d.Weight,
			Volume: d.Volume,
		})
	}
	return items
}

func convertObjIdToHex(objId *primitive.ObjectID) string {
	if objId == nil {
		return ""
	}

	return objId.Hex()
}
