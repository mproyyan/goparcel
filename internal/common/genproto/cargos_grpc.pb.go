// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: cargos.proto

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
	CargoService_GetCargos_FullMethodName         = "/protobuf.CargoService/GetCargos"
	CargoService_GetMatchingCargos_FullMethodName = "/protobuf.CargoService/GetMatchingCargos"
	CargoService_LoadShipment_FullMethodName      = "/protobuf.CargoService/LoadShipment"
	CargoService_MarkArrival_FullMethodName       = "/protobuf.CargoService/MarkArrival"
	CargoService_UnloadShipment_FullMethodName    = "/protobuf.CargoService/UnloadShipment"
)

// CargoServiceClient is the client API for CargoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CargoServiceClient interface {
	GetCargos(ctx context.Context, in *GetCargosRequest, opts ...grpc.CallOption) (*CargoResponse, error)
	GetMatchingCargos(ctx context.Context, in *GetMatchingCargosRequest, opts ...grpc.CallOption) (*CargoResponse, error)
	LoadShipment(ctx context.Context, in *LoadShipmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MarkArrival(ctx context.Context, in *MarkArrivalRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UnloadShipment(ctx context.Context, in *UnloadShipmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type cargoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCargoServiceClient(cc grpc.ClientConnInterface) CargoServiceClient {
	return &cargoServiceClient{cc}
}

func (c *cargoServiceClient) GetCargos(ctx context.Context, in *GetCargosRequest, opts ...grpc.CallOption) (*CargoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CargoResponse)
	err := c.cc.Invoke(ctx, CargoService_GetCargos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cargoServiceClient) GetMatchingCargos(ctx context.Context, in *GetMatchingCargosRequest, opts ...grpc.CallOption) (*CargoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CargoResponse)
	err := c.cc.Invoke(ctx, CargoService_GetMatchingCargos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cargoServiceClient) LoadShipment(ctx context.Context, in *LoadShipmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CargoService_LoadShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cargoServiceClient) MarkArrival(ctx context.Context, in *MarkArrivalRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CargoService_MarkArrival_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cargoServiceClient) UnloadShipment(ctx context.Context, in *UnloadShipmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CargoService_UnloadShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CargoServiceServer is the server API for CargoService service.
// All implementations must embed UnimplementedCargoServiceServer
// for forward compatibility.
type CargoServiceServer interface {
	GetCargos(context.Context, *GetCargosRequest) (*CargoResponse, error)
	GetMatchingCargos(context.Context, *GetMatchingCargosRequest) (*CargoResponse, error)
	LoadShipment(context.Context, *LoadShipmentRequest) (*emptypb.Empty, error)
	MarkArrival(context.Context, *MarkArrivalRequest) (*emptypb.Empty, error)
	UnloadShipment(context.Context, *UnloadShipmentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCargoServiceServer()
}

// UnimplementedCargoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCargoServiceServer struct{}

func (UnimplementedCargoServiceServer) GetCargos(context.Context, *GetCargosRequest) (*CargoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCargos not implemented")
}
func (UnimplementedCargoServiceServer) GetMatchingCargos(context.Context, *GetMatchingCargosRequest) (*CargoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchingCargos not implemented")
}
func (UnimplementedCargoServiceServer) LoadShipment(context.Context, *LoadShipmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadShipment not implemented")
}
func (UnimplementedCargoServiceServer) MarkArrival(context.Context, *MarkArrivalRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkArrival not implemented")
}
func (UnimplementedCargoServiceServer) UnloadShipment(context.Context, *UnloadShipmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnloadShipment not implemented")
}
func (UnimplementedCargoServiceServer) mustEmbedUnimplementedCargoServiceServer() {}
func (UnimplementedCargoServiceServer) testEmbeddedByValue()                      {}

// UnsafeCargoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CargoServiceServer will
// result in compilation errors.
type UnsafeCargoServiceServer interface {
	mustEmbedUnimplementedCargoServiceServer()
}

func RegisterCargoServiceServer(s grpc.ServiceRegistrar, srv CargoServiceServer) {
	// If the following call pancis, it indicates UnimplementedCargoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CargoService_ServiceDesc, srv)
}

func _CargoService_GetCargos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCargosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoServiceServer).GetCargos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CargoService_GetCargos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoServiceServer).GetCargos(ctx, req.(*GetCargosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CargoService_GetMatchingCargos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchingCargosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoServiceServer).GetMatchingCargos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CargoService_GetMatchingCargos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoServiceServer).GetMatchingCargos(ctx, req.(*GetMatchingCargosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CargoService_LoadShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoServiceServer).LoadShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CargoService_LoadShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoServiceServer).LoadShipment(ctx, req.(*LoadShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CargoService_MarkArrival_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkArrivalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoServiceServer).MarkArrival(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CargoService_MarkArrival_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoServiceServer).MarkArrival(ctx, req.(*MarkArrivalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CargoService_UnloadShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnloadShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoServiceServer).UnloadShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CargoService_UnloadShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoServiceServer).UnloadShipment(ctx, req.(*UnloadShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CargoService_ServiceDesc is the grpc.ServiceDesc for CargoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CargoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.CargoService",
	HandlerType: (*CargoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCargos",
			Handler:    _CargoService_GetCargos_Handler,
		},
		{
			MethodName: "GetMatchingCargos",
			Handler:    _CargoService_GetMatchingCargos_Handler,
		},
		{
			MethodName: "LoadShipment",
			Handler:    _CargoService_LoadShipment_Handler,
		},
		{
			MethodName: "MarkArrival",
			Handler:    _CargoService_MarkArrival_Handler,
		},
		{
			MethodName: "UnloadShipment",
			Handler:    _CargoService_UnloadShipment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cargos.proto",
}
