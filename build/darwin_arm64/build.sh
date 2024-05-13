#!/bin/bash

set -e
GOOS=darwin GOARCH=arm64 go build ../..
sha256sum easySeed > easySeed.sha256
