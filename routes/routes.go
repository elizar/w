package routes

import (
	"net/http"

	"github.com/gorilla/pat"
)

// RegisterRoutes ...
func RegisterRoutes(r *pat.Router) {
	// home and 404 handler
	r.HandleFunc("/", home).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notfound)

	// rest of the handlers
	r.Get("/signup", signup)
}
