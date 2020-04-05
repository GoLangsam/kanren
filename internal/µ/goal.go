package µ

// type E = *Stream // Eventualities
type Goal func(S) *Stream

// And

func (g Goal) and_composer(s *Stream) *Stream {
	if s == mZero {
		return mZero
	}

	return g(s.head).Concat(func() *Stream {
		a := s.tail()
		if a == mZero {
			return mZero
		}
		return g.and_composer(a)
	})
}

func (g Goal) and_base(g1 Goal) Goal {
	return func(s S) *Stream {
		ss := g1(s)
		return g.and_composer(ss)

	}
}

func And(gs ...Goal) Goal {
	if len(gs) == 0 {
		return Fail()
	}

	var g Goal = gs[0]
	for _, h := range gs[1:] {
		// g = g.and_base(h)
		g = func(s S) *Stream {
			return g.and_composer(h(s))
		}
	}
	return g
}

// OR

func (g Goal) or_base(h Goal) Goal {
	return func(s S) *Stream {
		return g(s).Interleave(h(s))
	}
}

func Or(gs ...Goal) Goal {
	if len(gs) == 0 {
		return Fail()
	}

	var g Goal = gs[0]
	for _, h := range gs[1:] {
		//	g = g.or_base(h)
		g = func(s S) *Stream {
			return g(s).Interleave(h(s))
		}
	}
	return g
}

// Fail

func Fail() Goal {
	return func(s S) *Stream {
		return mZero
	}
}
