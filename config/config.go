package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

// DatabaseString is a string alias containing a database connection string
type DatabaseString string

// WebPort is the port that the web service will run on
type WebPort int

// DebugMode is whether a debug state is set
type DebugMode bool

// InitTime is the program initialization time
type InitTime time.Time

// Config contains all the runtime configuration options
type Config struct {
	databaseString DatabaseString
	webPort        WebPort
	debugIsSet     DebugMode
	initTime       InitTime
}

// ProvideDatabaseStringFromEnvironment creates a DatabaseString from the environment, or errors when it's missing
func ProvideDatabaseStringFromEnvironment() (DatabaseString, error) {
	if envValue, ok := os.LookupEnv("DATABASE"); ok {
		return DatabaseString(envValue), nil
	}
	return "", errors.New("missing env var: DATABASE")

}

// ProvideWebPortFromEnvironment creates a WebPort from the environment, defaulting to 8080 when missing
func ProvideWebPortFromEnvironment() (WebPort, error) {
	portString, ok := os.LookupEnv("PORT")
	if !ok {
		portString = "8080"
	}
	port, err := strconv.ParseInt(portString, 0, 0)
	return WebPort(port), err
}

// ProvideDebugModeFromEnvironment creates a DebugMode based on the value in the DEBUG env var
func ProvideDebugModeFromEnvironment() (DebugMode, error) {
	debugString, ok := os.LookupEnv("DEBUG")
	if !ok {
		return false, nil
	}
	mode, err := strconv.ParseBool(debugString)
	return DebugMode(mode), err
}

// ProvideInitTimeFromCurrentTime initialized the program InitTime from the current time
func ProvideInitTimeFromCurrentTime() InitTime {
	return InitTime(time.Now())
}
