// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/K-Kizuku/plesio-server/app/adapter/controller"
	"github.com/K-Kizuku/plesio-server/app/adapter/gateway"
	"github.com/K-Kizuku/plesio-server/app/driver/ristretto"
	"github.com/K-Kizuku/plesio-server/app/usecase"
	"net"
)

// Injectors from wire.go:

func Init(ln *net.UDPConn) controller.IController {
	iInMemoryCacheRepository := ristretto.NewCacheClient()
	iClientRepository := gateway.NewClientRepository(iInMemoryCacheRepository)
	iRoomRepository := gateway.NewRoomRepository(iInMemoryCacheRepository)
	iMeetingUsecase := usecase.NewMeetingUsecase(iClientRepository, iRoomRepository)
	iMeetingController := controller.NewMeetingContrallor(iMeetingUsecase)
	iController := controller.NewController(ln, iMeetingController)
	return iController
}
