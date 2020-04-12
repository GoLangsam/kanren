package kanren

// CallFresh expects a function f that returns a Goal given an eXpression.
//
// CallFresh returns the Goal which, when evaluated, applies f to a fresh anonymous variable
// and evaluates the resulting Goal.
//
// CallFresh allows to introduce host-language-symbols as free variables when constructing
// some Goal, e.g. in order to model some relation. See Append, for example.
//
func CallFresh(f func(X) Goal) Goal {
	return func(s S) StreamOfStates {
		v := s.V()
		ss := s // TODO: .Clone()?
		return f(v)(ss)
	}
}
