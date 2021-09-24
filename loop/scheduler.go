package loop

import (
	"sync/atomic"
)

type Worker interface {
	Serve()
}
type funcWorker struct {
	f func()
}

func (f funcWorker) Serve() {
	if f.f != nil {
		f.f()
	}
}
func NewWorker(f func()) Worker {
	return funcWorker{
		f: f,
	}
}

type Scheduler interface {
	Go(w Worker)
	Stop()
}

var defaultScheduler Scheduler = goScheduler{}

type goScheduler struct{}

func (goScheduler) String() string {
	return `DefaultScheduler`
}
func (goScheduler) Go(w Worker) {
	go w.Serve()
}
func (g goScheduler) Stop() {}

type asyncScheduler struct {
	stoped int32
	stop   chan struct{}
	ch     chan Worker
}

func NewScheduler(idle int) Scheduler {
	if idle < 1 {
		return defaultScheduler
	}
	scheduler := &asyncScheduler{
		stop: make(chan struct{}),
		ch:   make(chan Worker),
	}
	for i := 0; i < idle; i++ {
		go scheduler.serve()
	}
	return scheduler
}
func (a *asyncScheduler) serve() {
	for {
		select {
		case <-a.stop:
			return
		case w := <-a.ch:
			w.Serve()
		}
	}
}
func (a *asyncScheduler) Go(w Worker) {
	select {
	case <-a.stop:
	case a.ch <- w:
	default:
		go w.Serve()
	}
}
func (a *asyncScheduler) Stop() {
	if a.stoped == 0 && atomic.SwapInt32(&a.stoped, 1) == 0 {
		close(a.stop)
	}
}
