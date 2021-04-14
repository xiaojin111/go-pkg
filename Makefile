SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export GOPATH := $(shell go env GOPATH)
export GO111MODULE := on
export GOPROXY := https://goproxy.io,direct
export GOPRIVATE := gitee.com/jt-heath/*

# Install all the build and lint dependencies
setup:
	# TODO: 官方 golangci-lint 发行包不兼容 go 1.13，需要使用手动编译的版本
	#curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	[ -d "./bin" ] || mkdir -p ./bin
	cp $(GOPATH)/bin/golangci-lint ./bin/
	# curl -L https://git.io/misspell | sh
	go mod download
.PHONY: setup

# Update go packages
go-update:
	@echo "Updating go packages..."
	@go get -u -t ./...
	@echo "go mod tidy..."
	@$(MAKE) go-mod-tidy
.PHONY: go-update

# Clean go.mod
go-mod-tidy:
	@go mod tidy -v
	# @git --no-pager diff HEAD
	# @git --no-pager diff-index --quiet HEAD
.PHONY: go-mod-tidy

# Reset go.mod
go-mod-reset:
	@rm -f go.sum
	@sed -i '' -e '/^require/,/^)/d' go.mod
	@go mod tidy -v
	# @git --no-pager diff HEAD
	# @git --no-pager diff-index --quiet HEAD
.PHONY: go-mod-tidy

generate:
	@go generate ./...
.PHONY: generate

# Format go files
format:
	@goimports -w ./
.PHONY: format

# Run all the linters
lint:
	@./bin/golangci-lint run
.PHONY: lint

# Go build all
build:
	@go build ./... > /dev/null
.PHONY: build

# Go test all
test:
	@go test ./...
.PHONY: test

# Run all code checks
ci: generate format lint build test
.PHONY: ci

.DEFAULT_GOAL := ci
