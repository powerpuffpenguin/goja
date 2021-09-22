package goja

type promiseState uint8

const (
	pending promiseState = iota + 1
	fulfilled
	rejected
)

func (s promiseState) String() string {
	switch s {
	case fulfilled:
		return `fulfilled`
	case rejected:
		return `rejected`
	}
	return `pending`
}

type promiseImpl struct {
	runtime   *Runtime
	ctor      *Object
	self      *Object
	state     promiseState
	callbacks []promiseCallback
	result    Value
	err       Value
	completed bool
}
type promiseCallback struct {
	onFulfilled, onRejected, onFinally Callable
}

func newPromiseImpl(runtime *Runtime, ctor *Object, self *Object) *promiseImpl {
	return &promiseImpl{
		runtime: runtime,
		ctor:    ctor,
		self:    self,
		state:   pending,
		result:  _undefined,
		err:     _undefined,
	}
}
func (impl *promiseImpl) register(executor Callable) {
	self := impl.self
	runtime := impl.runtime
	e := self.Set(`toString`, impl.toString)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	e = self.Set(`then`, impl.then)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	e = self.Set(`catch`, impl.catch)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	e = self.Set(`finally`, impl.finally)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	_, e = executor(_undefined,
		runtime.newNativeFunc(impl.resolve, nil, "resolve", nil, 1),
		runtime.newNativeFunc(impl.reject, nil, "reject", nil, 1),
	)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
}
func (impl *promiseImpl) String() string {
	return `Promise { <` + impl.state.String() + `> }`
}
func (impl *promiseImpl) toString(call FunctionCall) Value {
	return newStringValue(impl.String())
}
func (impl *promiseImpl) resolve(call FunctionCall) Value {
	impl.resolveHandle(call.Argument(0))
	return _undefined
}
func (impl *promiseImpl) reject(call FunctionCall) Value {
	impl.rejectHandle(call.Argument(0))
	return _undefined
}
func (impl *promiseImpl) resolveHandle(v Value) {
	if impl.state != pending {
		return
	}
	impl.state = fulfilled
	impl.result = v
	for _, callback := range impl.callbacks {
		if callback.onFulfilled != nil {
			callback.onFulfilled(_undefined, v)
		}
	}
	for _, callback := range impl.callbacks {
		if callback.onFinally != nil {
			callback.onFinally(_undefined)
		}
	}
	impl.completed = true
}
func (impl *promiseImpl) rejectHandle(v Value) {
	if impl.state != pending {
		return
	}
	impl.state = rejected
	impl.err = v
	for _, callback := range impl.callbacks {
		if callback.onRejected != nil {
			callback.onRejected(_undefined, v)
		}
	}
	impl.completed = true
}
func (impl *promiseImpl) then(call FunctionCall) Value {
	onFulfilled, _ := AssertFunction(call.Argument(0))
	onRejected, _ := AssertFunction(call.Argument(1))
	executor := newPromiseExecutor(impl.runtime, impl.ctor, onFulfilled, onRejected, nil)
	callback := promiseCallback{
		onFulfilled: executor.resolve,
		onRejected:  executor.reject,
	}
	if impl.completed {
		if impl.state == fulfilled {
			callback.onFulfilled(_undefined, impl.result)
		} else {
			callback.onRejected(_undefined, impl.err)
		}
	} else {
		impl.callbacks = append(impl.callbacks, callback)
	}
	return executor.result
}
func (impl *promiseImpl) catch(call FunctionCall) Value {
	onRejected, _ := AssertFunction(call.Argument(0))
	executor := newPromiseExecutor(impl.runtime, impl.ctor, nil, onRejected, nil)
	callback := promiseCallback{
		onFulfilled: executor.resolve,
		onRejected:  executor.reject,
	}
	if impl.completed {
		if impl.state == fulfilled {
			callback.onFulfilled(_undefined, impl.result)
		} else {
			callback.onRejected(_undefined, impl.err)
		}
	} else {
		impl.callbacks = append(impl.callbacks, callback)
	}
	return executor.result
}
func (impl *promiseImpl) finally(call FunctionCall) Value {
	onFinally, _ := AssertFunction(call.Argument(0))
	executor := newPromiseExecutor(impl.runtime, impl.ctor, nil, nil, onFinally)
	callback := promiseCallback{
		onFulfilled: executor.resolve,
		onRejected:  executor.reject,
	}
	if impl.completed {
		if impl.state == fulfilled {
			callback.onFulfilled(_undefined, impl.result)
		} else {
			callback.onRejected(_undefined, impl.err)
		}
	} else {
		impl.callbacks = append(impl.callbacks, callback)
	}
	return executor.result
}

