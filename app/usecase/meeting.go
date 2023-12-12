package usecase

import "github.com/K-Kizuku/plesio-server/app/adapter/gateway"

type MeetingUsecase struct {
	ClientRepository *gateway.ClientRepository
	RoomRepository   *gateway.RoomRepository
}

type IMeetingUsecase interface {
	JoinRoom()
	ExitRoom()
	CreateRoom()
	DeleteRoom()
}

func NewMeetingUsecase(clientRepository *gateway.ClientRepository, roomRepository *gateway.RoomRepository) IMeetingUsecase {
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
