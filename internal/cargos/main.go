package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mproyyan/goparcel/internal/cargos/adapter"
	"github.com/mproyyan/goparcel/internal/cargos/app"
	"github.com/mproyyan/goparcel/internal/cargos/port"
	"github.com/mproyyan/goparcel/internal/common/client"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/common/server"
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

	// Connect to shipment service
	shipmentServiceClient, close, err := client.NewShipmentServiceClient()
	if err != nil {
		log.Fatal("cannot connect to shipment service", err)
	}
	defer close()

	// Dependency
	database := databaseClient.Database(os.Getenv("MONGO_DATABASE"))
	cargoRepository := adapter.NewCargoRepository(database)
	carrierRepository := adapter.NewCarrierRepository(database)
	shipmentService := adapter.NewShipmentService(shipmentServiceClient)
	cargoService := app.NewCargoService(cargoRepository, carrierRepository, shipmentService)

	server.RunGrpcServer(func(server *grpc.Server) {
		service := port.NewGrpcServer(cargoService)
		genproto.RegisterCargoServiceServer(server, service)
	})
}
