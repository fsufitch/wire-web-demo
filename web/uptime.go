package web

import (
	"fmt"
	"net/http"
	"time"
)

// UptimeHandler is a handler that writes the server uptime in plaintext
type UptimeHandler http.Handler

// DefaultUptimeHandler is an implementation of UptimeHandler that remembers its init time
type DefaultUptimeHandler struct {
	initTime time.Time
}

// ProvideDefaultUptimeHandler creates a DefaultUptimeHandler based on the current time
func ProvideDefaultUptimeHandler() DefaultUptimeHandler {
	return DefaultUptimeHandler{time.Now()}
}

func (h DefaultUptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	timeDelta := time.Now().Sub(h.initTime).Seconds()
	w.Write([]byte(fmt.Sprintf("%.2f", timeDelta)))
}
