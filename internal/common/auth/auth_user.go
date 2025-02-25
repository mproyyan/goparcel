package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthUser struct {
	UserID  string
	ModelID string
}

func (a AuthUser) exists() bool {
	return a.UserID != "" || a.ModelID != ""
}

func SendAuthUser(ctx context.Context, userId, modelId string) context.Context {
	md := metadata.Pairs("user_id", userId, "model_id", modelId)
	return metadata.NewOutgoingContext(ctx, md)
}

func RetrieveAuthUser(ctx context.Context) (*AuthUser, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadta")
	}

	var authUser AuthUser
	if values := md.Get("user_id"); len(values) > 0 {
		authUser.UserID = values[0]
	}

	if values := md.Get("model_id"); len(values) > 0 {
		authUser.ModelID = values[0]
	}

	if !authUser.exists() {
		return nil, status.Error(codes.Unauthenticated, "Missing credentials")
	}

	return &authUser, nil
}
