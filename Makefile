proto:
	rm -rf internal/common/genproto/*.go

	protoc --proto_path=api/protobuf \
  		--go_out=internal/common/genproto/users --go_opt=paths=source_relative \
  		--go-grpc_out=internal/common/genproto/users --go-grpc_opt=paths=source_relative \
  		api/protobuf/users.proto

	protoc --proto_path=api/protobuf \
  		--go_out=internal/common/genproto/locations --go_opt=paths=source_relative \
  		--go-grpc_out=internal/common/genproto/locations --go-grpc_opt=paths=source_relative \
  		api/protobuf/locations.proto
	
	protoc --proto_path=api/protobuf \
  		--go_out=internal/common/genproto/shipments --go_opt=paths=source_relative \
  		--go-grpc_out=internal/common/genproto/shipments --go-grpc_opt=paths=source_relative \
  		api/protobuf/shipments.proto