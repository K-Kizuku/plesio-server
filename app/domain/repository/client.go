package repository

type IClientRepository interface {
	ReadMessage()
	WriteMessage()
}
