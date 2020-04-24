package µ

type Goal func(S) StreamOfStates

// Try evaluates the Goal with an empty state.
func (g Goal) Try() StreamOfStates {
	return g(NewS())
}

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

// =============================================================================

// Or is a goal that returns a logical OR of the goals.
//
// The implementation returns a non-deterministic
// interleave of the result streams;
// such search style is a characteristic of µKanren.
func (g Goal) Or(h Goal) Goal {
	return func(s S) StreamOfStates {
		sc := s.Clone() // we Clone S before we evaluate g
		return g(s).Plus(h(sc))
	}
}

// And is a goal that returns a logical AND of the goals.
func (g Goal) And(h Goal) Goal {
	return func(s S) StreamOfStates {
		return g(s).Bind(goal(h))
	}
}

// =============================================================================

// Once is a goal that returns
// the first success of g, if any
// (and discards further results, if any)
// or fails.
func (g Goal) Once() Goal {
	return func(s S) StreamOfStates {
		ss := g(s)
		head, ok := ss.Head()
		ss.Drop()
		if ok {
			return Unit(head)
		} else {
			return Zero()
		}
	}
}

// =============================================================================

// IfThenElse is a goal that upon evaluation probes Goal g
// (using a clone of the state), and if g evaluates successful,
// evaluates the THEN goal,
// or evaluates the ELSE goal otherwise.
func (g Goal) IfThenElse(THEN, ELSE Goal) Goal {
	return func(s S) StreamOfStates {
		sc := s.Clone()
		IFs := g(sc)
		head, ok := IFs.Head()
		IFs.Drop()

		if ok && head != nil {
			return THEN(s) // then
		} else {
			return ELSE(s) // else
		}
	}
}

// EitherOr is a goal that behaves like this Goal g
// unless this fails, when it behaves like the that Goal.
func (g Goal) EitherOr(that Goal) Goal {
	return g.IfThenElse(g, that)

}

// =============================================================================

// ForEver is a goal that keeps returning g forever.
func (g Goal) ForEver() Goal {
	return func(s S) StreamOfStates {
		return g.Or(g.ForEver())(s)
	}
}

// =============================================================================
