package repository

import (
	"context"
	"net"
)

type IClientRepository interface {
	JoinRoom(ctx context.Context, roomID string, client *net.UDPAddr) error
	ExitRoom(ctx context.Context, roomID string, client *net.UDPAddr) error
	GetPresenter(ctx context.Context, roomID string) (string, error)
	SelectPresenter(ctx context.Context, roomID string, presenter string) error
}
