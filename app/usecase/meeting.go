package usecase

import (
	"github.com/K-Kizuku/plesio-server/app/domain/repository"
)

type MeetingUsecase struct {
	ClientRepository repository.IClientRepository
	RoomRepository   repository.IRoomRepository
}

type IMeetingUsecase interface {
	JoinRoom()
	ExitRoom()
	CreateRoom()
	DeleteRoom()
}

func NewMeetingUsecase(clientRepository repository.IClientRepository, roomRepository repository.IRoomRepository) IMeetingUsecase {
	return &MeetingUsecase{
		ClientRepository: clientRepository,
		RoomRepository:   roomRepository,
	}
}

func (m *MeetingUsecase) JoinRoom() {

}

func (m *MeetingUsecase) CreateRoom() {

}

func (m *MeetingUsecase) ExitRoom() {

}

func (m *MeetingUsecase) DeleteRoom() {

}
