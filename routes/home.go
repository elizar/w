package routes

import (
	"html/template"
	"net/http"

	"github.com/elizar/w/tmpl"
)

func home(w http.ResponseWriter, r *http.Request) {
	lb, err := tmpl.Asset("layout.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	layout, _ := template.New("").Parse(string(lb))
	layout.Execute(w, struct{ Title, Name string }{"Home", "Universe"})
}
