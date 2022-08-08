package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

type Middleware struct{}

// Basic Middleware functions to add basic headers
func (m Middleware) BasicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

// Basic Server logger for every URI mock
func (m Middleware) Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serverLog := fmt.Sprintf("| %v | %v", r.Method, r.URL.Path)
		log.Println(serverLog)
		h.ServeHTTP(w, r)
	})
}
