package µ

// RunGoal calls a goal with an emptystate and n possible resulting states.
//
// Note: If n < 0 (e.g. = -1) then all possible states are returned.
func RunGoal(n int, g Goal) []S {
	ss := g(newState())
	return ss.Take(n)
}
