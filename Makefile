# Init vars.
MAKEFILE := $(lastword $(MAKEFILE_LIST))
BASENAME := $(shell basename "$(PWD)")

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Commands:"
	@echo
	@sed -n 's/^##//p' $< | sed -e 's/^/ /' | sort
	@echo

## test            Run gofmt, golint, staticcheck, go vet and go test.
test:
	$(eval FMT=$(shell find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | sort | uniq | xargs gofmt -l | wc -l | tr -d ' '))
	@if [ "$(FMT)" != "0" ]; then \
		echo "some files are not formatted, run 'make fmt'"; \
		exit 1; \
	fi

	$(eval LINT=$(shell find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | sort | uniq | xargs golint | wc -l | tr -d ' '))
	@if [ "$(LINT)" != "0" ]; then \
		echo "some files have linting errors, run 'make lint'"; \
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

## test-go         Run go test
test-go:
	@find . -type f -name '*.go' | xargs -L1 dirname | sort | uniq | xargs go test -v -race

## test-benchmarks Run go benchmarks
test-benchmarks:
	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | sort | uniq | xargs go test -benchmem -bench

## test-ui         Launch test UI
test-ui:
	$(eval GOCONVEY_PATH=$(shell which goconvey))
	@if [ -z "$(GOCONVEY_PATH)" ]; then \
		GO111MODULE=off go get github.com/smartystreets/goconvey; \
	fi
	goconvey -port 8088 -excludedDirs vendor,node_modules,assets

## test-clean      Clean test cache
test-clean:
	@go clean -testcache

## test-tools      Install test tools
test-tools:
	@# golint is deprecated and frozen.
	$(eval GOLINT_PATH=$(shell which golint))
	@if [ -z "$(GOLINT_PATH)" ]; then \
		GO111MODULE=off go get golang.org/x/lint/golint; \
	fi

	$(eval STATICCHECK_PATH=$(shell which staticcheck))
	@if [ -z "$(STATICCHECK_PATH)" ]; then \
		go install honnef.co/go/tools/cmd/staticcheck@v0.3.1; \
	fi

## fmt             Run formating
fmt:
	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | sort | uniq | xargs gofmt -l

## lint            Run linting
lint:
	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | sort | uniq | xargs golint

## staticcheck     Run staticcheck
staticcheck:
	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | sort | uniq | xargs staticcheck

## vet             Run vetting
vet:
	@find . -type f -name '*.go' | grep -v -E '^./vendor' | xargs -L1 dirname | sort | uniq | xargs go vet 2>&1

## release         Release a version
release:
	@if [ "$(shell echo \$${GIT_TAG:0:1})" != "v" ]; then \
		echo "invalid GIT_TAG (${GIT_TAG}). Try something like 'make release GIT_TAG=v1.0.0'"; \
		exit 1; \
	fi
	git tag -a $(GIT_TAG) -m "$(GIT_TAG)"
	git push --follow-tags

## build           Build
build:
	$(eval BIN_VERSION=$(shell git describe --tags --exact-match 2>>/dev/null || echo "$(shell git describe --abbrev=0 2>>/dev/null || echo "v0.0.0")+$(shell date '+%Y%m%d%H%M')"))
	$(eval GIT_COMMIT=$(shell git describe --match=NEVERMATCH --always --abbrev=7 --dirty))
	$(eval BUILD_INFO=$(shell go env | grep -E '^GOVERSION|^GOHOSTOS|^GOHOSTARCH|^GOOS|^GOARCH|^GOARM|^GO386|^GOPATH' | tr '\n' ' '))
	$(eval BUILD_TARGET_DIR=$(shell echo "$$(go env GOPATH)/bin/"))

	@if [ -z "$(BUILD_ARGS)" ]; then \
		BUILD_ARGS=$(eval BUILD_ARGS=-trimpath); \
	fi
	@echo BUILD_INFO=$(BUILD_INFO)
	go build -ldflags "-X main.version=$(BIN_VERSION) -X main.gitCommit=$(GIT_COMMIT) -s -w" $(BUILD_ARGS) -o $(BUILD_TARGET_DIR)gocmdbasic examples/basic/*.go

## build-force     Build (force rebuilding of packages).
build-force:
	@$(MAKE) -f $(MAKEFILE) build BUILD_ARGS="-a -trimpath"
