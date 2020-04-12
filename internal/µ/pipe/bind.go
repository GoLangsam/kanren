package pipe

// Bind is the monad bind function for goals.
func (s StreamOfStates) Bind(g func(S) StreamOfStates) StreamOfStates {
	if s == nil {
		s = Zero()
	}

	head, ok := s.Head()
	if !ok {
		return s.Drop()
	}

	return g(head).Plus(s.Bind(g))
}
