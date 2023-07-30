# Project
SHELL := /usr/bin/env bash -o pipefail
NAME := kreutzer
VERSION := 0.1.0

# Main
BINARY_NAME := kreutzer
MAIN := ./cmd/kreutzer/main.go

# Toolchain
GO := GO111MODULE=on GOPROXY=https://goproxy.cn,direct go
GO_VERSION := $(shell $(GO) version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/')

# Container
DOCKER := docker
DOCKERFILE := ci/docker/Dockerfile
REGISTRY := harbor.leryn.top/infra
IMAGE_NAME := kreutzer
FULL_IMAGE_NAME = $(REGISTRY)/$(IMAGE_NAME):$(VERSION)

##@ General

.PHONY: help
help: ## Print help info.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Developement

.PHONY: install
install: ## Install dependencies.
	$(GO) get -d -v ./...

.PHONY: check
check: ## Check
	$(GO) vet ./...

.PHONY: fmt
fmt: ## Format against code.
	$(GO) fmt ./...

.PHONY: clean
clean: ## Clean target artifact.
	$(GO) clean -r -x

.PHONY: test
test: ## Run test.
	$(GO) test ./...

##@ Build

.PHONY: build
build: ## Build target artifact.
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -ldflags '-extldflags "-static"' -o target/$(BINARY_NAME) $(MAIN)

