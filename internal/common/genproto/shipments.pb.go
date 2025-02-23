// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: shipments.proto

package genproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateShipmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Origin        string                 `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Sender        *Entity                `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient     *Entity                `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Package       []*Package             `protobuf:"bytes,4,rep,name=package,proto3" json:"package,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateShipmentRequest) Reset() {
	*x = CreateShipmentRequest{}
	mi := &file_shipments_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShipmentRequest) ProtoMessage() {}

func (x *CreateShipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipments_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShipmentRequest.ProtoReflect.Descriptor instead.
func (*CreateShipmentRequest) Descriptor() ([]byte, []int) {
	return file_shipments_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShipmentRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *CreateShipmentRequest) GetSender() *Entity {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *CreateShipmentRequest) GetRecipient() *Entity {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *CreateShipmentRequest) GetPackage() []*Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type Entity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber   string                 `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	ZipCode       string                 `protobuf:"bytes,3,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	StreetAddress string                 `protobuf:"bytes,4,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Entity) Reset() {
	*x = Entity{}
	mi := &file_shipments_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_shipments_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_shipments_proto_rawDescGZIP(), []int{1}
}

func (x *Entity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Entity) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Entity) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *Entity) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

type Package struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Amount        int32                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Weight        int32                  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Volume        *Volume                `protobuf:"bytes,4,opt,name=volume,proto3" json:"volume,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Package) Reset() {
	*x = Package{}
	mi := &file_shipments_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_shipments_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_shipments_proto_rawDescGZIP(), []int{2}
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Package) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Package) GetVolume() *Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

type Volume struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Length        int32                  `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Width         int32                  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height        int32                  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Volume) Reset() {
	*x = Volume{}
	mi := &file_shipments_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_shipments_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_shipments_proto_rawDescGZIP(), []int{3}
}

func (x *Volume) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Volume) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Volume) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type GetUnroutedShipmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LocationId    string                 `protobuf:"bytes,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUnroutedShipmentRequest) Reset() {
	*x = GetUnroutedShipmentRequest{}
	mi := &file_shipments_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUnroutedShipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnroutedShipmentRequest) ProtoMessage() {}

func (x *GetUnroutedShipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipments_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnroutedShipmentRequest.ProtoReflect.Descriptor instead.
func (*GetUnroutedShipmentRequest) Descriptor() ([]byte, []int) {
	return file_shipments_proto_rawDescGZIP(), []int{4}
}

func (x *GetUnroutedShipmentRequest) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

type ShipmentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shipment      []*Shipment            `protobuf:"bytes,1,rep,name=shipment,proto3" json:"shipment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShipmentResponse) Reset() {
	*x = ShipmentResponse{}
	mi := &file_shipments_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShipmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipmentResponse) ProtoMessage() {}

func (x *ShipmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shipments_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipmentResponse.ProtoReflect.Descriptor instead.
func (*ShipmentResponse) Descriptor() ([]byte, []int) {
	return file_shipments_proto_rawDescGZIP(), []int{5}
}

func (x *ShipmentResponse) GetShipment() []*Shipment {
	if x != nil {
		return x.Shipment
	}
	return nil
}

type ItineraryLog struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ActivityType  string                 `protobuf:"bytes,1,opt,name=activity_type,json=activityType,proto3" json:"activity_type,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LocationId    string                 `protobuf:"bytes,3,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ItineraryLog) Reset() {
	*x = ItineraryLog{}
	mi := &file_shipments_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItineraryLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItineraryLog) ProtoMessage() {}

func (x *ItineraryLog) ProtoReflect() protoreflect.Message {
	mi := &file_shipments_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItineraryLog.ProtoReflect.Descriptor instead.
func (*ItineraryLog) Descriptor() ([]byte, []int) {
	return file_shipments_proto_rawDescGZIP(), []int{6}
}

func (x *ItineraryLog) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

func (x *ItineraryLog) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ItineraryLog) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

