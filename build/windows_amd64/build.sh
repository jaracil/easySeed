#!/bin/bash

set -e
GOOS=windows GOARCH=amd64 go build ../..
sha256sum easySeed.exe > easySeed.sha256
