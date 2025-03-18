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
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.3
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/moon-monitor/stringer@latest

.PHONY: all
all:
	@echo "Initialization of moon project"
	@if [ -z "$(APP_NAME)" ]; then echo "app name is required"; echo "usage: make all app=<app_name>"; exit 1; fi
	make api
	make errors
	make conf
	make stringer-$(APP_NAME)
	make conf-$(APP_NAME)
	make wire-$(APP_NAME)
	make gen-$(APP_NAME)

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
	       --experimental_allow_proto3_optional \
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
           --experimental_allow_proto3_optional \
           ./proto/config/*.proto
.PHONY: gen
gen:
	make gen-palace

.PHONY: gen-palace
# generate gorm gen
gen-palace:
	go run cmd/palace/migrate/gen/gen.go

.PHONY: gen-system
# generate gorm gen system
gen-system:
	go run cmd/palace/migrate/gen/gen.go sys

.PHONY: gen-biz
gen-biz:
	go run cmd/palace/migrate/gen/gen.go biz

.PHONY: gen-event
gen-event:
	go run cmd/palace/migrate/gen/gen.go event


.PHONY: conf-palace
# generate palace-config
conf-palace:
	protoc --proto_path=./proto/config \
           --proto_path=./proto/third_party \
           --proto_path=./cmd/palace/internal/conf \
           --go_out=paths=source_relative:./cmd/palace/internal/conf \
           --experimental_allow_proto3_optional \
           ./cmd/palace/internal/conf/*.proto

.PPHONY: wire-palace
wire-palace:
	cd ./cmd/palace && wire

.PPHONY: stringer-palace
stringer-palace:
	cd ./cmd/palace/internal/biz/vobj && go generate

.PHONY: build
build:
	@if [ -z "$(APP_NAME)" ]; then echo "app name is required"; echo "usage: make build app=<app_name>"; exit 1; fi
	@echo "Building moon app=$(APPP_NAME)"
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/$(APP_NAME)

.PHONY: docker-build
docker-build:
	@if [ -z "$(APP_NAME)" ]; then echo "app name is required"; echo "usage: make docker-build app=<app_name>"; exit 1; fi
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