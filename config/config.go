package config

import (
	"errors"
	"os"
	"strconv"
)

// DatabaseString is a string alias containing a database connection string
type DatabaseString string

// WebPort is the port that the web service will run on
type WebPort int

// Config contains all the runtime configuration options
type Config struct {
	databaseString DatabaseString
	webPort        WebPort
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
