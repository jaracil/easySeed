#!/bin/bash

set -e
GOOS=linux GOARCH=amd64 go build ../..
sha256sum easySeed > easySeed.sha256
