package server

import (
	"fmt"
	"log"
	"net"
	"os"

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
		log.Fatal(err)
	}

	log.Fatal(grpcServer.Serve(listen))
}
