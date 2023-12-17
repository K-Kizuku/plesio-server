package main

import (
	"github.com/K-Kizuku/plesio-server/app/driver/udp"
	"github.com/K-Kizuku/plesio-server/utils/config"
)

func main() {
	config.LoadEnv()
	udp.Server()
}
