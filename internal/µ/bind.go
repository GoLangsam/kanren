package µ

// Bind is the monad bind function for goals.
func (g Goal) Bind(s StreamOfStates) StreamOfStates {
	if s == mZero {
		return mZero
	}
	head := s.Head()
	if head != nil { // not a suspension => procedure? == false
		return g(head).Plus(g.Bind(s))
	}
	return Suspend(func() StreamOfStates {
		return g.Bind(s)
	})
}
