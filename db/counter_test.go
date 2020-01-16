package db

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fsufitch/wire-web-demo/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestCounterDAOValue(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		assert.Failf(t, "failed to open stub db", "%v", err)
	}
	defer mockDB.Close()

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT value FROM counter").WillReturnRows(
		sqlmock.NewRows([]string{"value"}).AddRow(1234),
	)
	mock.ExpectRollback()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Tested code
	dao := PostgresCounterDAO{Log: logger, DB: mockDB}
	value, err := dao.Value(ctx)

	// Asserts
	assert.Nil(t, err)
	assert.Equal(t, 1234, value)

	if err := mock.ExpectationsWereMet(); err != nil {
		assert.Failf(t, "there were unfulfilled expectations", "%v", err)
	}
}

func TestCounterDAOIncrement(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		assert.Failf(t, "failed to open stub db", "%v", err)
	}
	defer mockDB.Close()

	mock.ExpectBegin()
	mock.ExpectQuery("UPDATE counter").WillReturnRows(
		sqlmock.NewRows([]string{"value"}).AddRow(1234),
	)
	mock.ExpectCommit()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Tested code
	dao := PostgresCounterDAO{Log: logger, DB: mockDB}
	value, err := dao.Increment(ctx)

	// Asserts
	assert.Nil(t, err)
	assert.Equal(t, 1234, value)

	if err := mock.ExpectationsWereMet(); err != nil {
		assert.Failf(t, "there were unfulfilled expectations", "%v", err)
	}
}
