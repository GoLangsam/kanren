package µ

import "github.com/GoLangsam/kanren/internal/µ/bind"

// reifyVarFromState is a curried function that reifies the input variable for the given bindings.
func reifyVarFromState(name string) func(s S) X {
	return func(s S) X {
		// TODO:
		// vari.Able pool implements Fresh & V();
		// bind.Ings implements newName and uses x := NewVariable, thus producing a X directly
		v := s.Fresh(name)
		x := v.Expr()
		xx := s.Walk(x)
		bb := bind.New()
		return bb.Reify(xx).Walk(xx)
	}
}

/*
// Reify reifies the input variable for the given input states.
func Reify(name string, ss []*State) []*ast.SExpr {
	return deriveFmapRs(reifyVarFromState(name), ss)
}
*/
