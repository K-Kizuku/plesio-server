package models

import (
	"net"
	"sync"
)

const (
	//バッファサイズは要検討
	BUFFER_SIZE = 512
)

type Client struct {
	room *Room

	conn *net.UDPConn

	// Buffered channel of broadcast.
	buf chan []byte

	displayName string

	mux sync.Mutex
}

func NewClient(room *Room, conn *net.UDPConn, displayName string) *Client {
	return &Client{
		room:        room,
		conn:        conn,
		buf:         make(chan []byte, BUFFER_SIZE),
		displayName: displayName,
		mux:         sync.Mutex{},
	}
}
