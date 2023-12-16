package repository

import "net"

type IRoomRepository interface {
	// ルームを作成し，RoomIDを返す
	CreateRoom(owner net.UDPAddr) (string, error)
	// ルームを削除する
	DeleteRoom(roomID string) error
	// 特定のルームに属するクライアントの情報を取得する
	GetClients(roomID string) []net.UDPAddr
}
