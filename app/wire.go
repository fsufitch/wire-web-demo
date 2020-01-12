package app

import "github.com/google/wire"

// ApplicationProviderSet is a provider set for building an application
var ApplicationProviderSet = wire.NewSet(
	ProvideApplicationRunFunc,
	ProvideInterruptChannel,
)
