install:
	go install -ldflags="-X main.version=$(shell git describe --always --long --dirty)"
