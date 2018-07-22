package routes

import "net/http"

func notfound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 - Not Found", http.StatusNotFound)
}
