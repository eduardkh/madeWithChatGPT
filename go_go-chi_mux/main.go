package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func handleAllMethods(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Received a GET request")
	case http.MethodPost:
		fmt.Fprintf(w, "Received a POST request")
	case http.MethodPut:
		fmt.Fprintf(w, "Received a PUT request")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func main() {
	// Create a new Chi router
	r := chi.NewRouter()

	// Add some middleware (optional)
	r.Use(middleware.Logger)

	// Handle the root path with specific HTTP methods
	r.HandleFunc("/", handleAllMethods)

	// Start the server on port 8090
	fmt.Println("Server is listening on port 8090...")
	http.ListenAndServe(":8090", r)
}
