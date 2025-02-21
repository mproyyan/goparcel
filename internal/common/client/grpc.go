package client

import (
	"errors"
	"os"

	"github.com/mproyyan/goparcel/internal/common/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient() (client genproto.UserServiceClient, close func() error, err error) {
	// Get address from env
	grpcAddr := os.Getenv("USER_SERVICE_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env USER_SERVICE_ADDR")
	}

	// Setup new client with insecure connection
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {
		return nil, func() error { return nil }, err
	}

	// Create user service client and return it
	return genproto.NewUserServiceClient(conn), conn.Close, nil
}

func NewLocationServiceClient() (client genproto.LocationServiceClient, close func() error, err error) {
	// Get address from env
	grpcAddr := os.Getenv("LOCATION_SERVICE_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env LOCATION_SERVICE_ADDR")
	}

	// Setup new client with insecure connection
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {
		return nil, func() error { return nil }, err
	}

	// Create user service client and return it
	return genproto.NewLocationServiceClient(conn), conn.Close, nil
}

func NewShipmentServiceClient() (client genproto.ShipmentServiceClient, close func() error, err error) {
	// Get address from env
	grpcAddr := os.Getenv("SHIPMENT_SERVICE_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env SHIPMENT_SERVICE_ADDR")
	}

	// Setup new client with insecure connection
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {
		return nil, func() error { return nil }, err
	}

	// Create shipment service client and return it
	return genproto.NewShipmentServiceClient(conn), conn.Close, nil
}

func NewCourierServiceClient() (client genproto.CourierServiceClient, close func() error, err error) {
	// Get address from env
	grpcAddr := os.Getenv("COURIER_SERVICE_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env COURIER_SERVICE_ADDR")
	}

	// Setup new client with insecure connection
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {
		return nil, func() error { return nil }, err
	}

	// Create courier service client and return it
	return genproto.NewCourierServiceClient(conn), conn.Close, nil
}
