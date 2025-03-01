// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: shipments.proto

package genproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ShipmentService_CreateShipment_FullMethodName       = "/protobuf.ShipmentService/CreateShipment"
	ShipmentService_GetUnroutedShipment_FullMethodName  = "/protobuf.ShipmentService/GetUnroutedShipment"
	ShipmentService_GetRoutedShipments_FullMethodName   = "/protobuf.ShipmentService/GetRoutedShipments"
	ShipmentService_RequestTransit_FullMethodName       = "/protobuf.ShipmentService/RequestTransit"
	ShipmentService_IncomingShipments_FullMethodName    = "/protobuf.ShipmentService/IncomingShipments"
	ShipmentService_GetShipments_FullMethodName         = "/protobuf.ShipmentService/GetShipments"
	ShipmentService_ScanArrivingShipment_FullMethodName = "/protobuf.ShipmentService/ScanArrivingShipment"
	ShipmentService_ShipPackage_FullMethodName          = "/protobuf.ShipmentService/ShipPackage"
	ShipmentService_AddItineraryHistory_FullMethodName  = "/protobuf.ShipmentService/AddItineraryHistory"
	ShipmentService_DeliverPackage_FullMethodName       = "/protobuf.ShipmentService/DeliverPackage"
)

// ShipmentServiceClient is the client API for ShipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShipmentServiceClient interface {
	CreateShipment(ctx context.Context, in *CreateShipmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUnroutedShipment(ctx context.Context, in *GetUnroutedShipmentRequest, opts ...grpc.CallOption) (*ShipmentResponse, error)
	GetRoutedShipments(ctx context.Context, in *GetRoutedShipmentsRequest, opts ...grpc.CallOption) (*ShipmentResponse, error)
	RequestTransit(ctx context.Context, in *RequestTransitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	IncomingShipments(ctx context.Context, in *IncomingShipmentRequest, opts ...grpc.CallOption) (*TransferRequestResponse, error)
	GetShipments(ctx context.Context, in *GetShipmentsRequest, opts ...grpc.CallOption) (*ShipmentResponse, error)
	ScanArrivingShipment(ctx context.Context, in *ScanArrivingShipmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ShipPackage(ctx context.Context, in *ShipPackageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddItineraryHistory(ctx context.Context, in *AddItineraryHistoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeliverPackage(ctx context.Context, in *DeliverPackageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type shipmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShipmentServiceClient(cc grpc.ClientConnInterface) ShipmentServiceClient {
	return &shipmentServiceClient{cc}
}

func (c *shipmentServiceClient) CreateShipment(ctx context.Context, in *CreateShipmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ShipmentService_CreateShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) GetUnroutedShipment(ctx context.Context, in *GetUnroutedShipmentRequest, opts ...grpc.CallOption) (*ShipmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShipmentResponse)
	err := c.cc.Invoke(ctx, ShipmentService_GetUnroutedShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) GetRoutedShipments(ctx context.Context, in *GetRoutedShipmentsRequest, opts ...grpc.CallOption) (*ShipmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShipmentResponse)
	err := c.cc.Invoke(ctx, ShipmentService_GetRoutedShipments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) RequestTransit(ctx context.Context, in *RequestTransitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ShipmentService_RequestTransit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) IncomingShipments(ctx context.Context, in *IncomingShipmentRequest, opts ...grpc.CallOption) (*TransferRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransferRequestResponse)
	err := c.cc.Invoke(ctx, ShipmentService_IncomingShipments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) GetShipments(ctx context.Context, in *GetShipmentsRequest, opts ...grpc.CallOption) (*ShipmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShipmentResponse)
	err := c.cc.Invoke(ctx, ShipmentService_GetShipments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) ScanArrivingShipment(ctx context.Context, in *ScanArrivingShipmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ShipmentService_ScanArrivingShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) ShipPackage(ctx context.Context, in *ShipPackageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ShipmentService_ShipPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) AddItineraryHistory(ctx context.Context, in *AddItineraryHistoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ShipmentService_AddItineraryHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) DeliverPackage(ctx context.Context, in *DeliverPackageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ShipmentService_DeliverPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShipmentServiceServer is the server API for ShipmentService service.
// All implementations must embed UnimplementedShipmentServiceServer
// for forward compatibility.
type ShipmentServiceServer interface {
	CreateShipment(context.Context, *CreateShipmentRequest) (*emptypb.Empty, error)
	GetUnroutedShipment(context.Context, *GetUnroutedShipmentRequest) (*ShipmentResponse, error)
	GetRoutedShipments(context.Context, *GetRoutedShipmentsRequest) (*ShipmentResponse, error)
	RequestTransit(context.Context, *RequestTransitRequest) (*emptypb.Empty, error)
	IncomingShipments(context.Context, *IncomingShipmentRequest) (*TransferRequestResponse, error)
	GetShipments(context.Context, *GetShipmentsRequest) (*ShipmentResponse, error)
	ScanArrivingShipment(context.Context, *ScanArrivingShipmentRequest) (*emptypb.Empty, error)
	ShipPackage(context.Context, *ShipPackageRequest) (*emptypb.Empty, error)
	AddItineraryHistory(context.Context, *AddItineraryHistoryRequest) (*emptypb.Empty, error)
	DeliverPackage(context.Context, *DeliverPackageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedShipmentServiceServer()
}

// UnimplementedShipmentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShipmentServiceServer struct{}

func (UnimplementedShipmentServiceServer) CreateShipment(context.Context, *CreateShipmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShipment not implemented")
}
func (UnimplementedShipmentServiceServer) GetUnroutedShipment(context.Context, *GetUnroutedShipmentRequest) (*ShipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnroutedShipment not implemented")
}
func (UnimplementedShipmentServiceServer) GetRoutedShipments(context.Context, *GetRoutedShipmentsRequest) (*ShipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoutedShipments not implemented")
}
func (UnimplementedShipmentServiceServer) RequestTransit(context.Context, *RequestTransitRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestTransit not implemented")
}
func (UnimplementedShipmentServiceServer) IncomingShipments(context.Context, *IncomingShipmentRequest) (*TransferRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncomingShipments not implemented")
}
func (UnimplementedShipmentServiceServer) GetShipments(context.Context, *GetShipmentsRequest) (*ShipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShipments not implemented")
}
func (UnimplementedShipmentServiceServer) ScanArrivingShipment(context.Context, *ScanArrivingShipmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanArrivingShipment not implemented")
}
func (UnimplementedShipmentServiceServer) ShipPackage(context.Context, *ShipPackageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShipPackage not implemented")
}
func (UnimplementedShipmentServiceServer) AddItineraryHistory(context.Context, *AddItineraryHistoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItineraryHistory not implemented")
}
func (UnimplementedShipmentServiceServer) DeliverPackage(context.Context, *DeliverPackageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverPackage not implemented")
}
func (UnimplementedShipmentServiceServer) mustEmbedUnimplementedShipmentServiceServer() {}
func (UnimplementedShipmentServiceServer) testEmbeddedByValue()                         {}

// UnsafeShipmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShipmentServiceServer will
// result in compilation errors.
type UnsafeShipmentServiceServer interface {
	mustEmbedUnimplementedShipmentServiceServer()
}

func RegisterShipmentServiceServer(s grpc.ServiceRegistrar, srv ShipmentServiceServer) {
	// If the following call pancis, it indicates UnimplementedShipmentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShipmentService_ServiceDesc, srv)
}

func _ShipmentService_CreateShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).CreateShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_CreateShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).CreateShipment(ctx, req.(*CreateShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_GetUnroutedShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnroutedShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetUnroutedShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_GetUnroutedShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetUnroutedShipment(ctx, req.(*GetUnroutedShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_GetRoutedShipments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoutedShipmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetRoutedShipments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_GetRoutedShipments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetRoutedShipments(ctx, req.(*GetRoutedShipmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_RequestTransit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTransitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).RequestTransit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_RequestTransit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).RequestTransit(ctx, req.(*RequestTransitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_IncomingShipments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncomingShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).IncomingShipments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_IncomingShipments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).IncomingShipments(ctx, req.(*IncomingShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_GetShipments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShipmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetShipments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_GetShipments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetShipments(ctx, req.(*GetShipmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_ScanArrivingShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanArrivingShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).ScanArrivingShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_ScanArrivingShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).ScanArrivingShipment(ctx, req.(*ScanArrivingShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_ShipPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShipPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).ShipPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_ShipPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).ShipPackage(ctx, req.(*ShipPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_AddItineraryHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItineraryHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).AddItineraryHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_AddItineraryHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).AddItineraryHistory(ctx, req.(*AddItineraryHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_DeliverPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).DeliverPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_DeliverPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).DeliverPackage(ctx, req.(*DeliverPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShipmentService_ServiceDesc is the grpc.ServiceDesc for ShipmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShipmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.ShipmentService",
	HandlerType: (*ShipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShipment",
			Handler:    _ShipmentService_CreateShipment_Handler,
		},
		{
			MethodName: "GetUnroutedShipment",
			Handler:    _ShipmentService_GetUnroutedShipment_Handler,
		},
		{
			MethodName: "GetRoutedShipments",
			Handler:    _ShipmentService_GetRoutedShipments_Handler,
		},
		{
			MethodName: "RequestTransit",
			Handler:    _ShipmentService_RequestTransit_Handler,
		},
		{
			MethodName: "IncomingShipments",
			Handler:    _ShipmentService_IncomingShipments_Handler,
		},
		{
			MethodName: "GetShipments",
			Handler:    _ShipmentService_GetShipments_Handler,
		},
		{
			MethodName: "ScanArrivingShipment",
			Handler:    _ShipmentService_ScanArrivingShipment_Handler,
		},
		{
			MethodName: "ShipPackage",
			Handler:    _ShipmentService_ShipPackage_Handler,
		},
		{
			MethodName: "AddItineraryHistory",
			Handler:    _ShipmentService_AddItineraryHistory_Handler,
		},
		{
			MethodName: "DeliverPackage",
			Handler:    _ShipmentService_DeliverPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipments.proto",
}
