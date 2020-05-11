package kanren

import "reflect"

// Call helps to construct recursive goals
func Call(constructor interface{}, args ...interface{}) Goal {
	foo := make([]reflect.Value, len(args))
	for i, e := range args {
		foo[i] = reflect.ValueOf(e)
	}
	fun := reflect.ValueOf(constructor)
	return func(s S) StreamOfStates {
		r := fun.Call(foo)
		x := r[0].Interface()
		g, ok := x.(Goal)
		if ok {
			return g(s)
		} else {
			panic("whoops")
		}

	}
}
