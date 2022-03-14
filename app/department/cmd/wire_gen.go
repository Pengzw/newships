// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"newships/app/department/internal/biz"
	"newships/app/department/internal/conf"
	"newships/app/department/internal/data"
	"newships/app/department/internal/server"
	"newships/app/department/internal/service"
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
	deptRepo := data.NewDeptRepo(dataData)
	deptUsecase := biz.NewDeptUsecase(deptRepo)
	deptService := service.NewDeptService(deptUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, tracerProvider, deptService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
