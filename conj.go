package kanren

// Conjunction is a goal that returns a logical AND of the input goals.
func Conjunction(gs ...Goal) Goal {
	if len(gs) == 0 {
		return GOAL
	}
	if len(gs) == 1 {
		return gs[0]
	}
	return gs[0].And(Conjunction(gs[1:]...))
}
