package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseString(t *testing.T) {
	// Setup
	os.Setenv("DATABASE", "test")

	// Tested code
	db, err := ProvideDatabaseStringFromEnvironment()

	// Asserts
	assert.Nil(t, err)
	assert.Equal(t, "test", string(db))
}

func TestDatabaseString_Missing(t *testing.T) {
	// Setup
	os.Unsetenv("DATABASE")

	// Tested code
	_, err := ProvideDatabaseStringFromEnvironment()

	// Asserts
	assert.NotNil(t, err)
}

func TestWebPort(t *testing.T) {
	// Setup
	os.Setenv("PORT", "1234")

	// Tested code
	port, err := ProvideWebPortFromEnvironment()

	// Asserts
	assert.Nil(t, err)
	assert.Equal(t, 1234, int(port))
}

func TestWebPort_Default(t *testing.T) {
	// Setup
	os.Unsetenv("PORT")

	// Tested code
	port, err := ProvideWebPortFromEnvironment()

	// Asserts
	assert.Nil(t, err)
	assert.Equal(t, 8080, int(port))
}

func TestWebPort_Invalid(t *testing.T) {
	// Setup
	os.Setenv("PORT", "Not a real port")

	// Tested code
	_, err := ProvideWebPortFromEnvironment()

	// Asserts
	assert.NotNil(t, err)
}

func TestDebugMode(t *testing.T) {
	// Setup
	os.Setenv("DEBUG", "1")

	// Tested code
	debug, err := ProvideDebugModeFromEnvironment()

	// Asserts
	assert.Nil(t, err)
	assert.True(t, bool(debug))
}

func TestDebugMode_Default(t *testing.T) {
	// Setup
	os.Unsetenv("DEBUG")

	// Tested code
	debug, err := ProvideDebugModeFromEnvironment()

	// Asserts
	assert.Nil(t, err)
	assert.False(t, bool(debug))
}

func TestDebugMode_Invalid(t *testing.T) {
	// Setup
	os.Setenv("DEBUG", "not a real bool")

	// Tested code
	_, err := ProvideDebugModeFromEnvironment()

	// Asserts
	assert.NotNil(t, err)
}

func TestInitTime(t *testing.T) {
	// Tested Code
	initTime := ProvideInitTimeFromCurrentTime()

	// Asserts
	assert.WithinDuration(t, time.Now(), time.Time(initTime), 5*time.Millisecond)
}
