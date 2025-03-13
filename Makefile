VERSION=$(shell git describe --tags --always)
APP_NAME ?= $(app)

.PHONY: init
init:
	@echo "Initializing moon environment"

.PHONY: all
all:
	@echo "Initialization of moon project "

.PHONY: build
build:
	@if [ -z "$(APP_NAME)" ]; then echo "app name is required"; exit 1; fi
	@echo "Building moon app=$(APPP_NAME)"
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/$(APP_NAME)

.PHONY: docker-build
docker-build:
	@if [ -z "$(APP_NAME)" ]; then echo "app name is required"; exit 1; fi
	@echo "Building moon app=$(APPP_NAME)"
	docker build -t ghcr.io/moon-monitor/palace:$(VERSION) \
      --build-arg APP_NAME=$(APP_NAME) \
      -f deploy/server/Dockerfile .

.PHONY: builder-image
builder-image:
	@echo "Building moon builder image"
	docker build -f deploy/base/DockerfileBuilder -t ghcr.io/moon-monitor/moon:builder .

.PHONY: deploy-image
deploy-image:
	@echo "Building moon deploy image"
	docker build -f deploy/base/DockerfileDeploy -t ghcr.io/moon-monitor/moon:deploy .