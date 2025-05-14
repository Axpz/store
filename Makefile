# --- Go parameters ---
GOCMD = go
GOBUILD = $(GOCMD) build
GORUN = $(GOCMD) run
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

BINARY_NAME = store
BINARY_OUT = $(BINARY_NAME)
BINARY_LINUX = $(BINARY_NAME)_linux
LDFLAGS = -ldflags="-s -w"

# --- Build environment ---
BUILD_ENV ?= dev

# --- Git commit ---
GIT_COMMIT := $(shell git rev-parse --short HEAD)

# --- Frontend ---
FRONTEND_DIR = frontend
NPM_CMD = pnpm

# --- Docker ---
IMAGE_NAME = store
DOCKER_TAG ?= $(if $(filter $(BUILD_ENV),dev),latest,$(GIT_COMMIT))
DOCKERFILE = Dockerfile

.PHONY: all build dev prod run tidy test clean get-deps \
        frontend-install frontend-build frontend-clean \
        build-all run-all backend-build \
        docker-build docker-buildx docker-push \
        test-coverage test-race test-bench deploy

# --- Main targets ---

all: test build-all

build:
	@echo "🚀 Building in production mode..."
	$(MAKE) build-all BUILD_ENV=dev
	$(MAKE) docker-build BUILD_ENV=dev

## Production build
buildprod:
	@echo "🚀 Building in production mode..."
	$(MAKE) build-all BUILD_ENV=production NODE_ENV=production
	$(MAKE) docker-build BUILD_ENV=production

# --- Build backend + frontend ---
build-all: frontend-install frontend-build backend-build

backend-build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_OUT) ./cmd/

run:
	$(GORUN) ./cmd/main.go

dev:
	cd $(FRONTEND_DIR) && pnpm run dev

clean: frontend-clean
	$(GOCLEAN)
	rm -f $(BINARY_OUT) $(BINARY_LINUX)

tidy:
	$(GOCMD) mod tidy

get-deps:
	$(GOGET) -v -t -d ./...

# --- Frontend tasks ---
frontend-install:
	cd $(FRONTEND_DIR) && $(NPM_CMD) install

frontend-build:
	cd $(FRONTEND_DIR) && $(NPM_CMD) run build

frontend-clean:
	cd $(FRONTEND_DIR) && rm -rf .next out

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BINARY_LINUX) ./cmd/

# --- Docker ---
docker-build:
	@echo "🐳 Building Docker image: $(IMAGE_NAME):$(DOCKER_TAG)"
	docker build -t $(IMAGE_NAME):$(DOCKER_TAG) -f $(DOCKERFILE) .
	docker build -t $(IMAGE_NAME)-$(FRONTEND_DIR):$(DOCKER_TAG) -f ./$(FRONTEND_DIR)/$(DOCKERFILE) .

docker-buildx:
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--push \
		-t $(IMAGE_NAME):$(DOCKER_TAG) \
		-f $(DOCKERFILE) .

docker-push:
	docker push $(IMAGE_NAME):$(DOCKER_TAG)

# --- Tests ---
test:
	$(GOTEST) -v ./...

test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

test-race:
	$(GOTEST) -v -race ./...

test-bench:
	$(GOTEST) -v -bench=. ./...
deploy:
	cd ./kind && bash deploy.sh
