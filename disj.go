package kanren

// Disjoint is a goal that returns a logical OR of the input goals.
func Disjoint(gs ...Goal) Goal {
	if len(gs) == 0 {
		return Failure()
	}
	if len(gs) == 1 {
		return gs[0]
	}
	g := gs[0]
	h := Disjoint(gs[1:]...)
	return func(s S) StreamOfStates {
		sc := s.Clone() // we Clone S before we evaluate g
		return g(s).Plus(h(sc))
	}
}
