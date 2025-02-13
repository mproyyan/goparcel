package port

import (
	"context"
	"fmt"

	"github.com/mproyyan/goparcel/internal/common/genproto/users"
	"github.com/mproyyan/goparcel/internal/users/app"
	"github.com/mproyyan/goparcel/internal/users/domain/operator"
	cuserr "github.com/mproyyan/goparcel/internal/users/errors"
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

func (g GrpcServer) RegisterAsOperator(context context.Context, request *users.RegisterAsOperatorRequest) (*emptypb.Empty, error) {
	// Check oneof request
	if request.Type == nil {
		return nil, cuserr.ErrInvalidOperatorType
	}

	// Check value of operator type request value to decide operator type
	operatorTypeRequestValue := fmt.Sprintf("%v", request.Type)
	operatorType, err := operator.OperatorTypeFromString(operatorTypeRequestValue)
	if err != nil {
		return nil, cuserr.ErrInvalidOperatorType
	}

	// Call user serice
	err = g.service.RegisterAsOperator(context, request.Name, request.Email, request.Password, request.Location, operatorType)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
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
