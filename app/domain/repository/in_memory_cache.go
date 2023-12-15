package repository

type IInMemoryCacheRepository interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, cost int64) bool
}
