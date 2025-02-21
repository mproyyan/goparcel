package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/common/server"
	"github.com/mproyyan/goparcel/internal/locations/adapter"
	"github.com/mproyyan/goparcel/internal/locations/app"
	"github.com/mproyyan/goparcel/internal/locations/port"
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
	client, err := mongo.Connect(ctxWithTimeout, clientOptions)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	// Wait for connection
	err = client.Ping(ctxWithTimeout, nil)
	if err != nil {
		log.Fatalf("MongoDB not responding: %v", err)
	}

	// Create http client
	httpClient := http.Client{Timeout: time.Second * 10}

	// Dependency
	database := client.Database(os.Getenv("MONGO_DATABASE"))
	locationRepository := adapter.NewLocationRepository(database)
	regionService := adapter.NewRegionService(&httpClient, os.Getenv("RAJAONGKIR_API_KEY"))
	locationService := app.NewLocationService(regionService, locationRepository)

	// Run grpc server
	server.RunGrpcServer(func(server *grpc.Server) {
		service := port.NewGrpcServer(locationService)
		genproto.RegisterLocationServiceServer(server, service)
	})
}
