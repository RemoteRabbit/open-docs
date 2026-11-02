# Self-documenting Makefile: run `make` to list targets.
# A `## ...` comment directly above a target documents it (several `##` lines
# are joined); targets without one are hidden from the help.
# Convention from https://gist.github.com/klmr/575726c7e05d8780505a

SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables --no-builtin-rules

BIN     := bin/open-docs
PKG     := ./...
EXAMPLE := ./examples/vpc
GO_SRC   = $(shell find . -name '*.go' -not -path './.devenv/*' -not -path './.git/*')
# Go reads GOFLAGS from the environment, e.g. `make build GOFLAGS=-v`.

.PHONY: help all build test vet fmt fmt-check smoke verify tidy clean

## Show this help
help:
	@awk -v c="$$([ -t 1 ] && [ -z "$${NO_COLOR:-}" ] && printf '\033[36m')" \
		-v r="$$([ -t 1 ] && [ -z "$${NO_COLOR:-}" ] && printf '\033[0m')" \
		'/^## / { doc = doc (doc ? " " : "") substr($$0, 4); next } \
		/^[A-Za-z0-9_.-]+:/ && doc { sub(/:.*/, ""); printf "  %s%-10s%s %s\n", c, $$0, r, doc } \
		{ doc = "" }' $(MAKEFILE_LIST)

## Run all checks, then build the binary
all: verify build

## Build the binary into bin/
build:
	@mkdir -p $(dir $(BIN))
	go build -o $(BIN) .

## Run unit tests with the race detector
test:
	go test -race -count=1 $(PKG)

## Run go vet
vet:
	go vet $(PKG)

## Format Go sources with gofmt
fmt:
	gofmt -w $(GO_SRC)

## Fail if any Go source needs gofmt
fmt-check:
	@files="$$(gofmt -l $(GO_SRC))"; \
	if [ -n "$$files" ]; then echo "Files need gofmt:" >&2; echo "$$files" >&2; exit 1; fi

## Render the example module and check the output is non-empty
smoke: build
	@out="$$(mktemp)"; trap 'rm -f "$$out"' EXIT; \
	./$(BIN) -o "$$out" $(EXAMPLE); \
	test -s "$$out"; echo "smoke: ok"

## Run every check: formatting, vet, tests, and smoke test
verify: fmt-check vet test smoke

## Tidy go.mod and go.sum
tidy:
	go mod tidy

## Remove build output
clean:
	rm -rf $(dir $(BIN))
