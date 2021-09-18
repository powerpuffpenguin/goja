package goja

import (
	go_context "context"

	"github.com/dop251/goja/zpp_expand/loop"
)

type Future struct {
	loop     *loop.Loop
	callable Callable
	self     Value
}

func NewFuture(loop *loop.Loop, callable Callable, self Value) *Future {
	return &Future{
		loop:     loop,
		callable: callable,
		self:     self,
	}
}
func (a *Future) Async(w loop.Worker) {
	a.loop.Async(w)
}
func (a *Future) Go(w loop.Worker) {
	a.loop.Go(w)
}
func (a *Future) Complete() {
	a.loop.Complete()
}
func (a *Future) Result(result loop.AsyncResult) {
	a.loop.Result(result)
}
func (a *Future) TryResult(result loop.AsyncResult) (ok bool) {
	return a.loop.TryResult(result)
}
func (a *Future) Done() <-chan struct{} {
	return a.loop.Done()
}
func (a *Future) Context() go_context.Context {
	return a.loop.Context()
}
func (a *Future) Call(args ...Value) {
	if a.callable != nil {
		a.callable(a.self, args...)
	}
}

func (r *Runtime) Serve() error {
	return r.loop.Serve()
}
func (r *Runtime) SetScheduler(scheduler loop.Scheduler) (old loop.Scheduler) {
	return r.loop.SetScheduler(scheduler)
}
func (r *Runtime) Loop() *loop.Loop {
	return r.loop
}

func (r *Runtime) RunScriptAndServe(filename, source string) (val Value, e error) {
	val, e = r.RunScript(filename, source)
	if e != nil {
		return
	}
	e = r.loop.Serve()
	return
}
func (r *Runtime) RunStringAndServe(source string) (val Value, e error) {
	val, e = r.RunString(source)
	if e != nil {
		return
	}
	e = r.loop.Serve()
	return
}

func (r *Runtime) pp_expand_init() {
	r.pp_expand_init_number()
	r.loop = loop.NewLoop()
}
