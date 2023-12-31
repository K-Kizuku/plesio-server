package udp

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/K-Kizuku/plesio-server/app/di"
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
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGKILL, syscall.SIGINT)
	errCh := make(chan error, 1)

	tcpAddr := &net.TCPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 8089,
	}
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	conn, err := ln.AcceptTCP()
	c := di.InitTCP(conn)

	log.Println("Starting udp server...")

	go func() {
		for {
			err = c.Run(ctx)
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
	case <-ctx.Done():
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
	}
	return nil
}
