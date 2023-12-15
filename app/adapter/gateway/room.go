package gateway

import (
	"github.com/K-Kizuku/plesio-server/app/domain/repository"
)

type RoomRepository struct {
	InMemoryRepo repository.IInMemoryCacheRepository
	DataStore    repository.IDataStoreRepository
}

func NewRoomRepository(inMemoryRepo repository.IInMemoryCacheRepository, dataStore repository.IDataStoreRepository) repository.IRoomRepository {
	return &RoomRepository{
		InMemoryRepo: inMemoryRepo,
		DataStore:    dataStore,
	}
}

func (r *RoomRepository) Run() {

}
