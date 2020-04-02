package µ

type Goal func(S) *Stream

// And

func and_composer(g1s *Stream, g2 Goal) *Stream {
	if g1s == mZero() {
		return mZero()
	} else {
		return g2(g1s.head).concat(func() *Stream {
			a := g1s.tail()
			if a == mZero() {
				return mZero()
			} else {
				return and_composer(a, g2)
			}
		})
	}
}

func and_base(g1, g2 Goal) Goal {
	return func(s S) *Stream {
		g1s := g1(s)
		return and_composer(g1s, g2)

	}
}

func And(gs ...Goal) Goal {
	var g Goal = gs[0]
	for _, e := range gs[1:] {
		g = and_base(g, e)
	}
	return g
}

// OR

func or_base(g1, g2 Goal) Goal {
	return func(s S) *Stream {
		g1s := g1(s)
		g2s := g2(s)
		return g1s.interleave(g2s)
	}
}

func Fail() Goal {
	return func(s S) *Stream {
		return mZero()
	}
}

func Or(gs ...Goal) Goal {
	var g Goal = gs[0]
	for _, e := range gs[1:] {
		g = or_base(g, e)
	}
	return g
}
