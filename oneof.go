package kanren

// OneOf is a goal that returns v to be equal to one of the given symbols.
func OneOf(v V, symbols ...X) Goal {

	gs := make([]Goal, len(symbols))

	for i, symbol := range symbols {
		gs[i] = Equal(symbol, v)
	}

	return Disjoint(gs...)
}

// AnyOf is a goal that returns v to be equal to any one of the atoms the given expression is composed of.
func AnyOf(v V, x X) Goal {

	if !x.IsPair() {
		return Equal(x, v)
	}

	gs := []Goal{}

	for head, tail := x.Pair.Car, x.Pair.Cdr; head != nil; head, tail = tail.Car(), tail.Cdr() {
		gs = append(gs, anyOf(v, head)...)
	}

	return Disjoint(gs...)
}

func anyOf(v V, x X) (goals []Goal) {

	if !x.IsPair() {
		return []Goal{Equal(x, v)}
	}

	for head, tail := x.Pair.Car, x.Pair.Cdr; head != nil; head, tail = tail.Car(), tail.Cdr() {
		goals = append(goals, anyOf(v, head)...)
	}

	return
}

// AtAnyPostitionOf is a goal that returns v to be at any position of an n-Tuple of variables.
func AtAnyPositionOf(v V, n int, q V) Goal {

	return func(s S) StreamOfStates {
		gs := make([]Goal, n)

		for p := 0; p < n; p++ {
			vs := make([]V, n)
			for i := 0; i < n; i++ {
				if p == i {
					vs[i] = v
				} else {
					vs[i] = s.V()
				}
			}
			gs[p] = Equal(NewList(vs...), q)
		}

		return Disjoint(gs...)(s)

	}
}
