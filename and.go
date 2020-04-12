package kanren

// And

func and_composer(g Goal, s StreamOfStates) StreamOfStates {
	if s == nil {
		return mZero()
	}

	head, ok := s.Head()
	if !ok {
		return mZero()
	}
	return g(head).Concat(func() StreamOfStates {
		return and_composer(g, s)
	})
}

// And ...
func And(gs ...Goal) Goal {
	if len(gs) == 0 {
		return Success()
	}

	g := gs[0]
	for _, h := range gs[1:] {
		gg, hh := g, h
		g = func(s S) StreamOfStates {
			return and_composer(gg, hh(s))
		}
	}
	return g
}
