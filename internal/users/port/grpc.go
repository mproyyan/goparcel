package port

import (
	"context"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/users/app"
	"github.com/mproyyan/goparcel/internal/users/domain/carrier"
	"github.com/mproyyan/goparcel/internal/users/domain/courier"
	"github.com/mproyyan/goparcel/internal/users/domain/operator"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
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

func (g GrpcServer) GetUser(ctx context.Context, request *genproto.GetUserRequest) (*genproto.User, error) {
	user, err := g.service.GetUser(ctx, request.Id)
	if err != nil {
		return nil, cuserr.Decorate(err, "GetUser service failed")
	}

	return userToProtoResponse(user), nil
}

func (g GrpcServer) GetUsers(ctx context.Context, request *genproto.GetUsersRequest) (*genproto.UserResponse, error) {
	users, err := g.service.GetUsers(ctx, request.Id)
	if err != nil {
		return nil, cuserr.Decorate(err, "GetUsers service failed")
	}

	return &genproto.UserResponse{Users: usersToProtoResponse(users)}, nil
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

func (g GrpcServer) GetCarriers(ctx context.Context, request *genproto.GetCarriersRequest) (*genproto.CarrierResponse, error) {
	carriers, err := g.service.GetCarriers(ctx, request.Ids)
	if err != nil {
		return nil, cuserr.Decorate(err, "user service failed to get carriers")
	}

	return &genproto.CarrierResponse{Carriers: carriersToProtoResponse(carriers)}, nil
}

// userToProtoResponse converts a domain User to a protobuf User response
func userToProtoResponse(u *user.User) *genproto.User {
	return &genproto.User{
		Id:      u.ID,
		ModelId: u.ModelID,
		Entity:  u.Entity.String(),
	}
}

// usersToProtoResponse converts a slice of domain Users to a slice of protobuf User responses
func usersToProtoResponse(users []*user.User) []*genproto.User {
	var protoUsers []*genproto.User
	for _, u := range users {
		protoUsers = append(protoUsers, userToProtoResponse(u))
	}
	return protoUsers
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

func carrierToProtoResponse(model *carrier.Carrier) *genproto.Carrier {
	return &genproto.Carrier{
		Id:         model.ID,
		UserId:     model.UserID,
		Name:       model.Name,
		Email:      model.Email,
		LocationId: model.LocationID,
		Status:     model.Status,
		CargoId:    model.CargoID,
	}
}

func carriersToProtoResponse(models []*carrier.Carrier) []*genproto.Carrier {
	var couriers []*genproto.Carrier
	for _, model := range models {
		c := carrierToProtoResponse(model)
		couriers = append(couriers, c)
	}

	return couriers
}
