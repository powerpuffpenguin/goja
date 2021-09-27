package goja_test

import (
	"testing"

	"github.com/powerpuffpenguin/goja"
)

func TestExpandPromiseErr(t *testing.T) {
	vm := goja.New()
	vm.Set(`print`, print)
	vm.Set(`err`, func(args ...interface{}) {
		t.Fatal(args...)
	})
	_, e := vm.RunScriptAndServe("promise.js", `
function check(ok,msg){
	if(!ok){
		if(msg){
			err("not pass -> "+msg)
		}else{
			err("not pass")
		}
	}
}
new Promise((resolve, reject) => {
    throw 123
}).catch((e) => {
    check(typeof e==="number")
})
Promise.resolve(123).then((v)=>{
	throw v.toString()
}).catch((e) => {
    check(typeof e==="string")
})
`)
	if e != nil {
		t.Fatal(e)
	}
}
func TestExpandPromise(t *testing.T) {
	vm := goja.New()
	vm.Set(`print`, print)
	vm.Set(`add`, func(val ...int) int {
		var result int
		for _, v := range val {
			result += v
		}
		return result
	})
	vm.Set(`err`, func(args ...interface{}) {
		t.Fatal(args...)
	})
	result, e := vm.RunScriptAndServe("promise.js", `
function check(ok,msg){
	if(!ok){
		if(msg){
			err("not pass -> "+msg)
		}else{
			err("not pass")
		}
	}
}
Promise.reject(123).then(() => {
    check(false,'unexpected then')
}).then((v) => {
	check(false,'err then')
}, (e) => {
	 check(e==123,'catch not equal')
})
Promise.reject(123).then(() => {
    check(false,'unexpected then')
},(e)=>{
	check(e==123,'catch not equal')
	return 456
}).then((v) => {
	check(v==456,'then not equal')
}, (e) => {
	check(false,'unexpected catch')
})

var result = {
	val:0
};


// Call the go function synchronously
check(add(1,2,3),6)

// The last parameter is passed into the Scheduler to asynchronously call the go function
add(1,2,3,DefaultScheduler).then(function(v){
	check(v==6,6)

	// test Promise
	var completer=new Completer();
	completer.resolve(1);
	completer.promise.then(function(v){
		check(v==1,"v1")
		return Promise.resolve(2)
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
})
result
`)
	if e != nil {
		t.Fatal(e)
	}
	obj := result.(*goja.Object)
	if obj.Get(`val`).ToInteger() != 5 {
		t.Fatal("promise result not equal 5")
	}
}
