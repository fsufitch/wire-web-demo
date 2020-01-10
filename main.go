package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
)

func main() {
	mainProcess, err := InitializeMainProcess()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during dependency injection: %v", err)
		os.Exit(1)
	}
	os.Exit(mainProcess.Run())
}

// MainProcess encapsulates the main background process
type MainProcess struct {
	webServer InjectableServer
}

// NewMainProcess creates a new MainProcess
func NewMainProcess(webServer InjectableServer, db InjectableDatabase) (*MainProcess, error) {
	if err := db.InitializeSchema(context.Background()); err != nil {
		return nil, errors.Wrap(err, "failed to initialize db schema")
	}
	return &MainProcess{webServer}, nil
}

// Run starts the server
func (m MainProcess) Run() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	errChan := make(chan error)

	go func() {
		errChan <- m.webServer.Run(ctx, 8080)
	}()

	select {
	case <-interrupt:
		fmt.Println("Interrupt received, shutting down")
	case err := <-errChan:
		fmt.Fprintf(os.Stderr, "Received error during server operation: %v\n", err)
		return 1
	}
	return 0
}
