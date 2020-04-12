package kanren

// Equal returns a Goal that unifies the input expressions in the output stream.
func Equal(x, y X) Goal {
	return func(s S) StreamOfStates {
		if s.Unify(x, y) {
			return Unit(s) // .Clone())
		}
		return mZero()
	}
}
