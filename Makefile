GOHOSTOS:=$(shell go env GOHOSTOS)
VERSION=$(shell git describe --tags --always)
APP_NAME ?= $(app)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find cmd -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find proto/api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find cmd -name *.proto)
	API_PROTO_FILES=$(shell find proto/api -name *.proto)
endif

.PHONY: init
init:
	@echo "Initializing moon environment"

.PHONY: all
all:
	@echo "Initialization of moon project "

.PHONY: api
# generate api proto
api:
	mkdir -p ./pkg/api
	protoc --proto_path=./proto/api \
	       --proto_path=./proto/third_party \
 	       --go_out=paths=source_relative:./pkg/api \
 	       --go-http_out=paths=source_relative:./pkg/api \
 	       --go-grpc_out=paths=source_relative:./pkg/api \
	       --openapi_out=fq_schema_naming=true,default_response=false:./swagger \
	       $(API_PROTO_FILES)

.PHONY: errors
# generate errors
errors:
	mkdir -p ./pkg/merr
	protoc --proto_path=./proto/merr \
           --proto_path=./proto/third_party \
           --go_out=paths=source_relative:./pkg/merr \
           --go-errors_out=paths=source_relative:./pkg/merr \
           ./proto/merr/*.proto

.PHONY: conf
# generate config
conf:
	mkdir -p ./pkg/config
	protoc --proto_path=./proto/config \
           --proto_path=./proto/third_party \
           --go_out=paths=source_relative:./pkg/config \
           ./proto/config/*.proto

.PHONY: palace-conf
# generate palace-config
palace-conf:
	protoc --proto_path=./proto/config \
           --proto_path=./proto/third_party \
           --proto_path=./cmd/palace/internal/conf \
           --go_out=paths=source_relative:./cmd/palace/internal/conf \
           ./cmd/palace/internal/conf/*.proto

.PPHONY: wire-palace
wire-palace:
	cd ./cmd/palace && wire

.PHONY: build
build:
	@if [ -z "$(APP_NAME)" ]; then echo "app name is required"; exit 1; fi
	@echo "Building moon app=$(APPP_NAME)"
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/$(APP_NAME)

.PHONY: docker-build
docker-build:
	@if [ -z "$(APP_NAME)" ]; then echo "app name is required"; exit 1; fi
	@echo "Building moon app=$(APP_NAME)"
	docker build -t ghcr.io/moon-monitor/$(APP_NAME):$(VERSION) \
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