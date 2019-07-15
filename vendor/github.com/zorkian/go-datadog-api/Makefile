TEST?=$$(go list ./... | grep -v '/go-datadog-api/vendor/')
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: test fmt

generate:
	go generate

# test runs the unit tests and vets the code
test:
	go test . $(TESTARGS) -v -timeout=30s -parallel=4
	@$(MAKE) vet

# testacc runs acceptance tests
testacc:
	go test integration/* -v $(TESTARGS) -timeout 90m

# testrace runs the race checker
testrace:
	go test -race $(TEST) $(TESTARGS)

fmt:
	gofmt -w $(GOFMT_FILES)

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@echo "go vet"
	@go vet; if [ $$? -ne 0 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

.PHONY: default test testacc updatedeps vet
