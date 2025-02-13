package port

import (
	"context"

	"github.com/mproyyan/goparcel/internal/common/genproto/users"
	"github.com/mproyyan/goparcel/internal/users/app"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	service app.UserService
	users.UnimplementedUserServiceServer
}

func NewGrpcServer(service app.UserService) GrpcServer {
	return GrpcServer{
		service: service,
	}
}

func (g GrpcServer) Login(ctx context.Context, request *users.LoginRequest) (*users.LoginResponse, error) {
	token, err := g.service.Login(ctx, request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	return &users.LoginResponse{Token: token}, nil
}

func (g GrpcServer) RegisterAsOperator(_ context.Context, _ *users.RegisterAsOperatorRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (g GrpcServer) RegisterAsCarrier(_ context.Context, _ *users.RegisterAsCarrierRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (g GrpcServer) RegisterAsCourier(_ context.Context, _ *users.RegisterAsCarrierRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (g GrpcServer) mustEmbedUnimplementedUserServiceServer() {
	panic("not implemented") // TODO: Implement
}
