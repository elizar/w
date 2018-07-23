package main

import (
	"log"
	"net/http"
	"os"

	"github.com/elizar/w/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/pat"
)

func main() {
	r := pat.New()

	// register routes
	routes.RegisterRoutes(r)

	// Use gorilla logger if server is not mounted
	// through `up start`
	if os.Getenv("UP_STAGE") == "" {
		h := handlers.LoggingHandler(os.Stdout, r)
		http.Handle("/", h)
	} else {
		http.Handle("/", r)
	}

	// port
	port := ":" + os.Getenv("PORT")

	// listen and serve
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
