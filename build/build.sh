#!/bin/bash

set -e
GOOS=linux GOARCH=amd64 go build -o easySeed_linux_amd64 ..
GOOS=linux GOARCH=arm64 go build -o easySeed_linux_arm64 ..
GOOS=darwin GOARCH=amd64 go build -o easySeed_darwin_amd64 ..
GOOS=darwin GOARCH=arm64 go build -o easySeed_darwin_arm64 ..
GOOS=windows GOARCH=amd64 go build -o easySeed_windows_amd64.exe ..
GOOS=windows GOARCH=arm64 go build -o easySeed_windows_arm64.exe ..
 
