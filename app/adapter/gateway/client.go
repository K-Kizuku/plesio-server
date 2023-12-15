package gateway

import (
	"github.com/K-Kizuku/plesio-server/app/domain/repository"
)

type ClientRepository struct {
	InMemoryRepo  repository.IInMemoryCacheRepository
	DataStoreRepo repository.IDataStoreRepository
}

func NewClientRepository(inMemoryRepo repository.IInMemoryCacheRepository, DataStore repository.IDataStoreRepository) repository.IClientRepository {
	return &ClientRepository{
		InMemoryRepo: inMemoryRepo,
	}
}

func (c *ClientRepository) ReadMessage() {
}

func (c *ClientRepository) WriteMessage() {

}
