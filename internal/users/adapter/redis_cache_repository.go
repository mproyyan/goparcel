package adapter

import (
	"context"
	"time"

	cuserr "github.com/mproyyan/goparcel/internal/common/errors"
	_ "github.com/mproyyan/goparcel/internal/common/logger"
	"github.com/mproyyan/goparcel/internal/users/domain/user"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
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
	logrus.WithField("permissions", permissions).Info("Caching user permissions")
	err := c.client.Set(ctx, userID, permissions, time.Hour).Err()
	if err != nil {
		logrus.WithError(err).Error("Failed to cache user permission")
		return cuserr.Decorate(err, "%s", err.Error())
	}

	return nil
}
