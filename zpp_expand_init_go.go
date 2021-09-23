package goja

import (
	"errors"
	"reflect"
)

func (r *Runtime) pp_expand_init_go() {
	r.addToGlobal(`isGoSlice`, r.newNativeFunc(r.builtinGo_isGoSlice, nil, "isGoSlice", nil, 1))
	r.addToGlobal(`isGoMap`, r.newNativeFunc(r.builtinGo_isGoMap, nil, "isGoMap", nil, 1))
	r.addToGlobal(`goLen`, r.newNativeFunc(r.builtinGo_goLen, nil, "goLen", nil, 1))
	r.addToGlobal(`goHasKey`, r.newNativeFunc(r.builtinGo_goHasKey, nil, "goHasKey", nil, 2))
}
func (r *Runtime) builtinGo_isGoSlice(call FunctionCall) Value {
	export := call.Argument(0).Export()
	kind := reflect.ValueOf(export).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		return valueTrue
	}
	return valueFalse
}
func (r *Runtime) builtinGo_isGoMap(call FunctionCall) Value {
	export := call.Argument(0).Export()
	kind := reflect.ValueOf(export).Kind()
	if kind == reflect.Map {
		return valueTrue
	}
	return valueFalse
}
func (r *Runtime) builtinGo_goLen(call FunctionCall) Value {
	export := call.Argument(0).Export()
	value := reflect.ValueOf(export)
	kind := value.Kind()
	var result int
	if kind == reflect.Slice ||
		kind == reflect.Array ||
		kind == reflect.String ||
		kind == reflect.Map ||
		kind == reflect.Chan {
		result = value.Len()
	} else {
		panic(r.NewGoError(errors.New(`type ` + kind.String() + ` not supported len`)))
	}
	return r.ToValue(NewInt(result))
}
func (r *Runtime) builtinGo_goHasKey(call FunctionCall) Value {
	export := call.Argument(0).Export()
	m := reflect.ValueOf(export)
	kind := m.Kind()
	if kind == reflect.Map {
		export = call.Argument(1).Export()
		key := reflect.ValueOf(export)
		if key.IsValid() && m.MapIndex(key).IsValid() {
			return valueTrue
		}
	} else {
		panic(r.NewGoError(errors.New(`type ` + kind.String() + ` not go a map`)))
	}
	return valueFalse
}
