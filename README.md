# Proto Nats

[![Build Status](https://img.shields.io/travis/faabiosr/proto-nats/master.svg?style=flat-square)](https://travis-ci.org/faabiosr/proto-nats)
[![Codecov branch](https://img.shields.io/codecov/c/github/faabiosr/proto-nats/master.svg?style=flat-square)](https://codecov.io/gh/faabiosr/proto-nats)
[![Release](https://img.shields.io/github/release/faabiosr/proto-nats.svg?style=flat-square)](https://github.com/faabiosr/proto-nats/releases/latest)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/faabiosr/proto-nats)
[![Go Report Card](https://goreportcard.com/badge/github.com/faabiosr/proto-nats?style=flat-square&)](https://goreportcard.com/report/github.com/faabiosr/proto-nats)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://github.com/faabiosr/proto-nats/blob/master/LICENSE)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)


## Description

`proto-nats` provides a plugin for the `protoc compiler`. It generates the client and server for working with NATS.


## Installation

Download the plugin from the [Releases](https://github.com/faabiosr/proto-nats/releases) page.

### Shell script:
```sh
$ curl -Ls https://git.io/proto-nats | bash
```


## Development

### Requirements

- Install [Go](https://golang.org)

### Makefile
```sh
# Build (dev only)
$ make build

# Clean up
$ make clean

# Run tests and generates html coverage file
$ make cover

# Format all go files
$ make fmt

# Pack the template using [Pkger](https://github.com/markbates/pkger)
$ make pack

# Run tests
$ make test
```

## License

This project is released under the MIT licence. See [LICENSE](https://github.com/faabiosr/proto-nats/blob/master/LICENSE) for more details.
