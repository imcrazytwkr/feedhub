#!/bin/sh

exec find "$PWD" -type f -iname '*_test.go' | xargs dirname | sort | uniq | xargs go test -timeout 30s
