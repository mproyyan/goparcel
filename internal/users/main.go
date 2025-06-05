package main

import (
	"context"
	"os"
	"time"

	"github.com/mproyyan/goparcel/internal/common/db"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/mproyyan/goparcel/internal/common/server"
	"github.com/mproyyan/goparcel/internal/users/adapter"
	"github.com/mproyyan/goparcel/internal/users/app"
	"github.com/mproyyan/goparcel/internal/users/port"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	logrus.Info("Starting user service...")

	// Define database options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_SERVER"))
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to database
	databaseClient, err := mongo.Connect(ctxWithTimeout, clientOptions)
	if err != nil {
		logrus.Fatalf("cannot connect to database: %v", err)
	}

	// Wait for connection
	err = databaseClient.Ping(ctxWithTimeout, nil)
	if err != nil {
		logrus.Fatalf("MongoDB not responding: %v", err)
	}

	// Dependency
	database := databaseClient.Database(os.Getenv("MONGO_DATABASE"))
	userRepository := adapter.NewUserRepository(database)
	userTypeRepository := adapter.NewUserTypeRepository(database)
	operatorRepository := adapter.NewOperatorRepository(database)
	carrierRepository := adapter.NewCarrierRepository(database)
	courierRepository := adapter.NewCourierRepository(database)
	transaction := db.NewMongoTransactionManager(databaseClient)
	userService := app.NewUserService(
		transaction,
		userRepository,
		userTypeRepository,
		operatorRepository,
		carrierRepository,
		courierRepository,
	)

	server.RunGrpcServer(func(server *grpc.Server) {
		service := port.NewGrpcServer(userService)
		genproto.RegisterUserServiceServer(server, service)
	})
}
