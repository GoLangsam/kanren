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
	done := make(chan struct{})
	cha := New(done)
	go func() {
		_ = cha.Provide(a)
		cha.Close()
	}()
	return cha
}
