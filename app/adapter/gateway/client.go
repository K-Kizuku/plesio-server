package gateway

import (
	"github.com/K-Kizuku/plesio-server/app/domain/repository"
)

type ClientRepository struct {
	InMemoryRepo  repository.IInMemoryCacheRepository
	DataStoreRepo repository.IDataStoreRepository
}

func NewClientRepository(inMemoryRepo repository.IInMemoryCacheRepository, dataStore repository.IDataStoreRepository) repository.IClientRepository {
	return &ClientRepository{
		InMemoryRepo:  inMemoryRepo,
		DataStoreRepo: dataStore,
	}
}

func (c *ClientRepository) ReadMessage() {
}

func (c *ClientRepository) WriteMessage() {

}
