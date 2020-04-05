package µ

//type head = S
//type tail func() *Stream

type Stream struct {
	head S
	tail func() *Stream
}

type StreamOfStates = *Stream

func (s Stream) Head() S {
	return s.head
}

func (s Stream) Tail() StreamOfStates {
	return s.tail()
}

var mZero StreamOfStates //not &Stream{}, not Cons(nil, nil)
var nilTail = func() StreamOfStates { return mZero }

func Unit(a S) StreamOfStates {
	return Cons(a, nilTail)
}

func Cons(head S, tail func() StreamOfStates) StreamOfStates {
	return &Stream{head, tail}
}

func (s Stream) Concat(tail func() StreamOfStates) StreamOfStates {
	return Cons(s.Head(), func() *Stream { return s.Tail().Concat(tail) })
}

func (s Stream) Interleave(s2 StreamOfStates) StreamOfStates {
	return Cons(s.Head(), func() StreamOfStates { return s2.Interleave(s.Tail()) })
}

// Suspend prepends a empty state in front of the given stream of states.
func Suspend(s func() StreamOfStates) StreamOfStates {
	return Cons(*Init(), func() StreamOfStates { return s() })
}
