package resolvers

import (
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/graphql/graph/model"
)

func itemInputToProtoRequest(items []*model.ItemInput) []*genproto.Package {
	var itemsProto []*genproto.Package
	for _, item := range items {
		i := &genproto.Package{
			Name:   item.Name,
			Amount: item.Amount,
			Weight: item.Weight,
			Volume: &genproto.Volume{
				Length: *item.Volume.Length,
				Width:  *item.Volume.Width,
				Height: *item.Volume.Height,
			},
		}
		itemsProto = append(itemsProto, i)
	}

	return itemsProto
}

func itemToGraphResponse(item *genproto.Item) *model.Item {
	return &model.Item{
		Name:   item.Name,
		Amount: item.Amount,
		Weight: item.Weight,
		Volume: &item.Volume,
	}
}

func itemsToGraphResponse(items []*genproto.Item) []*model.Item {
	var itemsResponse []*model.Item
	for _, item := range items {
		i := itemToGraphResponse(item)
		itemsResponse = append(itemsResponse, i)
	}

	return itemsResponse
}

func shipmentToGraphResponse(shipment *genproto.Shipment) *model.Shipment {
	return &model.Shipment{
		ID:              shipment.Id,
		AirwayBill:      shipment.AirwayBill,
		TransportStatus: shipment.TransportStatus,
		RoutingStatus:   shipment.RoutingStatus,
		Items:           itemsToGraphResponse(shipment.Items),
		SenderDetail:    entityToGraphResponse(shipment.Sender),
		RecipientDetail: entityToGraphResponse(shipment.Recipient),
		Origin:          &model.Location{ID: shipment.Origin},
		Destination:     &model.Location{ID: shipment.Destination},
		ItineraryLogs:   itineraryLogToGraphResponse(shipment.ItineraryLogs),
	}
}

func shipmentsToGraphResponse(shipmentModel []*genproto.Shipment) []*model.Shipment {
	var shipments []*model.Shipment
	for _, s := range shipmentModel {
		ss := shipmentToGraphResponse(s)
		shipments = append(shipments, ss)
	}

	return shipments
}

func entityToGraphResponse(entity *genproto.EntityDetail) *model.PartyDetail {
	return &model.PartyDetail{
		Name:        &entity.Name,
		Contact:     &entity.Contact,
		Province:    &entity.Address.Province,
		City:        &entity.Address.City,
		District:    &entity.Address.District,
		Subdistrict: &entity.Address.Subdistrict,
		Address:     &entity.Address.StreetAddress,
		ZipCode:     &entity.Address.ZipCode,
	}
}

func itineraryLogToGraphResponse(logs []*genproto.ItineraryLog) []*model.ItineraryLog {
	var itineraries []*model.ItineraryLog
	for _, log := range logs {
		l := &model.ItineraryLog{
			ActivityType: log.ActivityType,
			Timestamp:    log.Timestamp.AsTime(),
			Location:     &model.Location{ID: log.LocationId},
		}
		itineraries = append(itineraries, l)
	}

	return itineraries
}
