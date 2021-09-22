package goja

import (
	"time"
)

type timeoutImpl struct {
	*Future
	running bool
	timer   *time.Timer
}

func newTimeoutImpl(future *Future, duration time.Duration) (timeout *timeoutImpl) {
	timeout = &timeoutImpl{
		Future:  future,
		running: true,
	}
	timeout.serve(duration)
	return
}
func (t *timeoutImpl) serve(duration time.Duration) {
	t.Async(nil)
	if duration <= 0 {
		t.Result(t)
		return
	}
	t.timer = time.NewTimer(duration)
	t.loop.Go(t)
}
func (t *timeoutImpl) Serve() {
	select {
	case <-t.timer.C:
		t.Result(t)
	case <-t.Done():
	}
}
func (t *timeoutImpl) OnResult(closed bool) bool {
	if t.running {
		t.running = false
		t.Call()
	}
	return true
}
func (t *timeoutImpl) Stop() {
	if t.running {
		if t.timer != nil && t.timer.Stop() {
			t.Complete()
		}
		t.running = false
	}
}
func (r *Runtime) builtinGo_SetTimeout(call FunctionCall) Value {
	callable, ok := AssertFunction(call.Argument(0))
	if !ok {
		panic(r.NewTypeError(`"callback" argument must be a function`))
	}
	duration := time.Duration(call.Argument(1).ToInteger()) * time.Millisecond
	timeout := newTimeoutImpl(NewFuture(r.loop, callable, call.This),
		duration,
	)
	return r.ToValue(timeout)
}
func (r *Runtime) builtinGo_ClearTimeout(call FunctionCall) Value {
	if len(call.Arguments) > 0 {
		arg0 := call.Argument(0)
		export := arg0.Export()
		if timeout, ok := export.(*timeoutImpl); ok {
			timeout.Stop()
		}
	}
	return nil
}

type intervalImpl struct {
	*Future
	running bool
	ticker  *time.Ticker
	cancel  chan struct{}
	frame   chan bool
}

func newIntervalImpl(future *Future, duration time.Duration) (interval *intervalImpl) {
	interval = &intervalImpl{
		Future:  future,
		running: true,
	}
	interval.serve(duration)
	return
}

func (t *intervalImpl) serve(duration time.Duration) {
	t.Async(nil)
	if duration <= 0 {
		w := frameWorker{
			t:     t,
			frame: make(chan bool),
		}
		t.loop.Go(w)
		t.frame <- true
		return
	}
	t.cancel = make(chan struct{})
	t.ticker = time.NewTicker(duration)
	t.loop.Go(t)
}

type frameWorker struct {
	frame <-chan bool
	t     *intervalImpl
}

func (f frameWorker) Serve() {
	t := f.t
	for range f.frame {
		t.Result(t)
	}
}
func (t *intervalImpl) Serve() {
	for {
		select {
		case <-t.ticker.C:
			t.Result(t)
		case <-t.cancel:
			return
		}
	}
}
func (t *intervalImpl) OnResult(closed bool) bool {
	if t.running {
		t.Call()
		if closed {
			t.Stop()
		} else if t.frame != nil {
			t.frame <- true
		}
	}
	return false
}
func (t *intervalImpl) Stop() {
	if t.running {
		if t.ticker != nil {
			t.ticker.Stop()
		}
		if t.cancel != nil {
			close(t.cancel)
		}
		if t.frame != nil {
			close(t.frame)
			t.frame = nil
		}
		t.Complete()
		t.running = false
	}
}

func (r *Runtime) builtinGo_SetInterval(call FunctionCall) Value {
	callable, ok := AssertFunction(call.Argument(0))
	if !ok {
		panic(r.NewTypeError(`"callback" argument must be a function`))
	}
	duration := time.Duration(call.Argument(1).ToInteger()) * time.Millisecond
	interval := newIntervalImpl(NewFuture(r.loop, callable, call.This),
		duration,
	)
	return r.ToValue(interval)
}
func (r *Runtime) builtinGo_ClearInterval(call FunctionCall) Value {
	if len(call.Arguments) > 0 {
		arg0 := call.Argument(0)
		export := arg0.Export()
		if interval, ok := export.(*intervalImpl); ok {
			interval.Stop()
		}
	}
	return nil
}
func (r *Runtime) pp_expand_init_timer() {
	r.addToGlobal(`setTimeout`, r.newNativeFunc(r.builtinGo_SetTimeout, nil, "setTimeout", nil, 2))
	r.addToGlobal(`clearTimeout`, r.newNativeFunc(r.builtinGo_ClearTimeout, nil, "clearTimeout", nil, 1))
	r.addToGlobal(`setInterval`, r.newNativeFunc(r.builtinGo_SetInterval, nil, "setInterval", nil, 2))
	r.addToGlobal(`clearInterval`, r.newNativeFunc(r.builtinGo_ClearInterval, nil, "clearInterval", nil, 1))
}
