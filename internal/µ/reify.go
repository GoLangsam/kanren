package µ

import "github.com/GoLangsam/kanren/internal/µ/bind"

// reifyVarFromState is a curried function that reifies the input variable for the given bindings.
func reifyVarFromState(v V) func(s S) X {
	return func(s S) X {
		x := s.Walk(v.Expr())
		b := bind.New()
		return b.Reify(x).Walk(x)
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
