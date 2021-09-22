package goja_test

import (
	"testing"

	"github.com/dop251/goja"
)

func TestExpandPromise(t *testing.T) {
	vm := goja.New()
	vm.Set(`print`, print)
	result, e := vm.RunScriptAndServe("promise.js", `
function check(ok,msg){
	if(!ok){
		if(msg){
			throw new Error("not pass -> "+msg)
		}else{
			throw new Error("not pass")
		}
	}
}
var result = {
	val:0
};
new Promise(function(resolve,reject){
	resolve(1)
}).then(function(v){
	check(v==1,"v1")
	return new Promise(function(resolve){
		resolve(2)
	})
}).then(function(v){
	check(v==2,"v2")
	return new Promise(function(resolve){
		resolve(3)
	})
}).then(function(v){
	check(v==3,"v3")
	return 4
}).then(function(v){
	check(v==4,"v4")
	return 5
}).then(function(v){
	check(v==5,"v5")
	result.val = v
})
result
`)
	if e != nil {
		t.Fatal(e)
	}
	if e != nil {
		t.Fatal(e)
	}
	obj := result.(*goja.Object)
	if obj.Get(`val`).ToInteger() != 5 {
		t.Fatal("promise result not equal 5")
	}
}
