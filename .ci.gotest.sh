#!/bin/bash

set -e

main() {
    mkdir -p ./coverage

    gotestsum --rerun-fails="1" \
        --packages="$(go list ./... | grep -v '/mocks/' | grep -v '/example')" \
        --format-icons="hivis" \
        -- \
        -race \
        -timeout="600s" \
        -parallel="4" \
        --tags="" \
        -coverprofile="./coverage/coverage.txt" \
        -covermode="atomic"

    go tool cover -html="./coverage/coverage.txt" -o="./coverage/coverage.html"
}

main
