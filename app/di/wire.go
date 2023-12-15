//go:build wireinject
// +build wireinject

package di

import (
	"net"

	"github.com/K-Kizuku/plesio-server/app/adapter/controller"
	"github.com/K-Kizuku/plesio-server/app/adapter/gateway"
	"github.com/K-Kizuku/plesio-server/app/driver/ristretto"
	"github.com/K-Kizuku/plesio-server/app/usecase"
	"github.com/google/wire"
)

func Init(ln *net.UDPConn) controller.IController {
	wire.Build(
		gateway.NewClientRepository,
		gateway.NewRoomRepository,
		ristretto.NewCacheClient,
		usecase.NewMeetingUsecase,
		controller.NewMeetingContrallor,
		controller.NewController,
		// wire.Bind(new(repository.IClientRepository), new(*gateway.ClientRepository)),
		// wire.Bind(new(repository.IRoomRepository), new(*gateway.RoomRepository)),
		// wire.Bind(new(repository.IInMemoryCacheRepository), new(*ristretto.Client)),
		// wire.Bind(new(usecase.IMeetingUsecase), new(*usecase.MeetingUsecase)),
		// wire.Bind(new(controller.IController), new(*controller.Controller)),
		// wire.Bind(new(controller.IMeetingController), new(*controller.MeetingController)),
	)
	return &controller.Controller{}
}
