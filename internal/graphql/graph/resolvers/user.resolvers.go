package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"

	"github.com/mproyyan/goparcel/internal/common/genproto"
	"github.com/mproyyan/goparcel/internal/graphql/graph/generated"
	"github.com/mproyyan/goparcel/internal/graphql/graph/model"
)

// Login is the resolver for the Login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (string, error) {
	result, err := r.userService.Login(ctx, &genproto.LoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return result.Token, nil
}

// RegisterAsOperator is the resolver for the RegisterAsOperator field.
func (r *mutationResolver) RegisterAsOperator(ctx context.Context, input model.RegisterAsOperatorInput) (string, error) {
	_, err := r.userService.RegisterAsOperator(ctx, &genproto.RegisterAsOperatorRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Location: input.LocationID,
		Type:     input.Type.String(),
	})

	if err != nil {
		return "", err
	}

	return "registration as an operator has been successful", nil
}

// RegisterAsCourier is the resolver for the RegisterAsCourier field.
func (r *mutationResolver) RegisterAsCourier(ctx context.Context, input model.RegisterAsCourierInput) (string, error) {
	_, err := r.userService.RegisterAsCourier(ctx, &genproto.RegisterAsCourierRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Location: input.LocationID,
	})

	if err != nil {
		return "", err
	}

	return "registration as a courier has been successful", nil
}

// RegisterAsCarrier is the resolver for the RegisterAsCarrier field.
func (r *mutationResolver) RegisterAsCarrier(ctx context.Context, input model.RegisterAsCarrierInput) (string, error) {
	_, err := r.userService.RegisterAsCarrier(ctx, &genproto.RegisterAsCarrierRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Location: input.LocationID,
	})

	if err != nil {
		return "", err
	}

	return "registration as a carrier has been successful", nil
}

// GetUser is the resolver for the GetUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	user, err := r.userService.GetUser(ctx, &genproto.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return userToGraphResponse(user), nil
}

// Entity is the resolver for the entity field.
func (r *userResolver) Entity(ctx context.Context, obj *model.User) (*model.UserEntity, error) {
	if obj.ModelID == "" {
		return nil, nil
	}

	key := fmt.Sprintf("%s:%s", obj.Type, obj.ModelID)
	return r.entityLoader.Load(ctx, key)
}

// Location is the resolver for the location field.
func (r *userEntityResolver) Location(ctx context.Context, obj *model.UserEntity) (*model.Location, error) {
	if obj.Location != nil && obj.Location.ID == "" {
		return nil, nil
	}

	return r.locationLoader.Load(ctx, obj.Location.ID)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserEntity returns generated.UserEntityResolver implementation.
func (r *Resolver) UserEntity() generated.UserEntityResolver { return &userEntityResolver{r} }

type userResolver struct{ *Resolver }
type userEntityResolver struct{ *Resolver }
