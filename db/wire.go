package db

import "github.com/google/wire"

// ProvidePostgresCounterDAO is a provider set for building a PostgresCounterDAO
var ProvidePostgresCounterDAO = wire.NewSet(
	wire.Struct(new(PostgresCounterDAO), "DB", "Log"),
	wire.Bind(new(CounterDAO), new(*PostgresCounterDAO)),
)

// PostgresDBProviderSet is a provider set for building a Postgres database
var PostgresDBProviderSet = wire.NewSet(
	ProvidePostgresDBConn,
	ProvidePreInitPostgresDBConn,
	ProvidePostgresCounterDAO,
)
