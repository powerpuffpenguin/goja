package goja_test

import (
	"testing"

	"github.com/dop251/goja"
)

func TestExpandTimer(t *testing.T) {
	vm := goja.New()
	vm.Set(`print`, print)
	result, e := vm.RunScriptAndServe("timer.js", `
var result = {
	val:0
};
setTimeout(function(){
	result.val++
	var interval=setInterval(function(){
		result.val++
		if(result.val>5){
			clearInterval(interval)
		}
	},10)
},10)
result
`)
	if e != nil {
		t.Fatal(e)
	}
	obj := result.(*goja.Object)
	if obj.Get(`val`).ToInteger() != 6 {
		t.Fatal("result not equal 6")
	}
}
