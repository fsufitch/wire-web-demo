package main

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// InjectableDatabase is an abstract interface for database interaction
type InjectableDatabase interface {
	InitializeSchema(ctx context.Context) error
	IncreaseAndReturnCounter(ctx context.Context) (int, error)
}

type livePQDatabase struct {
	db *sql.DB
}

// NewPQDatabase creates an InjectableDatabase interacting with a PostgreSQL DB
func NewPQDatabase(config InjectableConfig) (InjectableDatabase, error) {
	db, err := sql.Open("postgres", config.Database())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &livePQDatabase{db}, nil
}

func (d livePQDatabase) InitializeSchema(ctx context.Context) error {
	tx, err := d.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS counter (value int NOT NULL DEFAULT 0);
		INSERT INTO counter (value)
			SELECT 0 WHERE NOT EXISTS (SELECT * FROM counter);
	`)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (d livePQDatabase) IncreaseAndReturnCounter(ctx context.Context) (int, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return 0, err
	}

	row := tx.QueryRowContext(ctx, `
		UPDATE counter SET value = value + 1 WHERE 1=1 RETURNING value
	`)
	var value int

	if err := row.Scan(&value); err != nil {
		return 0, err
	}
	return value, tx.Commit()
}
