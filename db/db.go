package db

import (
	"context"
	"database/sql"

	"github.com/fsufitch/wire-web-demo/config"
	"github.com/fsufitch/wire-web-demo/log"
	_ "github.com/lib/pq" // inject PQ database driver
)

// PostgresDBConn is a database connection to a Postgres DB
type PostgresDBConn *sql.DB

// PreInitPostgresDBConn is a database connection to a Postgres DB which may not have initialized schema
type PreInitPostgresDBConn *sql.DB

var sqlOpen = sql.Open

// ProvidePreInitPostgresDBConn provides a PostgresDBConn by connecting to a database
func ProvidePreInitPostgresDBConn(logger *log.MultiLogger, dbString config.DatabaseString) (PreInitPostgresDBConn, func(), error) {
	logger.Infof("connecting to postgres database")
	db, err := sqlOpen("postgres", string(dbString))
	cleanup := func() { db.Close() }

	if err != nil {
		return nil, nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, cleanup, err
	}
	return db, cleanup, nil
}

const setupTableQuery = `
		CREATE TABLE IF NOT EXISTS counter (value int NOT NULL DEFAULT 0);
		`
const setupRowQuery = `
		INSERT INTO counter (value)
		SELECT 0 WHERE NOT EXISTS (SELECT * FROM counter);`

// ProvidePostgresDBConn performs schema initialization
func ProvidePostgresDBConn(logger *log.MultiLogger, db PreInitPostgresDBConn) (PostgresDBConn, error) {
	logger.Infof("initializing db schema")
	tx, err := (*sql.DB)(db).BeginTx(context.Background(), nil)
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	if _, err = tx.Exec(setupTableQuery); err != nil {
	} else if _, err = tx.Exec(setupRowQuery); err != nil {
	}

	if err != nil {
		return nil, err
	}

	return PostgresDBConn(db), tx.Commit()
}
