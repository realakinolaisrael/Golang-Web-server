package handler

import (
	"fmt"
	"net/http"
)

// HomeHandler handles requests to the root path
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// AboutHandler handles requests to the /about path
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page")
}

// NotFoundHandler handles requests to undefined paths
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}
