# http2curl
:triangular_ruler: Convert Golang's http.Request to CURL commands for either:

- Printing in logs or for reference
- Executing via os/exec's exec.Command()

[![Build Status](https://travis-ci.org/dkoston/http2curl.svg?branch=master)](https://travis-ci.org/dkoston/http2curl)
[![GoDoc](https://godoc.org/github.com/dkoston/http2curl?status.svg)](https://godoc.org/github.com/dkoston/http2curl)
[![Coverage Status](https://coveralls.io/repos/dkoston/http2curl/badge.svg)](https://coveralls.io/github/dkoston/http2curl)


To do the reverse, check out [mholt/curl-to-go](https://github.com/mholt/curl-to-go).

## Examples

### Printing the curl command

For code see: [./examples/printcurl/main.go](./examples/printcurl/main.go)

To run:

```bash
go run examples/printcurl/main.go
```

### exec.Command() execution of curl command

For code see: [./examples/execcommand/main.go](./examples/execcommand/main.go)

To run:

```bash
go run examples/execcommand/main.go
```

## Using within your golang code / Install

```bash
$ go get github.com/dkoston/http2curl
```

## License

MIT

## Attribution

Forked from [https://github.com/moul/http2curl](https://github.com/moul/http2curl)

motivation: moul's library was focused on printing the command only. I also 
wanted to be able to pass the command to exec.Command() easily
