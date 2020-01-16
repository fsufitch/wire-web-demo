package db

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fsufitch/wire-web-demo/internal/testutil"
	"github.com/stretchr/testify/assert"
)

const testDBString = "postgres://postgres@test.dummy:5432/test?sslmode=disable"

func TestPreInitProviderConnectsAndPings(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		assert.Failf(t, "failed to open stub db", "%v", err)
	}
	defer mockDB.Close()
	sqlOpen = func(string, string) (*sql.DB, error) {
		return mockDB, nil
	}
	mock.ExpectPing()

	// Tested code
	db, cleanup, err := ProvidePreInitPostgresDBConn(logger, testDBString)

	// Asserts
	assert.Nil(t, err)
	assert.NotNil(t, db)
	assert.NotNil(t, cleanup)

	mock.ExpectClose()
	cleanup()

	if err := mock.ExpectationsWereMet(); err != nil {
		assert.Failf(t, "there were unfulfilled expectations", "%v", err)
	}
}

func TestPreInitProviderConnectsAndPings_ConnectFail(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	sqlOpen = func(string, string) (*sql.DB, error) {
		return nil, errors.New("expected connection failure")
	}

	// Tested code
	db, cleanup, err := ProvidePreInitPostgresDBConn(logger, testDBString)

	// Asserts
	assert.Nil(t, db)
	assert.Nil(t, cleanup)
	assert.Error(t, err)
}

func TestPreInitProviderConnectsAndPings_PingFail(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		assert.Failf(t, "failed to open stub db", "%v", err)
	}
	defer mockDB.Close()
	sqlOpen = func(string, string) (*sql.DB, error) {
		return mockDB, nil
	}
	mock.ExpectPing().WillReturnError(errors.New("expected ping failure"))

	// Tested code
	db, cleanup, err := ProvidePreInitPostgresDBConn(logger, testDBString)

	// Asserts
	assert.EqualError(t, err, "expected ping failure")
	assert.Nil(t, db)
	assert.NotNil(t, cleanup)

	mock.ExpectClose()
	cleanup()

	if err := mock.ExpectationsWereMet(); err != nil {
		assert.Failf(t, "there were unfulfilled expectations", "%v", err)
	}
}

func TestProvidePostgresDBConn(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		assert.Failf(t, "failed to open stub db", "%v", err)
	}
	defer mockDB.Close()

	mock.ExpectBegin()
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS counter").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("INSERT INTO counter").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Tested code
	db, err := ProvidePostgresDBConn(logger, mockDB)

	// Asserts
	assert.Nil(t, err)
	assert.NotNil(t, db)

	if err := mock.ExpectationsWereMet(); err != nil {
		assert.Failf(t, "there were unfulfilled expectations", "%v", err)
	}
}

func TestProvidePostgresDBConn_QueryError(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		assert.Failf(t, "failed to open stub db", "%v", err)
	}
	defer mockDB.Close()

	mock.ExpectBegin()
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS counter").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("INSERT INTO counter").WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	// Tested code
	db, err := ProvidePostgresDBConn(logger, mockDB)

	// Asserts
	assert.NotNil(t, err)
	assert.Nil(t, db)

	if err := mock.ExpectationsWereMet(); err != nil {
		assert.Failf(t, "there were unfulfilled expectations", "%v", err)
	}
}
