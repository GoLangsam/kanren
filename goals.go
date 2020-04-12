package kanren

// Fail is an alias for Failure.
var Fail = Failure

// Failure is a goal that always returns an empty stream of states.
func Failure() Goal {
	return func(s S) StreamOfStates {
		return Zero()
	}
}

// Success is a goal that always returns the input state in the resulting stream of states.
func Success() Goal {
	return func(s S) StreamOfStates {
		return Unit(s)
	}
}

// Never is a Goal that returns a never ending evaluation of itself.
//
// Note: This is a joke.
// Use on Your own risk!
func Never() Goal {
	return func(s S) StreamOfStates {
		return Never()(s)
	}
}

// Always is a goal that always returns a never ending stream of success.
func Always() Goal {
	return func(s S) StreamOfStates {
		return Disjoint(
			Success(),
			Always(),
		)(s)
	}
}
