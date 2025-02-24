package port

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/users/app"
	"github.com/mproyyan/goparcel/internal/users/domain/courier"
	"github.com/mproyyan/goparcel/internal/users/domain/operator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	service app.UserService
	genproto.UnimplementedUserServiceServer
}

func NewGrpcServer(service app.UserService) GrpcServer {
	return GrpcServer{
		service: service,
	}
}

func (g GrpcServer) Login(ctx context.Context, request *genproto.LoginRequest) (*genproto.LoginResponse, error) {
	token, err := g.service.Login(ctx, request.Email, request.Password)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to login")
	}

	return &genproto.LoginResponse{Token: token}, nil
}

func (g GrpcServer) RegisterAsOperator(context context.Context, request *genproto.RegisterAsOperatorRequest) (*emptypb.Empty, error) {
	// Check value of operator type request value to decide operator type
	operatorTypeRequestValue := request.Type
	operatorType := operator.OperatorTypeFromString(operatorTypeRequestValue)
	if operatorType == operator.OperatorUnknown {
		return nil, status.Error(codes.InvalidArgument, "invalid operator type, must be depot_operator or warehouse_operator")
	}

	// Call user serice
	err := g.service.RegisterAsOperator(context, request.Name, request.Email, request.Password, request.Location, operatorType)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to register as operator")
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) RegisterAsCarrier(ctx context.Context, request *genproto.RegisterAsCarrierRequest) (*emptypb.Empty, error) {
	err := g.service.RegisterAsCarrier(ctx, request.Name, request.Email, request.Password, request.Location)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to register as carrier")
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) RegisterAsCourier(ctx context.Context, request *genproto.RegisterAsCourierRequest) (*emptypb.Empty, error) {
	err := g.service.RegisterAsCourier(ctx, request.Name, request.Email, request.Password, request.Location)
	if err != nil {
		return nil, cuserr.Decorate(err, "failed to register as operator")
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcServer) GetOperators(ctx context.Context, request *genproto.GetOperatorsRequest) (*genproto.OperatorResponse, error) {
	operators, err := g.service.GetOperators(ctx, request.Ids)
	if err != nil {
		return nil, cuserr.Decorate(err, "user service failed to get operators")
	}

	return &genproto.OperatorResponse{Operators: operatorsToProtoResponse(operators)}, nil
}

func (g GrpcServer) GetCouriers(ctx context.Context, request *genproto.GetCouriersRequest) (*genproto.CourierResponse, error) {
	couriers, err := g.service.GetCouriers(ctx, request.Ids)
	if err != nil {
		return nil, cuserr.Decorate(err, "user service failed to get couriers")
	}

	return &genproto.CourierResponse{Couriers: couriersToProtoResponse(couriers)}, nil
}

func operatorToProtoResponse(model *operator.Operator) *genproto.Operator {
	return &genproto.Operator{
		Id:         model.ID,
		UserId:     model.UserID,
		Type:       model.Type.String(),
		Name:       model.Name,
		Email:      model.Email,
		LocationId: model.LocationID,
	}
}

func operatorsToProtoResponse(models []*operator.Operator) []*genproto.Operator {
	var operators []*genproto.Operator
	for _, model := range models {
		op := operatorToProtoResponse(model)
		operators = append(operators, op)
	}

	return operators
}

func courierToProtoResponse(model *courier.Courier) *genproto.Courier {
	return &genproto.Courier{
		Id:         model.ID,
		UserId:     model.UserID,
		Name:       model.Name,
		Email:      model.Email,
		LocationId: model.LocationID,
	}
}

func couriersToProtoResponse(models []*courier.Courier) []*genproto.Courier {
	var couriers []*genproto.Courier
	for _, model := range models {
		c := courierToProtoResponse(model)
		couriers = append(couriers, c)
	}

	return couriers
}
