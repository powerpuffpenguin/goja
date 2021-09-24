package goja_test

import (
	"testing"
	"time"

	"github.com/powerpuffpenguin/goja"
)

func TestExpandChan(t *testing.T) {
	vm := goja.New()
	vm.Set(`print`, print)
	vm.Set(`printType`, printType)
	vm.Set(`Second`, time.Second)
	vm.Set(`make`, func(i int) chan time.Duration {
		if i == 0 {
			return make(chan time.Duration)
		}
		return make(chan time.Duration, i)
	})
	v, e := vm.RunScriptAndServe("chan.js", `
function check(ok,msg){
	if(!ok){
		if(msg){
			throw new Error("not pass -> "+msg)
		}else{
			throw new Error("not pass")
		}
	}
}
var obj={
	v:0,
}
var ch=make()
var result=goTryRecv(ch)
check(!result[1])
result=goTrySend(ch,Second)
check(!result)
goRecv(ch,DefaultScheduler).then(function(result){
	try{
		check(result[1])
		check(result[0].String()==Second.String())
		
		var r=goRecv(ch)
		check(!r[1])

		obj.v=result[0]
	}catch(e){
		print(e)
	}
});
goSend(ch,Second);
goClose(ch);
obj;
`)
	if e != nil {
		t.Fatal(e)
	}
	obj := v.(*goja.Object)
	export := obj.Get("v").Export().(time.Duration)
	if export != time.Second {
		t.Fatal(`chan not equal`)
	}
}
func TestExpandChanNumber(t *testing.T) {
	vm := goja.New()
	vm.Set(`print`, print)
	vm.Set(`printType`, printType)
	vm.Set(`make`, func(i int) chan int {
		if i == 0 {
			return make(chan int)
		}
		return make(chan int, i)
	})
	v, e := vm.RunScriptAndServe("chan_number.js", `
function check(ok,msg){
	if(!ok){
		if(msg){
			throw new Error("not pass -> "+msg)
		}else{
			throw new Error("not pass")
		}
	}
}
var obj={
	v:0,
}
var ch=make()
var result=goTryRecv(ch)
check(!result[1])
result=goTrySend(ch,NewInt(1))
check(!result)
goRecv(ch,DefaultScheduler).then(function(result){
	try{
		check(result[1])
		check(result[0].String()=="1")
		
		var r=goRecv(ch)
		check(!r[1])
	
		obj.v=result[0]
	}catch(e){
		print(e)
	}
});
goSend(ch,NewInt(1));
goClose(ch);
obj;
`)
	if e != nil {
		t.Fatal(e)
	}
	obj := v.(*goja.Object)
	export := obj.Get("v").Export().(goja.Int)
	if export != 1 {
		t.Fatal(`chan number not equal`)
	}
}
func TestExpandChanSelect(t *testing.T) {
	vm := goja.New()
	vm.Set(`print`, print)
	vm.Set(`printType`, printType)
	vm.Set(`make`, func(i int) chan int {
		if i == 0 {
			return make(chan int)
		}
		return make(chan int, i)
	})
	_, e := vm.RunScriptAndServe("chan_select.js", `
function check(ok,msg){
	if(!ok){
		if(msg){
			throw new Error("not pass -> "+msg)
		}else{
			throw new Error("not pass")
		}
	}
}
var c0 = make(1);
var result = goSelect(
	NewRecvCase(c0),
	DefaultCase
)
check(result[0].String()=='1')
check(!result[2])
result = goSelect(
	NewSendCase(c0,NewInt(2)),
	DefaultCase
);
check(result[0].String()=='0')
check(!result[2])

var result = goSelect(
	DefaultCase,
	NewRecvCase(c0)
)
check(result[0].String()=='1')
check(result[2])
check(result[1].String()=='2')
`)
	if e != nil {
		t.Fatal(e)
	}

}
