package pipe

type StreamOfStates <-chan S

func Zero() StreamOfStates {
	zero := make(chan S)
	close(zero)
	return zero
}

func (s StreamOfStates) Head() (a S, ok bool) {
	a, ok = <-s
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

// Dummy for lazy channels
func (s StreamOfStates) Drop() StreamOfStates {
	return s
}
