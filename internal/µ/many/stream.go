package many

type Stream struct {
	head // = S
	tail // func() *Stream
}

type head = S
type tail func() *Stream
type StreamOfStates = *Stream

func (s StreamOfStates) Head() S {
	return s.head
}

func (s StreamOfStates) Tail() StreamOfStates {
	return s.tail()
}

var Zero StreamOfStates //=nil, not &Stream{}, not Cons(nil, nil)
var nilTail = func() StreamOfStates { return Zero }

func Unit(a S) StreamOfStates {
	return cons(a, Zero)
}

func conc(head head, tail tail) StreamOfStates {
	return &Stream{head, tail}
}

// Note: Cons builds the closure for tail internally.
func cons(head head, tail StreamOfStates) StreamOfStates {
	return conc(head, func() StreamOfStates { return tail })
}

// Suspend prepends a empty state in front of the given stream-of-states thunk.
func Suspend(thunk func() StreamOfStates) StreamOfStates {
	return conc(newState(), thunk)
}

func (s StreamOfStates) Suspend() StreamOfStates {
	return cons(newState(), s)
}

func (s StreamOfStates) Concat(thunk func() StreamOfStates) StreamOfStates {
	return cons(s.Head(), s.Tail().Concat(thunk))
}

func (s StreamOfStates) Interleave(ss StreamOfStates) StreamOfStates {
	return cons(s.Head(), ss.Interleave(s.Tail()))
}
