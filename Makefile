VERSION=$(shell git describe --tags --always)
APP_NAME ?= $(app)

.PHONY: init
init:
	echo "Initializing moon environment"

.PHONY: all
all:
	echo "Initialization of moon project "

.PHONY: build
build:
	echo "Building moon app=$(APPP_NAME)"
	@if [ -z "$(APP_NAME)" ]; then echo "app name is required"; exit 1; fi
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/$(APP_NAME)