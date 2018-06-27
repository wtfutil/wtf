.PHONY: fmt vet check-vendor lint check clean test build
PACKAGES = $(shell go list ./...)
PACKAGE_DIRS = $(shell go list -f '{{ .Dir }}' ./...)

check: test vet lint

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

vet:
	go vet $(PACKAGES) || (go clean $(PACKAGES); go vet $(PACKAGES))

lint:
	gometalinter --config gometalinter.json ./...

fmt:
	go fmt $(PACKAGES)
	goimports -w $(PACKAGE_DIRS)

deps:
	go get -t -v ./...
	go get github.com/axw/gocov/gocov
	go get golang.org/x/tools/cmd/cover
	[ -f $(GOPATH)/bin/gometalinter ] || go get -u github.com/alecthomas/gometalinter
	[ -f $(GOPATH)/bin/goimports ] || go get golang.org/x/tools/cmd/goimports
	gometalinter --install
