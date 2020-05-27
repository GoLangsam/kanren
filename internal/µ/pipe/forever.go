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
			//r state := range ss {
			for state, ok := ss.Head(); ok; state, ok = ss.Head() {
				cha.Provide(state)
			}
		}
	}()
	return cha
}

// TODO: ForEver needs to respect done/abort!

// =============================================================================
