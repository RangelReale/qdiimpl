// Code generated by "qdiimpl"; DO NOT EDIT.
package main

import (
	"fmt"
	"io"
	"runtime"
	"sync"
)

type ReaderContext struct {
	MethodName     string
	ExecCount      int
	CallerFunc     string
	CallerFile     string
	CallerLine     int
	isNotSupported bool
}

// NotSupported should be called if the current callback don't support the passed arguments.
// The function return values will be ignored.
func (c *ReaderContext) NotSupported() {
	c.isNotSupported = true
}

type Reader struct {
	lock                   sync.Mutex
	execCount              map[string]int
	fallback               io.Reader
	onMethodNotImplemented func(qdCtx *ReaderContext, hasCallbacks bool) error
	implRead               []func(qdCtx *ReaderContext, p []byte) (n int, err error)
}

var _ io.Reader = (*Reader)(nil)

type ReaderOption func(*Reader)

func NewReader(options ...ReaderOption) io.Reader {
	ret := &Reader{execCount: map[string]int{}}
	for _, opt := range options {
		opt(ret)
	}
	return ret
}

// Read implements [io.Reader.Read].
func (d *Reader) Read(p []byte) (n int, err error) {
	const methodName = "Read"
	for _, impl := range d.implRead {
		qctx := d.createContext(methodName)
		r0, r1 := impl(qctx, p)
		if !qctx.isNotSupported {
			d.addCallMethod(methodName)
			return r0, r1
		}
	}
	if d.fallback != nil {
		return d.fallback.Read(p)
	}
	panic(d.methodNotImplemented(d.createContext(methodName), len(d.implRead) > 0))
}

func (d *Reader) getCallerFuncName(skip int) (funcName string, file string, line int) {
	counter, file, line, success := runtime.Caller(skip)
	if !success {
		panic("runtime.Caller failed")
	}
	return runtime.FuncForPC(counter).Name(), file, line
}

func (d *Reader) addCallMethod(methodName string) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.execCount[methodName]++
}

func (d *Reader) createContext(methodName string) *ReaderContext {
	callerFunc, callerFile, callerLine := d.getCallerFuncName(3)
	d.lock.Lock()
	defer d.lock.Unlock()
	return &ReaderContext{
		MethodName: methodName,
		ExecCount:  d.execCount[methodName],
		CallerFunc: callerFunc,
		CallerFile: callerFile,
		CallerLine: callerLine,
	}
}

func (d *Reader) methodNotImplemented(qdCtx *ReaderContext, hasCallbacks bool) error {
	if d.onMethodNotImplemented != nil {
		if err := d.onMethodNotImplemented(qdCtx, hasCallbacks); err != nil {
			return err
		}
	}
	msg := "not implemented"
	if hasCallbacks {
		msg = "not supported by any callbacks"
	}
	return fmt.Errorf("[Reader] method '%s' %s", qdCtx.MethodName, msg)
}

// Options

func WithFallback(fallback io.Reader) ReaderOption {
	return func(d *Reader) {
		d.fallback = fallback
	}
}

func WithOnMethodNotImplemented(m func(qdCtx *ReaderContext, hasCallbacks bool) error) ReaderOption {
	return func(d *Reader) {
		d.onMethodNotImplemented = m
	}
}

// WithRead implements [io.Reader.Read].
func WithRead(implRead func(qdCtx *ReaderContext, p []byte) (n int, err error)) ReaderOption {
	return func(d *Reader) {
		d.implRead = append(d.implRead, implRead)
	}
}
