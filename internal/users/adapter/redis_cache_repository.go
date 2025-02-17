package adapter

import (
	"context"
	"time"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{
		client: client,
	}
}

func (c *CacheRepository) CacheUserPermissions(ctx context.Context, userID string, permissions user.Permissions) error {
	err := c.client.Set(ctx, userID, permissions, time.Hour).Err()
	if err != nil {
		return cuserr.Decorate(err, "%s", err.Error())
	}

	return nil
}
