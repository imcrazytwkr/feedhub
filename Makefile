# patsubst is used to remove trailing slash from path
PROJECT_ROOT := $(patsubst %/,%,$(dir $(realpath $(lastword $(MAKEFILE_LIST)))))

all: clean build

clean:
	rm -f '$(PROJECT_ROOT)/feedhub'
	find '$(PROJECT_ROOT)' -type f -iname '*.pb.go' -delete
	go clean -testcache

build:
	go build -o '$(PROJECT_ROOT)/feedhub'

test:
	find '$(PROJECT_ROOT)' -type f -iname '*_test.go' | xargs dirname | sort | uniq | xargs go test -timeout 30s
