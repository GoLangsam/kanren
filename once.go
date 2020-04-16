package kanren

// Once is a goal that returns the first success of g,
// if any, and discards further results, if any.
func Once(g Goal) Goal {
	return g.Once()
}
