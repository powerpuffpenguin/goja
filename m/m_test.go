package m

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

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

func TestMain(t *testing.T) {
	var x []byte
	fmt.Println(x[0:])
	return
	b, e := ioutil.ReadFile("a.js")
	if e != nil {
		t.Fatal(e)
	}

	vm := goja.New()
	obj := vm.NewObject()
	vm.Set(`console`, obj)
	vm.Set(`make`, func() []byte {
		x := make([]byte, 10)
		x[0] = 1
		return x
		// return nil
	})
	vm.Set(`fmt`, func(b []byte) {
		fmt.Println(b)
	})
	vm.Set(`makeb`, func() byte {
		return 1
	})
	vm.Set(`print`, print)
	vm.Set(`printType`, printType)

	_, e = vm.RunScript("a.js", string(b))
	if e != nil {
		t.Fatal(e)
	}
	time.Sleep(time.Second)
}
