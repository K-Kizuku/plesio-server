package repository

import "context"

type IDataStoreRepository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) (string, error)
}
