# Server-Sent Events (SSE) Example

## Overview

This project demonstrates how to implement Server-Sent Events (SSE) in Go and reflect those events in a simple HTML page. SSE is a simple and efficient protocol for creating a unidirectional stream from the server to the browser over a single HTTP connection.

## Why Use SSE?

- **Real-time Updates**: SSE is perfect for applications that require real-time updates from the server, such as notifications or live feed updates.
- **Simplicity**: It's easier to implement compared to WebSockets and works over existing HTTP protocols.
- **Efficiency**: SSE maintains a single open connection for all events, reducing the overhead of opening multiple HTTP connections for polling.

## How to Run

1. Run the Go server with `go run main.go`.
2. Open `index.html` in your browser.

## Files

- `main.go`: The Go server that sends timestamp events to the client.
- `index.html`: HTML file that uses JavaScript to listen for SSE and updates the page accordingly.

## Further Reading

For more information on SSE, check the [Mozilla Developer Network documentation](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events).
