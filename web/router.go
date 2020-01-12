package web

import "github.com/gorilla/mux"

// Router is a router containing the routes a server is to serve
type Router *mux.Router

// Handlers is a bundled set of all the application's HTTP handlers
type Handlers struct {
	Uptime  UptimeHandler
	Counter CounterHandler
}

// ProvideDefaultRouter creates a router using the available handlers
func ProvideDefaultRouter(handlers Handlers) Router {
	router := mux.NewRouter()
	router.Handle("/uptime", handlers.Uptime)
	router.Handle("/counter", handlers.Counter)
	return router
}
