package kanren

// CallFresh expects a function that expects an expression and returns a Goal
// and retuns the Goal which applies this to a fresh anonymous variable in a cloned state.
func CallFresh(f func(X) Goal) Goal {
	return func(s S) StreamOfStates {
		v := s.V()
		ss := s // .Clone()
		return f(v)(ss)
	}
}
