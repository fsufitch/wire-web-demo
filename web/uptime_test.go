package web

import (
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/fsufitch/wire-web-demo/config"
	"github.com/fsufitch/wire-web-demo/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestUptimeHandler_ServeHTTP(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()
	initTime := time.Now()
	h := DefaultUptimeHandler{
		InitTime: config.InitTime(initTime),
		Logger:   logger,
	}
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/uptime", strings.NewReader(""))

	// Tested code
	h.ServeHTTP(resp, req)

	// Asserts
	delta, err := strconv.ParseFloat(resp.Body.String(), 64)
	assert.Nil(t, err)
	assert.InDelta(t, 0, delta, 0.1) // should be within 0.1 second of 0
}
