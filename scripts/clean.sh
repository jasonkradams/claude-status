#!/usr/bin/env bash
set -euo pipefail

GOBIN="${GOBIN:-$(pwd)/bin}"

echo "==> Cleaning build artifacts..."
rm -rf "$GOBIN"
