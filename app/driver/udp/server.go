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
		log.Print(err)
	}
}

var mux sync.RWMutex
var buf = make([]byte, 70000)

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()
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
	c := di.InitUDP(ln)

	log.Println("Starting udp server...")

	go func() {
		for {
			err = c.Run(ctx)
			if err != nil {
				errCh <- err
			}
		}
	}()

	for {
		select {
		case <-termCh:
			return errors.New("terminated by signal")
		case err = <-errCh:
			log.Println(err)
		case <-ctx.Done():
			_, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return nil
		}
	}
}
