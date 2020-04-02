package kanren

import "github.com/GoLangsam/sexpr"

// eXpression

type Variable = sexpr.Variable

type X = *sexpr.Expression

//type v = *Variable

// ============================================================================

func (s S) walk(x X) X {
	if x.IsVariable() {
		return x
	} else {
		term, found := s.Val_at(x.Atom.Var)
		if found {
			return s.walk(term)
		} else {
			return x
		}
	}
}
