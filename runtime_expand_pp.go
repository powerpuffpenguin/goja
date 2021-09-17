package goja

import (
	"fmt"
	"reflect"
)

func (r *Runtime) wrapReflectFunc_pp(value reflect.Value) func(FunctionCall) Value {
	return func(call FunctionCall) Value {
		typ := value.Type()
		nargs := typ.NumIn()
		var in []reflect.Value
		args := call.Arguments
		async := false
		if len(args) != 0 {
			offset := len(args) - 1
			arg := args[offset]
			if str, ok := arg.Export().(string); ok && str == "async" {
				call.Arguments = args[:offset]
				async = true
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

		if async {
			return r.wrapReflectFunc_ppCallAsync(callSlice, value, in)
		}
		return r.wrapReflectFunc_ppCall(callSlice, value, in)
	}
}
func (r *Runtime) wrapReflectFunc_ppCallAsync(callSlice bool, value reflect.Value, in []reflect.Value) (result Value) {
	go func() {
		fmt.Println("async call")
		var out []reflect.Value
		if callSlice {
			out = value.CallSlice(in)
		} else {
			out = value.Call(in)
		}
		result, e := r.wrapReflectFunc_ppResult(out)
		if e == nil {
			fmt.Println("resolve", result)
		} else {
			fmt.Println("reject", e)
		}
	}()
	// return promise
	return _undefined
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
		result = r.ToValue(out[0].Interface())
	default:
		s := make([]interface{}, len(out))
		for i, v := range out {
			s[i] = v.Interface()
		}
		result = r.ToValue(s)
	}
	return
}
