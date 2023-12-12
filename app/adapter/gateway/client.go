package gateway

import (
	"github.com/K-Kizuku/plesio-server/app/domain/repository"
	"github.com/K-Kizuku/plesio-server/app/driver/ristretto"
)

type ClientRepository struct {
	InMemoryRepo *ristretto.Client
}

func NewClientRepository(inMemoryRepo *ristretto.Client) repository.IClientRepository {
	return &ClientRepository{
		InMemoryRepo: inMemoryRepo,
	}
}

func (c *ClientRepository) ReadMessage() {
}

func (c *ClientRepository) WriteMessage() {

}
