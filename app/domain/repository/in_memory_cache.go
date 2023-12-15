package repository

type IInMemoryCacheRepository interface {
	Get(key string) (interface{}, bool)
	Set(key, value string, cost int64) bool
}
