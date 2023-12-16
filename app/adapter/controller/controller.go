package controller

import (
	"context"
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
	n, addr, err := c.LnUDP.ReadFromUDP(buf)
	if err != nil {
		return err
	}
	c.LnUDP.WriteToUDP([]byte(fmt.Sprintf("%d, ok\n", 0)), addr)
	log.Println(n)
	return nil
}

func (c *TCPController) Run(ctx context.Context) error {
	mux.Lock()
	defer mux.Unlock()
	n, err := c.LnTCP.Read(buf)
	if err != nil {
		return err
	}
	c.LnTCP.Write([]byte(fmt.Sprintf("%d, ok\n", 0)))
	log.Println(n)
	return nil
}
