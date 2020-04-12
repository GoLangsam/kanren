package kanren

// Once is a goal that returns the first success of g,
// if any, and discards further results, if any.
func Once(g Goal) Goal {
	return func(s S) StreamOfStates {
		ss := g(s)
		head, ok := ss.Head()
		ss.Drop()
		if !ok {
			return Zero()
		} else {
			return Unit(head)
		}
	}
}
