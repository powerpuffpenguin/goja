package goja

import (
	"errors"
	"fmt"
	"reflect"
)

type goNative interface {
	Native() interface{}
}

var errNotReadChan = errors.New(`not a <-chan`)
var errNotWriteChan = errors.New(`not a chan<-`)

func (r *Runtime) pp_expand_init_chan() {
	r.addToGlobal(`goRecv`, r.pp_expand_wrapReflectFunc_pp(r.builtinGo_goRecv))
	r.addToGlobal(`goTryRecv`, r.pp_expand_wrapReflectFunc_pp(r.builtinGo_goTryRecv))
	r.addToGlobal(`goSend`, r.pp_expand_wrapReflectFunc_pp(r.builtinGo_goSend))
	r.addToGlobal(`goTrySend`, r.pp_expand_wrapReflectFunc_pp(r.builtinGo_goTrySend))
	r.addToGlobal(`goClose`, r.pp_expand_wrapReflectFunc_pp(r.builtinGo_goClose))

	self := r.GlobalObject().self
	self.setOwnStr("SelectSend", &valueProperty{
		configurable: true,
		getterFunc:   r.newNativeFunc(r.pp_expand_get_SelectSend, nil, "SelectSend", nil, 0),
		accessor:     true,
	}, false)
	self.setOwnStr("SelectRecv", &valueProperty{
		configurable: true,
		getterFunc:   r.newNativeFunc(r.pp_expand_get_SelectRecv, nil, "SelectRecv", nil, 0),
		accessor:     true,
	}, false)
	self.setOwnStr("SelectDefault", &valueProperty{
		configurable: true,
		getterFunc:   r.newNativeFunc(r.pp_expand_get_SelectDefault, nil, "SelectDefault", nil, 0),
		accessor:     true,
	}, false)
	self.setOwnStr("DefaultCase", &valueProperty{
		configurable: true,
		getterFunc:   r.newNativeFunc(r.pp_expand_get_defaultCase, nil, "DefaultCase", nil, 0),
		accessor:     true,
	}, false)
	r.addToGlobal(`NewSendCase`, r.pp_expand_wrapReflectFunc_pp(r.builtinGo_NewSendCase))
	r.addToGlobal(`NewRecvCase`, r.pp_expand_wrapReflectFunc_pp(r.builtinGo_NewRecvCase))
	r.addToGlobal(`goSelect`, r.pp_expand_wrapReflectFunc_pp(r.builtinGo_goSelect))
}
func (r *Runtime) builtinGo_goSelect(cases ...reflect.SelectCase) (chosen int, result interface{}, ok bool, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()
	chosen, value, ok := reflect.Select(cases)
	if ok {
		result = r.wrapReflectFunc_ppResultWrap(value)
	}
	return
}
func (r *Runtime) builtinGo_NewSendCase(ch, x interface{}) (reflect.SelectCase, error) {
	value := reflect.ValueOf(ch)
	if value.Kind() != reflect.Chan {
		return reflect.SelectCase{}, errNotWriteChan
	}
	if native, ok := x.(goNative); ok {
		x = native.Native()
	}
	return reflect.SelectCase{
		Dir:  reflect.SelectSend,
		Chan: value,
		Send: reflect.ValueOf(x),
	}, nil
}
func (r *Runtime) builtinGo_NewRecvCase(ch interface{}) (reflect.SelectCase, error) {
	value := reflect.ValueOf(ch)
	if value.Kind() != reflect.Chan {
		return reflect.SelectCase{}, errNotReadChan
	}
	return reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: value,
	}, nil
}

var defaultCase = reflect.SelectCase{
	Dir: reflect.SelectDefault,
}

func (r *Runtime) pp_expand_get_defaultCase(call FunctionCall) Value {
	return r.ToValue(defaultCase)
}
func (r *Runtime) pp_expand_get_SelectDefault(call FunctionCall) Value {
	return r.ToValue(reflect.SelectDefault)
}
func (r *Runtime) pp_expand_get_SelectRecv(call FunctionCall) Value {
	return r.ToValue(reflect.SelectRecv)
}
func (r *Runtime) pp_expand_get_SelectSend(call FunctionCall) Value {
	return r.ToValue(reflect.SelectSend)
}
func (r *Runtime) builtinGo_goRecv(ch interface{}) (result interface{}, ok bool, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()
	value := reflect.ValueOf(ch)
	if value.Kind() == reflect.Chan {
		value, ok = value.Recv()
		if ok {
			result = r.wrapReflectFunc_ppResultWrap(value)
		}
	} else {
		e = errNotReadChan
	}
	return
}
func (r *Runtime) builtinGo_goTryRecv(ch interface{}) (result interface{}, ok bool, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()
	value := reflect.ValueOf(ch)
	if value.Kind() == reflect.Chan {
		value, ok = value.TryRecv()
		if ok {
			result = r.wrapReflectFunc_ppResultWrap(value)
		}
	} else {
		e = errNotReadChan
	}
	return
}
func (r *Runtime) builtinGo_goSend(ch interface{}, x interface{}) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()
	value := reflect.ValueOf(ch)
	if value.Kind() == reflect.Chan {
		if native, ok := x.(goNative); ok {
			x = native.Native()
		}
		value.Send(reflect.ValueOf(x))
	} else {
		e = errNotWriteChan
	}
	return
}
func (r *Runtime) builtinGo_goTrySend(ch interface{}, x interface{}) (ok bool, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()
	value := reflect.ValueOf(ch)
	if value.Kind() == reflect.Chan {
		if native, ok := x.(goNative); ok {
			x = native.Native()
		}
		ok = value.TrySend(reflect.ValueOf(x))
	} else {
		e = errNotWriteChan
	}
	return
}
func (r *Runtime) builtinGo_goClose(ch interface{}) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()
	value := reflect.ValueOf(ch)
	if value.Kind() == reflect.Chan {
		value.Close()
	} else {
		e = errNotWriteChan
	}
	return
}
