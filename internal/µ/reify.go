package µ

// reifyVarFromState is a curried function that reifies the input variable for the given bindings.
func reifyVarFromState(v X) func(s S) X {
	return func(s S) X {
		x := s.Walk(v)
		b := s.Clone()
		r := b.Reify(x)
		z := r.Walk(x)
		return z
	}
}

// Reify reifies the input variable for the given input states.
func Reify(v V, ss []S) []X {
	return deriveFmapRs(reifyVarFromState(v), ss)
}

// deriveFmapRs returns a list where each element of the input list has been morphed by the input function.
func deriveFmapRs(f func(S) X, list []S) []X {
	out := make([]X, len(list))
	for i, elem := range list {
		out[i] = f(elem)
	}
	return out
}
