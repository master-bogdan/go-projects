# Binary name
BINARY_NAME=ephermal-notes

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

# Make commands
.PHONY: all build clean run test

all: clean build

build:
	@echo "Building..."
	@go build -o $(GOBIN)/$(BINARY_NAME) ./...

clean:
	@echo "Cleaning..."
	@rm -rf $(GOBIN)

run:
	@go run ./...

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Running linter..."
	@golangci-lint run

.DEFAULT_GOAL := build