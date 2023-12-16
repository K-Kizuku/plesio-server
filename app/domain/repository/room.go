package repository

import (
	"context"
	"net"
)

type IRoomRepository interface {
	// ルームを作成し，RoomIDを返す
	CreateRoom(ctx context.Context) (string, error)
	// ルームを削除する
	DeleteRoom(ctx context.Context, roomID string) error
	// 特定のルームに属するクライアントの情報を取得する
	GetClients(ctx context.Context, roomID string) []net.UDPAddr
}
