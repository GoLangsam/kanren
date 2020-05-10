package pipe

func zero() StreamOfStates {
	zero := New(make(chan struct{}))
	zero.Close()
	return zero
}

var ZERO StreamOfStates = zero()

func (s StreamOfStates) Head() (a S, ok bool) {
	return s.Get()
}

func Unit(a S) StreamOfStates {
	// TODO: done := a.Done()
	cha := New(make(chan struct{}))
	go func() {
		cha.Provide(a)
		cha.Close()
	}()
	return cha
}
