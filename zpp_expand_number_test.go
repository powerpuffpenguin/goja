package goja_test

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/powerpuffpenguin/goja"
)

func print(call goja.FunctionCall) goja.Value {
	for i, arg := range call.Arguments {
		if i != 0 {
			fmt.Print(",")
		}
		fmt.Print(arg.String())
	}
	fmt.Println()
	return goja.Undefined()
}
func printType(call goja.FunctionCall) goja.Value {
	for i, arg := range call.Arguments {
		if i != 0 {
			fmt.Print(",")
		}
		fmt.Print(arg.ExportType())
	}
	fmt.Println()
	return goja.Undefined()
}
func TestExpandNumber(t *testing.T) {
	vm := goja.New()
	vm.Set(`print`, print)
	vm.Set(`printType`, printType)
	vm.Set(`nativeInt64`, func(x int64) error {
		if x != math.MaxInt64 {
			return errors.New(`nativeInt64 err`)
		}
		return nil
	})
	vm.Set(`nativeUint64`, func(x uint64) error {
		if x != math.MaxUint64 {
			return errors.New(`nativeUint64 err`)
		}
		return nil
	})
	vm.Set(`nativeFloat64`, func(x float64) error {
		if x != math.MaxFloat64 {
			return errors.New(`nativeFloat64 err`)
		}
		return nil
	})
	vm.Set(`make`, func() map[string]string {
		return map[string]string{
			"id": "1",
		}
	})
	_, e := vm.RunScript("number.js", `
function check(ok,msg){
	if(!ok){
		if(msg){
			throw new Error("not pass -> "+msg)
		}else{
			throw new Error("not pass")
		}
	}
}
nativeInt64(MaxInt64)
nativeUint64(MaxUint64)
nativeFloat64(MaxFloat64)
check(MaxInt64.String()=="9223372036854775807","MaxInt64")
check(MaxInt32.String()=="2147483647","MaxInt32")
check(MaxInt16.String()=="32767","MaxInt16")
check(MaxInt8.String()=="127","MaxInt8")
check(MinInt64.String()=="-9223372036854775808","MinInt64")
check(MinInt32.String()=="-2147483648","MinInt32")
check(MinInt16.String()=="-32768","MinInt16")
check(MinInt8.String()=="-128","MinInt8")

check(MaxUint64.String()=="18446744073709551615","MaxUint64")
check(MaxUint32.String()=="4294967295","MaxUint32")
check(MaxUint16.String()=="65535","MaxUint16")
check(MaxUint8.String()=="255","MaxUint8")

check(MaxFloat64.String()=="1.7976931348623157e+308","MaxFloat64")
check(MaxFloat32.String()=="3.4028235e+38","MaxFloat32")

check(MaxInt64.Sub(1).String()=="9223372036854775806","MaxInt64")
check(MaxInt32.Sub(1).String()=="2147483646","MaxInt32")
check(MaxInt16.Sub(1).String()=="32766","MaxInt16")
check(MaxInt8.Sub(1).String()=="126","MaxInt8")
check(MinInt64.Add(1).String()=="-9223372036854775807","MinInt64")
check(MinInt32.Add(1).String()=="-2147483647","MinInt32")
check(MinInt16.Add(1).String()=="-32767","MinInt16")
check(MinInt8.Add(1).String()=="-127","MinInt8")

check(MaxUint64.Sub(NewInt(2)).String()=="18446744073709551613","MaxUint64")
check(MaxUint32.Sub(NewInt8(2)).String()=="4294967293","MaxUint32")
check(MaxUint16.Sub(NewUint(2)).String()=="65533","MaxUint16")
check(MaxUint8.Sub(NewFloat32(2)).String()=="253","MaxUint8")
// printType(MinInt8.Add(1))

var a = NewInt64Array(2,4)
var b = NewInt64Array(2)
a.Set(0,1)
a.Set(1,2)
b.Set(0,3)
b.Set(1,MaxInt64)
check(a.Append(b).String()=="[1 2 3 9223372036854775807]","Append")
check(a.Append(3,MaxInt64).String()=="[1 2 3 9223372036854775807]","Append")
check(a.Append(3,MaxInt64).Join(",")=="1,2,3,9223372036854775807","Join")
check(goLen(a)==2,'goLen')
var m =make()
check(goLen(m)==1,'goLen map')
check(goHasKey(m,"id"),'goLen has id')
check(!goHasKey(m,"name"),'goLen has name')
`)
	if e != nil {
		t.Fatal(e)
	}
}
