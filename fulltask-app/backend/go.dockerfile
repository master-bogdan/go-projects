# Build stage
FROM golang:1.24-alpine AS builder

# Disable CGO for static binary, set target OS and architecture
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary with optimizations to reduce size
RUN go build -ldflags="-s -w" -o api

# Final stage - minimal runtime image
FROM alpine:3.21

# Install ca-certificates for HTTPS support if needed
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the statically compiled binary from builder stage
COPY --from=builder /app/api .

# Expose port if your app listens on one (adjust as needed)
EXPOSE 8000

# Run the binary
CMD ["./api"]
