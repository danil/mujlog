# log0

[![Build Status](https://cloud.drone.io/api/badges/kvlog/log0/status.svg)](https://cloud.drone.io/kvlog/log0)
[![Go Reference](https://pkg.go.dev/badge/github.com/kvlog/log0.svg)](https://pkg.go.dev/github.com/kvlog/log0)

JSON logging for Go.  
Source files are distributed under the BSD-style license
found in the [LICENSE](./LICENSE) file.

## About

The software is considered to be at a alpha level of readiness -
its extremely slow and allocates a lots of memory)

## Install

    go get github.com/kvlog/log0@v0.173.0

## Usage

Set log0 as global logger

```go
package main

import (
    "os"
    "log"

    "github.com/kvlog/log0"
)

func main() {
    l := &log0.Log{
        Output:  os.Stdout,
        Keys:    [4]encoding.TextMarshaler{log0.String("message"), log0.String("excerpt")},
        Trunc:   12,
        Marks:   [3][]byte{[]byte("…")},
        Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
    }
    log.SetFlags(0)
    log.SetOutput(l)

    log.Print("Hello,\nWorld!")
}
```

Output:

```json
{
    "message":"Hello,\nWorld!",
    "excerpt":"Hello, World…"
}
```

## Use as GELF formater

```go
package main

import (
    "log"
    "os"

    "github.com/kvlog/log0"
)

func main() {
    l := log0.GELF()
    l.Output = os.Stdout
    log.SetFlags(0)
    log.SetOutput(l)
    log.Print("Hello,\nGELF!")
}
```

Output:

```json
{
    "version":"1.1",
    "short_message":"Hello, GELF!",
    "full_message":"Hello,\nGELF!",
    "timestamp":1602785340
}
```

## Caveat: numeric types appears in the message as a string

```go
package main

import (
    "log"
    "os"

    "github.com/kvlog/log0"
)

func main() {
    l := log0.Log{
        Output: os.Stdout,
        Keys:   [4]encoding.TextMarshaler{log0.String("message")},
    }
    log.SetFlags(0)
    log.SetOutput(l)

    log.Print(123)
    log.Print(3.21)
}
```

Output 1:

```json
{
    "message":"123"
}
```

Output 2:

```json
{
    "message":"3.21"
}
```

## Benchmark

```
go test -bench=. ./...
goos: linux
goarch: amd64
pkg: github.com/kvlog/log0
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkLog0/io.Writer_77-8         	  163160	      7221 ns/op	    1722 B/op	      53 allocs/op
BenchmarkLog0/fmt.Fprint_io.Writer_1127-8         	  102769	     12399 ns/op	    3401 B/op	      61 allocs/op
PASS
ok  	github.com/kvlog/log0	3.835s
```
