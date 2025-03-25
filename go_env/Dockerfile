# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy files
COPY . .

# Build the binary with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM scratch

# Copy the binary from builder
COPY --from=builder /app/main /main

# Expose port 8080
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/main"]