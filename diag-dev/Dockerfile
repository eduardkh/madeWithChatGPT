# Dockerfile - Development environment for network diagnostic portal
# Base image with Go toolchain and minimal footprint
FROM golang:1.24-alpine

# Install network diagnostic & CLI tools
RUN apk add --no-cache \
    iputils \
    bind-tools \
    drill \
    mtr \
    iproute2 \
    curl \
    jq \
    liboping \
    bash

# iputils: ping
# bind-tools: dig, nslookup
# drill: DNS troubleshooting
# mtr: ping + traceroute
# iproute2: ip, ss for interface & route inspection
# curl: HTTP connectivity checks
# jq: JSON processing
# liboping: library/API for ICMP “ping” in programs (optional)
# bash: shell for development

# Set working directory for mounting host code
WORKDIR /app

# Expose application port (adjust if needed)
EXPOSE 8080

# Default to bash shell for interactive development
ENTRYPOINT ["/bin/bash"]

# Build & Run instructions:
# 1. Build the image:
#    docker build -t diag-dev .
# 2. Run for development (mount host project into /app):
#    docker run -it --rm -p 8080:8080 -v $(pwd):/app diag-dev
# Inside the container, develop and run your Go+Echo app:
#    go run main.go
# Use HTMX via CDN in your templates for zero-build frontend.
