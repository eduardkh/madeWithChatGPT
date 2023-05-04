package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go name")
		os.Exit(1)
	}
	name := os.Args[1]
	fmt.Printf("Hello, %s!\n", name)

	// Signal handling code
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	fmt.Println("Shutting down gracefully...")
}
