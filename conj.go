package kanren

// Conjunction is a goal that returns a logical AND of the input goals.
func Conjunction(gs ...Goal) Goal {
	if len(gs) == 0 {
		return Success()
	}
	if len(gs) == 1 {
		return gs[0]
	}
	g := gs[0]
	h := Conjunction(gs[1:]...)
	return func(s S) StreamOfStates {
		return g(s).Bind(h)
	}
}
