proto:
	rm -rf internal/common/genproto/*.go

	protoc --proto_path=api/protobuf \
  		--go_out=internal/common/genproto --go_opt=paths=source_relative \
  		--go-grpc_out=internal/common/genproto --go-grpc_opt=paths=source_relative \
  		api/protobuf/*.proto