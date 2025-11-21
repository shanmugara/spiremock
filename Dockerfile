# Multi-stage Dockerfile for building spiremock
# Stage 1: build the Go binary
FROM golang:1.25.3-alpine AS builder

# Install git (needed if modules are fetched) and ca-certificates for TLS
RUN apk add --no-cache git

WORKDIR /src

# Cache modules
COPY go.mod go.sum ./

# Copy the rest of the source
COPY . .

# If the repository contains a checked-in vendor/ directory it can cause
# "inconsistent vendoring" errors inside the container. Remove vendor here
# so the container build uses module mode and fetches dependencies normally.
RUN rm -rf vendor || true

# Ensure module mode (ignore vendor) when running go commands in this image
ENV GOFLAGS=-mod=mod

# Download modules explicitly to cache them in Docker layer
RUN go mod download

RUN mkdir "/app"

# Build only the main package to avoid building unrelated packages (and to match `go build` usage)
# Build a small static binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags "-s -w" -o /app/spiremock .

# Stage 2: minimal runtime image
FROM alpine:3.18
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/spiremock /usr/local/bin/spiremock

# Drop privileges
USER 65532:65532

ENTRYPOINT ["/usr/local/bin/spiremock"]
