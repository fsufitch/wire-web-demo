package db

import (
	"context"
	"database/sql"

	"github.com/fsufitch/wire-web-demo/log"
)

// CounterDAO is a generic interface for a DAO accessing a counter
type CounterDAO interface {
	Value(ctx context.Context) (int, error)
	Increment(ctx context.Context) (int, error)
}

// PostgresCounterDAO implements a CounterDAO for a Postgres database
type PostgresCounterDAO struct {
	DB  PostgresDBConn
	Log *log.MultiLogger
}

// Value retrieves the value of the counter
func (dao PostgresCounterDAO) Value(ctx context.Context) (int, error) {
	dao.Log.Debugf("retrieving counter value")
	tx, err := (*sql.DB)(dao.DB).BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()
	if err != nil {
		return 0, err
	}

	row := tx.QueryRow(`
		SELECT value FROM counter LIMIT 1
	`)
	var value int

	if err := row.Scan(&value); err != nil {
		return 0, err
	}
	return value, tx.Commit()
}

// Increment increases the counter and returns its new value
func (dao PostgresCounterDAO) Increment(ctx context.Context) (int, error) {
	dao.Log.Debugf("incrementing counter value")
	tx, err := (*sql.DB)(dao.DB).BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return 0, err
	}

	row := tx.QueryRow(`
		UPDATE counter SET value = value + 1 WHERE 1=1 RETURNING value
	`)
	var value int

	if err := row.Scan(&value); err != nil {
		return 0, err
	}
	return value, tx.Commit()
}
