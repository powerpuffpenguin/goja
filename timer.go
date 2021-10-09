package goja

import (
	"time"

	"github.com/powerpuffpenguin/goja/loop"
)

type timeoutImpl struct {
	*Future
	running bool
	timer   *time.Timer
}

func newTimeoutImpl(future *Future, duration time.Duration) (*timeoutImpl, error) {
	timeout := &timeoutImpl{
		Future:  future,
		running: true,
	}
	e := timeout.serve(duration)
	if e != nil {
		return nil, e
	}
	return timeout, nil
}
func (t *timeoutImpl) serve(duration time.Duration) (e error) {
	e = t.Async(nil)
	if e != nil {
		return
	}
	if duration <= 0 {
		t.Go(loop.NewWorker(func() {
			t.Result(t)
		}))
		return nil
	}
	t.timer = time.NewTimer(duration)
	return t.Go(t)
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
		if !closed {
			t.Call()
		}
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
	timeout, e := newTimeoutImpl(NewFuture(r.loop, callable, call.This),
		duration,
	)
	if e != nil {
		panic(r.NewGoError(e))
	}
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

func newIntervalImpl(future *Future, duration time.Duration) (*intervalImpl, error) {
	interval := &intervalImpl{
		Future:  future,
		running: true,
	}
	e := interval.serve(duration)
	if e != nil {
		return nil, e
	}
	return interval, nil
}

func (t *intervalImpl) serve(duration time.Duration) (e error) {
	e = t.Async(nil)
	if e != nil {
		return
	}
	if duration <= 0 {
		w := frameWorker{
			t:     t,
			frame: make(chan bool),
		}
		e = t.loop.Go(w)
		if e != nil {
			t.Stop()
			return
		}
		t.frame <- true
		return
	}
	t.cancel = make(chan struct{})
	t.ticker = time.NewTicker(duration)
	e = t.loop.Go(t)
	if e != nil {
		t.Stop()
		return
	}
	return
}

type frameWorker struct {
	frame <-chan bool
	t     *intervalImpl
}

func (f frameWorker) Serve() {
	t := f.t
	for range f.frame {
		if !t.Result(t) {
			break
		}
	}
}
func (t *intervalImpl) Serve() {
	for {
		select {
		case <-t.ticker.C:
			if !t.Result(t) {
				t.ticker.Stop()
				return
			}
		case <-t.cancel:
			return
		}
	}
}
func (t *intervalImpl) OnResult(closed bool) bool {
	if t.running {
		if closed {
			t.Stop()
		} else {
			t.Call()
			if t.frame != nil {
				t.frame <- true
			}
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
		err := t.Complete()
		if err != nil {
			return
		}
		t.running = false
	}
}

func (r *Runtime) builtinGo_SetInterval(call FunctionCall) Value {
	callable, ok := AssertFunction(call.Argument(0))
	if !ok {
		panic(r.NewTypeError(`"callback" argument must be a function`))
	}
	duration := time.Duration(call.Argument(1).ToInteger()) * time.Millisecond
	interval, e := newIntervalImpl(NewFuture(r.loop, callable, call.This),
		duration,
	)
	if e != nil {
		panic(r.NewGoError(e))
	}
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
func (r *Runtime) initTimer() {
	r.addToGlobal(`setTimeout`, r.newNativeFunc(r.builtinGo_SetTimeout, nil, "setTimeout", nil, 2))
	r.addToGlobal(`clearTimeout`, r.newNativeFunc(r.builtinGo_ClearTimeout, nil, "clearTimeout", nil, 1))
	r.addToGlobal(`setInterval`, r.newNativeFunc(r.builtinGo_SetInterval, nil, "setInterval", nil, 2))
	r.addToGlobal(`clearInterval`, r.newNativeFunc(r.builtinGo_ClearInterval, nil, "clearInterval", nil, 1))
}
