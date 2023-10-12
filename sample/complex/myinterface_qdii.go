// Code generated by "qdiimpl"; DO NOT EDIT.
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

type QDMyInterfaceContext struct {
	ExecCount  int
	CallerFunc string
	CallerFile string
	CallerLine int
	Data       any
}

type qdMyInterface[T any, X II] struct {
	DataQDII any

	lock            sync.Mutex
	execCount       map[string]int
	implCloseNotify func(qdCtx *QDMyInterfaceContext) <-chan bool
	implData        func(qdCtx *QDMyInterfaceContext)
	implGet         func(qdCtx *QDMyInterfaceContext, ctx context.Context, name string) (T, error)
	implOther       func(qdCtx *QDMyInterfaceContext, si SecondInterface) int
	implOther2      func(qdCtx *QDMyInterfaceContext, ti ThirdInterface[T]) int
	implSet         func(qdCtx *QDMyInterfaceContext, ctx context.Context, name string, value T) error
	implUnnamed     func(qdCtx *QDMyInterfaceContext, p0 bool, p1 string)
	implXGet        func(qdCtx *QDMyInterfaceContext, ss *SI) *SI
	implinternal    func(qdCtx *QDMyInterfaceContext) bool
}

type QDMyInterfaceOption[T any, X II] func(*qdMyInterface[T, X])

func NewQDMyInterface[T any, X II](options ...QDMyInterfaceOption[T, X]) MyInterface[T, X] {
	ret := &qdMyInterface[T, X]{execCount: map[string]int{}}
	for _, opt := range options {
		opt(ret)
	}
	return ret
}

// CloseNotify implements [main.MyInterface.CloseNotify].
func (d *qdMyInterface[T, X]) CloseNotify() <-chan bool {
	return d.implCloseNotify(d.createContext("CloseNotify", d.implCloseNotify == nil))
}

// Data implements [main.MyInterface.Data].
func (d *qdMyInterface[T, X]) Data() {
	d.implData(d.createContext("Data", d.implData == nil))
}

// Get implements [main.MyInterface.Get].
func (d *qdMyInterface[T, X]) Get(ctx context.Context, name string) (T, error) {
	return d.implGet(d.createContext("Get", d.implGet == nil), ctx, name)
}

// Other implements [main.MyInterface.Other].
func (d *qdMyInterface[T, X]) Other(si SecondInterface) int {
	return d.implOther(d.createContext("Other", d.implOther == nil), si)
}

// Other2 implements [main.MyInterface.Other2].
func (d *qdMyInterface[T, X]) Other2(ti ThirdInterface[T]) int {
	return d.implOther2(d.createContext("Other2", d.implOther2 == nil), ti)
}

// Set implements [main.MyInterface.Set].
func (d *qdMyInterface[T, X]) Set(ctx context.Context, name string, value T) error {
	return d.implSet(d.createContext("Set", d.implSet == nil), ctx, name, value)
}

// Unnamed implements [main.MyInterface.Unnamed].
func (d *qdMyInterface[T, X]) Unnamed(p0 bool, p1 string) {
	d.implUnnamed(d.createContext("Unnamed", d.implUnnamed == nil), p0, p1)
}

// XGet implements [main.MyInterface.XGet].
func (d *qdMyInterface[T, X]) XGet(ss *SI) *SI {
	return d.implXGet(d.createContext("XGet", d.implXGet == nil), ss)
}

// internal implements [main.MyInterface.internal].
func (d *qdMyInterface[T, X]) internal() bool {
	return d.implinternal(d.createContext("internal", d.implinternal == nil))
}

func (d *qdMyInterface[T, X]) getCallerFuncName(skip int) (funcName string, file string, line int) {
	counter, file, line, success := runtime.Caller(skip)
	if !success {
		panic("runtime.Caller failed")
	}
	return runtime.FuncForPC(counter).Name(), file, line
}

