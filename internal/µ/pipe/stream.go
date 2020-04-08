package pipe

type StreamOfStates <-chan S

var Zero StreamOfStates //=nil, not &Stream{}, not Cons(nil, nil)

func (s StreamOfStates) Head() (e S, ok bool) {
	e, ok = <-s
	return
}

func Unit(a S) StreamOfStates {
	cha := make(chan S)
	go func() {
		cha <- a
		close(cha)
	}()
	return cha
}

// Prepend prepends a given state in front of the given stream-of-states.
func Prepend(a S, thunk func() StreamOfStates) StreamOfStates {
	cha := make(chan S)
	go func() {
		cha <- a
		for e := range thunk() {
			cha <- e
		}
		close(cha)
	}()
	return cha
}

// Suspend prepends a empty state in front of the given stream-of-states thunk.
func Suspend(thunk func() StreamOfStates) StreamOfStates {
	return Prepend(newState(), thunk)
}

func (s StreamOfStates) Suspend() StreamOfStates {
	cha := make(chan S)
	go func() {
		cha <- newState()
		for e := range s {
			cha <- e
		}
		close(cha)
	}()
	return cha
}

func (s StreamOfStates) Concat(thunk func() StreamOfStates) StreamOfStates {
	cha := make(chan S)
	go func() {
		for e := range s {
			cha <- e
		}
		for e := range thunk() {
			cha <- e
		}
		close(cha)
	}()
	return cha
	//return cons(s.Head(), s.Tail().Concat(thunk))
}

func (s StreamOfStates) Interleave(ss StreamOfStates) StreamOfStates {
	return anyThingFanIn2(s, ss)
}
