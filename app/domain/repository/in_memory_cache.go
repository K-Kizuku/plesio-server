package repository

type IInMemoryCacheRepository interface {
	Get(key string) string
	Set(key, value string, cost int64) bool
}
