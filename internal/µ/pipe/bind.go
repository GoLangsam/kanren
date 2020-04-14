package pipe

func (s StreamOfStates) bind(out chan<- S, g func(S) StreamOfStates) {
	defer close(out)
	if s == nil {
		return
	}

	//r head, ok := s.Head(); ok; head, ok = s.Head() {
	for head := range s {
		for gs := range g(head) {
			out <- gs
		}
	}
	s.Drop()
}

// Bind is the monad bind function for goals.
func (s StreamOfStates) Bind(g func(S) StreamOfStates) StreamOfStates {
	cha := make(chan anyThing)
	go s.bind(cha, g)
	return cha
}
