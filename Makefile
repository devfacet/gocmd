# Init vars
MAKEFILE := $(lastword $(MAKEFILE_LIST))
BASENAME := $(shell basename "$(PWD)")
SHELL := /bin/bash

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Commands:"
	@echo
	@sed -n 's/^##//p' $< | sed -e 's/^/ /' | sort
	@echo

## init-tools       Initialize tools
init-tools:
	$(eval STATICCHECK_PATH=$(shell which staticcheck))
	@if [ -z "$(STATICCHECK_PATH)" ]; then \
		go install honnef.co/go/tools/cmd/staticcheck@latest; \
	fi

## check-tools      Check tools
check-tools:
	$(eval STATICCHECK_PATH=$(shell which staticcheck))
	@if [ -z "$(STATICCHECK_PATH)" ]; then \
		echo "staticcheck binary not found"; \
		exit 1; \
	fi

## test             Run tests, formatting, etc.
test:
	@$(MAKE) -f $(MAKEFILE) check-tools

	$(eval FMT=$(shell find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | sort | uniq | xargs gofmt -l | wc -l | tr -d ' '))
	@if [ "$(FMT)" != "0" ]; then \
		echo "some files are not formatted, run 'make fmt'"; \
		exit 1; \
	fi

	$(eval STATICCHECK=$(shell find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | sort | uniq | xargs staticcheck | wc -l | tr -d ' '))
	@if [ "$(STATICCHECK)" != "0" ]; then \
		echo "some files have staticcheck errors, run 'make staticcheck'"; \
		exit 1; \
	fi

	$(eval GOVET=$(shell find . -type f -name '*.go' | grep -v -E '^./vendor' | xargs -L1 dirname | sort | uniq | xargs go vet 2>&1 | wc -l | tr -d ' '))
	@if [ "$(GOVET)" != "0" ]; then \
		echo "some files have vetting errors, run 'make vet'"; \
		exit 1; \
	fi

	@$(MAKE) -f $(MAKEFILE) test-go

## test-go          Run go test
test-go:
	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party|^./_examples' | xargs -L1 dirname | sort | uniq | xargs go test -v -race

## test-benchmarks  Run go benchmarks
test-benchmarks:
	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party|^./_examples' | xargs -L1 dirname | sort | uniq | xargs go test -benchmem -bench

## test-ui          Launch test UI
test-ui:
	$(eval GOCONVEY_PATH=$(shell which goconvey))
	@if [ -z "$(GOCONVEY_PATH)" ]; then \
		GO111MODULE=off go get github.com/smartystreets/goconvey; \
	fi
	goconvey -port 8088 -excludedDirs vendor,node_modules,assets,tmp,_examples

## test-clean       Clean test cache
test-clean:
	@go clean -testcache

## fmt              Run formatting
fmt:
	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | sort | uniq | xargs gofmt -l

## staticcheck      Run staticcheck
staticcheck:
	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | sort | uniq | xargs staticcheck

## vet              Run vetting
vet:
	@find . -type f -name '*.go' | grep -v -E '^./vendor' | xargs -L1 dirname | sort | uniq | xargs go vet 2>&1
