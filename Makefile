include .env
export RAJAONGKIR_API_KEY := $(shell grep '^RAJAONGKIR_API_KEY=' .env | cut -d '=' -f2)

proto:
	rm -rf internal/common/genproto/*.go

	protoc --proto_path=api/protobuf \
  		--go_out=internal/common/genproto --go_opt=paths=source_relative \
  		--go-grpc_out=internal/common/genproto --go-grpc_opt=paths=source_relative \
  		api/protobuf/*.proto

user-service:
	REDIS_SERVER=localhost:6379 MONGO_SERVER=mongodb://172.26.192.1:27017/?replicaSet=rs0 MONGO_DATABASE=goparcel JWT_SECRET_KEY=huh GRPC_PORT=7777 \
	go run internal/users/main.go

shipment-service:
	CARGO_SERVICE_ADDR=127.0.0.1:5555 LOCATION_SERVICE_ADDR=127.0.0.1:8888 MONGO_SERVER=mongodb://172.26.192.1:27017/?replicaSet=rs0 MONGO_DATABASE=goparcel GRPC_PORT=9999 \
	go run internal/shipments/main.go

location-service:
	MONGO_SERVER=mongodb://172.26.192.1:27017/?replicaSet=rs0 MONGO_DATABASE=goparcel GRPC_PORT=8888 \
	go run internal/locations/main.go

courier-service:
	MONGO_SERVER=mongodb://172.26.192.1:27017/?replicaSet=rs0 MONGO_DATABASE=goparcel GRPC_PORT=6666 \
	go run internal/couriers/main.go

cargo-service:
	SHIPMENT_SERVICE_ADDR=127.0.0.1:9999 MONGO_SERVER=mongodb://172.26.192.1:27017/?replicaSet=rs0 MONGO_DATABASE=goparcel GRPC_PORT=5555 \
	go run internal/cargos/main.go

api-gateway:
	REDIS_SERVER=localhost:6379 JWT_SECRET_KEY=huh CARGO_SERVICE_ADDR=127.0.0.1:5555 COURIER_SERVICE_ADDR=127.0.0.1:6666 USER_SERVICE_ADDR=127.0.0.1:7777 LOCATION_SERVICE_ADDR=127.0.0.1:8888 SHIPMENT_SERVICE_ADDR=127.0.0.1:9999 \
	go run internal/graphql/server.go