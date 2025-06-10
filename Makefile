GO_VERSION=1.24
GOLINT_VERSION=v2.1.6
PROJECT ?= eip-server



ROOT_DIR=$(shell git rev-parse --show-toplevel)
TOOLS_DIR=$(ROOT_DIR)/.tools

ALL_GO_FILES=$(shell find $(ROOT_DIR) -type f -name "*.go")

LINT := $(TOOLS_DIR)/golangci-lint

GOCMD ?= go
GO_ENV=$(shell CGO_ENABLED=0)


$(TOOLS_DIR):
	mkdir -p $@

.PHONY: check-fmt
check-fmt: fmt
	@git diff -s --exit-code *.go || (echo "Build failed: a go file is not formated correctly. Run 'make fmt' and update your PR." && exit 1)

.PHONY: fmt
fmt:
	$(GOCMD) fmt ./...

.PHONY: govet
govet:
	$(GOCMD) vet ./...

.PHONY: test
test: govet 
	$(GO_ENV) $(GOCMD) test -v ./... -failfast

.PHONY: gomoddownload
gomoddownload:
	$(GOCMD) mod download -x

.PHONY: tools
tools: $(TOOLS_DIR)
	GOBIN=$(TOOLS_DIR) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLINT_VERSION)

.PHONY: lint
lint:
	$(LINT) run --verbose --allow-parallel-runners --timeout=10m 

.PHONY: tidy
tidy:
	$(GOCMD) mod tidy -compat=$(GO_VERSION)

.PHONY: vendor
vendor: 
	$(GOCMD) mod vendor

.PHONY: build
build: tools
	$(GOCMD) build -o bin/eip-server cmd/server/main.go

.PHONY: exec
exec: gofmt build 
	./bin/eip-server

.PHONY: run
run:
	$(GOCMD) run cmd/server/main.go