type promiseExecutor struct {
	runtime                                  *Runtime
	onFulfilled, onRejected, onFinally       Callable
	argResolve, argReject                    Value
	selfResolve, selfReject, resolve, reject Callable
	result                                   Value
}

func (p *promiseExecutor) handle(call FunctionCall) Value {
	p.argResolve = call.Argument(0)
	p.argReject = call.Argument(1)
	p.selfResolve, _ = AssertFunction(p.argResolve)
	p.selfReject, _ = AssertFunction(p.argReject)
	return _undefined
}
func (p *promiseExecutor) resolveHandle(call FunctionCall) Value {
	var (
		runtime       = p.runtime
		result  Value = _undefined
		e       error
	)
	if p.onFinally != nil {
		result, e = p.onFinally(_undefined)
		if e != nil {
			p.selfReject(_undefined, runtime.NewGoError(e))
			return _undefined
		}
	} else if p.onFulfilled != nil {
		result, e = p.onFulfilled(_undefined, call.Argument(0))
		if e != nil {
			p.selfReject(_undefined, runtime.NewGoError(e))
			return _undefined
		}
	}
	p.resolveResult(result)
	return _undefined
}
func (p *promiseExecutor) rejectHandle(call FunctionCall) Value {
	var (
		runtime       = p.runtime
		result  Value = _undefined
		e       error
	)
	if p.onFinally != nil {
		result, e = p.onFinally(_undefined)
		if e != nil {
			p.selfReject(_undefined, runtime.NewGoError(e))
			return _undefined
		}
	} else if p.onRejected != nil {
		result, e = p.onRejected(_undefined, call.Argument(0))
		if e != nil {
			p.selfReject(_undefined, runtime.NewGoError(e))
			return _undefined
		}
	}
	p.resolveResult(result)
	return _undefined
}
func (p *promiseExecutor) resolveResult(val Value) {
	if obj, ok := val.(*Object); ok {
		if callable, ok := AssertFunction(obj.Get(`then`)); ok {
			_, e := callable(_undefined,
				p.argResolve,
				p.argReject,
			)
			if e != nil {
				p.selfReject(_undefined, p.runtime.NewGoError(e))
			}
			return
		}
	}
	p.selfResolve(_undefined, val)
}
func newPromiseExecutor(runtime *Runtime, ctor *Object,
	onFulfilled, onRejected, onFinally Callable,
) (executor *promiseExecutor) {
	executor = &promiseExecutor{
		runtime:     runtime,
		onFulfilled: onFulfilled,
		onRejected:  onRejected,
		onFinally:   onFinally,
	}
	executor.resolve, _ = AssertFunction(runtime.newNativeFunc(executor.resolveHandle, nil, "onFulfilled", nil, 1))
	executor.reject, _ = AssertFunction(runtime.newNativeFunc(executor.rejectHandle, nil, "onRejected", nil, 1))
	executor.result, _ = runtime.New(ctor,
		runtime.newNativeFunc(executor.handle, nil, "executor", nil, 2),
	)
	return executor
}

