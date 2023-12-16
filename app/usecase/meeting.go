package usecase

import (
	"context"
	"net"

	"github.com/K-Kizuku/plesio-server/app/domain/repository"
)

type MeetingUsecase struct {
	ClientRepository repository.IClientRepository
	RoomRepository   repository.IRoomRepository
}

type IMeetingUsecase interface {
	CreateRoom(ctx context.Context) (string, error)
	JoinRoom(ctx context.Context, roomID string, client *net.UDPAddr) error
	ExitRoom(ctx context.Context, roomID string, client *net.UDPAddr) error
}

func NewMeetingUsecase(clientRepository repository.IClientRepository, roomRepository repository.IRoomRepository) IMeetingUsecase {
	return &MeetingUsecase{
		ClientRepository: clientRepository,
		RoomRepository:   roomRepository,
	}
}

func (m *MeetingUsecase) CreateRoom(ctx context.Context) (string, error) {
	return "", nil
}
func (m *MeetingUsecase) JoinRoom(ctx context.Context, roomID string, client *net.UDPAddr) error {
	return nil
}
func (m *MeetingUsecase) ExitRoom(ctx context.Context, roomID string, client *net.UDPAddr) error {
	return nil
}
