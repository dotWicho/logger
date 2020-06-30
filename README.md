# logger

[![Go](https://github.com/dotWicho/logger/workflows/Go/badge.svg?branch=master)](https://github.com/dotWicho/logger)
[![Quality Report](https://goreportcard.com/badge/github.com/dotWicho/logger)](https://goreportcard.com/badge/github.com/dotWicho/logger)
[![GoDoc](https://godoc.org/github.com/dotWicho/logger?status.svg)](https://pkg.go.dev/github.com/dotWicho/logger?tab=doc)

## Library to manage a logger easily

## Getting started

- API documentation is available via [godoc](https://godoc.org/github.com/dotWicho/logger).
- The [standard_test.go](./standard_test.go) contains some examples of application.

## Installation

To install Marathon package, you need to install Go and set your Go workspace first.

1 - The first need [Go](https://golang.org/) installed (**version 1.13+ is required**).
Then you can use the below Go command to install logger

```bash
$ go get -u github.com/dotWicho/logger
```

And then Import it in your code:

``` go
package main

import "github.com/dotWicho/logger"
```
Or

2 - Use as module in you project (go.mod file):

``` go
module myapp

go 1.13

require (
	github.com/dotWicho/logger v1.0.0
)
```
## Usage

``` go
package main

import "github.com/dotWicho/logger"

func main () {

	// We define some vars
	buffer := &bytes.Buffer{}

	// Try to create Logger
	_logger := NewLogger(true)
	_logger.SetLevel(logrus.InfoLevel)

	// Some lines to _logger
	_logger.Info("This line appears on console? %v", true)
	_logger.Info("Another line with number %d", 22)

	_logger.Warn("This line appears on console? %v", true)
	_logger.Warn("Another line with number %d", 22)
}

Console Output:
INFO[2020-06-30 17:20:29] This line appears on console? true           
INFO[2020-06-30 17:20:29] Another line with number 22    
WARN[2020-06-30 17:20:29] This line appears on console? true           
WARN[2020-06-30 17:20:29] Another line with number 22          

```

## Contributing

- Get started by checking our [contribution guidelines](https://github.com/dotWicho/logger/blob/master/CONTRIBUTING.md).
- If you have any questions, just ask!

