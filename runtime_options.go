package goja

import (
	"reflect"

	"github.com/powerpuffpenguin/goja/loop"
)

var defaultOptions = options{}

type Option interface {
	apply(*options)
}
type Caller interface {
	// Before the native function is called, calling Before can be used as an interceptor, and returning errr will panic(r.NewGoError(err))
	Before(r *Runtime, call *FunctionCall) (err error)
	// Control how to call native functions, for example, you can start a new goroutine to call native functions
	Call(r *Runtime, callSlice bool, callable reflect.Value, in []reflect.Value) (out []reflect.Value, err error)
	// Called before returning the function call result out to js, used to convert the return value to js or filter the function return value
	After(r *Runtime, out []reflect.Value) (result Value, err error)
}
type CallerFactory interface {
	// Get a Caller
	Get() Caller
	// If you want to reuse, you can implement the Put function, otherwise save it as an empty implementation.
	Put(caller Caller)
}

type options struct {
	callerFactory CallerFactory
	scheduler     loop.Scheduler
	fieldGetter   func(reflect.Value) reflect.Value
}
type funcOption struct {
	f func(*options)
}

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}
func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}
func WithCallerFactory(callerFactory CallerFactory) Option {
	return newFuncOption(func(o *options) {
		o.callerFactory = callerFactory
	})
}
func WithScheduler(scheduler loop.Scheduler) Option {
	return newFuncOption(func(o *options) {
		o.scheduler = scheduler
	})
}
func WithFieldGetter(getter func(reflect.Value) reflect.Value) Option {
	return newFuncOption(func(o *options) {
		o.fieldGetter = getter
	})
}
