GOOS ?= linux
GOARCH ?= amd64
CGO_ENABLED ?= 1
LDFLAGS += -s -w
BIN_NAME ?= bump-version

ifeq ($(NOGIT),1)
  GIT_SUMMARY ?= Unknown
  GIT_BRANCH ?= Unknown
  GIT_MERGE ?= Unknown
else
  GIT_SUMMARY := $(shell git describe --tags --dirty --always)
  GIT_BRANCH := $(shell git symbolic-ref -q --short HEAD)
  GIT_MERGE := $(shell git rev-list --count --merges main)
endif

LDFLAGS += -X main.GitBranch=${GIT_BRANCH} -X main.GitSummary=${GIT_SUMMARY} -X main.GitMerge=${GIT_MERGE}

default: help

## build: builds the binaries
build:
	@echo GOOS       : $(GOOS)
	@echo GOARCH     : $(GOARCH)
	@echo LDFLAGS    : $(LDFLAGS)
	@echo CGO_ENABLED: $(CGO_ENABLED)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="${LDFLAGS}" -o bin/${BIN_NAME} main.go

## clean: cleans bin folder
clean:
	@rm -rf bin/*

## deps: downloads mod dependencies
deps:
	@go mod download

## test: test all files recursively
test:
	@go test ./...

## vet: check for linting issues
vet:
	@go vet ./...

## install: build and install binaries inside /usr/local/bin
install:
	@echo GOOS       : $(GOOS)
	@echo GOARCH     : $(GOARCH)
	@echo LDFLAGS    : $(LDFLAGS)
	@echo CGO_ENABLED: $(CGO_ENABLED)
	@echo GIT_SUMMARY: $(GIT_SUMMARY)
	@echo GIT_BRANCH : $(GIT_BRANCH)
	@echo GIT_MERGE  : $(GIT_MERGE)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="${LDFLAGS}" -o /usr/local/bin/${BIN_NAME} main.go

## fmt: format project
fmt:
	@gofumpt -w -l .

## help: show this help
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: build help
