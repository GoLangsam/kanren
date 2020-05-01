package pipe

// Plus is the monadic append.
func (s StreamOfStates) Plus(ss StreamOfStates) StreamOfStates {
	return StreamOfStates{s.FanIn2(ss.StreamOfStates)}
}
