#!/usr/bin/env bash
set -euo pipefail

VERSION="${VERSION:-1.1.2}"
mkdir -p dist
export CGO_ENABLED=1

go build \
  -buildvcs=false \
  -buildmode=c-shared \
  -trimpath \
  -ldflags="-s -w" \
  -o dist/grok-manager.so \
  .

cp -f dist/grok-manager.so dist/grok-manager-linux-amd64.so
cp -f dist/grok-manager.so "dist/grok-manager-v${VERSION}.so"

echo "Built dist/grok-manager.so (v${VERSION})"
ls -lh dist/grok-manager*.so
