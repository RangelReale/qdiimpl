# qdiimpl - Quick'n'Dirty Interface Implementation (for Golang)

`qdiimpl` is a Go generator cli that generates "Quick and **Dirty** Interface Implementations" meant for quick 
debugging, absolutely not production-ready.

This is for that time when you want to test one feature that depends on lots of external service interfaces like
databases, cloud storage, message queues, and you don't want to use the real thing just to test something that has
nothing to do with these interfaces.

Another option would be using a mock, however outside of tests they are cumbersome to use because mocks need to set 
expectations and usually are limited by execution amounts. 

## Usage

```shell
$ cd app/pkg/client
$ go run github.com/RangelReale/qdiimpl/cmd/qdiimpl -type=StorageClient
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
  -force-package string
        force package name
  -name-prefix string
        interface name prefix; default is 'debug' (default "debug")
  -name-suffix string
        interface name suffix; default is blank
  -output string
        output file name; default srcdir/<type>_qdiimpl.go
  -overwrite
        overwrite file if exists
  -package string
        package name; if not set, use package from dir
  -same-package
        output package should be the same as the source (default true)
  -tags string
        comma-separated list of build tags to apply
  -type string
        type name; must be set
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
)

type DebugReaderContext struct {
    ExecCount  int
    CallerFunc string
    CallerFile string
    CallerLine int
    Data       any
}

type DebugReader struct {
    Data any

    execCount map[string]int
    implRead  func(debugCtx *DebugReaderContext, p []byte) (n int, err error)
}

var _ io.Reader = (*DebugReader)(nil)

type DebugReaderOption func(*DebugReader)

func NewDebugReader(options ...DebugReaderOption) *DebugReader {
    ret := &DebugReader{execCount: map[string]int{}}
    for _, opt := range options {
        opt(ret)
    }
    return ret
}

func (d *DebugReader) Read(p []byte) (n int, err error) {
    return d.implRead(d.createContext("Read", d.implRead == nil), p)
}

func (d *DebugReader) getCallerFuncName(skip int) (funcName string, file string, line int) {
    counter, file, line, success := runtime.Caller(skip)
    if !success {
        panic("runtime.Caller failed")
    }
    return runtime.FuncForPC(counter).Name(), file, line
}

func (d *DebugReader) checkCallMethod(methodName string, implIsNil bool) (count int) {
    if implIsNil {
        panic(fmt.Errorf("[DebugReader] method '%s' not implemented", methodName))
    }
    d.execCount[methodName]++
    return d.execCount[methodName]
}

func (d *DebugReader) createContext(methodName string, implIsNil bool) *DebugReaderContext {
    callerFunc, callerFile, callerLine := d.getCallerFuncName(3)
    return &DebugReaderContext{ExecCount: d.checkCallMethod(methodName, implIsNil), CallerFunc: callerFunc, CallerFile: callerFile, CallerLine: callerLine, Data: d.Data}
}

// Options

func WithDebugReaderData(data any) DebugReaderOption {
    return func(d *DebugReader) {
        d.Data = data
    }
}

func WithDebugReaderRead(implRead func(debugCtx *DebugReaderContext, p []byte) (n int, err error)) DebugReaderOption {
    return func(d *DebugReader) {
        d.implRead = implRead
    }
}
```

Usage:

```go
func main() {
    reader := NewDebugReader(
        WithDebugReaderRead(func(debugCtx *DebugReaderContext, p []byte) (n int, err error) {
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

    fmt.Printf("%d: %v", n, b)
}
```

# Details

### Debug Context

Each method is passed a "DebugContext" which contains these fields:

- `ExecCount`: times this method was called since the execution start.
- `CallerFunc`: fully-qualified function name that called the interface method.
- `CallerFile`: source file name of the function that called the interface method.
- `CallerLine`: line number of the source file of the function that called the interface method.
- `Data`: a custom data field set by `WithDebugTYPEData` option.

Use these properties to help detect where from the method was called from and return different responses if needed.

# Author

Rangel Reale (rangelreale@gmail.com)
