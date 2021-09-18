package m

import (
	"fmt"
	"io/ioutil"
	"math"
	"testing"
	"time"

	"github.com/dop251/goja"
)

func consoleLog(call goja.FunctionCall) goja.Value {

	for i, arg := range call.Arguments {
		if i != 0 {
			fmt.Print(",")
		}
		// if str, ok := arg.Export().(interface {
		// 	String() string
		// }); ok {
		// 	fmt.Print(str.String())
		// } else {
		// 	fmt.Print(arg.String())
		// }
		fmt.Print(arg.String())
	}
	fmt.Println()
	return goja.Undefined()
}

type MyInt64 struct {
	Value int64
}

func (v MyInt64) String() string {
	return fmt.Sprint(v.Value)
}
func (v MyInt64) Native() int64 {
	return v.Value
}
func (v MyInt64) Add(x int64) MyInt64 {
	return MyInt64{v.Value + x}
}

func TestMain(t *testing.T) {
	b, e := ioutil.ReadFile("a.js")
	if e != nil {
		t.Fatal(e)
	}

	vm := goja.New()
	obj := vm.NewObject()
	vm.Set(`console`, obj)
	obj.Set(`log`, consoleLog)
	vm.Set(`make`, func() []int64 {
		x := make([]int64, 10)
		return x
	})
	vm.Set(`makei64`, func() MyInt64 {
		var x MyInt64
		x.Value = math.MaxInt64
		x.Value -= 2
		fmt.Println(`make`, x)
		vm.ToValue(x)
		return x
	})
	vm.Set(`make0`, func() MyInt64 {
		var x MyInt64
		return x
	})
	// vm.Set(`print`, func(call goja.FunctionCall) goja.Value {
	// 	arg := call.Argument(0)
	// 	fmt.Println(arg.ExportType(), arg.Export())
	// 	return goja.Undefined()
	// })
	vm.Set(`print`, func(x []int64) {
		fmt.Println(`print`, x)
	})
	vm.Set(`printInt`, func(x int64) {
		fmt.Println(`print`, x)
	})
	vm.Set(`makeAll`, func() (int, string) {
		return 123, "ok"
	})

	_, e = vm.RunScript("a.js", string(b))
	if e != nil {
		t.Fatal(e)
	}
	time.Sleep(time.Second)
}
