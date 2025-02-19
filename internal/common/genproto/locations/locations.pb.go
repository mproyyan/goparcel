// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: locations.proto

package locations

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type GetLocationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LocationID    string                 `protobuf:"bytes,1,opt,name=LocationID,proto3" json:"LocationID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLocationRequest) Reset() {
	*x = GetLocationRequest{}
	mi := &file_locations_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationRequest) ProtoMessage() {}

func (x *GetLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_locations_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationRequest.ProtoReflect.Descriptor instead.
func (*GetLocationRequest) Descriptor() ([]byte, []int) {
	return file_locations_proto_rawDescGZIP(), []int{0}
}

func (x *GetLocationRequest) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

type Location struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	WarehouseId   string                 `protobuf:"bytes,4,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id,omitempty"`
	Address       *Address               `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Location) Reset() {
	*x = Location{}
	mi := &file_locations_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_locations_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_locations_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Location) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Location) GetWarehouseId() string {
	if x != nil {
		return x.WarehouseId
	}
	return ""
}

func (x *Location) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Province      string                 `protobuf:"bytes,1,opt,name=province,proto3" json:"province,omitempty"`
	City          string                 `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	District      string                 `protobuf:"bytes,3,opt,name=district,proto3" json:"district,omitempty"`
	Subdistrict   string                 `protobuf:"bytes,4,opt,name=subdistrict,proto3" json:"subdistrict,omitempty"`
	Latitude      float64                `protobuf:"fixed64,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     float64                `protobuf:"fixed64,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
	StreetAddress string                 `protobuf:"bytes,7,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	ZipCode       string                 `protobuf:"bytes,8,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_locations_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_locations_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_locations_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *Address) GetSubdistrict() string {
	if x != nil {
		return x.Subdistrict
	}
	return ""
}

func (x *Address) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Address) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Address) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

type CreateLocationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	WarehouseId   string                 `protobuf:"bytes,3,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id,omitempty"`
	ZipCode       string                 `protobuf:"bytes,4,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	Latitude      float64                `protobuf:"fixed64,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     float64                `protobuf:"fixed64,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
	StreetAddress string                 `protobuf:"bytes,7,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateLocationRequest) Reset() {
	*x = CreateLocationRequest{}
	mi := &file_locations_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLocationRequest) ProtoMessage() {}

func (x *CreateLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_locations_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLocationRequest.ProtoReflect.Descriptor instead.
func (*CreateLocationRequest) Descriptor() ([]byte, []int) {
	return file_locations_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLocationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLocationRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateLocationRequest) GetWarehouseId() string {
	if x != nil {
		return x.WarehouseId
	}
	return ""
}

func (x *CreateLocationRequest) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *CreateLocationRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *CreateLocationRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *CreateLocationRequest) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

var File_locations_proto protoreflect.FileDescriptor

var file_locations_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x92,
	0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0xf3, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x75, 0x62, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x75, 0x62, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xde, 0x01, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xa1, 0x01, 0x0a, 0x0f, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x41,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x70, 0x72,
	0x6f, 0x79, 0x79, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_locations_proto_rawDescOnce sync.Once
	file_locations_proto_rawDescData []byte
)

func file_locations_proto_rawDescGZIP() []byte {
	file_locations_proto_rawDescOnce.Do(func() {
		file_locations_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_locations_proto_rawDesc), len(file_locations_proto_rawDesc)))
	})
	return file_locations_proto_rawDescData
}

var file_locations_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_locations_proto_goTypes = []any{
	(*GetLocationRequest)(nil),    // 0: protobuf.GetLocationRequest
	(*Location)(nil),              // 1: protobuf.Location
	(*Address)(nil),               // 2: protobuf.Address
	(*CreateLocationRequest)(nil), // 3: protobuf.CreateLocationRequest
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_locations_proto_depIdxs = []int32{
	2, // 0: protobuf.Location.address:type_name -> protobuf.Address
	0, // 1: protobuf.LocationService.GetLocation:input_type -> protobuf.GetLocationRequest
	3, // 2: protobuf.LocationService.CreateLocation:input_type -> protobuf.CreateLocationRequest
	1, // 3: protobuf.LocationService.GetLocation:output_type -> protobuf.Location
	4, // 4: protobuf.LocationService.CreateLocation:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_locations_proto_init() }
func file_locations_proto_init() {
	if File_locations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_locations_proto_rawDesc), len(file_locations_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_locations_proto_goTypes,
		DependencyIndexes: file_locations_proto_depIdxs,
		MessageInfos:      file_locations_proto_msgTypes,
	}.Build()
	File_locations_proto = out.File
	file_locations_proto_goTypes = nil
	file_locations_proto_depIdxs = nil
}
