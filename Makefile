# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=store
BINARY_UNIX=$(BINARY_NAME)_unix

# Build flags
LDFLAGS=-ldflags "-s -w"

.PHONY: all build test clean get-deps

all: test build

build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) ./cmd/

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

get-deps:
	$(GOGET) -v -t -d ./...

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BINARY_UNIX) ./cmd/
	chmod +x $(BINARY_UNIX)

# Run tests with coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Run tests with race detector
test-race:
	$(GOTEST) -v -race ./...

# Run tests with benchmarks
test-bench:
	$(GOTEST) -v -bench=. ./... 