# Use the official Golang 1.24 Alpine image as the build stage
FROM golang:1.25-alpine AS builder

# Set environment variables for static build
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory inside container
WORKDIR /build

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary
RUN go build -o /app ./cmd/server

# Use a minimal Alpine image for the final stage
FROM alpine:3.21

RUN apk add --no-cache ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app /app

# Expose the port your app listens on (adjust as needed)
EXPOSE 8000

# Run the application
CMD ["/app"]
