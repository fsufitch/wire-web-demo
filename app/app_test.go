package app

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/fsufitch/wire-web-demo/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestAppInterrupted(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	runServer := func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(50 * time.Millisecond):
			return errors.New("server completed without interruption")
		}
	}
	interrupt := make(chan os.Signal, 2)
	interrupt <- os.Interrupt

	// Tested code
	runFunc := ProvideApplicationRunFunc(logger, runServer, interrupt)
	appErrorChan := make(chan error)
	go func() {
		appErrorChan <- runFunc()
	}()

	// Asserts
	select {
	case err := <-appErrorChan:
		assert.Nil(t, err)
	case <-time.After(500 * time.Millisecond):
		t.Error("server never shut down")
	}
}

func TestAppFatalError(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	runServer := func(ctx context.Context) error {
		return errors.New("server had an (expected) fatal error")
	}
	interrupt := make(chan os.Signal, 2)

	// Tested code
	runFunc := ProvideApplicationRunFunc(logger, runServer, interrupt)
	appErrorChan := make(chan error)
	go func() {
		appErrorChan <- runFunc()
	}()

	// Asserts
	select {
	case err := <-appErrorChan:
		assert.EqualError(t, err, "server had an (expected) fatal error")
	case <-time.After(500 * time.Millisecond):
		t.Error("server never shut down")
	}
}
