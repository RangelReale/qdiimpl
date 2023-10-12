# qdiimpl - Quick'n'Dirty Interface Implementation (for Golang)

`qdiimpl` is a Go generator cli that generates "Quick and **Dirty** Interface Implementations" meant for quick 
debugging, absolutely not production-ready.

This is for that time when you want to test one feature that depends on lots of external service interfaces like
databases, cloud storage, message queues, and you don't want to use the real thing just to test something that has
nothing to do with these interfaces.

Another option would be using a mock, however outside of tests they are cumbersome to use because mocks need to set 
expectations and usually are limited by execution amounts. 

## Install

```shell
$ go install github.com/RangelReale/qdiimpl/cmd/qdiimpl@master
```

## Usage

```shell
$ cd app/pkg/client
$ go run github.com/RangelReale/qdiimpl/cmd/qdiimpl@master -type=StorageClient
Writing file storageclient_qdii.go...
```

There is an option for each interface method called `WithDebugTYPEMETHOD` to set a function that will be called when
the method is called. If a method is called when a function is not set, the implementation panics with a useful
message.

## Command line parameters

```
Usage of qdiimpl:
        qdiimpl [flags] -type T [directory]
Flags:
  -data-type any
        add a data member of this type (e.g.: any, `package.com/data.XData`)
  -force-package-name string
        force generated package name
  -name-prefix string
        interface name prefix (default "QD")
  -name-suffix string
        interface name suffix (default blank)
  -output string
        output file name; default srcdir/<type>_qdii.go
  -overwrite
        overwrite file if exists
  -same-package
        if false will import source package and qualify the types (default true)
  -sync
        use mutex to prevent concurrent accesses (default true)
  -tags string
        comma-separated list of build tags to apply
  -type string
        type name; must be set
  -type-package string
        type package path if not the current directory
```

# Samples

### io.Reader

```shell
$ go run github.com/RangelReale/qdiimpl/cmd/qdiimpl -type=Reader -package=io -force-package=main
```

File: `reader_qdii.go`

```go
// Code generated by "qdiimpl"; DO NOT EDIT.
package main

import (
    "fmt"
    "io"
    "runtime"
    "sync"
)

type QDReaderContext struct {
    ExecCount  int
    CallerFunc string
    CallerFile string
    CallerLine int
}

type QDReader struct {
    lock      sync.Mutex
    execCount map[string]int
    implRead  func(qdCtx *QDReaderContext, p []byte) (n int, err error)
}

var _ io.Reader = (*QDReader)(nil)

type QDReaderOption func(*QDReader)

func NewQDReader(options ...QDReaderOption) *QDReader {
    ret := &QDReader{execCount: map[string]int{}}
    for _, opt := range options {
        opt(ret)
    }
    return ret
}

// Read implements [io.Reader.Read].
func (d *QDReader) Read(p []byte) (n int, err error) {
    return d.implRead(d.createContext("Read", d.implRead == nil), p)
}

func (d *QDReader) getCallerFuncName(skip int) (funcName string, file string, line int) {
    counter, file, line, success := runtime.Caller(skip)
    if !success {
        panic("runtime.Caller failed")
    }
    return runtime.FuncForPC(counter).Name(), file, line
}

func (d *QDReader) checkCallMethod(methodName string, implIsNil bool) (count int) {
    if implIsNil {
        panic(fmt.Errorf("[QDReader] method '%s' not implemented", methodName))
    }
    d.lock.Lock()
    defer d.lock.Unlock()
    d.execCount[methodName]++
    return d.execCount[methodName]
}

func (d *QDReader) createContext(methodName string, implIsNil bool) *QDReaderContext {
    callerFunc, callerFile, callerLine := d.getCallerFuncName(3)
    return &QDReaderContext{ExecCount: d.checkCallMethod(methodName, implIsNil), CallerFunc: callerFunc, CallerFile: callerFile, CallerLine: callerLine}
}

// Options

// WithQDReaderRead implements [io.Reader.Read].
func WithQDReaderRead(implRead func(qdCtx *QDReaderContext, p []byte) (n int, err error)) QDReaderOption {
    return func(d *QDReader) {
        d.implRead = implRead
    }
}
```

Usage:

```go
func main() {
    reader := NewQDReader(
        WithQDReaderRead(func(qdCtx *QDReaderContext, p []byte) (n int, err error) {
            n = copy(p, []byte("test"))
            return n, nil
        }),
    )

    readInterface(reader)
}

func readInterface(r io.Reader) {
    b := make([]byte, 10)

    n, err := r.Read(b)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%d: %v\n", n, b)
}
```

# Details

### QD Context

Each method is passed a "QDContext" which contains these fields:

- `ExecCount`: times this method was called since the execution start.
- `CallerFunc`: fully-qualified function name that called the interface method.
- `CallerFile`: source file name of the function that called the interface method.
- `CallerLine`: line number of the source file of the function that called the interface method.
- `Data`: a custom data field set by `WithDebugTYPEData` option. (only when `data-type` command line parameter is set)

Use these properties to help detect where the method was called from and return different responses if needed.

# Author

Rangel Reale (rangelreale@gmail.com)
