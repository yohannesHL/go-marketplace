#!/bin/bash

if [ -z "${VERSION}" ]; then
    echo "VERSION must be set"
    exit 1
fi

# Disable C code, enable Go modules
export CGO_ENABLED=0
export GO111MODULE=on
export GOFLAGS="-mod=vendor"
export GOBIN="$PWD/bin"

go install -v -ldflags="-X $(go list -m)/pkg/version.VERSION=${VERSION}" ./...