func (f *factoryPromise) constructor(call ConstructorCall) *Object {
	runtime := f.runtime
	executor, ok := AssertFunction(call.Argument(0))
	if !ok {
		panic(runtime.NewTypeError(`Promise executor is not a function`))
	}
	newPromiseImpl(runtime, f.ctor, call.This).register(executor)
	return nil
}
func (f *factoryPromise) register() {
	self := f.ctor
	runtime := f.runtime
	e := self.Set(`resolve`, f.resolve)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	e = self.Set(`reject`, f.reject)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
}
func (f *factoryPromise) resolve(call FunctionCall) Value {
	runtime := f.runtime
	completer, e := NewCompleter(runtime)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	completer.Resolve(call.Argument(0))
	return completer.promise
}
func (f *factoryPromise) reject(call FunctionCall) Value {
	runtime := f.runtime
	completer, e := NewCompleter(runtime)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	completer.Reject(call.Argument(0))
	return completer.promise
}
func (f *factoryPromise) completer(call ConstructorCall) *Object {
	runtime := f.runtime
	completer, e := NewCompleter(runtime)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	completer.register(call.This)
	return nil
}

type factoryPromise struct {
	runtime *Runtime
	ctor    *Object
}
type Completer struct {
	runtime         *Runtime
	resolve, reject Callable
	promise         Value
	completed       bool
}

func NewCompleter(runtime *Runtime) (completer *Completer, e error) {
	completer = &Completer{
		runtime: runtime,
	}
	promise, e := runtime.New(runtime.Get(`Promise`),
		runtime.newNativeFunc(completer.handle, nil, "executor", nil, 2),
	)
	if e != nil {
		completer = nil
		return
	}
	completer.promise = promise
	return
}
func (c *Completer) handle(call FunctionCall) Value {
	c.resolve, _ = AssertFunction(call.Argument(0))
	c.reject, _ = AssertFunction(call.Argument(1))
	return _undefined
}
func (c *Completer) Resolve(v Value) {
	if !c.completed {
		c.resolve(_undefined, v)
		c.completed = true
	}
}
func (c *Completer) Reject(reason Value) {
	if !c.completed {
		c.reject(_undefined, reason)
		c.completed = true
	}
}
func (c *Completer) register(self *Object) {
	runtime := c.runtime
	e := self.Set(`toString`, c.toString)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	e = self.Set(`resolve`, c.jsResolve)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	e = self.Set(`reject`, c.jsReject)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	e = self.DefineAccessorProperty(`promise`,
		runtime.ToValue(c.getPromise),
		nil,
		FLAG_TRUE, FLAG_TRUE,
	)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	e = self.DefineAccessorProperty(`completed`,
		runtime.ToValue(c.getCompleted),
		nil,
		FLAG_TRUE, FLAG_TRUE,
	)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
}

func (c *Completer) String() string {
	if c.completed {
		return `Completer { completed:true }`
	} else {
		return `Completer { completed:false }`
	}
}
func (c *Completer) toString(call FunctionCall) Value {
	return newStringValue(c.String())
}
func (c *Completer) jsResolve(call FunctionCall) Value {
	c.Resolve(call.Argument(0))
	return _undefined
}
func (c *Completer) jsReject(call FunctionCall) Value {
	c.Reject(call.Argument(0))
	return _undefined
}
func (c *Completer) getPromise(call FunctionCall) Value {
	return c.promise
}
func (c *Completer) getCompleted(call FunctionCall) Value {
	if c.completed {
		return valueTrue
	}
	return valueFalse
}
func (r *Runtime) pp_expand_get_scheduler(call FunctionCall) Value {
	return r.ToValue(r.Loop().GetScheduler())
}
func (r *Runtime) pp_expand_init_promise() {
	var factory factoryPromise
	factory.runtime = r
	factory.ctor = r.newNativeConstructor(factory.constructor, "Promise", 1)
	r.addToGlobal(`Promise`, factory.ctor)
	factory.register()
	r.addToGlobal(`Completer`, r.newNativeConstructor(factory.completer, "Completer", 0))
	r.GlobalObject().self.setOwnStr("defaultScheduler", &valueProperty{
		configurable: true,
		getterFunc:   r.newNativeFunc(r.pp_expand_get_scheduler, nil, "defaultScheduler", nil, 0),
		accessor:     true,
	}, false)
}
