package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Serve the HTML file when root URL is accessed
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Handle the SSE request
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		// Set necessary headers for SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Infinite loop to keep the connection open and send data
		for {
			// Send the current time to the client
			fmt.Fprintf(w, "data: %s\n\n", time.Now().Format(time.RFC1123))

			// Flush the data immediately instead of buffering it
			w.(http.Flusher).Flush()

			// Wait for a second before the next iteration
			time.Sleep(time.Second)
		}
	})

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
