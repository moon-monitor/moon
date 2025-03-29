//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/moon-monitor/moon/cmd/houyi/internal/conf"
	"github.com/moon-monitor/moon/cmd/houyi/internal/server"
	"github.com/moon-monitor/moon/cmd/houyi/internal/service"
)

// wireApp init wired
func wireApp(bc *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		service.ProviderSetService,
		server.ProviderSetServer,
		newApp,
	))
}
