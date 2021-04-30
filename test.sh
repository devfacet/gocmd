#!/bin/bash

set -e

# Init vars
#GO111MODULE=off go get golang.org/x/lint/golint
#GO111MODULE=off go get github.com/smartystreets/goconvey
CGO_ENABLED=0
if [ "${CLEAN_CACHE}" == "1" ] ; then
  BUILD_ARGS=-a
fi

# Format, lint, check
FMT=`(gofmt -l . | grep -v -E '^vendor/') || true`; if [ "$FMT" ]; then echo -e "fmt:\n$FMT"; fi
LINT=`find . -type f -name '*.go' | grep -v -E '^./vendor|^./third_party' | xargs -L1 dirname | uniq | xargs golint`; if [ "$LINT" ]; then echo -e "lint:\n$LINT"; fi
VET=`find . -type f -name '*.go' | grep -v -E '^./vendor' | xargs -L1 dirname | uniq | xargs go vet`; if [ "$VET" ]; then echo -e "vet:\n$VET"; fi
if [[ "$FMT" ]] || [[ "$LINT" ]] || [[ "$VET" ]]; then
  echo format: $FMT
  echo lint: $LINT
  echo vet: $VET
  echo "failed"
  exit 1
fi

# Test
go test -v $BUILD_ARGS ./...
result="$?"
if [[ "$result" -ne 0 ]]; then
  echo "failed"
  exit 1
fi
echo "succeed"
