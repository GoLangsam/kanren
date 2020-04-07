package pipe

// Take returns the first n non-nil states from the stream of states as a slice.
// If n < 0 the whole stream is returned.
func (s StreamOfStates) Take(n int) []S {
	if n == 0 {
		return nil
	}
	if s == Zero {
		return nil
	}
	head := s.Head()
	if head == nil {
		return s.Take(n)
	}
	return append([]S{head}, s.Take(n-1)...)
}
