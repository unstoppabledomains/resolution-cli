#!/usr/bin/env bash

set -euo pipefail

go build -o ./test-cli ./resolution
go test
rm ./test-cli