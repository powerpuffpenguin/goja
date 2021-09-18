package goja

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Int struct {
	Value int
}

func NewInt(val int) Int {
	return Int{val}
}
func (v Int) String() string {
	return fmt.Sprint(v.Value)
}
func (v Int) Native() int64 {
	return int64(v.Value)
}
func (v Int) ToNumber() Value {
	return intToValue(int64(v.Value))
}
func (v Int) ToInt() int {
	return v.Value
}
func (v Int) ToInt64() int64 {
	return int64(v.Value)
}
func (v Int) ToInt32() int32 {
	return int32(v.Value)
}
func (v Int) ToInt16() int16 {
	return int16(v.Value)
}
func (v Int) ToInt8() int8 {
	return int8(v.Value)
}
func (v Int) ToUint() uint {
	return uint(v.Value)
}
func (v Int) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Int) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Int) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Int) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Int) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Int) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Int) ABS() int {
	if v.Value < 0 {
		return -v.Value
	}
	return v.Value
}
func (v Int) Negate() int {
	return -v.Value
}
func (v Int) Add(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Int) Sub(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Int) Mul(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Int) Div(vals ...int) (int, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Int) Mod(vals ...int) (int, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Int) And(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Int) AndNot(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Int) Not() int {
	return ^v.Value
}
func (v Int) Or(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Int) Xor(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Int) ShiftLeft(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Int) ShiftRight(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Int) Compare(val int) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Int) Max(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Int) Min(vals ...int) int {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewInt(call FunctionCall) Value {
	var (
		result int
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseInt(s, base, 64)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = int(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewInt(result))
}

type Int64 struct {
	Value int64
}

func NewInt64(val int64) Int64 {
	return Int64{val}
}
func (v Int64) String() string {
	return fmt.Sprint(v.Value)
}
func (v Int64) Native() int64 {
	return v.Value
}
func (v Int64) ToNumber() Value {
	return intToValue(v.Value)
}
func (v Int64) ToInt() int {
	return int(v.Value)
}
func (v Int64) ToInt64() int64 {
	return v.Value
}
func (v Int64) ToInt32() int32 {
	return int32(v.Value)
}
func (v Int64) ToInt16() int16 {
	return int16(v.Value)
}
func (v Int64) ToInt8() int8 {
	return int8(v.Value)
}
func (v Int64) ToUint() uint {
	return uint(v.Value)
}
func (v Int64) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Int64) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Int64) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Int64) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Int64) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Int64) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Int64) ABS() int64 {
	if v.Value < 0 {
		return -v.Value
	}
	return v.Value
}
func (v Int64) Negate() int64 {
	return -v.Value
}
func (v Int64) Add(vals ...int64) int64 {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Int64) Sub(vals ...int64) int64 {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Int64) Mul(vals ...int64) int64 {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Int64) Div(vals ...int64) (int64, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Int64) Mod(vals ...int64) (int64, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Int64) And(vals ...int64) int64 {
	result := v.Value
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Int64) AndNot(vals ...int64) int64 {
	result := v.Value
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Int64) Not() int64 {
	return ^v.Value
}
func (v Int64) Or(vals ...int64) int64 {
	result := v.Value
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Int64) Xor(vals ...int64) int64 {
	result := v.Value
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Int64) ShiftLeft(vals ...int) int64 {
	result := v.Value
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Int64) ShiftRight(vals ...int) int64 {
	result := v.Value
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Int64) Compare(val int64) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Int64) Max(vals ...int64) int64 {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Int64) Min(vals ...int64) int64 {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewInt64(call FunctionCall) Value {
	var (
		result int64
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseInt(s, base, 64)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = val
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewInt64(result))
}

type Int32 struct {
	Value int32
}

func NewInt32(val int32) Int32 {
	return Int32{val}
}
func (v Int32) String() string {
	return fmt.Sprint(v.Value)
}
func (v Int32) Native() int64 {
	return int64(v.Value)
}
func (v Int32) ToNumber() Value {
	return intToValue(int64(v.Value))
}
func (v Int32) ToInt() int {
	return int(v.Value)
}
func (v Int32) ToInt64() int64 {
	return int64(v.Value)
}
func (v Int32) ToInt32() int32 {
	return v.Value
}
func (v Int32) ToInt16() int16 {
	return int16(v.Value)
}
func (v Int32) ToInt8() int8 {
	return int8(v.Value)
}
func (v Int32) ToUint() uint {
	return uint(v.Value)
}
func (v Int32) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Int32) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Int32) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Int32) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Int32) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Int32) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Int32) ABS() int32 {
	if v.Value < 0 {
		return -v.Value
	}
	return v.Value
}
func (v Int32) Negate() int32 {
	return -v.Value
}
func (v Int32) Add(vals ...int32) int32 {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Int32) Sub(vals ...int32) int32 {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Int32) Mul(vals ...int32) int32 {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Int32) Div(vals ...int32) (int32, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Int32) Mod(vals ...int32) (int32, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Int32) And(vals ...int32) int32 {
	result := v.Value
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Int32) AndNot(vals ...int32) int32 {
	result := v.Value
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Int32) Not() int32 {
	return ^v.Value
}
func (v Int32) Or(vals ...int32) int32 {
	result := v.Value
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Int32) Xor(vals ...int32) int32 {
	result := v.Value
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Int32) ShiftLeft(vals ...int) int32 {
	result := v.Value
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Int32) ShiftRight(vals ...int) int32 {
	result := v.Value
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Int32) Compare(val int32) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Int32) Max(vals ...int32) int32 {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Int32) Min(vals ...int32) int32 {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewInt32(call FunctionCall) Value {
	var (
		result int32
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseInt(s, base, 32)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = int32(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewInt32(result))
}

type Int16 struct {
	Value int16
}

func NewInt16(val int16) Int16 {
	return Int16{val}
}
func (v Int16) String() string {
	return fmt.Sprint(v.Value)
}
func (v Int16) Native() int64 {
	return int64(v.Value)
}
func (v Int16) ToNumber() Value {
	return intToValue(int64(v.Value))
}
func (v Int16) ToInt() int {
	return int(v.Value)
}
func (v Int16) ToInt64() int64 {
	return int64(v.Value)
}
func (v Int16) ToInt32() int32 {
	return int32(v.Value)
}
func (v Int16) ToInt16() int16 {
	return v.Value
}
func (v Int16) ToInt8() int8 {
	return int8(v.Value)
}
func (v Int16) ToUint() uint {
	return uint(v.Value)
}
func (v Int16) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Int16) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Int16) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Int16) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Int16) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Int16) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Int16) ABS() int16 {
	if v.Value < 0 {
		return -v.Value
	}
	return v.Value
}
func (v Int16) Negate() int16 {
	return -v.Value
}
func (v Int16) Add(vals ...int16) int16 {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Int16) Sub(vals ...int16) int16 {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Int16) Mul(vals ...int16) int16 {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Int16) Div(vals ...int16) (int16, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Int16) Mod(vals ...int16) (int16, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Int16) And(vals ...int16) int16 {
	result := v.Value
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Int16) AndNot(vals ...int16) int16 {
	result := v.Value
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Int16) Not() int16 {
	return ^v.Value
}
func (v Int16) Or(vals ...int16) int16 {
	result := v.Value
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Int16) Xor(vals ...int16) int16 {
	result := v.Value
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Int16) ShiftLeft(vals ...int) int16 {
	result := v.Value
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Int16) ShiftRight(vals ...int) int16 {
	result := v.Value
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Int16) Compare(val int16) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Int16) Max(vals ...int16) int16 {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Int16) Min(vals ...int16) int16 {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewInt16(call FunctionCall) Value {
	var (
		result int16
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseInt(s, base, 16)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = int16(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewInt16(result))
}

type Int8 struct {
	Value int8
}

func NewInt8(val int8) Int8 {
	return Int8{val}
}
func (v Int8) String() string {
	return fmt.Sprint(v.Value)
}
func (v Int8) Native() int64 {
	return int64(v.Value)
}
func (v Int8) ToNumber() Value {
	return intToValue(int64(v.Value))
}
func (v Int8) ToInt() int {
	return int(v.Value)
}
func (v Int8) ToInt64() int64 {
	return int64(v.Value)
}
func (v Int8) ToInt32() int32 {
	return int32(v.Value)
}
func (v Int8) ToInt16() int16 {
	return int16(v.Value)
}
func (v Int8) ToInt8() int8 {
	return v.Value
}
func (v Int8) ToUint() uint {
	return uint(v.Value)
}
func (v Int8) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Int8) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Int8) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Int8) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Int8) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Int8) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Int8) ABS() int8 {
	if v.Value < 0 {
		return -v.Value
	}
	return v.Value
}
func (v Int8) Negate() int8 {
	return -v.Value
}
func (v Int8) Add(vals ...int8) int8 {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Int8) Sub(vals ...int8) int8 {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Int8) Mul(vals ...int8) int8 {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Int8) Div(vals ...int8) (int8, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Int8) Mod(vals ...int8) (int8, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Int8) And(vals ...int8) int8 {
	result := v.Value
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Int8) AndNot(vals ...int8) int8 {
	result := v.Value
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Int8) Not() int8 {
	return ^v.Value
}
func (v Int8) Or(vals ...int8) int8 {
	result := v.Value
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Int8) Xor(vals ...int8) int8 {
	result := v.Value
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Int8) ShiftLeft(vals ...int) int8 {
	result := v.Value
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Int8) ShiftRight(vals ...int) int8 {
	result := v.Value
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Int8) Compare(val int8) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Int8) Max(vals ...int8) int8 {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Int8) Min(vals ...int8) int8 {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewInt8(call FunctionCall) Value {
	var (
		result int8
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseInt(s, base, 8)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = int8(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewInt8(result))
}

type Uint struct {
	Value uint
}

func NewUint(val uint) Uint {
	return Uint{val}
}
func (v Uint) String() string {
	return fmt.Sprint(v.Value)
}
func (v Uint) Native() int64 {
	return int64(v.Value)
}
func (v Uint) ToNumber() Value {
	return intToValue(int64(v.Value))
}
func (v Uint) ToInt() int {
	return int(v.Value)
}
func (v Uint) ToInt64() int64 {
	return int64(v.Value)
}
func (v Uint) ToInt32() int32 {
	return int32(v.Value)
}
func (v Uint) ToInt16() int16 {
	return int16(v.Value)
}
func (v Uint) ToInt8() int8 {
	return int8(v.Value)
}
func (v Uint) ToUint() uint {
	return v.Value
}
func (v Uint) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Uint) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Uint) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Uint) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Uint) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Uint) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Uint) Add(vals ...uint) uint {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Uint) Sub(vals ...uint) uint {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Uint) Mul(vals ...uint) uint {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Uint) Div(vals ...uint) (uint, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Uint) Mod(vals ...uint) (uint, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Uint) And(vals ...uint) uint {
	result := v.Value
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Uint) AndNot(vals ...uint) uint {
	result := v.Value
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Uint) Not() uint {
	return ^v.Value
}
func (v Uint) Or(vals ...uint) uint {
	result := v.Value
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Uint) Xor(vals ...uint) uint {
	result := v.Value
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Uint) ShiftLeft(vals ...int) uint {
	result := v.Value
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Uint) ShiftRight(vals ...int) uint {
	result := v.Value
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Uint) Compare(val uint) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Uint) Max(vals ...uint) uint {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Uint) Min(vals ...uint) uint {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewUint(call FunctionCall) Value {
	var (
		result uint
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseUint(s, base, 64)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = uint(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewUint(result))
}

type Uint64 struct {
	Value uint64
}

func NewUint64(val uint64) Uint64 {
	return Uint64{val}
}
func (v Uint64) String() string {
	return fmt.Sprint(v.Value)
}
func (v Uint64) Native() int64 {
	return int64(v.Value)
}
func (v Uint64) ToNumber() Value {
	return intToValue(int64(v.Value))
}
func (v Uint64) ToInt() int {
	return int(v.Value)
}
func (v Uint64) ToInt64() int64 {
	return int64(v.Value)
}
func (v Uint64) ToInt32() int32 {
	return int32(v.Value)
}
func (v Uint64) ToInt16() int16 {
	return int16(v.Value)
}
func (v Uint64) ToInt8() int8 {
	return int8(v.Value)
}
func (v Uint64) ToUint() uint {
	return uint(v.Value)
}
func (v Uint64) ToUint64() uint64 {
	return v.Value
}
func (v Uint64) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Uint64) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Uint64) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Uint64) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Uint64) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Uint64) Add(vals ...uint64) uint64 {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Uint64) Sub(vals ...uint64) uint64 {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Uint64) Mul(vals ...uint64) uint64 {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Uint64) Div(vals ...uint64) (uint64, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Uint64) Mod(vals ...uint64) (uint64, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Uint64) And(vals ...uint64) uint64 {
	result := v.Value
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Uint64) AndNot(vals ...uint64) uint64 {
	result := v.Value
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Uint64) Not() uint64 {
	return ^v.Value
}
func (v Uint64) Or(vals ...uint64) uint64 {
	result := v.Value
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Uint64) Xor(vals ...uint64) uint64 {
	result := v.Value
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Uint64) ShiftLeft(vals ...int) uint64 {
	result := v.Value
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Uint64) ShiftRight(vals ...int) uint64 {
	result := v.Value
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Uint64) Compare(val uint64) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Uint64) Max(vals ...uint64) uint64 {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Uint64) Min(vals ...uint64) uint64 {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewUint64(call FunctionCall) Value {
	var (
		result uint64
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseUint(s, base, 64)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = val
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewUint64(result))
}

type Uint32 struct {
	Value uint32
}

func NewUint32(val uint32) Uint32 {
	return Uint32{val}
}
func (v Uint32) String() string {
	return fmt.Sprint(v.Value)
}
func (v Uint32) Native() int64 {
	return int64(v.Value)
}
func (v Uint32) ToNumber() Value {
	return intToValue(int64(v.Value))
}
func (v Uint32) ToInt() int {
	return int(v.Value)
}
func (v Uint32) ToInt64() int64 {
	return int64(v.Value)
}
func (v Uint32) ToInt32() int32 {
	return int32(v.Value)
}
func (v Uint32) ToInt16() int16 {
	return int16(v.Value)
}
func (v Uint32) ToInt8() int8 {
	return int8(v.Value)
}
func (v Uint32) ToUint() uint {
	return uint(v.Value)
}
func (v Uint32) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Uint32) ToUint32() uint32 {
	return v.Value
}
func (v Uint32) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Uint32) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Uint32) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Uint32) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Uint32) Add(vals ...uint32) uint32 {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Uint32) Sub(vals ...uint32) uint32 {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Uint32) Mul(vals ...uint32) uint32 {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Uint32) Div(vals ...uint32) (uint32, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Uint32) Mod(vals ...uint32) (uint32, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Uint32) And(vals ...uint32) uint32 {
	result := v.Value
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Uint32) AndNot(vals ...uint32) uint32 {
	result := v.Value
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Uint32) Not() uint32 {
	return ^v.Value
}
func (v Uint32) Or(vals ...uint32) uint32 {
	result := v.Value
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Uint32) Xor(vals ...uint32) uint32 {
	result := v.Value
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Uint32) ShiftLeft(vals ...int) uint32 {
	result := v.Value
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Uint32) ShiftRight(vals ...int) uint32 {
	result := v.Value
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Uint32) Compare(val uint32) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Uint32) Max(vals ...uint32) uint32 {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Uint32) Min(vals ...uint32) uint32 {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewUint32(call FunctionCall) Value {
	var (
		result uint32
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseUint(s, base, 32)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = uint32(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewUint32(result))
}

type Uint16 struct {
	Value uint16
}

func NewUint16(val uint16) Uint16 {
	return Uint16{val}
}
func (v Uint16) String() string {
	return fmt.Sprint(v.Value)
}
func (v Uint16) Native() int64 {
	return int64(v.Value)
}
func (v Uint16) ToNumber() Value {
	return intToValue(int64(v.Value))
}
func (v Uint16) ToInt() int {
	return int(v.Value)
}
func (v Uint16) ToInt64() int64 {
	return int64(v.Value)
}
func (v Uint16) ToInt32() int32 {
	return int32(v.Value)
}
func (v Uint16) ToInt16() int16 {
	return int16(v.Value)
}
func (v Uint16) ToInt8() int8 {
	return int8(v.Value)
}
func (v Uint16) ToUint() uint {
	return uint(v.Value)
}
func (v Uint16) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Uint16) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Uint16) ToUint16() uint16 {
	return v.Value
}
func (v Uint16) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Uint16) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Uint16) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Uint16) Add(vals ...uint16) uint16 {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Uint16) Sub(vals ...uint16) uint16 {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Uint16) Mul(vals ...uint16) uint16 {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Uint16) Div(vals ...uint16) (uint16, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Uint16) Mod(vals ...uint16) (uint16, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Uint16) And(vals ...uint16) uint16 {
	result := v.Value
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Uint16) AndNot(vals ...uint16) uint16 {
	result := v.Value
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Uint16) Not() uint16 {
	return ^v.Value
}
func (v Uint16) Or(vals ...uint16) uint16 {
	result := v.Value
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Uint16) Xor(vals ...uint16) uint16 {
	result := v.Value
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Uint16) ShiftLeft(vals ...int) uint16 {
	result := v.Value
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Uint16) ShiftRight(vals ...int) uint16 {
	result := v.Value
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Uint16) Compare(val uint16) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Uint16) Max(vals ...uint16) uint16 {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Uint16) Min(vals ...uint16) uint16 {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewUint16(call FunctionCall) Value {
	var (
		result uint16
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseUint(s, base, 16)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = uint16(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewUint16(result))
}

type Uint8 struct {
	Value uint8
}

func NewUint8(val uint8) Uint8 {
	return Uint8{val}
}
func (v Uint8) String() string {
	return fmt.Sprint(v.Value)
}
func (v Uint8) Native() int64 {
	return int64(v.Value)
}
func (v Uint8) ToNumber() Value {
	return intToValue(int64(v.Value))
}
func (v Uint8) ToInt() int {
	return int(v.Value)
}
func (v Uint8) ToInt64() int64 {
	return int64(v.Value)
}
func (v Uint8) ToInt32() int32 {
	return int32(v.Value)
}
func (v Uint8) ToInt16() int16 {
	return int16(v.Value)
}
func (v Uint8) ToInt8() int8 {
	return int8(v.Value)
}
func (v Uint8) ToUint() uint {
	return uint(v.Value)
}
func (v Uint8) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Uint8) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Uint8) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Uint8) ToUint8() uint8 {
	return v.Value
}
func (v Uint8) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Uint8) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Uint8) Add(vals ...uint8) uint8 {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Uint8) Sub(vals ...uint8) uint8 {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Uint8) Mul(vals ...uint8) uint8 {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Uint8) Div(vals ...uint8) (uint8, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Uint8) Mod(vals ...uint8) (uint8, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Uint8) And(vals ...uint8) uint8 {
	result := v.Value
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Uint8) AndNot(vals ...uint8) uint8 {
	result := v.Value
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Uint8) Not() uint8 {
	return ^v.Value
}
func (v Uint8) Or(vals ...uint8) uint8 {
	result := v.Value
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Uint8) Xor(vals ...uint8) uint8 {
	result := v.Value
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Uint8) ShiftLeft(vals ...int) uint8 {
	result := v.Value
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Uint8) ShiftRight(vals ...int) uint8 {
	result := v.Value
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Uint8) Compare(val uint8) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Uint8) Max(vals ...uint8) uint8 {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Uint8) Min(vals ...uint8) uint8 {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewUint8(call FunctionCall) Value {
	var (
		result uint8
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseUint(s, base, 8)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = uint8(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewUint8(result))
}

type Float64 struct {
	Value float64
}

func NewFloat64(val float64) Float64 {
	return Float64{val}
}
func (v Float64) String() string {
	return fmt.Sprint(v.Value)
}
func (v Float64) Native() float64 {
	return v.Value
}
func (v Float64) ToNumber() Value {
	return floatToValue(v.Value)
}
func (v Float64) ToInt() int {
	return int(v.Value)
}
func (v Float64) ToInt64() int64 {
	return int64(v.Value)
}
func (v Float64) ToInt32() int32 {
	return int32(v.Value)
}
func (v Float64) ToInt16() int16 {
	return int16(v.Value)
}
func (v Float64) ToInt8() int8 {
	return int8(v.Value)
}
func (v Float64) ToUint() uint {
	return uint(v.Value)
}
func (v Float64) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Float64) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Float64) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Float64) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Float64) ToFloat64() float64 {
	return v.Value
}
func (v Float64) ToFloat32() float32 {
	return float32(v.Value)
}
func (v Float64) ABS() float64 {
	if v.Value < 0 {
		return -v.Value
	}
	return v.Value
}
func (v Float64) Negate() float64 {
	return -v.Value
}
func (v Float64) Add(vals ...float64) float64 {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Float64) Sub(vals ...float64) float64 {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Float64) Mul(vals ...float64) float64 {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Float64) Div(vals ...float64) (float64, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Float64) Sqrt() float64 {
	result := math.Sqrt(v.Value)
	return result
}
func (v Float64) Compare(val float64) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Float64) Max(vals ...float64) float64 {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Float64) Min(vals ...float64) float64 {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewFloat64(call FunctionCall) Value {
	var (
		result float64
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			val, e := strconv.ParseFloat(s, 64)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = val
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewFloat64(result))
}

type Float32 struct {
	Value float32
}

func NewFloat32(val float32) Float32 {
	return Float32{val}
}
func (v Float32) String() string {
	return fmt.Sprint(v.Value)
}
func (v Float32) Native() float64 {
	return float64(v.Value)
}
func (v Float32) ToNumber() Value {
	return floatToValue(float64(v.Value))
}
func (v Float32) ToInt() int {
	return int(v.Value)
}
func (v Float32) ToInt64() int64 {
	return int64(v.Value)
}
func (v Float32) ToInt32() int32 {
	return int32(v.Value)
}
func (v Float32) ToInt16() int16 {
	return int16(v.Value)
}
func (v Float32) ToInt8() int8 {
	return int8(v.Value)
}
func (v Float32) ToUint() uint {
	return uint(v.Value)
}
func (v Float32) ToUint64() uint64 {
	return uint64(v.Value)
}
func (v Float32) ToUint32() uint32 {
	return uint32(v.Value)
}
func (v Float32) ToUint16() uint16 {
	return uint16(v.Value)
}
func (v Float32) ToUint8() uint8 {
	return uint8(v.Value)
}
func (v Float32) ToFloat64() float64 {
	return float64(v.Value)
}
func (v Float32) ToFloat32() float32 {
	return v.Value
}
func (v Float32) ABS() float32 {
	if v.Value < 0 {
		return -v.Value
	}
	return v.Value
}
func (v Float32) Negate() float32 {
	return -v.Value
}
func (v Float32) Add(vals ...float32) float32 {
	result := v.Value
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Float32) Sub(vals ...float32) float32 {
	result := v.Value
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Float32) Mul(vals ...float32) float32 {
	result := v.Value
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Float32) Div(vals ...float32) (float32, error) {
	result := v.Value
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Float32) Sqrt() float32 {
	result := math.Sqrt(float64(v.Value))
	return float32(result)
}
func (v Float32) Compare(val float32) Value {
	if v.Value == val {
		return intToValue(0)
	} else if v.Value < val {
		return intToValue(-1)
	}
	return intToValue(1)
}
func (v Float32) Max(vals ...float32) float32 {
	result := v.Value
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Float32) Min(vals ...float32) float32 {
	result := v.Value
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (r *Runtime) builtinGo_NewFloat32(call FunctionCall) Value {
	var (
		result float32
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			val, e := strconv.ParseFloat(s, 32)
			if e !=nil {
				panic(r.NewGoError(e))
			}
			result = float32(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(NewFloat32(result))
}

type IntArray struct {
	Value []int
}

type Int64Array struct {
	Value []int64
}

type Int32Array struct {
	Value []int32
}

type Int16Array struct {
	Value []int16
}

type Int8Array struct {
	Value []int8
}

type UintArray struct {
	Value []uint
}

type Uint64Array struct {
	Value []uint64
}

type Uint32Array struct {
	Value []uint32
}

type Uint16Array struct {
	Value []uint16
}

type Uint8Array struct {
	Value []uint8
}

type Float64Array struct {
	Value []float64
}

type Float32Array struct {
	Value []float32
}

func (r *Runtime) pp_expand_init_number() {
	r.addToGlobal(`NewInt`, r.newNativeFunc(r.builtinGo_NewInt, nil, "NewInt", nil, 2))
	r.addToGlobal(`NewInt64`, r.newNativeFunc(r.builtinGo_NewInt64, nil, "NewInt64", nil, 2))
	r.addToGlobal(`NewInt32`, r.newNativeFunc(r.builtinGo_NewInt32, nil, "NewInt32", nil, 2))
	r.addToGlobal(`NewInt16`, r.newNativeFunc(r.builtinGo_NewInt16, nil, "NewInt16", nil, 2))
	r.addToGlobal(`NewInt8`, r.newNativeFunc(r.builtinGo_NewInt8, nil, "NewInt8", nil, 2))
	r.addToGlobal(`NewUint`, r.newNativeFunc(r.builtinGo_NewUint, nil, "NewUint", nil, 2))
	r.addToGlobal(`NewUint64`, r.newNativeFunc(r.builtinGo_NewUint64, nil, "NewUint64", nil, 2))
	r.addToGlobal(`NewUint32`, r.newNativeFunc(r.builtinGo_NewUint32, nil, "NewUint32", nil, 2))
	r.addToGlobal(`NewUint16`, r.newNativeFunc(r.builtinGo_NewUint16, nil, "NewUint16", nil, 2))
	r.addToGlobal(`NewUint8`, r.newNativeFunc(r.builtinGo_NewUint8, nil, "NewUint8", nil, 2))
	r.addToGlobal(`NewFloat64`, r.newNativeFunc(r.builtinGo_NewFloat64, nil, "NewFloat64", nil, 1))
	r.addToGlobal(`NewFloat32`, r.newNativeFunc(r.builtinGo_NewFloat32, nil, "NewFloat32", nil, 1))
	r.addToGlobal(`MaxInt64`, r.ToValue(NewInt64(math.MaxInt64)))
	r.addToGlobal(`MaxInt32`, r.ToValue(NewInt32(math.MaxInt32)))
	r.addToGlobal(`MaxInt16`, r.ToValue(NewInt16(math.MaxInt16)))
	r.addToGlobal(`MaxInt8`, r.ToValue(NewInt8(math.MaxInt8)))
	r.addToGlobal(`MaxUint64`, r.ToValue(NewUint64(math.MaxUint64)))
	r.addToGlobal(`MaxUint32`, r.ToValue(NewUint32(math.MaxUint32)))
	r.addToGlobal(`MaxUint16`, r.ToValue(NewUint16(math.MaxUint16)))
	r.addToGlobal(`MaxUint8`, r.ToValue(NewUint8(math.MaxUint8)))
	r.addToGlobal(`MaxFloat64`, r.ToValue(NewFloat64(math.MaxFloat64)))
	r.addToGlobal(`MaxFloat32`, r.ToValue(NewFloat32(math.MaxFloat32)))
	r.addToGlobal(`MinInt64`, r.ToValue(NewInt64(math.MinInt64)))
	r.addToGlobal(`MinInt32`, r.ToValue(NewInt32(math.MinInt32)))
	r.addToGlobal(`MinInt16`, r.ToValue(NewInt16(math.MinInt16)))
	r.addToGlobal(`MinInt8`, r.ToValue(NewInt8(math.MinInt8)))
	r.addToGlobal(`IntSize`, intToValue(strconv.IntSize))
}
