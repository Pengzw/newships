// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"newships/app/user/internal/biz"
	"newships/app/user/internal/conf"
	"newships/app/user/internal/data"
	"newships/app/user/internal/server"
	"newships/app/user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

func initApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	userUsecase := biz.NewUserUsecase(userRepo)
	userService := service.NewUserService(userUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, tracerProvider, userService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}