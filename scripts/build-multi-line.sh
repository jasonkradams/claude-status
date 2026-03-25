#!/usr/bin/env bash
set -euo pipefail

GOOS="${GOOS:-$(go env GOOS)}"
GOARCH="${GOARCH:-$(go env GOARCH)}"
GOBIN="${GOBIN:-$(pwd)/bin}"

echo "==> Formatting..."
go fmt ./...

echo "==> Vetting..."
go vet ./...

echo "==> Fixing..."
go fix ./...

echo "==> Building multi-line..."
mkdir -p "$GOBIN"
GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$GOBIN/multi-line" ./cmd/multi-line
