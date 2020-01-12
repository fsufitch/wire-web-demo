package web

import (
	"fmt"
	"net/http"

	"github.com/fsufitch/testable-web-demo/db"
)

// CounterHandler is a handler that increments the counter and displays the new value
type CounterHandler http.Handler

// DefaultCounterHandler is a default implementation of CounterHandler
type DefaultCounterHandler struct {
	CounterDAO db.CounterDAO
}

func (h DefaultCounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count, err := h.CounterDAO.Increment(r.Context())
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
