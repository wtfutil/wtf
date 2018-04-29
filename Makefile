BRANCH := `git rev-parse --abbrev-ref HEAD`

install:
	which wtf | xargs rm
	go install -ldflags="-X main.version=$(shell git describe --always --abbrev=6)_$(BRANCH) -X main.builtat=$(shell date +%FT%T%z)"
