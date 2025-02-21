package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mproyyan/goparcel/internal/common/client"
	"github.com/mproyyan/goparcel/internal/gateway/services"
)

func main() {
	// Get api gateway address
	addr := os.Getenv("API_GATEWAY_ADDR")

	// Setup fiber
	app := fiber.New()
	api := app.Group("/api")

	// Connect to user service
	userServiceClient, close, err := client.NewUserServiceClient()
	if err != nil {
		log.Fatal("cannot connect to user service", err)
	}
	defer close()

	// Connect to location service
	locationServiceClient, close, err := client.NewLocationServiceClient()
	if err != nil {
		log.Fatal("cannot connect to location service", err)
	}
	defer close()

	// Connect to shipment service
	shipmentServiceClient, close, err := client.NewShipmentServiceClient()
	if err != nil {
		log.Fatal("cannot connect to shipment service", err)
	}
	defer close()

	// Connect to courier service
	courierServiceClient, close, err := client.NewCourierServiceClient()
	if err != nil {
		log.Fatal("cannot connect to user service", err)
	}
	defer close()

	// Bootstrap services
	userService := services.NewUserService(api, userServiceClient)
	locationService := services.NewLocationService(api, locationServiceClient)
	shipmentService := services.NewShipmentService(api, shipmentServiceClient)
	courierService := services.NewCourierService(api, courierServiceClient)

	userService.Bootstrap()
	locationService.Bootstrap()
	shipmentService.Bootstrap()
	courierService.Bootstrap()

	// Run server
	log.Fatal(app.Listen(addr))
}
