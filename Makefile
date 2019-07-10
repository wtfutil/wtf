.PHONY: build contrib_check install binary_msg lint run size test uninstall

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
export GO111MODULE = on
export GOPROXY = https://gocenter.io

# Determines the path to this Makefile
THIS_FILE := $(lastword $(MAKEFILE_LIST))

build:
	go build -o bin/wtfutil
	@$(MAKE) -f $(THIS_FILE) binary_msg

contrib_check:
	npx all-contributors-cli check

install:
	@echo "Installing wtfutil..."
	@go clean
	@go install -ldflags="-s -w -X main.version=$(shell git describe --always --abbrev=6) -X main.date=$(shell date +%FT%T%z)"
	@mv ~/go/bin/wtf ~/go/bin/wtfutil
	@$(MAKE) -f $(THIS_FILE) binary_msg

binary_msg:
	@echo "\n\033[1mIMPORTANT NOTE\033[0m:"
	@echo "    The \033[0;33mwtf\033[0m binary has been renamed to \033[0;33mwtfutil\033[0m."
	@echo "    Executing \033[0;33mwtf\033[0m will no longer work. You \033[1mmust\033[0m use \033[0;33mwtfutil\033[0m.\n"
	@echo "Install path: "
	@which wtfutil || echo "Could not find wtfutil in PATH" && exit 0

lint:
	structcheck ./...
	varcheck ./...

run: build
	bin/wtfutil

size:
	loc --exclude _sample_configs/ _site/ docs/ Makefile *.md

test: build
	go test ./...

uninstall:
	@rm ~/go/bin/wtfutil