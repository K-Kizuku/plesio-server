package redis

import (
	"context"

	"github.com/K-Kizuku/plesio-server/app/domain/repository"
	"github.com/K-Kizuku/plesio-server/utils/config"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Con *redis.Client
}

func NewDataStoreClient() repository.IDataStoreRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
		PoolSize: config.RedisPoolSize,
	})
	return &Client{
		Con: rdb,
	}
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	val, err := c.Con.Get(ctx, key).Result()
	return val, err
}

func (c *Client) Set(ctx context.Context, key, value string) (string, error) {
	val, err := c.Con.Set(ctx, key, value, 0).Result()
	return val, err
}
