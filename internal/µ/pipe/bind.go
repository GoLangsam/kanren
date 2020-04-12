package pipe

// Bind is the monad bind function for goals.
func (s StreamOfStates) Bind(g func(S) StreamOfStates) StreamOfStates {
	if s == nil {
		return Zero()
	}

	head, ok := s.Head()
	if !ok {
		return Zero()
	}

	if head != nil { // not a suspension => procedure? == false
		return g(head).Plus(
			s.Bind(g),
		)
	}
	{
		return Suspend(func() StreamOfStates {
			return s.Bind(g)
		})
	}
}
