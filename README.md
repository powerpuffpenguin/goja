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
* New supports some option parameters to define how native code works
```
package goja

import (
	"testing"
)

func TestPromise(t *testing.T) {
	r := New()
	r.Set(`err`, func(args ...interface{}) {
		t.Fatal(args...)
	})
	_, err := r.RunString(`
function check(ok,msg){
	if(!ok){
		if(msg){
			err("not pass -> "+msg)
		}else{
			err("not pass")
		}
	}
}
var p = new Promise((resolve,reject)=>{
	resolve(1)
}).then((v)=>{
	check(v==1,'then not equal')
},(e)=>{
	check(false,'unexpected catch')
})
check(p instanceof Promise,'instanceof false')

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
	if err != nil {
		t.Fatal(err)
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