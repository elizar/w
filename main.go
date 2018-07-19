package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/elizar/w/tmpl"
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
	data, err := tmpl.Asset("layout.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	layout, _ := template.New("").Parse(string(data))
	layout.Execute(w, struct{ Title, Name string }{"home", "tibur"})
}
