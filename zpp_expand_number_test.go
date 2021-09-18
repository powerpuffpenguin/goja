package goja_test

import (
	"fmt"
	"testing"

	"github.com/dop251/goja"
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

	_, e := vm.RunScript("a.js", `
function check(ok,msg){
	if(!ok){
		if(msg){
			throw new Error("not pass -> "+msg)
		}else{
			throw new Error("not pass")
		}
	}
}
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
`)
	if e != nil {
		t.Fatal(e)
	}
}
