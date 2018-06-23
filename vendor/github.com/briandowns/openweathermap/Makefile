GOCMD = go
GOBUILD = $(GOCMD) build
GOGET = $(GOCMD) get -v
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install
GOTEST = $(GOCMD) test

.PHONY: all

all: test

test:
	$(GOTEST) -v -covermode=count -coverprofile=coverage.out ./...

build: test
	$(GOBUILD)

install: test
	$(GOINSTALL)
