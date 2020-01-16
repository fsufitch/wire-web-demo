package web

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/fsufitch/wire-web-demo/config"
	"github.com/fsufitch/wire-web-demo/internal/testutil"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func findOpenPort() (int, error) {
	min := 10000
	max := 65535
	attempts := 10

	for i := 0; i < attempts; i++ {
		port := rand.Intn(max-min) + min
		if ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port)); err != nil {
			// Port unavailable
			continue
		} else if err := ln.Close(); err != nil {
			return 0, err
		}
		return port, nil
	}
	return 0, fmt.Errorf("could not find port to use for testing (%d attempts)", attempts)
}

func TestProvideServerRunFunc(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	port, err := findOpenPort()
	if err != nil {
		assert.Fail(t, "could not find a testing port")
	}
	t.Log("Using port", port)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Tested code
	runFunc, _ := ProvideServerRunFunc(logger, config.WebPort(port), mux.NewRouter())
	serverErrChan := make(chan error)
	go func() {
		serverErrChan <- runFunc(ctx)
	}()

	resp, queryErr := http.Get(fmt.Sprintf("http://localhost:%d", port))
	respStatus := resp.StatusCode
	cancel()

	// Asserts

	assert.Nil(t, queryErr)
	assert.Equal(t, http.StatusNotFound, respStatus)
	select {
	case serverErr := <-serverErrChan:
		assert.EqualError(t, serverErr, "server interrupted through context")
	case <-time.After(1 * time.Second):
		assert.Fail(t, "server never quit")
	}
}
