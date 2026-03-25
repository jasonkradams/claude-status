#!/usr/bin/env bash
set -euo pipefail

NAME="${1:?usage: build-widget.sh <name> <pkg>}"
PKG="${2:?usage: build-widget.sh <name> <pkg>}"

GOOS="${GOOS:-$(go env GOOS)}"
GOARCH="${GOARCH:-$(go env GOARCH)}"
GOBIN="${GOBIN:-$(pwd)/bin}"

echo "==> Building $NAME..."
mkdir -p "$GOBIN"
GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$GOBIN/$NAME" "$PKG"
