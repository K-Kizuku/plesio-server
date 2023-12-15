package repository

type IDBRepository interface {
	Connect()
	Close()
	Query()
}
