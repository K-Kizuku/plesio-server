package gateway

import (
	"github.com/K-Kizuku/plesio-server/app/domain/repository"
	"github.com/K-Kizuku/plesio-server/app/driver/ristretto"
)

type RoomRepository struct {
	InMemoryRepo *ristretto.Client
}

func NewRoomRepository(inMemoryRepo *ristretto.Client) repository.IRoomRepository {
	return &RoomRepository{
		InMemoryRepo: inMemoryRepo,
	}
}

func (r *RoomRepository) Run() {

}
