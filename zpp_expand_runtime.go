package goja

import (
	"reflect"

	"github.com/powerpuffpenguin/goja/loop"
)

func (r *Runtime) wrapReflectFunc_pp(value reflect.Value) func(FunctionCall) Value {
	return func(call FunctionCall) Value {
		typ := value.Type()
		nargs := typ.NumIn()
		var in []reflect.Value
		args := call.Arguments
		var scheduler loop.Scheduler
		if len(args) != 0 {
			offset := len(args) - 1
			arg := args[offset]
			if tmp, ok := arg.Export().(loop.Scheduler); ok {
				call.Arguments = args[:offset]
				scheduler = tmp
			}
		}

		if l := len(call.Arguments); l < nargs {
			// fill missing arguments with zero values
			n := nargs
			if typ.IsVariadic() {
				n--
			}
			in = make([]reflect.Value, n)
			for i := l; i < n; i++ {
				in[i] = reflect.Zero(typ.In(i))
			}
		} else {
			if l > nargs && !typ.IsVariadic() {
				l = nargs
			}
			in = make([]reflect.Value, l)
		}

		callSlice := false
		for i, a := range call.Arguments {
			var t reflect.Type

			n := i
			if n >= nargs-1 && typ.IsVariadic() {
				if n > nargs-1 {
					n = nargs - 1
				}

				t = typ.In(n).Elem()
			} else if n > nargs-1 { // ignore extra arguments
				break
			} else {
				t = typ.In(n)
			}

			// if this is a variadic Go function, and the caller has supplied
			// exactly the number of JavaScript arguments required, and this
			// is the last JavaScript argument, try treating the it as the
			// actual set of variadic Go arguments. if that succeeds, break
			// out of the loop.
			if typ.IsVariadic() && len(call.Arguments) == nargs && i == nargs-1 {
				v := reflect.New(typ.In(n)).Elem()
				if err := r.toReflectValue(a, v, &objectExportCtx{}); err == nil {
					in[i] = v
					callSlice = true
					break
				}
			}
			v := reflect.New(t).Elem()
			err := r.toReflectValue(a, v, &objectExportCtx{})
			if err != nil {
				panic(r.newError(r.global.TypeError, "could not convert function call parameter %v to %v", a, t))
			}
			in[i] = v
		}

		if scheduler == nil {
			return r.wrapReflectFunc_ppCall(callSlice, value, in)
		}
		return r.wrapReflectFunc_ppCallAsync(scheduler, callSlice, value, in)
	}
}
func (r *Runtime) wrapReflectFunc_ppCallAsync(scheduler loop.Scheduler, callSlice bool, value reflect.Value, in []reflect.Value) Value {
	completer, e := NewCompleter(r)
	if e != nil {
		panic(r.NewGoError(e))
	}
	r.loop.Async(nil)
	scheduler.Go(newAsyncImpl(r, completer, callSlice, value, in))
	return completer.promise
}

type asyncImpl struct {
	runtime   *Runtime
	completer *Completer

	callSlice bool
	value     reflect.Value
	in        []reflect.Value

	result []reflect.Value
}

