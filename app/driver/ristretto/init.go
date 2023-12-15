package ristretto

import (
	"log"

	"github.com/K-Kizuku/plesio-server/app/domain/repository"
	"github.com/dgraph-io/ristretto"
)

type Client struct {
	Con *ristretto.Cache
}

func NewCacheClient() repository.IInMemoryCacheRepository {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Client{
		Con: cache,
	}
}

func (c *Client) Init() (*ristretto.Cache, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil, err
	}
	return cache, nil
}

func (c *Client) Get(key string) (interface{}, bool) {
	val, founded := c.Con.Get(key)
	return val, founded
}

func (c *Client) Set(key string, value interface{}, cost int64) bool {
	added := c.Con.Set(key, value, 1)
	return added
}
