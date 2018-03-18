#!/bin/bash

set -e

# Requirements
go get github.com/golang/lint/golint
go get github.com/smartystreets/goconvey

# Format, lint, check
FMT=`find . -type f -name '*.go' | grep -v -E '^./vendor' | xargs -L1 dirname | uniq | xargs gofmt -l`; if [ "$FMT" ]; then echo -e "fmt:\n$FMT"; fi
LINT=`find . -type f -name '*.go' | grep -v -E '^./vendor' | xargs -L1 dirname | uniq | xargs golint`; if [ "$LINT" ]; then echo -e "lint:\n$LINT"; fi
VET=`find . -type f -name '*.go' | grep -v -E '^./vendor' | xargs -L1 dirname | uniq | xargs go vet`; if [ "$VET" ]; then echo -e "vet:\n$VET"; fi
if [[ "$FMT" || "$LINT" || "$VET" ]]; then
  exit 1
fi

# Test
go test -v ./...