func newAsyncImpl(runtime *Runtime,
	completer *Completer,
	callSlice bool, value reflect.Value, in []reflect.Value,
) *asyncImpl {
	return &asyncImpl{
		runtime:   runtime,
		completer: completer,
		callSlice: callSlice,
		value:     value,
		in:        in,
	}
}
func (a *asyncImpl) Serve() {
	if a.callSlice {
		a.result = a.value.CallSlice(a.in)
	} else {
		a.result = a.value.Call(a.in)
	}
	a.runtime.loop.Result(a)
}
func (a *asyncImpl) OnResult(closed bool) (completed bool) {
	completed = true
	runtime := a.runtime
	result, e := runtime.wrapReflectFunc_ppResult(a.result)
	if e == nil {
		a.completer.Resolve(result)
	} else {
		a.completer.Reject(runtime.ToValue(e))
	}
	return
}
func (r *Runtime) wrapReflectFunc_ppCall(callSlice bool, value reflect.Value, in []reflect.Value) (result Value) {
	var out []reflect.Value
	if callSlice {
		out = value.CallSlice(in)
	} else {
		out = value.Call(in)
	}
	result, e := r.wrapReflectFunc_ppResult(out)
	if e != nil {
		panic(e)
	}
	return result
}
func (r *Runtime) wrapReflectFunc_ppResult(out []reflect.Value) (result Value, e interface{}) {
	if len(out) == 0 {
		result = _undefined
		return
	}

	if last := out[len(out)-1]; last.Type().Name() == "error" {
		if !last.IsNil() {
			err := last.Interface()
			if _, ok := err.(*Exception); ok {
				e = err
				return
			}
			e = r.NewGoError(last.Interface().(error))
			return
		}
		out = out[:len(out)-1]
	}

	switch len(out) {
	case 0:
		result = _undefined
	case 1:
		result = r.ToValue(r.wrapReflectFunc_ppResultWrap(out[0]))
	default:
		s := make([]interface{}, len(out))
		for i, v := range out {
			s[i] = r.wrapReflectFunc_ppResultWrap(v)
		}
		result = r.ToValue(s)
	}
	return
}
func (r *Runtime) wrapReflectFunc_ppResultWrap(v reflect.Value) interface{} {
	result := v.Interface()
	switch v.Kind() {
	case reflect.Int:
		if val, ok := result.(int); ok {
			return NewInt(val)
		}
	case reflect.Int8:
		if val, ok := result.(int8); ok {
			return NewInt8(val)
		}
	case reflect.Int16:
		if val, ok := result.(int16); ok {
			return NewInt16(val)
		}
	case reflect.Int32:
		if val, ok := result.(int32); ok {
			return NewInt32(val)
		}
	case reflect.Int64:
		if val, ok := result.(int64); ok {
			return NewInt64(val)
		}
	case reflect.Uint:
		if val, ok := result.(uint); ok {
			return NewUint(val)
		}
	case reflect.Uint8:
		if val, ok := result.(uint8); ok {
			return NewUint8(val)
		}
	case reflect.Uint16:
		if val, ok := result.(uint16); ok {
			return NewUint16(val)
		}
	case reflect.Uint32:
		if val, ok := result.(uint32); ok {
			return NewUint32(val)
		}
	case reflect.Uint64:
		if val, ok := result.(uint64); ok {
			return NewUint64(val)
		}
	case reflect.Float32:
		if val, ok := result.(float32); ok {
			return NewFloat32(val)
		}
	case reflect.Float64:
		if val, ok := result.(float64); ok {
			return NewFloat64(val)
		}

	case reflect.Slice:
		switch v.Type().Elem().Kind() {
		case reflect.Int:
			if val, ok := result.([]int); ok {
				return NewIntArray(val)
			}
		case reflect.Int8:
			if val, ok := result.([]int8); ok {
				return NewInt8Array(val)
			}
		case reflect.Int16:
			if val, ok := result.([]int16); ok {
				return NewInt16Array(val)
			}
		case reflect.Int32:
			if val, ok := result.([]int32); ok {
				return NewInt32Array(val)
			}
		case reflect.Int64:
			if val, ok := result.([]int64); ok {
				return NewInt64Array(val)
			}
		case reflect.Uint:
			if val, ok := result.([]uint); ok {
				return NewUintArray(val)
			}
		case reflect.Uint8:
			if val, ok := result.([]uint8); ok {
				return NewUint8Array(val)
			}
		case reflect.Uint16:
			if val, ok := result.([]uint16); ok {
				return NewUint16Array(val)
			}
		case reflect.Uint32:
			if val, ok := result.([]uint32); ok {
				return NewUint32Array(val)
			}
		case reflect.Uint64:
			if val, ok := result.([]uint64); ok {
				return NewUint64Array(val)
			}
		case reflect.Float32:
			if val, ok := result.([]float32); ok {
				return NewFloat32Array(val)
			}
		case reflect.Float64:
			if val, ok := result.([]float64); ok {
				return NewFloat64Array(val)
			}
		}
	}
	return result
}
