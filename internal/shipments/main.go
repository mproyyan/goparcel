package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mproyyan/goparcel/internal/common/client"
	"github.com/mproyyan/goparcel/internal/common/db"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/common/server"
	"github.com/mproyyan/goparcel/internal/shipments/adapter"
	"github.com/mproyyan/goparcel/internal/shipments/app"
	"github.com/mproyyan/goparcel/internal/shipments/port"
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

	// Connect to location service
	grpcLocationServiceClient, close, err := client.NewLocationServiceClient()
	if err != nil {
		log.Fatal("cannot connect to location service", err)
	}
	defer close()

	// Dependency
	database := databaseClient.Database(os.Getenv("MONGO_DATABASE"))
	transactionManager := db.NewMongoTransactionManager(databaseClient)
	shipmentRepository := adapter.NewShipementRepository(database)
	locationService := adapter.NewLocationService(grpcLocationServiceClient)
	shipmentService := app.NewShipmentService(
		transactionManager,
		shipmentRepository,
		locationService,
	)

	server.RunGrpcServer(func(server *grpc.Server) {
		service := port.NewGrpcServer(shipmentService)
		genproto.RegisterShipmentServiceServer(server, service)
	})
}
