package µ

// Success is a goal that always returns the input state in the resulting stream of states.
func Success() Goal {
	return func(s S) StreamOfStates {
		return Unit(s)
	}
}

// Failure is a goal that always returns an empty stream of states.
func Failure() Goal {
	return func(s S) StreamOfStates {
		return mZero
	}
}

// Equal returns a Goal that unifies the input expressions in the output stream.
func Equal(u, v X) Goal {
	return func(s S) StreamOfStates {
		if s.Unify(u, v) {
			return Unit(s.Clone())
		}
		return mZero
	}
}

// Never is a Goal that returns a never ending stream of suspensions.
func Never() Goal {
	return func(s S) StreamOfStates {
		return Suspend(func() StreamOfStates {
			return Never()(s)
		})
	}
}

// Always is a goal that returns a never ending stream of success.
func Always() Goal {
	return func(s S) StreamOfStates {
		return Suspend(func() StreamOfStates {
			return Or(
				Success(),
				Always(),
			)(s)
		})
	}
}
