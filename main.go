package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// setup port and shit
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	addr := fmt.Sprintf(":%s", PORT)

	// main handler
	http.HandleFunc("/", ok)

	// serve
	log.Fatal(http.ListenAndServe(addr, nil))
}

func ok(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK!")
}
