#!/bin/bash

set -e
GOOS=darwin GOARCH=amd64 go build ../..
sha256sum easySeed > easySeed.sha256
