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
	"github.com/mproyyan/goparcel/internal/users/adapter"
	"github.com/mproyyan/goparcel/internal/users/app"
	"github.com/mproyyan/goparcel/internal/users/port"
	"github.com/redis/go-redis/v9"
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

	// Connect to redis
	redisClient := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Cant connect to redis: %v", err)
	}

	// Connect to location service
	grpcLocationServiceClient, close, err := client.NewLocationServiceClient()
	if err != nil {
		log.Fatal("cannot connect to user service", err)
	}
	defer close()

	// Dependency
	database := databaseClient.Database(os.Getenv("MONGO_DATABASE"))
	userRepository := adapter.NewUserRepository(database)
	userTypeRepository := adapter.NewUserTypeRepository(database)
	operatorRepository := adapter.NewOperatorRepository(database)
	carrierRepository := adapter.NewCarrierRepository(database)
	courierRepository := adapter.NewCourierRepository(database)
	cacheRepository := adapter.NewCacheRepository(redisClient)
	transaction := db.NewMongoTransactionManager(databaseClient)
	locationService := adapter.NewLocationService(grpcLocationServiceClient)
	userService := app.NewUserService(
		transaction,
		userRepository,
		userTypeRepository,
		operatorRepository,
		carrierRepository,
		courierRepository,
		cacheRepository,
		locationService,
	)

	server.RunGrpcServer(func(server *grpc.Server) {
		service := port.NewGrpcServer(userService)
		genproto.RegisterUserServiceServer(server, service)
	})
}
