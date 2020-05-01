package kanren

// CallFresh expects a function f that returns a Goal given an eXpression.
//
// CallFresh returns the Goal which, when evaluated,
// applies f to a fresh anonymous variable
// and evaluates the resulting Goal.
//
// CallFresh allows to introduce a host-language-symbol as a free variable
// when constructing some Goal, e.g. in order to model some relation.
// See `Append`, for example.
func CallFresh(f func(V) Goal) Goal {
	return func(s S) StreamOfStates {
		v := s.V()
		return f(v)(s)
	}
}

func Fresh1(f func(V) Goal) Goal {
	return func(s S) StreamOfStates {
		v := s.V()
		return f(v)(s)
	}
}

func Fresh2(f func(V, V) Goal) Goal {
	return func(s S) StreamOfStates {
		v1, v2 := s.V(), s.V()
		return f(v1, v2)(s)
	}
}

func Fresh3(f func(V, V, V) Goal) Goal {
	return func(s S) StreamOfStates {
		v1, v2, v3 := s.V(), s.V(), s.V()
		return f(v1, v2, v3)(s)
	}
}

func Fresh4(f func(V, V, V, V) Goal) Goal {
	return func(s S) StreamOfStates {
		v1, v2, v3, v4 := s.V(), s.V(), s.V(), s.V()
		return f(v1, v2, v3, v4)(s)
	}
}

func Fresh5(f func(V, V, V, V, V) Goal) Goal {
	return func(s S) StreamOfStates {
		v1, v2, v3, v4, v5 := s.V(), s.V(), s.V(), s.V(), s.V()
		return f(v1, v2, v3, v4, v5)(s)
	}
}

func Fresh6(f func(V, V, V, V, V, V) Goal) Goal {
	return func(s S) StreamOfStates {
		v1, v2, v3, v4, v5, v6 := s.V(), s.V(), s.V(), s.V(), s.V(), s.V()
		return f(v1, v2, v3, v4, v5, v6)(s)
	}
}

func Fresh7(f func(V, V, V, V, V, V, V) Goal) Goal {
	return func(s S) StreamOfStates {
		v1, v2, v3, v4, v5, v6, v7 := s.V(), s.V(), s.V(), s.V(), s.V(), s.V(), s.V()
		return f(v1, v2, v3, v4, v5, v6, v7)(s)
	}
}

func Fresh8(f func(V, V, V, V, V, V, V, V) Goal) Goal {
	return func(s S) StreamOfStates {
		v1, v2, v3, v4, v5, v6, v7, v8 := s.V(), s.V(), s.V(), s.V(), s.V(), s.V(), s.V(), s.V()
		return f(v1, v2, v3, v4, v5, v6, v7, v8)(s)
	}
}
