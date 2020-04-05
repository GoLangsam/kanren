package µ

import "github.com/GoLangsam/kanren/internal/µ/bind"

// reifyVarFromState is a curried function that reifies the input variable for the given bindings.
func reifyVarFromState(name string) func(s S) X {
	return func(s S) X {
		// TODO:
		// vari.Able pool implements Fresh & V();
		// bind.Ings implements newName and uses x := NewVariable, thus producing a X directly
		x := s.Walk(s.Fresh(name).Expr())
		b := bind.New()
		return b.Reify(x).Walk(x)
	}
}

/*
// Reify reifies the input variable for the given input states.
func Reify(name string, ss []*State) []*ast.SExpr {
	return deriveFmapRs(reifyVarFromState(name), ss)
}
*/
