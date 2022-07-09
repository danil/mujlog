# plog

[![Build Status](https://cloud.drone.io/api/badges/gorelib/plog/status.svg)](https://cloud.drone.io/gorelib/plog)
[![Go Reference](https://pkg.go.dev/badge/github.com/gorelib/plog.svg)](https://pkg.go.dev/github.com/gorelib/plog)

JSON logging for Go.  
Source files are distributed under the BSD-style license.

## About

The software is considered to be at a alpha level of readiness -
its extremely slow and allocates a lots of memory)

## Usage

Set plog as global logger

```go
package main

import (
    "os"
    "log"

    "github.com/gorelib/plog"
)

func main() {
    l := &plog.Log{
        Output:  os.Stdout,
        Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
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

    "github.com/gorelib/plog"
)

func main() {
    l := plog.GELF()
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

    "github.com/gorelib/plog"
)

func main() {
    l := plog.Log{
        Output: os.Stdout,
        Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
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

```sh
$ go test -count=1 -race -bench ./... 
goos: linux
goarch: amd64
pkg: github.com/gorelib/plog
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkPlog/plog_test.go:90/io.Writer-8         	   42229	     27608 ns/op
BenchmarkPlog/plog_test.go:1138/fmt.Fprint_io.Writer-8         	   15640	     76490 ns/op
PASS
ok  	github.com/gorelib/plog	3.547s
```
