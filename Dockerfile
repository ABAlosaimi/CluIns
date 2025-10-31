# Start from the official Golang image for building
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o cluins ./main

# Use a minimal image for running
FROM alpine:latest
WORKDIR /app

# Copy the built binary from the builder
COPY --from=builder /app/cluins .

# Run the binary
CMD ["./cluins"]