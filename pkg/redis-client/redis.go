package pkgredisclient

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	config *baseconfig.Config
	client *redis.Client
}

func NewRedisClient(
	config *baseconfig.Config,
) *RedisClient {
	return &RedisClient{
		config: config,
	}
}

func (c *RedisClient) ConnectRedis() {
	c.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.config.RedisHost, c.config.RedisPort),
		Password: "",
		DB:       0,
	})
}

func (c *RedisClient) SetValue(ctx context.Context, key string, value any) error {
	convertedValue, convertedErr := json.Marshal(value)
	if convertedErr != nil {
		return convertedErr
	}
	expireTime := time.Duration(0)
	return c.client.Set(ctx, key, convertedValue, expireTime).Err()
}

func (c *RedisClient) GetValue(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *RedisClient) SetHashValue(ctx context.Context, key string, hashKey string, value any) error {
	convertedValue, convertedErr := json.Marshal(value)
	if convertedErr != nil {
		return convertedErr
	}
	return c.client.HSet(ctx, key, hashKey, convertedValue).Err()
}

func (c *RedisClient) GetHashValue(ctx context.Context, key string, hashKey string) (string, error) {
	return c.client.HGet(ctx, key, hashKey).Result()
}
