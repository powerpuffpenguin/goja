package goja

import (
	"reflect"
)

func (o *objectGoReflect) _zpp_getFieldMethodImpl(jsName string) interface{} {
	if info, exists := o.valueTypeInfo.Fields[jsName]; exists {
		v := o.value.FieldByIndex(info.Index)
		return o.proto().runtime.wrapReflectFunc_ppResultWrap(v)
	}
	return nil
}
func (o *objectGoReflect) _zpp_getFieldMethod(jsName string) reflect.Value {
	if jsName == `get` {
		return reflect.ValueOf(o._zpp_getFieldMethodImpl)
	}
	return reflect.Value{}
}
