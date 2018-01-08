#!/bin/bash

set -e

# Requirements
go get github.com/golang/lint/golint
go get -t ./...

# Format, lint, check
OUT=`gofmt -l . | (grep -v '^vendor\/' || true)`; if [ "$OUT" ]; then echo "gofmt: $OUT"; exit 1; fi
OUT=`golint ./... | (grep -v '^vendor\/' || true)`; if [ "$OUT" ]; then echo "golint: $OUT"; exit 1; fi
OUT=`find . -type f -name '*.go' | grep -v -E '^./vendor' | xargs -L1 dirname | uniq | xargs go vet`; if [ "$OUT" ]; then echo "govet: $OUT"; exit 1; fi

# Test
go test -v ./...
