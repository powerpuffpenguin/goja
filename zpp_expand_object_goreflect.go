package goja

import "reflect"

func (o *objectGoReflect) _toNumber() Value {
	switch o.value.Kind() {
	case reflect.Int:
		if v, ok := o.value.Interface().(int); ok {
			return intToValue(int64(v))
		}
		return valueInt(o.value.Int())
	case reflect.Int8:
		if v, ok := o.value.Interface().(int8); ok {
			return intToValue(int64(v))
		}
		return valueInt(o.value.Int())
	case reflect.Int16:
		if v, ok := o.value.Interface().(int16); ok {
			return intToValue(int64(v))
		}
		return valueInt(o.value.Int())
	case reflect.Int32:
		if v, ok := o.value.Interface().(int32); ok {
			return intToValue(int64(v))
		}
		return valueInt(o.value.Int())
	case reflect.Int64:
		if v, ok := o.value.Interface().(int64); ok {
			return intToValue(v)
		}
		return valueInt(o.value.Int())
	case reflect.Uint:
		if v, ok := o.value.Interface().(uint); ok {
			return intToValue(int64(v))
		}
		return valueInt(o.value.Uint())
	case reflect.Uint8:
		if v, ok := o.value.Interface().(uint8); ok {
			return intToValue(int64(v))
		}
		return valueInt(o.value.Uint())
	case reflect.Uint16:
		if v, ok := o.value.Interface().(uint16); ok {
			return intToValue(int64(v))
		}
		return valueInt(o.value.Uint())
	case reflect.Uint32:
		if v, ok := o.value.Interface().(uint32); ok {
			return intToValue(int64(v))
		}
		return valueInt(o.value.Uint())
	case reflect.Uint64:
		if v, ok := o.value.Interface().(uint64); ok {
			return intToValue(int64(v))
		}
		return valueInt(o.value.Uint())
	case reflect.Bool:
		if o.value.Bool() {
			return intToValue(1)
		} else {
			return intToValue(0)
		}
	case reflect.Float32:
		if v, ok := o.value.Interface().(float32); ok {
			return floatToValue(float64(v))
		}
		return valueFloat(o.value.Float())
	case reflect.Float64:
		if v, ok := o.value.Interface().(float64); ok {
			return floatToValue(v)
		}
		return valueFloat(o.value.Float())
	}
	return nil
}
