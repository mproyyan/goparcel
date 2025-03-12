include .env
export RAJAONGKIR_API_KEY := $(shell grep '^RAJAONGKIR_API_KEY=' .env | cut -d '=' -f2)

proto:
	rm -rf internal/common/genproto/*.go

	protoc --proto_path=api/protobuf \
  		--go_out=internal/common/genproto --go_opt=paths=source_relative \
  		--go-grpc_out=internal/common/genproto --go-grpc_opt=paths=source_relative \
  		api/protobuf/*.proto

run-all:
	@echo "Starting all services..."
	@trap 'kill 0' SIGINT; \
	$(MAKE) user-service & \
	$(MAKE) shipment-service & \
	$(MAKE) location-service & \
	$(MAKE) courier-service & \
	$(MAKE) cargo-service & \
	$(MAKE) api-gateway & \
	wait

user-service:
	@echo "Starting User Service..." && \
	REDIS_SERVER=localhost:6379 MONGO_SERVER=mongodb://172.26.192.1:27017/?replicaSet=rs0 MONGO_DATABASE=goparcel JWT_SECRET_KEY=huh GRPC_PORT=7777 \
	go run internal/users/main.go &

shipment-service:
	@echo "Starting Shipment Service..." && \
	CARGO_SERVICE_ADDR=127.0.0.1:5555 LOCATION_SERVICE_ADDR=127.0.0.1:8888 MONGO_SERVER=mongodb://172.26.192.1:27017/?replicaSet=rs0 MONGO_DATABASE=goparcel GRPC_PORT=9999 \
	go run internal/shipments/main.go &

location-service:
	@echo "Starting Location Service..." && \
	MONGO_SERVER=mongodb://172.26.192.1:27017/?replicaSet=rs0 MONGO_DATABASE=goparcel GRPC_PORT=8888 \
	go run internal/locations/main.go &

courier-service:
	@echo "Starting Courier Service..." && \
	MONGO_SERVER=mongodb://172.26.192.1:27017/?replicaSet=rs0 MONGO_DATABASE=goparcel GRPC_PORT=6666 \
	go run internal/couriers/main.go &

cargo-service:
	@echo "Starting Cargo Service..." && \
	SHIPMENT_SERVICE_ADDR=127.0.0.1:9999 MONGO_SERVER=mongodb://172.26.192.1:27017/?replicaSet=rs0 MONGO_DATABASE=goparcel GRPC_PORT=5555 \
	go run internal/cargos/main.go &

api-gateway:
	@echo "Starting API Gateway..." && \
	REDIS_SERVER=localhost:6379 JWT_SECRET_KEY=huh CARGO_SERVICE_ADDR=127.0.0.1:5555 COURIER_SERVICE_ADDR=127.0.0.1:6666 USER_SERVICE_ADDR=127.0.0.1:7777 LOCATION_SERVICE_ADDR=127.0.0.1:8888 SHIPMENT_SERVICE_ADDR=127.0.0.1:9999 \
	go run internal/graphql/server.go &