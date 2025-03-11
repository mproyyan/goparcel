package server

import (
	"fmt"
	"net"
	"os"

	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunGrpcServer(registerServer func(server *grpc.Server)) {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "8080"
	}

	grpcEndpoint := fmt.Sprintf(":%s", port)
	grpcServer := grpc.NewServer()
	registerServer(grpcServer)

	listen, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.WithField("address", grpcEndpoint).Info("Service running")
	logrus.Fatal(grpcServer.Serve(listen))
}
