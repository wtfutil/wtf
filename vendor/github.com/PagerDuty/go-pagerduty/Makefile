# SOURCEDIR=.
# SOURCES = $(shell find $(SOURCEDIR) -name '*.go')
# VERSION=$(git describe --always --tags)
# BINARY=bin/pd

# bin: $(BINARY)

# $(BINARY): $(SOURCES)
# 	go build -o $(BINARY) command/*

.PHONY: build
build:
	go get ./...
	# go test -v -race -cover ./...
	# go tool vet $(SOURCES)

.PHONY: test
test:
	go test ./...

deploy:
	- curl -sL https://git.io/goreleaser | bash

