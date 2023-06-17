# patsubst is used to remove trailing slash from path
PROJECT_ROOT := $(patsubst %/,%,$(dir $(realpath $(lastword $(MAKEFILE_LIST)))))

all: clean proto build

clean:
	rm -f '$(PROJECT_ROOT)/feedhub'
	find '$(PROJECT_ROOT)' -type f -iname '*.pb.go' -delete

proto:
	find '$(PROJECT_ROOT)' -type f -iname '*.proto' | xargs protoc -I='$(PROJECT_ROOT)' --go_out='$(PROJECT_ROOT)'

build:
	go build -o '$(PROJECT_ROOT)/feedhub'
