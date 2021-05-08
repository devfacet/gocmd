# Init vars.
MAKEFILE := $(lastword $(MAKEFILE_LIST))
BASENAME := $(shell basename "$(PWD)")

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Commands:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## test: Run gofmt, golint, staticcheck, go vet and go test.
test:
	$(eval FMT=$(shell gofmt -l . | grep -v -E '^vendor/'))
	@if [ ! -z "$(FMT)" ]; then \
		echo "some files are not formatted, run 'make fmt'"; \
	fi

	$(eval GOLINT_PATH=$(shell which golint))
	@if [ -z $(GOLINT_PATH) ]; then \
		GO111MODULE=off go get golang.org/x/lint/golint; \
	fi
	$(eval LINT=$(shell find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | uniq | sort | xargs golint | wc -l | tr -d ' '))
	@if [ "$(LINT)" != "0" ]; then \
		echo "some files have linting errors, run 'make lint'"; \
	fi

	$(eval STATICCHECK_PATH=$(shell which staticcheck))
	@if [ -z $(STATICCHECK_PATH) ]; then \
		GO111MODULE=off go get honnef.co/go/tools/cmd/staticcheck; \
	fi
	$(eval STATICCHECK=$(shell find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | uniq | sort | xargs staticcheck | wc -l | tr -d ' '))
	@if [ "$(STATICCHECK)" != "0" ]; then \
		echo "some files have staticcheck errors, run 'make staticcheck'"; \
	fi

	$(eval GOVET=$(shell find . -type f -name '*.go' | grep -v -E '^./vendor' | xargs -L1 dirname | uniq | xargs go vet 2>&1 | wc -l | tr -d ' '))
	@if [ "$(GOVET)" != "0" ]; then \
		echo "some files have vetting errors, run 'make vet'"; \
	fi

	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | uniq | sort | xargs go test -v

## test-all: Run tests including vendor and third party modules.
test-all:
	@$(MAKE) -f $(MAKEFILE) test
	@find . -type f -name '*.go' | grep -E '^./vendor|^./third_party' | xargs -L1 dirname | uniq | sort | xargs go test -v

## test-ui: Launch test UI
test-ui:
	$(eval GOCONVEY_PATH=$(shell which goconvey))
	@if [ -z $(GOCONVEY_PATH) ]; then \
		GO111MODULE=off go get github.com/smartystreets/goconvey; \
	fi
	goconvey

## test-clean: Clean test cache
test-clean:
	@go clean -testcache

## fmt: Run formating
fmt:
	@gofmt -l . | grep -v -E '^vendor/' | xargs gofmt -w

## lint: Run linting
lint:
	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | uniq | xargs golint

## staticcheck: Run staticcheck
staticcheck:
	@find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | uniq | xargs staticcheck

## vet: Run vetting
vet:
	@find . -type f -name '*.go' | grep -v -E '^./vendor' | xargs -L1 dirname | uniq | xargs go vet 2>&1

## build: Build binaries
build:
	$(eval BIN_VERSION=$(shell git describe --tags --exact-match 2>>/dev/null || echo "v0.0.0-$(shell date '+%Y%m%dT%H%M')"))
	$(eval GIT_COMMIT=$(shell git diff-index --quiet HEAD -- || echo "$$(git rev-parse --short HEAD)-dirty"; if [ -z "$$GIT_COMMIT" ]; then GIT_COMMIT=$$(git rev-parse --short HEAD); fi))
	$(eval BUILD_INFO=$(shell go env | grep -E '^GOHOSTOS|^GOHOSTARCH|^GOOS|^GOARCH|^GOARM|^GO386' | tr '\n' ' '))
	@if [ ! -z $(GOPATH) ]; then \
		$(eval TARGET_DIR=$(GOPATH)/bin/) \
		TARGET_DIR=$(TARGET_DIR); \
	fi
	@if [ -z "$(BUILD_ARGS)" ]; then \
		$(eval BUILD_ARGS=-trimpath) \
		BUILD_ARGS=$(BUILD_ARGS); \
	fi
	@echo BUILD_INFO=$(BUILD_INFO)
	go build -ldflags "-X main.version=$(BIN_VERSION) -X main.gitCommit=$(GIT_COMMIT) -s -w" $(BUILD_ARGS) -o $(TARGET_DIR)gocmdbasic examples/basic/*.go

## build-force: Build binaries (force rebuilding of packages).
build-force:
	@$(MAKE) -f $(MAKEFILE) build BUILD_ARGS="-a -trimpath"

## release: Release a version
release:
	@if [ -z "$(GIT_TAG)" ]; then \
		echo "invalid GIT_TAG. Try something like 'make release GIT_TAG=v1.0.0'"; \
		exit 1; \
	fi
	git commit -m "$(GIT_TAG)"
	git tag -a $(GIT_TAG) -m "$(GIT_TAG)"
	git push --follow-tags
	git ls-remote --tags