type Shipment struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AirwayBill      string                 `protobuf:"bytes,2,opt,name=airway_bill,json=airwayBill,proto3" json:"airway_bill,omitempty"`
	TransportStatus string                 `protobuf:"bytes,3,opt,name=transport_status,json=transportStatus,proto3" json:"transport_status,omitempty"`
	RoutingStatus   string                 `protobuf:"bytes,4,opt,name=routing_status,json=routingStatus,proto3" json:"routing_status,omitempty"`
	Items           []*Item                `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	Sender          *EntityDetail          `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient       *EntityDetail          `protobuf:"bytes,7,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Origin          string                 `protobuf:"bytes,8,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination     string                 `protobuf:"bytes,9,opt,name=destination,proto3" json:"destination,omitempty"`
	ItineraryLogs   []*ItineraryLog        `protobuf:"bytes,10,rep,name=itinerary_logs,json=itineraryLogs,proto3" json:"itinerary_logs,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Shipment) Reset() {
	*x = Shipment{}
	mi := &file_shipments_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Shipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shipment) ProtoMessage() {}

func (x *Shipment) ProtoReflect() protoreflect.Message {
	mi := &file_shipments_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shipment.ProtoReflect.Descriptor instead.
func (*Shipment) Descriptor() ([]byte, []int) {
	return file_shipments_proto_rawDescGZIP(), []int{7}
}

func (x *Shipment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Shipment) GetAirwayBill() string {
	if x != nil {
		return x.AirwayBill
	}
	return ""
}

func (x *Shipment) GetTransportStatus() string {
	if x != nil {
		return x.TransportStatus
	}
	return ""
}

func (x *Shipment) GetRoutingStatus() string {
	if x != nil {
		return x.RoutingStatus
	}
	return ""
}

func (x *Shipment) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Shipment) GetSender() *EntityDetail {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Shipment) GetRecipient() *EntityDetail {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *Shipment) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Shipment) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Shipment) GetItineraryLogs() []*ItineraryLog {
	if x != nil {
		return x.ItineraryLogs
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Amount        int32                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Weight        int32                  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Volume        int32                  `protobuf:"varint,4,opt,name=volume,proto3" json:"volume,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_shipments_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_shipments_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_shipments_proto_rawDescGZIP(), []int{8}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Item) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Item) GetVolume() int32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type EntityDetail struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Contact       string                 `protobuf:"bytes,2,opt,name=contact,proto3" json:"contact,omitempty"`
	Address       *Address               `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntityDetail) Reset() {
	*x = EntityDetail{}
	mi := &file_shipments_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityDetail) ProtoMessage() {}

func (x *EntityDetail) ProtoReflect() protoreflect.Message {
	mi := &file_shipments_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityDetail.ProtoReflect.Descriptor instead.
func (*EntityDetail) Descriptor() ([]byte, []int) {
	return file_shipments_proto_rawDescGZIP(), []int{9}
}

func (x *EntityDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EntityDetail) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *EntityDetail) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_shipments_proto protoreflect.FileDescriptor

var file_shipments_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x77, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x22, 0x4e, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x3d, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x64, 0x53,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x42, 0x0a, 0x10, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x49, 0x74, 0x69, 0x6e, 0x65, 0x72, 0x61, 0x72,
	0x79, 0x4c, 0x6f, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x92, 0x03, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x77, 0x61, 0x79, 0x5f, 0x62, 0x69, 0x6c, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x69, 0x72, 0x77, 0x61, 0x79, 0x42, 0x69,
	0x6c, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0e, 0x69, 0x74,
	0x69, 0x6e, 0x65, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x74,
	0x69, 0x6e, 0x65, 0x72, 0x61, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x0d, 0x69, 0x74, 0x69, 0x6e,
	0x65, 0x72, 0x61, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x62, 0x0a, 0x04, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x69, 0x0a,
	0x0c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xb9, 0x01, 0x0a, 0x0f, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x55, 0x6e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x64, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x6e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x64, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x70, 0x72, 0x6f, 0x79, 0x79, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x70, 0x61,
	0x72, 0x63, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_shipments_proto_rawDescOnce sync.Once
	file_shipments_proto_rawDescData []byte
)

func file_shipments_proto_rawDescGZIP() []byte {
	file_shipments_proto_rawDescOnce.Do(func() {
		file_shipments_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shipments_proto_rawDesc), len(file_shipments_proto_rawDesc)))
	})
	return file_shipments_proto_rawDescData
}

var file_shipments_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_shipments_proto_goTypes = []any{
	(*CreateShipmentRequest)(nil),      // 0: protobuf.CreateShipmentRequest
	(*Entity)(nil),                     // 1: protobuf.Entity
	(*Package)(nil),                    // 2: protobuf.Package
	(*Volume)(nil),                     // 3: protobuf.Volume
	(*GetUnroutedShipmentRequest)(nil), // 4: protobuf.GetUnroutedShipmentRequest
	(*ShipmentResponse)(nil),           // 5: protobuf.ShipmentResponse
	(*ItineraryLog)(nil),               // 6: protobuf.ItineraryLog
	(*Shipment)(nil),                   // 7: protobuf.Shipment
	(*Item)(nil),                       // 8: protobuf.Item
	(*EntityDetail)(nil),               // 9: protobuf.EntityDetail
	(*timestamppb.Timestamp)(nil),      // 10: google.protobuf.Timestamp
	(*Address)(nil),                    // 11: protobuf.Address
	(*emptypb.Empty)(nil),              // 12: google.protobuf.Empty
}
var file_shipments_proto_depIdxs = []int32{
	1,  // 0: protobuf.CreateShipmentRequest.sender:type_name -> protobuf.Entity
	1,  // 1: protobuf.CreateShipmentRequest.recipient:type_name -> protobuf.Entity
	2,  // 2: protobuf.CreateShipmentRequest.package:type_name -> protobuf.Package
	3,  // 3: protobuf.Package.volume:type_name -> protobuf.Volume
	7,  // 4: protobuf.ShipmentResponse.shipment:type_name -> protobuf.Shipment
	10, // 5: protobuf.ItineraryLog.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 6: protobuf.Shipment.items:type_name -> protobuf.Item
	9,  // 7: protobuf.Shipment.sender:type_name -> protobuf.EntityDetail
	9,  // 8: protobuf.Shipment.recipient:type_name -> protobuf.EntityDetail
	6,  // 9: protobuf.Shipment.itinerary_logs:type_name -> protobuf.ItineraryLog
	11, // 10: protobuf.EntityDetail.address:type_name -> protobuf.Address
	0,  // 11: protobuf.ShipmentService.CreateShipment:input_type -> protobuf.CreateShipmentRequest
	4,  // 12: protobuf.ShipmentService.GetUnroutedShipment:input_type -> protobuf.GetUnroutedShipmentRequest
	12, // 13: protobuf.ShipmentService.CreateShipment:output_type -> google.protobuf.Empty
	5,  // 14: protobuf.ShipmentService.GetUnroutedShipment:output_type -> protobuf.ShipmentResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_shipments_proto_init() }
func file_shipments_proto_init() {
	if File_shipments_proto != nil {
		return
	}
	file_locations_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shipments_proto_rawDesc), len(file_shipments_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipments_proto_goTypes,
		DependencyIndexes: file_shipments_proto_depIdxs,
		MessageInfos:      file_shipments_proto_msgTypes,
	}.Build()
	File_shipments_proto = out.File
	file_shipments_proto_goTypes = nil
	file_shipments_proto_depIdxs = nil
}
