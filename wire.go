// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/fsufitch/testable-web-demo/app"
	"github.com/fsufitch/testable-web-demo/config"
	"github.com/fsufitch/testable-web-demo/db"
	"github.com/fsufitch/testable-web-demo/web"
	"github.com/google/wire"
)

var defaultProviderSet = wire.NewSet(
	app.ApplicationProviderSet,
	config.EnvironmentProviderSet,
	db.PostgresDBProviderSet,
	web.DefaultWebProviderSet,
)

func InitializeApplicationRunFunc() (app.ApplicationRunFunc, func(), error) {
	panic(wire.Build(defaultProviderSet))
}
