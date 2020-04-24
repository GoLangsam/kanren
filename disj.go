package kanren

// Disjoint is a goal that returns a logical OR of the input goals.
func Disjoint(gs ...Goal) Goal {
	if len(gs) == 0 {
		return FAIL
	}
	if len(gs) == 1 {
		return gs[0]
	}
	g := gs[0]
	h := Disjoint(gs[1:]...)
	return g.Or(h)
}
