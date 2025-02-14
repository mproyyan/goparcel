package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mproyyan/goparcel/internal/common/genproto/users"
	"github.com/mproyyan/goparcel/internal/common/server"
	"github.com/mproyyan/goparcel/internal/users/adapter"
	"github.com/mproyyan/goparcel/internal/users/app"
	"github.com/mproyyan/goparcel/internal/users/port"
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
		panic(err)
	}

	// Wait for connection
	err = client.Ping(ctxWithTimeout, nil)
	if err != nil {
		log.Fatal("MongoDB not responding:", err)
	}

	// Dependency
	database := client.Database(os.Getenv("MONGO_DATABASE"))
	userRepository := adapter.NewUserRepository(database)
	userTypeRepository := adapter.NewUserTypeRepository(database)
	operatorRepository := adapter.NewOperatorRepository(database)
	carrierRepository := adapter.NewCarrierRepository(database)
	userService := app.NewUserService(userRepository, userTypeRepository, operatorRepository, carrierRepository)

	server.RunGrpcServer(func(server *grpc.Server) {
		service := port.NewGrpcServer(userService)
		users.RegisterUserServiceServer(server, service)
	})
}
