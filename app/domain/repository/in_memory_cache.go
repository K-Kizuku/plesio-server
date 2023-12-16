package repository

type IInMemoryCacheRepository interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}) bool
}
