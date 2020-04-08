package kanren

// Or ...
func Or(gs ...Goal) Goal {
	if len(gs) == 0 {
		return Failure()
	}

	g := gs[0]
	for _, h := range gs[1:] {
		gg, hh := g, h
		g = func(s S) StreamOfStates {
			return gg(s).Interleave(hh(s))
		}
	}
	return g
}
