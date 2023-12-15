package redis

import (
	"context"

	"github.com/K-Kizuku/plesio-server/app/domain/repository"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Con *redis.Client
}

func NewDataStoreClient() repository.IDataStoreRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 1000,
	})
	return &Client{
		Con: rdb,
	}
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (c *Client) Set(ctx context.Context, key, value string) (string, error) {
	return "", nil
}
