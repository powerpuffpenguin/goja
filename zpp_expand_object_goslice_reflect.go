package goja

import (
	"reflect"
)

func (o *objectGoSliceReflect) _getIdx_zpp(v reflect.Value) Value {
	result := v.Interface()
	if v.Kind() == reflect.Slice {
		result = o.val.runtime.wrapReflectFunc_ppResultWrapArray(v.Type().Elem().Kind(), result)
	}
	return o.val.runtime.ToValue(result)
}
