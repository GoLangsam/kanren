package pipe

// ForEver evaluates the given goal forever.
func ForEver(s S, goal func(S) StreamOfStates) StreamOfStates {
	// TODO: done := a.Done()
	done := make(chan struct{})
	cha := New(done)
	go func() {
		defer cha.Close()
		for {
			sc := s.Clone()
			ss := goal(sc)
			for state, ok := ss.Receive(); ok; state, ok = ss.Receive() {
				//r state := range ss {
				cha.Provide(state)
			}
			ss.Drop()
		}
	}()
	return cha
}

// TODO: ForEver needs to respect done/abort!

// =============================================================================
