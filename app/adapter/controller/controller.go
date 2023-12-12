package controller

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var mux sync.RWMutex
var buf = make([]byte, 5)

type IController interface {
	Run() error
}
type Controller struct {
	Ln                *net.UDPConn
	MeetingController *MeetingController
}

func NewController(ln *net.UDPConn, meetingController *MeetingController) IController {
	return &Controller{
		Ln:                ln,
		MeetingController: meetingController,
	}
}

func (c *Controller) Run() error {
	mux.Lock()
	defer mux.Unlock()
	n, addr, err := c.Ln.ReadFromUDP(buf)
	if err != nil {
		return err
	}
	c.Ln.WriteToUDP([]byte(fmt.Sprintf("%d, ok\n", 0)), addr)
	log.Println(n)
	return nil
}
