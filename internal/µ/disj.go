package µ

// Disjoint is a goal that returns a logical OR of the input goals.
func Disjoint(gs ...Goal) Goal {
	if len(gs) == 0 {
		return Failure()
	}
	if len(gs) == 1 {
		return gs[0]
	}
	g1 := gs[0]
	g2 := Disjoint(gs[1:]...)
	return func(s S) StreamOfStates {
		return g1(s).Plus(g2(s))
	}
}
