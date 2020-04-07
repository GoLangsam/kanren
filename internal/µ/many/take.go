package many

// Take returns the first n non-nil states from the stream of states as a slice.
// If n < 0 the whole stream is returned.
func (s StreamOfStates) Take(n int) []S {
	if n == 0 {
		return nil
	}
	if s == Zero {
		return nil
	}
	if s.head == nil {
		return s.tail().Take(n)
	}
	return append([]S{s.head}, s.tail().Take(n-1)...)
}
