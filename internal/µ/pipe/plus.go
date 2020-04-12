package pipe

// Plus is the monadic append.
func (s StreamOfStates) Plus(ss StreamOfStates) StreamOfStates {
	return anyThingFanIn2(s, ss)
}
