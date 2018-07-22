package routes

import (
	"fmt"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sign Up")
}
