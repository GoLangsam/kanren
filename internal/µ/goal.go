package µ

type Goal func(S) StreamOfStates

// Equal returns a Goal that unifies the input expressions in the output stream.
func Equal(x, y X) Goal {
	return func(s S) StreamOfStates {
		if s.Unify(x, y) {
			return Unit(s.Clone())
		}
		return mZero
	}
}

// CallFresh expects a function that expects a variable and returns a Goal.
func CallFresh(f func(V) Goal) Goal {
	return func(s S) StreamOfStates {
		v := s.V()
		ss := s.Clone()
		return f(v)(ss)
	}
}

// And

func (g Goal) and_composer(s StreamOfStates) StreamOfStates {
	if s == mZero {
		return mZero
	}

	head, ok := s.Head()
	if !ok {
		return mZero
	}
	return g(head).Concat(func() StreamOfStates {
		return g.and_composer(s)
	})
}

// And ...
func And(gs ...Goal) Goal {
	if len(gs) == 0 {
		return Failure()
	}

	var g Goal = gs[0]
	for _, h := range gs[1:] {
		gg, hh := g, h
		g = func(s S) StreamOfStates {
			return gg.and_composer(hh(s))
		}
	}
	return g
}

// Or ...
func Or(gs ...Goal) Goal {
	if len(gs) == 0 {
		return Failure()
	}

	var g Goal = gs[0]
	for _, h := range gs[1:] {
		gg, hh := g, h
		g = func(s S) StreamOfStates {
			return gg(s).Interleave(hh(s))
		}
	}
	return g
}
