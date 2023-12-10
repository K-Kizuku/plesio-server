package models

import (
	"sync"
)

const (
	ROOM_CAPACITY = 20
)

type Room struct {
	clients map[*Client]bool

	broadcast chan []byte

	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	mux sync.Mutex
}

func NewRoom() *Room {
	return &Room{
		clients:    make(map[*Client]bool, ROOM_CAPACITY),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mux:        sync.Mutex{},
	}
}
