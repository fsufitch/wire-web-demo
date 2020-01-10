//+build wireinject

package main

import "github.com/google/wire"

var DefaultProviderSet = wire.NewSet(NewDefaultServer, NewPQDatabase, NewMainProcess, NewConfigFromEnvironment)

func InitializeMainProcess() (*MainProcess, error) {
	wire.Build(DefaultProviderSet)
	return &MainProcess{}, nil
}
