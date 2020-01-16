package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fsufitch/wire-web-demo/config"
	"github.com/fsufitch/wire-web-demo/log"
)

// UptimeHandler is a handler that writes the server uptime in plaintext
type UptimeHandler http.Handler

// DefaultUptimeHandler is an implementation of UptimeHandler that remembers its init time
type DefaultUptimeHandler struct {
	InitTime config.InitTime
	Logger   *log.MultiLogger
}

func (h DefaultUptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	timeDelta := time.Now().Sub(time.Time(h.InitTime)).Seconds()
	h.Logger.Infof("200 %s %s %s", r.Method, r.URL.String(), r.UserAgent())
	w.Write([]byte(fmt.Sprintf("%.2f", timeDelta)))
}
