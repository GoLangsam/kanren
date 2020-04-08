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
