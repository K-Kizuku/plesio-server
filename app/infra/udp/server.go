package udp

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Server() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

var mux sync.RWMutex
var buf = make([]byte, 5)
var i = 1

func run() error {
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGKILL, syscall.SIGINT)
	errCh := make(chan error, 1)

	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 8088,
	}
	ln, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	log.Println("Starting udp server...")

	go func() {
		for {
			err = handle(ln)
			if err != nil {
				errCh <- err
			}
		}
	}()

	select {
	case <-termCh:
		return errors.New("terminated by signal")
	case err = <-errCh:
		return err
	}
}

func handle(ln *net.UDPConn) error {
	mux.Lock()
	defer mux.Unlock()
	n, addr, err := ln.ReadFromUDP(buf)
	if err != nil {
		return err
	}
	ln.WriteToUDP([]byte(fmt.Sprintf("%d, ok\n", i)), addr)
	log.Println(n)
	i++
	return nil
}
