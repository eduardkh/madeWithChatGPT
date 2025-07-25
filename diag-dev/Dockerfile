# Dockerfile - MVP production container for network diagnostic portal

# --- Builder Stage ---------------------------------------------------------
FROM golang:1.24-alpine AS builder

WORKDIR /app
# Git for module fetching
RUN apk add --no-cache git

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY cmd cmd
COPY handlers handlers
COPY templates templates
COPY static static

# Build the static binary
RUN go build -o diagportal cmd/main.go


# --- Runtime Stage ---------------------------------------------------------
FROM alpine:3.18

# Install network diagnostic & CLI tools
RUN apk add --no-cache \
    iputils \
    bind-tools \
    drill \
    mtr \
    iproute2 \
    curl \
    jq \
    bash

# Create app directory and copy artifacts
WORKDIR /app
COPY --from=builder /app/diagportal .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Expose port and run
EXPOSE 8080
ENTRYPOINT ["/app/diagportal"]

# --- Build & Run Instructions ------------------------------------------------
# Build the image:
#   docker build -t eduardkh/diag-mvp:latest .
# Run the container:
#   docker run -d --rm -p 8080:8080 eduardkh/diag-mvp:latest
