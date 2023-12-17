package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

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

func (c *ClientRepository) JoinRoom(ctx context.Context, roomID string, client *net.UDPAddr) error {
	clients := &Client{}
	value, err := c.DataStoreRepo.Get(ctx, roomID)
	go c.InMemoryRepo.Delete(ctx, roomID)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(value), clients); err != nil {
		return err
	}
	clients.Clients = append(clients.Clients, *client)
	b, err := json.Marshal(clients)
	if err != nil {
		fmt.Print(clients)
		return err
	}
	if _, err := c.DataStoreRepo.Set(ctx, roomID, string(b)); err != nil {
		return err
	}
	return nil
}

func (c *ClientRepository) ExitRoom(ctx context.Context, roomID string, client *net.UDPAddr) error {
	clients := &Client{}
	value, err := c.DataStoreRepo.Get(ctx, roomID)
	go c.InMemoryRepo.Delete(ctx, roomID)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(value), clients); err != nil {
		return err
	}
	clients.Clients = remove(clients.Clients, *client)
	b, err := json.Marshal(clients)
	if err != nil {
		log.Println(err)
	}
	c.DataStoreRepo.Set(ctx, roomID, string(b))
	return nil
}

func remove(s []net.UDPAddr, sv net.UDPAddr) []net.UDPAddr {
	result := make([]net.UDPAddr, 0, 20)
	for _, v := range s {
		if v.String() != sv.String() {
			result = append(result, v)
		}
	}
	return result
}
func (c *ClientRepository) SelectPresenter(ctx context.Context, roomID string, presenter string) error {
	key := "pst" + roomID
	go c.InMemoryRepo.Set(ctx, key, presenter)
	_, err := c.DataStoreRepo.Set(ctx, key, presenter)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientRepository) GetPresenter(ctx context.Context, roomID string) (string, error) {
	key := "pst" + roomID
	var presenter string
	value, found := c.InMemoryRepo.Get(ctx, key)
	if !found {
		data, err := c.DataStoreRepo.Get(ctx, key)
		if err != nil {
			return "", err
		}
		value = data
	}
	presenter = value.(string)
	return presenter, nil
}
