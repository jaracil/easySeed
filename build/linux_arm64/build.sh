#!/bin/bash

set -e
GOOS=linux GOARCH=arm64 go build ../..
sha256sum easySeed > easySeed.sha256
