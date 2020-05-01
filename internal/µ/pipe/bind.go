package pipe

func (s StreamOfStates) bind(out StreamOfStates, goal func(S) StreamOfStates) {

	//r head, ok := s.Head(); ok; head, ok = s.Head() {
	for head, ok := s.Head(); ok; head, ok = s.Head() {
		s := goal(head)
		for head, ok := s.Head(); ok; head, ok = s.Head() {
			out.Provide(head)
		}
		s.Drop()
	}
	s.Drop()
	out.Close()
}

// Bind is the monad bind function for goals.
func (s StreamOfStates) Bind(goal func(S) StreamOfStates) StreamOfStates {
	cha := s.New()
	go s.bind(cha, goal)
	return cha
}

// TODO: Bind needs to respect done/abort!

// TODO: pipe/s knows only about ...PipeFunc( ..., acts ...func(a anyThing) anyThing)
