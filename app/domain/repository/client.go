package repository

import "net"

type IClientRepository interface {
	JoinRoom(roomID string, client *net.UDPAddr) error
	ExitRoom(roomID string, client *net.UDPAddr) error
}
