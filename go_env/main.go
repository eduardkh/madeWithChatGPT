package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", handleEnv)
	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func handleEnv(w http.ResponseWriter, r *http.Request) {
	// Get all environment variables
	env := os.Environ()

	// Create a formatted string of all environment variables
	var result strings.Builder
	result.WriteString("<html><body><h1>Environment Variables</h1><pre>")
	for _, e := range env {
		result.WriteString(e + "\n")
	}
	result.WriteString("</pre></body></html>")

	// Set content type and send response
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, result.String())
}
