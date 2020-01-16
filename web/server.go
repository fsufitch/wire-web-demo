package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/fsufitch/wire-web-demo/config"
	"github.com/fsufitch/wire-web-demo/log"
	"github.com/gorilla/mux"
)

// ServerRunFunc is a function that starts a blocking server; it returns an error if the server crashed
type ServerRunFunc func(context.Context) error

// ProvideServerRunFunc provides a ServerRunFunc
func ProvideServerRunFunc(logger *log.MultiLogger, port config.WebPort, router Router) (ServerRunFunc, func()) {
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: (*mux.Router)(router),
	}

	cleanup := func() {
		logger.Debugf("shutting down web server")
		httpServer.Close()
	}

	return func(ctx context.Context) error {
		errChan := make(chan error)
		go func() {
			logger.Infof("Now serving at %s\n", httpServer.Addr)
			errChan <- httpServer.ListenAndServe()
		}()

		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			httpServer.Close()
			return errors.New("server interrupted through context")
		}
	}, cleanup
}
