package controller

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"strings"
	"sync"
)

var mux sync.RWMutex

type IController interface {
	Run(ctx context.Context) error
}
type UDPController struct {
	LnUDP             *net.UDPConn
	MeetingController IMeetingController
}

type TCPController struct {
	LnTCP             *net.TCPConn
	MeetingController IMeetingController
}

func NewUDPController(lnUDP *net.UDPConn, meetingController IMeetingController) IController {
	return &UDPController{
		LnUDP:             lnUDP,
		MeetingController: meetingController,
	}
}

func NewTCPController(lnTCP *net.TCPConn, meetingController IMeetingController) IController {
	return &TCPController{
		LnTCP:             lnTCP,
		MeetingController: meetingController,
	}
}

func (c *UDPController) Run(ctx context.Context) error {
	mux.Lock()
	defer mux.Unlock()
	buf := make([]byte, 70000)
	_, addr, err := c.LnUDP.ReadFromUDP(buf)
	if err != nil {
		return err
	}
	req := &Protocol{}
	b := strings.Trim(string(buf), "\x00")

	// buf := strings.Replace(string(buf), "\000", "", -1)
	if err := json.Unmarshal([]byte(b), req); err != nil {
		log.Print(buf)
		return err
	}
	switch req.Type {
	case "AA":
		if err := c.MeetingController.BroadcastMessage(ctx, req.Header.RoomID, req.Body.Content, addr, c.LnUDP); err != nil {
			return err
		}
	case "audio":
		if err := c.MeetingController.BroadcastAudio(ctx, req.Header.RoomID, req.Body.Content, addr, c.LnUDP); err != nil {
			return err
		}
	case "comment":
		return nil
	case "create_room":
		if err := c.MeetingController.CreateRoom(ctx, addr, c.LnUDP); err != nil {
			return err
		}
	case "join_room":
		if err := c.MeetingController.JoinRoom(ctx, req.Header.RoomID, addr, c.LnUDP); err != nil {
			return err
		}
	case "exit_room":
		if err := c.MeetingController.ExitRoom(ctx, req.Header.RoomID, addr, c.LnUDP); err != nil {
			return err
		}
	case "select_presenter":
		if err := c.MeetingController.SelectPresenter(ctx, req.Header.RoomID, req.Header.WantClientID, c.LnUDP); err != nil {
			return err
		}

	}
	return nil
}

func (c *TCPController) Run(ctx context.Context) error {
	// mux.Lock()
	// defer mux.Unlock()
	// n, err := c.LnTCP.Read(buf)
	// if err != nil {
	// 	return err
	// }
	// c.LnTCP.Write([]byte("しめさばくんありがとう"))
	// log.Println(n)
	return nil
}
