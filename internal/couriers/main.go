package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/common/server"
	"github.com/mproyyan/goparcel/internal/couriers/adapter"
	"github.com/mproyyan/goparcel/internal/couriers/app"
	"github.com/mproyyan/goparcel/internal/couriers/port"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	// Define database options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_SERVER"))
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to database
	databaseClient, err := mongo.Connect(ctxWithTimeout, clientOptions)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	// Wait for connection
	err = databaseClient.Ping(ctxWithTimeout, nil)
	if err != nil {
		log.Fatalf("MongoDB not responding: %v", err)
	}

	// Dependency
	database := databaseClient.Database(os.Getenv("MONGO_DATABASE"))
	courierRepository := adapter.NewCourierRepository(database)
	courierService := app.NewCourierService(courierRepository)

	server.RunGrpcServer(func(server *grpc.Server) {
		service := port.NewGrpcServer(courierService)
		genproto.RegisterCourierServiceServer(server, service)
	})
}
