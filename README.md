# FeedHub

[![license](https://img.shields.io/github/license/imcrazytwkr/feedhub)](LICENSE)

A simple RSS2 and Atom feed provider for sites and services that do not provide any out of box.
The project is heavily inspired by [RSSHub](https://github.com/DIYgod/RSSHub).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine
for development and testing purposes. See deployment for notes on how to deploy the project
on a live system.

### Prerequisites

Requires Go v1.16 or newer (tested on 1.19 and 1.20 only).

### Building

Build process is automated using [make](https://en.wikipedia.org/wiki/Make_(software)).
Assuming you already have `make` installed on your system, you can build `feedhub` using it:

```sh
$ git clone 'https://github.com/imcrazytwkr/feedhub'
$ cd feedhub

# Build using make
$ make

# Alternative (manual) build command
$ go build -o feedhub
```

## Running

Feedhub can be run by simply executing the binary file produced by `go build (running
it from the root user is possible but is highly discouraged):

```sh
$ ./feedhub
```

## Configuration

At current stage, Feedhub is only configured through the environment variables. If the need for the
config file arises, it will be added.

### Common

- `HOST` sets an IP address for ingressor to listen on. Can be left empty for listening on all
  available interfaces.
- `PORT` sets a port for ingressor to listen in. Must be within range of 0-65535 if specified.

## Built With

* [gin](https://github.com/gin-gonic/gin) - HTTP web framework written in Go with Martini-like API
* [fastjson](https://github.com/valyala/fastjson) - fastest JSON parser for Go
* [zerolog](https://github.com/rs/zerolog) - Zero Allocation JSON Logger

## License

[MIT © Denis Chernov](./LICENSE)
