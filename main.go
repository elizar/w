package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// main handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Beep, Boop!", 200)
	})

	// setup port and shit
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	addr := fmt.Sprintf(":%s", PORT)

	// serve
	log.Println("Server running on port", PORT)
	log.Fatal(http.ListenAndServe(addr, nil))
}
