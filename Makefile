.PHONY: contrib_check dependencies install run size test

# detect GOPATH if not set
ifndef $(GOPATH)
		$(info GOPATH is not set, autodetecting..)
		TESTPATH := $(dir $(abspath ../../..))
		DIRS := bin pkg src
		# create a ; separated line of tests and pass it to shell
		MISSING_DIRS := $(shell $(foreach entry,$(DIRS),test -d "$(TESTPATH)$(entry)" || echo "$(entry)";))
		ifeq ($(MISSING_DIRS),)
				$(info Found GOPATH: $(TESTPATH))
				export GOPATH := $(TESTPATH)
		else
				$(info ..missing dirs "$(MISSING_DIRS)" in "$(TESTDIR)")
				$(info GOPATH autodetection failed)
		endif
endif

# Set go modules to on and use GoCenter for immutable modules
GO111MODULE=on
GOPROXY=https://gocenter.io

build:
	go build -o bin/wtf

contrib_check:
	npx all-contributors-cli check

install:
	@go clean
	@go install -ldflags="-s -w -X main.version=$(shell git describe --always --abbrev=6) -X main.date=$(shell date +%FT%T%z)"
	@which wtf || echo "Could not find wtf in PATH" && exit 0

lint:
	structcheck ./...
	varcheck ./...

run: build
	bin/wtf

size:
	loc --exclude _sample_configs/ _site/ docs/ Makefile *.md

test: build
	go test ./...
