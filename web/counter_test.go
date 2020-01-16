package web

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fsufitch/wire-web-demo/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDefaultCounterHandler_ServeHTTP(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()

	dao := testutil.MockCounterDAO{Int: 1234, Error: nil}
	h := DefaultCounterHandler{
		CounterDAO: dao,
		Logger:     logger,
	}
	resp := httptest.NewRecorder()

	// Tested code
	h.ServeHTTP(resp, httptest.NewRequest("GET", "/counter", strings.NewReader("")))

	// Asserts
	assert.Equal(t, "1234", string(resp.Body.Bytes()))
}

func TestDefaultCounterHandler_ServeHTTP_Error(t *testing.T) {
	// Setup
	logger, _, _ := testutil.HeadlessMultiLogger()

	dao := testutil.MockCounterDAO{Int: 0, Error: errors.New("test error")}
	h := DefaultCounterHandler{
		CounterDAO: dao,
		Logger:     logger,
	}
	resp := httptest.NewRecorder()

	// Tested code
	h.ServeHTTP(resp, httptest.NewRequest("GET", "/counter", strings.NewReader("")))

	// Asserts
	assert.Contains(t, string(resp.Body.Bytes()), "Internal error")
}
