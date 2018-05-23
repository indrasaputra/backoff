[![Build Status](https://travis-ci.org/indrasaputra/backoff.svg?branch=master)](https://travis-ci.org/indrasaputra/backoff)
[![Test Coverage](https://api.codeclimate.com/v1/badges/d13382387166b72007db/test_coverage)](https://codeclimate.com/github/indrasaputra/backoff/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/d13382387166b72007db/maintainability)](https://codeclimate.com/github/indrasaputra/backoff/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/indrasaputra/backoff)](https://goreportcard.com/report/github.com/indrasaputra/backoff)
[![Documentation](https://godoc.org/github.com/indrasaputra/backoff?status.svg)](http://godoc.org/github.com/indrasaputra/backoff)

# Backoff

Backoff strategy written in Go

## Description

Backoff is an implementation of the popular backoff strategy. It is written in Go.

## Installation

```sh
go get github.com/indrasaputra/backoff
```

## Usage

There are two actual backoff implementations. The first one is `ConstantBackoff` and the second one is `ExponentialBackoff`.

```go
package main

import (
  "time"
  "github.com/indrasaputra/backoff"
)

func main() {
  b := &backoff.ConstantBackoff{
    BackoffInterval: 200 * time.Millisecond,
    JitterInterval:  50 * time.Millisecond,
  }

  // use NextInterval() to get the next interval
  interval := b.NextInterval()

  // use Reset() to reset the backoff
  b.Reset()
}
```

If you want to use `ExponentialBackoff`, simply follow this code

```go
package main

import (
  "time"
  "github.com/indrasaputra/backoff"
)

func main() {
  b := backoff.ExponentialBackoff{
    BackoffInterval: 300 * time.Millisecond,
    JitterInterval: 100 * time.Millisecond,
    MaxInterval: 3 * time.Second,
    Multiplier: 2,
  }
}
```