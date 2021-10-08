package loop

import (
	"context"
	"errors"
	"sync/atomic"
)

var ErrLoopClosed = errors.New(`loop already closed`)

type AsyncResult interface {
	OnResult(closed bool) (completed bool)
}
type Loop struct {
	wait      uint64
	ch        chan AsyncResult
	scheduler Scheduler
	closed    int32
	cancel    context.CancelFunc
	ctx       context.Context
}

func New(scheduler Scheduler) *Loop {
	if scheduler == nil {
		scheduler = defaultScheduler
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Loop{
		ch:        make(chan AsyncResult, 128),
		cancel:    cancel,
		ctx:       ctx,
		scheduler: scheduler,
	}
}
func (l *Loop) Close() (e error) {
	if l.closed != 0 && atomic.SwapInt32(&l.closed, 1) == 0 {
		l.cancel()
		l.scheduler.Stop()
	} else {
		e = ErrLoopClosed
	}
	return
}

func (l *Loop) Serve() (e error) {
	for l.wait != 0 {
		select {
		case <-l.ctx.Done():
			e = ErrLoopClosed
			l.exitServe()
			return
		case result := <-l.ch:
			if result.OnResult(false) {
				l.wait--
			}
		}
	}
	return
}
func (l *Loop) exitServe() {
	for l.wait != 0 {
		result := <-l.ch
		if result.OnResult(true) {
			l.wait--
		}
	}
}

// Create an asynchronous event, can only be called in the goroutine running by js
func (l *Loop) Async(w Worker) {
	l.wait++
	if w != nil {
		l.scheduler.Go(w)
	}
}

// go w.Serve()
func (l *Loop) Go(w Worker) {
	l.scheduler.Go(w)
}

// Complete an asynchronous event, can only be called in the goroutine running by js
func (l *Loop) Complete() {
	l.wait--
}

// Send the result of an asynchronous event to js
func (l *Loop) Result(result AsyncResult) {
	l.ch <- result
}
func (l *Loop) TryResult(result AsyncResult) (ok bool) {
	select {
	case l.ch <- result:
		ok = true
	default:
	}
	return
}
func (l *Loop) Done() <-chan struct{} {
	return l.ctx.Done()
}
func (l *Loop) Context() context.Context {
	return l.ctx
}
func (l *Loop) SetScheduler(scheduler Scheduler) Scheduler {
	if scheduler == nil {
		scheduler = defaultScheduler
	}

	old := l.scheduler
	if old == scheduler {
		return nil
	}
	l.scheduler = scheduler
	return old
}
func (l *Loop) GetScheduler() Scheduler {
	return l.scheduler
}
