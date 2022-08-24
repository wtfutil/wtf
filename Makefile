.PHONY: build clean contrib_check coverage docker-build docker-install help install isntall lint run size test uninstall

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

GOBIN := $(GOPATH)/bin

APP=wtfutil

define HEADER
____    __    ____ .___________. _______
\   \  /  \  /   / |           ||   ____|
 \   \/    \/   /  `---|  |----`|  |__
  \            /       |  |     |   __|
   \    /\    /        |  |     |  |
    \__/  \__/         |__|     |__|

endef
export HEADER

# -------------------- Actions -------------------- #

## build: builds a local version
build:
	@echo "$$HEADER"
	@echo "Building..."
	go build -o bin/${APP}
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

## docker-build: builds in docker
docker-build:
	@echo "Building ${APP} in Docker..."
	docker build -t wtfutil:build --build-arg=version=master -f Dockerfile.build .
	@echo "Done with docker build"

## docker-install: installs a local version of the app from docker build
docker-install:
	@echo "Installing..."
	docker create --name wtf_build wtfutil:build
	docker cp wtf_build:/usr/local/bin/wtfutil ~/.local/bin/
	$(eval INSTALLPATH = $(shell which ${APP}))
	@echo "${APP} installed into ${INSTALLPATH}"
	docker rm wtf_build

## gosec: runs the gosec static security scanner against the source code
gosec: $(GOBIN)/gosec
	gosec -tests ./...

$(GOBIN)/gosec:
	cd && go install github.com/securego/gosec/v2/cmd/gosec@latest

## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## isntall: an alias for 'install'
isntall:
	@$(MAKE) -f $(THIS_FILE) install

## install: installs a local version of the app
install:
	$(eval GOVERS = $(shell go version))
	@echo "$$HEADER"
	@echo "Installing ${APP} with ${GOVERS}..."
	@go clean
	@go install -ldflags="-s -w"
	@mv $(GOBIN)/wtf $(GOBIN)/${APP}
	$(eval INSTALLPATH = $(shell which ${APP}))
	@echo "${APP} installed into ${INSTALLPATH}"

## lint: runs a number of code quality checks against the source code
lint: $(GOBIN)/golangci-lint
	golangci-lint cache clean
	golangci-lint run

$(GOBIN)/golangci-lint:
	cd && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# lint:
# 	@echo "\033[35mhttps://github.com/kisielk/errcheck\033[0m"
# 	errcheck ./app
# 	errcheck ./cfg
# 	errcheck ./flags
# 	errcheck ./help
# 	errcheck ./logger
# 	errcheck ./modules/...
# 	errcheck ./utils
# 	errcheck ./view
# 	errcheck ./wtf
# 	errcheck ./main.go

# 	@echo "\033[35mhttps://golang.org/cmd/vet/k\033[0m"
# 	go vet ./app
# 	go vet ./cfg
# 	go vet ./flags
# 	go vet ./help
# 	go vet ./logger
# 	go vet ./modules/...
# 	go vet ./utils
# 	go vet ./view
# 	go vet ./wtf
# 	go vet ./main.go

# 	@echo "\033[35m# https://staticcheck.io/docs/k\033[0m"
# 	staticcheck ./app
# 	staticcheck ./cfg
# 	staticcheck ./flags
# 	staticcheck ./help
# 	staticcheck ./logger
# 	staticcheck ./modules/...
# 	staticcheck ./utils
# 	staticcheck ./view
# 	staticcheck ./wtf
# 	staticcheck ./main.go

# 	@echo "\033[35m# https://github.com/mdempsky/unconvert\033[0m"
# 	unconvert ./...

## loc: displays the lines of code (LoC) count
loc:
	@loc --exclude _sample_configs/ _site/ docs/ Makefile *.md

## run: executes the locally-installed version
run: build
	@echo "$$HEADER"
	bin/${APP}

## test: runs the test suite
test: build
	@echo "$$HEADER"
	go test ./...

## uninstall: uninstals a locally-installed version
uninstall:
	@rm $(GOBIN)/${APP}
