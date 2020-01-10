package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// InjectableServer is an abstract interface encapsulating the web server functionality
type InjectableServer interface {
	Run(ctx context.Context, port int) error
}

type defaultServer struct {
	startTime time.Time
	port      int
	router    *mux.Router
	db        InjectableDatabase
}

// NewDefaultServer creates the default web server for running the demo
func NewDefaultServer(config InjectableConfig, db InjectableDatabase) (InjectableServer, error) {
	port, err := config.WebPort()
	if err != nil {
		return nil, err
	}

	server := defaultServer{
		router: mux.NewRouter(),
		port:   port,
		db:     db,
	}
	server.router.HandleFunc("/uptime", server.handleUptime)
	server.router.HandleFunc("/counter", server.handleCounter)
	return &server, nil
}

func (s *defaultServer) Run(ctx context.Context, port int) error {
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.router,
	}

	errChan := make(chan error)
	go func() {
		s.startTime = time.Now()
		fmt.Printf("Now serving at %s\n", httpServer.Addr)
		errChan <- httpServer.ListenAndServe()
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		httpServer.Close()
		return errors.New("server terminated via interrupt")
	}
}

func (s *defaultServer) handleUptime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	timeDelta := time.Now().Sub(s.startTime).Seconds()
	w.Write([]byte(fmt.Sprintf("%.2f", timeDelta)))
}

func (s *defaultServer) handleCounter(w http.ResponseWriter, r *http.Request) {
	count, err := s.db.IncreaseAndReturnCounter(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Internal error: %v", err)))
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", count)))
}
