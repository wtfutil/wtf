.PHONY: build clean contrib_check coverage help install isntall lint run size test uninstall

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
export GOPROXY = https://proxy.golang.org,direct

# Determines the path to this Makefile
THIS_FILE := $(lastword $(MAKEFILE_LIST))

APP=wtfutil

# -------------------- Actions -------------------- # 

## build: builds a local version
build:
	go build -o bin/${APP} -mod=vendor
	@echo "Done building"

## clean: removes old build cruft
clean:
	rm -rf ./dist
	rm -rf ./bin/${APP}
	@echo "Done cleaning"

## contrib-check: checks for any contributors who have not been given due credit
contrib-check:
	npx all-contributors-cli check

## coverage: figures out and displays test code coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## isntall: an alias for 'install'
isntall:
	@$(MAKE) -f $(THIS_FILE) install

## install: installs a local version of the app
install:
	@echo "Installing ${APP}..."
	@go clean
	@go install -ldflags="-s -w -X main.version=$(shell git describe --always --abbrev=6) -X main.date=$(shell date +%FT%T%z)"
	@mv ~/go/bin/wtf ~/go/bin/${APP}
	$(eval INSTALLPATH = $(shell which ${APP}))
	@echo "${APP} installed into ${INSTALLPATH}"

## lint: runs a number of code quality checks against the source code
lint:
	go vet ./...
	structcheck ./...
	varcheck ./...

## run: executes the locally-installed version
run: build
	bin/${APP}

## size: displays the lines of code (LoC) count
size:
	@loc --exclude _sample_configs/ _site/ docs/ Makefile *.md

## test: runs the test suite
test: build
	go test ./...

## uninstall: uninstals a locally-installed version
uninstall:
	@rm ~/go/bin/${APP}
