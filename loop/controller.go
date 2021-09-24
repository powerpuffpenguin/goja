package loop

import "sync/atomic"

type AsyncController struct {
	loop  *Loop
	state int32
}

func NewAsyncController(loop *Loop) *AsyncController {
	return &AsyncController{
		loop: loop,
	}
}
func (a *AsyncController) Async() bool {
	if a.state == 0 && atomic.CompareAndSwapInt32(&a.state, 0, 1) {
		a.loop.Async(nil)
		return true
	}
	return false
}
func (a *AsyncController) Complete() bool {
	if a.state == 1 && atomic.CompareAndSwapInt32(&a.state, 1, 2) {
		a.loop.Complete()
		return true
	}
	return false
}
func (a *AsyncController) Call(f func()) {
	a.loop.Result(callbackResult(f))
}

type callbackResult func()

func (f callbackResult) OnResult(closed bool) (completed bool) {
	f()
	return false
}
