# patsubst is used to remove trailing slash from path
PROJECT_ROOT := $(patsubst %/,%,$(dir $(realpath $(lastword $(MAKEFILE_LIST)))))

.PHONY: all clean build format test

all: clean build

clean:
	rm -f '$(PROJECT_ROOT)/feedhub'
	find '$(PROJECT_ROOT)' -type f -iname '*.pb.go' -delete
	go clean -testcache

build:
	go build -o '$(PROJECT_ROOT)/feedhub'

format:
	go fmt '$(PROJECT_ROOT)/...'
	go mod tidy

test: format
	go test '$(PROJECT_ROOT)/...' | sed '/^?/d'
