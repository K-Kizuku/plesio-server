package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
)

var mux sync.RWMutex
var buf = make([]byte, 5)

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
	_, addr, err := c.LnUDP.ReadFromUDP(buf)
	if err != nil {
		return err
	}
	req := &Protocol{}
	if err := json.Unmarshal(buf, req); err != nil {
		return err
	}
	switch req.Type {
	case "AA":
		if err := c.MeetingController.BroadcastMessage(ctx, req.Header.RoomID, "", addr, c.LnUDP); err != nil {
			return err
		}
	case "audio":
		if err := c.MeetingController.BroadcastAudio(ctx, req.Header.RoomID, "", addr, c.LnUDP); err != nil {
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

	}
	c.LnUDP.WriteToUDP([]byte(fmt.Sprintf("%d, ok\n", 0)), addr)
	return nil
}

func (c *TCPController) Run(ctx context.Context) error {
	mux.Lock()
	defer mux.Unlock()
	n, err := c.LnTCP.Read(buf)
	if err != nil {
		return err
	}
	c.LnTCP.Write([]byte("しめさばくんありがとう"))
	log.Println(n)
	return nil
}
