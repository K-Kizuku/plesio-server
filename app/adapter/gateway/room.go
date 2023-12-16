package gateway

import (
	"net"

	"github.com/google/uuid"

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

func (r *RoomRepository) CreateRoom(owner net.UDPAddr) (string, error) {
	uuid, err := uuid.NewV7()
	roomID := uuid.String()
	// clients := make([]net.UDPAddr, 0, 20)
	// r.InMemoryRepo.Set(roomID, clients)
	// r.DataStore.Set(context.TODO(), roomID, )

	return roomID, err
}

func (r *RoomRepository) DeleteRoom(roomID string) error {

}

func (r *RoomRepository) GetClients(roomID string) []net.UDPAddr {

}
