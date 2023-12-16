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
	GetClients(ctx context.Context, roomID string) []net.UDPAddr
}

func NewMeetingUsecase(clientRepository repository.IClientRepository, roomRepository repository.IRoomRepository) IMeetingUsecase {
	return &MeetingUsecase{
		ClientRepository: clientRepository,
		RoomRepository:   roomRepository,
	}
}

func (m *MeetingUsecase) CreateRoom(ctx context.Context) (string, error) {
	roomID, err := m.RoomRepository.CreateRoom(ctx)
	if err != nil {
		return "", err
	}
	return roomID, nil
}

func (m *MeetingUsecase) JoinRoom(ctx context.Context, roomID string, client *net.UDPAddr) error {
	if err := m.ClientRepository.JoinRoom(ctx, roomID, client); err != nil {
		return err
	}
	return nil
}
func (m *MeetingUsecase) ExitRoom(ctx context.Context, roomID string, client *net.UDPAddr) error {
	if err := m.ClientRepository.ExitRoom(ctx, roomID, client); err != nil {
		return err
	}
	return nil
}

func (m *MeetingUsecase) GetClients(ctx context.Context, roomID string) []net.UDPAddr {
	return m.RoomRepository.GetClients(ctx, roomID)
}

// func (m *MeetingUsecase) Broadcast(ctx context.Context, message any)
