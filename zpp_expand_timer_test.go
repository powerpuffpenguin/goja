package goja_test

import (
	"testing"
	"time"

	"github.com/powerpuffpenguin/goja"
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
func TestAsyncController(t *testing.T) {
	vm := goja.New()
	vm.Set(`print`, print)
	vm.Set(`after`, func(f func()) {
		time.AfterFunc(time.Millisecond, f)
	})
	result, e := vm.RunScriptAndServe("async_controller.js", `
var result = {
	val:0
};
var c = NewAsyncController()
c.Async()
after(function(){
	c.Call(function(){
		result.val++
		c.Complete()
	})
})
result
`)
	if e != nil {
		t.Fatal(e)
	}
	obj := result.(*goja.Object)
	if obj.Get(`val`).ToInteger() != 1 {
		t.Fatal("result not equal 1")
	}
}
