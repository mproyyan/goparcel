package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mproyyan/goparcel/internal/common/client"
	"github.com/mproyyan/goparcel/internal/graphql/graph/generated"
	"github.com/mproyyan/goparcel/internal/graphql/graph/middlewares"
	"github.com/mproyyan/goparcel/internal/graphql/graph/resolvers"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "1234"
const defultHost = "127.0.0.1"

func main() {
	port := os.Getenv("API_GATEWAY_PORT")
	if port == "" {
		port = defaultPort
	}

	host := os.Getenv("API_GATEWAY_HOST")
	if host == "" {
		host = defultHost
	}

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

	// Connect to user service
	userServiceClient, close, err := client.NewUserServiceClient()
	if err != nil {
		log.Fatal("cannot connect to user service", err)
	}
	defer close()

	// Connect to cargo service
	cargoServiceClient, close, err := client.NewCargoServiceClient()
	if err != nil {
		log.Fatal("cannot connect to cargo service", err)
	}
	defer close()

	resolver := resolvers.NewResolver(
		locationServiceClient,
		shipmentServiceClient,
		courierServiceClient,
		userServiceClient,
		cargoServiceClient,
	)

	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	// Add middleware
	srv.AroundOperations(middlewares.AuthMiddleware())

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
