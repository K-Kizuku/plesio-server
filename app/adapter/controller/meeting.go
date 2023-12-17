package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net"

	"github.com/K-Kizuku/plesio-server/app/usecase"
	"golang.org/x/sync/errgroup"
)

type MeetingController struct {
	MeetingUsecase usecase.IMeetingUsecase
}

type IMeetingController interface {
	CreateRoom(ctx context.Context, client *net.UDPAddr, ln *net.UDPConn) error
	JoinRoom(ctx context.Context, roomID string, client *net.UDPAddr, ln *net.UDPConn) error
	ExitRoom(ctx context.Context, roomID string, client *net.UDPAddr, ln *net.UDPConn) error
	BroadcastMessage(ctx context.Context, roomID string, message string, me *net.UDPAddr, ln *net.UDPConn) error
	BroadcastAudio(ctx context.Context, roomID string, message string, me *net.UDPAddr, ln *net.UDPConn) error
	SelectPresenter(ctx context.Context, roomID string, presenter string, ln *net.UDPConn) error
}

func NewMeetingContrallor(meetingUsecase usecase.IMeetingUsecase) IMeetingController {
	return &MeetingController{
		MeetingUsecase: meetingUsecase,
	}
}

func (m *MeetingController) CreateRoom(ctx context.Context, client *net.UDPAddr, ln *net.UDPConn) error {
	roomID, err := m.MeetingUsecase.CreateRoom(ctx)
	if err != nil {
		return err
	}
	res := &Protocol{
		Type: "create_room",
		Header: Header{
			RoomID:       roomID,
			WantClientID: "",
		},
		Body: Body{
			Content: roomID,
		},
	}
	b, err := json.Marshal(res)
	if err != nil {
		return err
	}
	if _, err := ln.WriteToUDP(b, client); err != nil {
		return err
	}
	return nil
}

func (m *MeetingController) JoinRoom(ctx context.Context, roomID string, client *net.UDPAddr, ln *net.UDPConn) error {
	if err := m.MeetingUsecase.JoinRoom(ctx, roomID, client); err != nil {
		return err
	}
	clients := m.MeetingUsecase.GetClients(ctx, roomID)
	roomData := make([]string, 0, 30)
	for _, v := range clients {
		roomData = append(roomData, v.String())
	}
	b, err := json.Marshal(roomData)
	if err != nil {
		return err
	}
	res := &Protocol{
		Type: "join_room",
		Header: Header{
			RoomID:       roomID,
			WantClientID: "",
		},
		Body: Body{
			Content: string(b),
		},
	}
	if err := m.broadcast(ctx, *res, clients, ln); err != nil {
		return err
	}
	return nil
}

func (m *MeetingController) ExitRoom(ctx context.Context, roomID string, client *net.UDPAddr, ln *net.UDPConn) error {
	if err := m.MeetingUsecase.ExitRoom(ctx, roomID, client); err != nil {
		return err
	}
	clients := m.MeetingUsecase.GetClients(ctx, roomID)
	res := &Protocol{
		Type: "exit_room",
		Header: Header{
			RoomID:       roomID,
			WantClientID: "",
		},
		Body: Body{
			Content: roomID,
		},
	}
	if err := m.broadcast(ctx, *res, clients, ln); err != nil {
		return err
	}
	return nil
}

func (m *MeetingController) BroadcastMessage(ctx context.Context, roomID string, message string, me *net.UDPAddr, ln *net.UDPConn) error {
	clients := m.MeetingUsecase.GetClients(ctx, roomID)
	presenter, err := m.MeetingUsecase.GetPresenter(ctx, roomID, me)
	if err != nil {
		return err
	}
	if presenter != me.String() {
		return nil
	}
	res := &Protocol{
		Type: "AA",
		Header: Header{
			RoomID:       roomID,
			WantClientID: presenter,
		},
		Body: Body{
			Content: message,
		},
	}
	b, err := json.Marshal(res)
	if err != nil {
		return err
	}
	for _, client := range clients {
		client := client
		if _, err := ln.WriteToUDP(b, &client); err != nil {
			return err
		}
	}

	return errors.New("no broadcast message in room:" + roomID)
}

func (m *MeetingController) BroadcastAudio(ctx context.Context, roomID string, message string, me *net.UDPAddr, ln *net.UDPConn) error {
	eg, ctx := errgroup.WithContext(ctx)
	clients := m.MeetingUsecase.GetClients(ctx, roomID)
	res := &Protocol{
		Type: "audio",
		Header: Header{
			RoomID:       roomID,
			WantClientID: "",
		},
		Body: Body{
			Content: message,
		},
	}
	b, err := json.Marshal(res)
	if err != nil {
		return err
	}
	for _, client := range clients {
		client := client
		eg.Go(func() error {
			if me.String() == client.String() {
				return nil
			}
			if _, err := ln.WriteToUDP(b, &client); err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func (m *MeetingController) SelectPresenter(ctx context.Context, roomID string, presenter string, ln *net.UDPConn) error {
	if err := m.MeetingUsecase.SelectPresenter(ctx, roomID, presenter); err != nil {
		return err
	}
	return nil
}

func (m *MeetingController) sendMe(ctx context.Context, data Protocol, me *net.UDPAddr, ln *net.UDPConn) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := ln.WriteToUDP(b, me); err != nil {
		return err
	}
	return nil
}

func (m *MeetingController) broadcast(ctx context.Context, data Protocol, target []net.UDPAddr, ln *net.UDPConn) error {
	eg, ctx := errgroup.WithContext(ctx)
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	for _, v := range target {
		v := v
		eg.Go(func() error {
			if _, err := ln.WriteToUDP(b, &v); err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
