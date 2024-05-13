#!/bin/bash

set -e
GOOS=windows GOARCH=arm64 go build ../..
sha256sum easySeed.exe > easySeed.sha256
