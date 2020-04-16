package kanren

// Equal is a relation: it reports whether x unifies with y.
//
// Note: In Scheme, Equal is often spelled "==".
func Equal(x, y X) Goal {
	return func(s S) StreamOfStates {
		if s.Unify(x, y) {
			return Unit(s) // TODO: .Clone()) ?
		}
		return Zero()
	}
}
