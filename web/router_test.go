package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	// Setup
	var uptimeCalled, counterCalled bool
	stubUptimeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uptimeCalled = true
	})
	stubCounterHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counterCalled = true
	})
	handlers := Handlers{
		Uptime:  stubUptimeHandler,
		Counter: stubCounterHandler,
	}
	router := (*mux.Router)(ProvideDefaultRouter(handlers))

	// Tested code 1
	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/uptime", strings.NewReader("")))

	// Asserts 1
	assert.True(t, uptimeCalled)
	assert.False(t, counterCalled)

	// Tested code 2
	uptimeCalled = false
	counterCalled = false
	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/counter", strings.NewReader("")))

	// Asserts 2
	assert.False(t, uptimeCalled)
	assert.True(t, counterCalled)
}
