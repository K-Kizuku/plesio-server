package ristretto

import (
	"context"
	"log"
	"time"

	"github.com/K-Kizuku/plesio-server/app/domain/repository"
	"github.com/dgraph-io/ristretto"
)

type Client struct {
	Con *ristretto.Cache
}

func NewCacheClient() repository.IInMemoryCacheRepository {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 80000,  // number of keys to track frequency of (10M).
		MaxCost:     1 << 6, // maximum cost of cache (1GB).
		BufferItems: 64,     // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Client{
		Con: cache,
	}
}

func (c *Client) Get(ctx context.Context, key string) (interface{}, bool) {
	val, founded := c.Con.Get(key)
	return val, founded
}

func (c *Client) Set(ctx context.Context, key string, value interface{}) bool {
	added := c.Con.SetWithTTL(key, value, 1, 5*time.Second) //5秒だけキャッシュ
	return added
}

func (c *Client) Delete(ctx context.Context, key string) {
	c.Con.Del(key)
}
