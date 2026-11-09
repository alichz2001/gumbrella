//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"log/slog"

	"github.com/alichz2001/gombrella/cmd/app/server"
	"github.com/alichz2001/gombrella/internal/biz"
	"github.com/alichz2001/gombrella/internal/conf"
	"github.com/alichz2001/gombrella/services/todo/data"
	"github.com/alichz2001/gombrella/services/todo/service"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
