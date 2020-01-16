package web

import "github.com/google/wire"

// ProvideDefaultCounterHandler is a provider set for wiring a default counter handler
var ProvideDefaultCounterHandler = wire.NewSet(
	wire.Struct(new(DefaultCounterHandler), "*"),
	wire.Bind(new(CounterHandler), new(*DefaultCounterHandler)),
)

// ProvideDefaultUptimeHandler is a provider set for wiring a default uptime handler
var ProvideDefaultUptimeHandler = wire.NewSet(
	wire.Struct(new(DefaultUptimeHandler), "*"),
	wire.Bind(new(UptimeHandler), new(*DefaultUptimeHandler)),
)

// ProvideDefaultHandlers is a provider set for wiring handlers to default implementations
var ProvideDefaultHandlers = wire.NewSet(
	wire.Struct(new(Handlers), "*"),
)

// DefaultWebProviderSet is a provider set for building a web server
var DefaultWebProviderSet = wire.NewSet(
	ProvideServerRunFunc,
	ProvideDefaultRouter,
	ProvideDefaultHandlers,
	ProvideDefaultUptimeHandler,
	ProvideDefaultCounterHandler,
)
