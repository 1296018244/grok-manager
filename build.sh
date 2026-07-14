#!/usr/bin/env bash
set -euo pipefail

mkdir -p dist
export CGO_ENABLED=1

go build \
  -buildvcs=false \
  -buildmode=c-shared \
  -trimpath \
  -ldflags="-s -w" \
  -o dist/grok-manager.so \
  .

echo "Built dist/grok-manager.so"
ls -lh dist/grok-manager.so
