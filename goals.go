package kanren

// Failure is a goal that always returns an empty stream of states.
func Failure() Goal {
	return func(s S) StreamOfStates {
		return mZero()
	}
}

// Success is a goal that always returns the input state in the resulting stream of states.
func Success() Goal {
	return func(s S) StreamOfStates {
		return Unit(s)
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
