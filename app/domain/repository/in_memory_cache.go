package repository

import "context"

type IInMemoryCacheRepository interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, value interface{}) bool
	Delete(ctx context.Context, key string)
}
