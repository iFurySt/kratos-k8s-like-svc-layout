//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/ifuryst/kratos-k8s-like-layout/internal/biz"
	"github.com/ifuryst/kratos-k8s-like-layout/internal/conf"
	"github.com/ifuryst/kratos-k8s-like-layout/internal/data"
	"github.com/ifuryst/kratos-k8s-like-layout/internal/server"
	"github.com/ifuryst/kratos-k8s-like-layout/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
