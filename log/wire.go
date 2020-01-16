package log

import "github.com/google/wire"

// StdOutErrMultiLoggerProviderSet provides a MultiLogger to stdout/stderr
var StdOutErrMultiLoggerProviderSet = wire.NewSet(
	ProvideStdOutErrMultiLogger,
)
