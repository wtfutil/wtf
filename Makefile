.PHONY: contrib_check dependencies install run size

build:
	go build -race -o bin/wtf

contrib_check:
	npx all-contributors-cli check

install:
	go clean
	go install -ldflags="-s -w -X main.version=$(shell git describe --always --abbrev=6) -X main.date=$(shell date +%FT%T%z)"
	which wtf

lint:
	structcheck ./...
	varcheck ./...

run: build
	bin/wtf

size:
	loc --exclude vendor/ _sample_configs/ _site/ docs/ Makefile *.md *.toml
