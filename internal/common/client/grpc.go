package client

import (
	"errors"
	"os"

	"github.com/mproyyan/goparcel/internal/common/genproto/locations"
	"github.com/mproyyan/goparcel/internal/common/genproto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient() (client users.UserServiceClient, close func() error, err error) {
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
	return users.NewUserServiceClient(conn), conn.Close, nil
}

func NewLocationServiceClient() (client locations.LocationServiceClient, close func() error, err error) {
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
	return locations.NewLocationServiceClient(conn), conn.Close, nil
}
