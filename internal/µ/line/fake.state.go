package micro

type State struct{}

func EmptyState() *State {
	return &State{}
}

func (*State) Unify(x, y X) (xUy X, ok bool) {
	return
}

// end of fakes
