package main

import (
	"fmt"
	"net/http"
)

// Handler function to handle different HTTP methods
func handleAllMethods(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Received a GET request")
	case http.MethodPost:
		fmt.Fprintf(w, "Received a POST request")
	case http.MethodPut:
		fmt.Fprintf(w, "Received a PUT request")
	case http.MethodDelete:
		fmt.Fprintf(w, "Received a DELETE request")
	case http.MethodPatch:
		fmt.Fprintf(w, "Received a PATCH request")
	case http.MethodOptions:
		fmt.Fprintf(w, "Received an OPTIONS request")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func main() {
	// Create a new ServeMux multiplexer
	mux := http.NewServeMux()

	// Register the handler function with the multiplexer
	mux.HandleFunc("/", handleAllMethods)

	// Start the server on port 8090 using the new multiplexer
	fmt.Println("Server is listening on port 8090...")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
