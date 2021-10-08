package goja_test

import (
	"errors"
	"reflect"
	"sync"
	"testing"

	"github.com/powerpuffpenguin/goja"
)

type caller struct {
}

func (*caller) Before(runtime *goja.Runtime, call *goja.FunctionCall) (err error) {
	return nil
}

func (*caller) Call(runtime *goja.Runtime, callSlice bool, callable reflect.Value, in []reflect.Value) (out []reflect.Value, err error) {
	if callSlice {
		out = callable.CallSlice(in)
	} else {
		out = callable.Call(in)
	}
	return
}

func (c *caller) After(runtime *goja.Runtime, out []reflect.Value) (result goja.Value, err error) {
	switch len(out) {
	case 0:
		result = goja.Undefined()
	case 1:
		result = runtime.ToValue(out[0].Interface())
	default:
		s := make([]interface{}, len(out))
		for i, v := range out {
			s[i] = v.Interface()
		}
		result = runtime.ToValue(s)
	}
	return
}

type callerFactory sync.Pool

func newFactory() *callerFactory {
	pool := &sync.Pool{
		New: func() interface{} {
			return &caller{}
		},
	}
	return (*callerFactory)(pool)
}
func (f *callerFactory) Get() goja.Caller {
	return ((*sync.Pool)(f)).Get().(*caller)
}
func (f *callerFactory) Put(caller goja.Caller) {
	((*sync.Pool)(f)).Put(caller)
}

func TestOptionCallerFactory(t *testing.T) {
	r := goja.New(goja.WithCallerFactory(newFactory()))

	r.Set(`make`, func(str string) (string, error) {
		return str, errors.New(str)
	})
	r.Set(`checkErr`, func(call goja.FunctionCall) goja.Value {
		var result bool
		if e, ok := call.Argument(0).Export().(error); ok {
			result = e.Error() == call.Argument(1).String()
		}
		return r.ToValue(result)
	})
	_, e := r.RunString(`
var s0="cerberus is an idea"
var [str,e] = make(s0)
if(str!=s0){
	throw new Error("not equal")
}
if(!checkErr(e,s0)){
	throw new Error("not error")
}

`)
	if e != nil {
		t.Fatal(e)
	}
}
