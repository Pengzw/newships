//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"newships/app/interface/internal/biz"
	"newships/app/interface/internal/conf"
	"newships/app/interface/internal/data"
	"newships/app/interface/internal/server"
	"newships/app/interface/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(
	*conf.Server,
	*conf.Data,
	*conf.Registry,
	*conf.Auth,
	log.Logger,
	*tracesdk.TracerProvider,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp,
	))
}