func (d *qdMyInterface[T, X]) checkCallMethod(methodName string, implIsNil bool) (count int) {
	if implIsNil {
		panic(fmt.Errorf("[qdMyInterface] method '%s' not implemented", methodName))
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	d.execCount[methodName]++
	return d.execCount[methodName]
}

func (d *qdMyInterface[T, X]) createContext(methodName string, implIsNil bool) *QDMyInterfaceContext {
	callerFunc, callerFile, callerLine := d.getCallerFuncName(3)
	return &QDMyInterfaceContext{ExecCount: d.checkCallMethod(methodName, implIsNil), CallerFunc: callerFunc, CallerFile: callerFile, CallerLine: callerLine, Data: d.DataQDII}
}

// Options

func WithQDMyInterfaceDataQDII[T any, X II](data any) QDMyInterfaceOption[T, X] {
	return func(d *qdMyInterface[T, X]) {
		d.DataQDII = data
	}
}

// WithqdMyInterfaceCloseNotify implements [main.MyInterface.CloseNotify].
func WithQDMyInterfaceCloseNotify[T any, X II](implCloseNotify func(qdCtx *QDMyInterfaceContext) <-chan bool) QDMyInterfaceOption[T, X] {
	return func(d *qdMyInterface[T, X]) {
		d.implCloseNotify = implCloseNotify
	}
}

// WithqdMyInterfaceData implements [main.MyInterface.Data].
func WithQDMyInterfaceData[T any, X II](implData func(qdCtx *QDMyInterfaceContext)) QDMyInterfaceOption[T, X] {
	return func(d *qdMyInterface[T, X]) {
		d.implData = implData
	}
}

// WithqdMyInterfaceGet implements [main.MyInterface.Get].
func WithQDMyInterfaceGet[T any, X II](implGet func(qdCtx *QDMyInterfaceContext, ctx context.Context, name string) (T, error)) QDMyInterfaceOption[T, X] {
	return func(d *qdMyInterface[T, X]) {
		d.implGet = implGet
	}
}

// WithqdMyInterfaceOther implements [main.MyInterface.Other].
func WithQDMyInterfaceOther[T any, X II](implOther func(qdCtx *QDMyInterfaceContext, si SecondInterface) int) QDMyInterfaceOption[T, X] {
	return func(d *qdMyInterface[T, X]) {
		d.implOther = implOther
	}
}

// WithqdMyInterfaceOther2 implements [main.MyInterface.Other2].
func WithQDMyInterfaceOther2[T any, X II](implOther2 func(qdCtx *QDMyInterfaceContext, ti ThirdInterface[T]) int) QDMyInterfaceOption[T, X] {
	return func(d *qdMyInterface[T, X]) {
		d.implOther2 = implOther2
	}
}

// WithqdMyInterfaceSet implements [main.MyInterface.Set].
func WithQDMyInterfaceSet[T any, X II](implSet func(qdCtx *QDMyInterfaceContext, ctx context.Context, name string, value T) error) QDMyInterfaceOption[T, X] {
	return func(d *qdMyInterface[T, X]) {
		d.implSet = implSet
	}
}

// WithqdMyInterfaceUnnamed implements [main.MyInterface.Unnamed].
func WithQDMyInterfaceUnnamed[T any, X II](implUnnamed func(qdCtx *QDMyInterfaceContext, p0 bool, p1 string)) QDMyInterfaceOption[T, X] {
	return func(d *qdMyInterface[T, X]) {
		d.implUnnamed = implUnnamed
	}
}

// WithqdMyInterfaceXGet implements [main.MyInterface.XGet].
func WithQDMyInterfaceXGet[T any, X II](implXGet func(qdCtx *QDMyInterfaceContext, ss *SI) *SI) QDMyInterfaceOption[T, X] {
	return func(d *qdMyInterface[T, X]) {
		d.implXGet = implXGet
	}
}

// WithqdMyInterfaceinternal implements [main.MyInterface.internal].
func WithQDMyInterfaceinternal[T any, X II](implinternal func(qdCtx *QDMyInterfaceContext) bool) QDMyInterfaceOption[T, X] {
	return func(d *qdMyInterface[T, X]) {
		d.implinternal = implinternal
	}
}
