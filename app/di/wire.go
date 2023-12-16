//go:build wireinject
// +build wireinject

package di

import (
	"net"

	"github.com/K-Kizuku/plesio-server/app/adapter/controller"
	"github.com/K-Kizuku/plesio-server/app/adapter/gateway"
	"github.com/K-Kizuku/plesio-server/app/driver/redis"
	"github.com/K-Kizuku/plesio-server/app/driver/ristretto"
	"github.com/K-Kizuku/plesio-server/app/usecase"
	"github.com/google/wire"
)

func InitUDP(ln *net.UDPConn) controller.IController {
	wire.Build(
		gateway.NewClientRepository,
		gateway.NewRoomRepository,
		ristretto.NewCacheClient,
		redis.NewDataStoreClient,
		usecase.NewMeetingUsecase,
		controller.NewMeetingContrallor,
		controller.NewUDPController,
	)
	return &controller.UDPController{}
}

func InitTCP(ln *net.TCPConn) controller.IController {
	wire.Build(
		gateway.NewClientRepository,
		gateway.NewRoomRepository,
		ristretto.NewCacheClient,
		redis.NewDataStoreClient,
		usecase.NewMeetingUsecase,
		controller.NewMeetingContrallor,
		controller.NewTCPController,
	)
	return &controller.TCPController{}
}
