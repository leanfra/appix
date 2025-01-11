//go:build wireinject
// +build wireinject

package main

// The build tag makes sure the stub is not built in the final build.

import (
	"appix/internal/biz"
	"appix/internal/conf"
	"appix/internal/data"
	"appix/internal/server"
	"appix/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Admin, *conf.Authz, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp))
}
