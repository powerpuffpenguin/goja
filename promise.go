package goja

// Callable represents a JavaScript function that can be called from Go.
type JSCallable func(this Value, args ...Value) (result Value, err Value)

// AssertFunction checks if the Value is a function and returns a Callable.
func AssertJSFunction(v Value) (JSCallable, bool) {
	if obj, ok := v.(*Object); ok {
		if f, ok := obj.self.assertCallable(); ok {
			return func(this Value, args ...Value) (ret, err Value) {
				defer func() {
					if x := recover(); x != nil {
						if ex, ok := x.(*uncatchableException); ok {
							err = obj.runtime.NewGoError(ex.err)
						} else {
							panic(x)
						}
					}
				}()
				ex := obj.runtime.vm.try(func() {
					ret = f(FunctionCall{
						This:      this,
						Arguments: args,
					})
				})
				if ex != nil {
					err = ex.val
				}
				vm := obj.runtime.vm
				vm.clearStack()
				if len(vm.callStack) == 0 {
					obj.runtime.leave()
				}
				return
			}, true
		}
	}
	return nil, false
}
