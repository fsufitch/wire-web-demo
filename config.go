package main

import (
	"os"
	"strconv"
)

// InjectableConfig is an abstract interface providing configuration to the program
type InjectableConfig interface {
	Database() string
	WebPort() (int, error)
}

// NewConfigFromEnvironment creates an InjectableConfig based on environment variables
func NewConfigFromEnvironment() InjectableConfig {
	return &envConfig{}
}

type envConfig struct{}

func (c envConfig) Database() string {
	return os.Getenv("DATABASE")
}

func (c envConfig) WebPort() (int, error) {
	portString, ok := os.LookupEnv("PORT")
	if !ok {
		portString = "8080"
	}
	port, err := strconv.ParseInt(portString, 0, 0)
	return int(port), err
}
