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