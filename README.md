goja
====

ECMAScript 5.1(+) implementation in Go.

Minimum required Go version is 1.14.

This is a version of the extended function after [github.com/dop251/goja](https://github.com/dop251/goja) fork. The usage method is consistent with the original version. For more api details, please refer to the original project. The following text only explains the features of extensions and additional extensions .

Features
====

* Added event loop, call runtime.Serve event run loop in go code.
* Implemented timer setTimeout clearTimeout setInterval clearInterval on event loop
* Provide support for Promise Promise.resolve Promise.reject
* When js calls the go function through reflection, if the last parameter is a Scheduler, the go function can be called asynchronously (call it in a goroutine other than the js main goroutine and pass the result back to the main goroutine)

```
package goja_test

import (
	"testing"

	"github.com/powerpuffpenguin/goja"
)

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
```
.d.ts
====

**zpp_expand/generation/globals.d.ts** declares the functions that can be called in js

GoJS
====

goja is just a js engine. If there is no api extension, it will be difficult to use it in other places other than the calculation logic. For this reason, I also started a new project to import the go standard library into goja for js to call

[https://github.com/powerpuffpenguin/gojs](https://github.com/powerpuffpenguin/gojs)

NodeJS Compatibility
====

The style of NodeJS is very different from golang, but it may be useful due to its popular compatibility with NodeJS, but this is not a small project. If I have free time after perfecting GoJS, I may implement a set of NodeJS compatible APIs, but I can expect it This will not be achieved for a long time.