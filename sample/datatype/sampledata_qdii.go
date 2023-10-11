// Code generated by "qdiimpl"; DO NOT EDIT.
package main

import (
	"fmt"
	idata "github.com/RangelReale/qdiimpl/sample/datatype/idata"
	"runtime"
)

type DebugSampleDataContext struct {
	ExecCount  int
	CallerFunc string
	CallerFile string
	CallerLine int
	Data       *idata.IData
}

type DebugSampleData struct {
	Data *idata.IData

	execCount map[string]int
	implGet   func(debugCtx *DebugSampleDataContext, name string) (any, error)
}

var _ SampleData = (*DebugSampleData)(nil)

type DebugSampleDataOption func(*DebugSampleData)

func NewDebugSampleData(options ...DebugSampleDataOption) *DebugSampleData {
	ret := &DebugSampleData{execCount: map[string]int{}}
	for _, opt := range options {
		opt(ret)
	}
	return ret
}

func (d *DebugSampleData) Get(name string) (any, error) {
	return d.implGet(d.createContext("Get", d.implGet == nil), name)
}

func (d *DebugSampleData) getCallerFuncName(skip int) (funcName string, file string, line int) {
	counter, file, line, success := runtime.Caller(skip)
	if !success {
		panic("runtime.Caller failed")
	}
	return runtime.FuncForPC(counter).Name(), file, line
}

func (d *DebugSampleData) checkCallMethod(methodName string, implIsNil bool) (count int) {
	if implIsNil {
		panic(fmt.Errorf("[DebugSampleData] method '%s' not implemented", methodName))
	}
	d.execCount[methodName]++
	return d.execCount[methodName]
}

func (d *DebugSampleData) createContext(methodName string, implIsNil bool) *DebugSampleDataContext {
	callerFunc, callerFile, callerLine := d.getCallerFuncName(3)
	return &DebugSampleDataContext{ExecCount: d.checkCallMethod(methodName, implIsNil), CallerFunc: callerFunc, CallerFile: callerFile, CallerLine: callerLine, Data: d.Data}
}

// Options

func WithDebugSampleDataData(data *idata.IData) DebugSampleDataOption {
	return func(d *DebugSampleData) {
		d.Data = data
	}
}

func WithDebugSampleDataGet(implGet func(debugCtx *DebugSampleDataContext, name string) (any, error)) DebugSampleDataOption {
	return func(d *DebugSampleData) {
		d.implGet = implGet
	}
}