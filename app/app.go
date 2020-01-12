package app

import (
	"context"
	"fmt"
	"os"

	"github.com/fsufitch/wire-web-demo/web"
)

// ApplicationRunFunc is a plain function that runs an application and returns an error
type ApplicationRunFunc func() error

// ProvideApplicationRunFunc creates an ApplicationRunFunc that runs a webserver and stops on interrupt
func ProvideApplicationRunFunc(runServer web.ServerRunFunc, interrupt InterruptChannel) ApplicationRunFunc {
	return func() error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		errChan := make(chan error)

		go func() {
			errChan <- runServer(ctx)
		}()

		select {
		case <-interrupt:
			fmt.Println("Interrupt received, shutting down")
			cancel()
		case err := <-errChan:
			fmt.Fprintf(os.Stderr, "fatal server error: %v\n", err)
			return err
		}
		return nil
	}
}
