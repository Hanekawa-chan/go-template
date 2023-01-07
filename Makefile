PROJECT_NAME := project-name
PROJECT := github.com/someone/project-name

VERSION := $(shell git describe --tags)
COMMIT := $(shell git rev-parse --short HEAD)

LDFLAGS := "-s -w -X $(PROJECT)/internal/version.Version=$(VERSION) -X $(PROJECT)/internal/version.Commit=$(COMMIT)"
build:
	CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o ./bin/$(PROJECT_NAME) ./cmd/$(PROJECT_NAME)

test:
	@go test -v -cover -gcflags=-l --race ./...

GOLANGCI_LINT_VERSION := v1.24.0
lint:
	@golangci-lint run -v

dep:
	@go mod download

models_linux:
	cp -f go-models/*.go internal/services/models/

models_win:
	copy go-models\*.go internal\services\models\ /Y

modules:
	git submodule update --remote

update_modules_win: modules models_win

update_modules_linux: modules models_linux

update_deps:
	go get -u ./...
