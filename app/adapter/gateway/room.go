package gateway

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/K-Kizuku/plesio-server/app/domain/repository"
	"github.com/K-Kizuku/plesio-server/utils/uuid"
)

type RoomRepository struct {
	InMemoryRepo  repository.IInMemoryCacheRepository
	DataStoreRepo repository.IDataStoreRepository
}

func NewRoomRepository(inMemoryRepo repository.IInMemoryCacheRepository, dataStoreRepo repository.IDataStoreRepository) repository.IRoomRepository {
	return &RoomRepository{
		InMemoryRepo:  inMemoryRepo,
		DataStoreRepo: dataStoreRepo,
	}
}

type Client struct {
	Clients []net.UDPAddr `json:"client"`
}

func (r *RoomRepository) CreateRoom(ctx context.Context) (string, error) {
	roomID, err := uuid.Generate()
	if err != nil {
		return "", err
	}

	clients := &Client{
		Clients: make([]net.UDPAddr, 0, 20),
	}

	go func(ctx context.Context) {
		r.InMemoryRepo.Set(ctx, roomID, clients)
	}(ctx)
	go func(ctx context.Context) {
		b, err := json.Marshal(clients)
		if err != nil {
			log.Println(err)
		}
		r.DataStoreRepo.Set(ctx, roomID, string(b))
	}(ctx)
	return roomID, err
}

func (r *RoomRepository) DeleteRoom(ctx context.Context, roomID string) error {
	return nil
}

func (r *RoomRepository) GetClients(ctx context.Context, roomID string) []net.UDPAddr {
	clients := &Client{
		Clients: make([]net.UDPAddr, 0, 20),
	}
	i, found := r.InMemoryRepo.Get(ctx, roomID)
	if found {
		clients.Clients = i.([]net.UDPAddr)
	} else {
		value, err := r.DataStoreRepo.Get(ctx, roomID)
		if err != nil {
			return nil
		}
		if err := json.Unmarshal([]byte(value), clients); err != nil {
			return nil
		}
	}
	return clients.Clients
}
