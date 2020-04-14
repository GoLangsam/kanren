package kanren

// IfThenElse is a goal that upon evaluation probes the IF goal and,
// using a clone of the state, evaluates the THEN goal,
// if IF evaluates successful
// and evaluates the ELSE goal otherwise.
func IfThenElse(IF, THEN, ELSE Goal) Goal {
	return func(s S) StreamOfStates {
		sc := s.Clone()
		IFs := IF(s)
		head, ok := IFs.Head()
		IFs.Drop()

		if ok && head != nil {
			return THEN(sc) // then
		} else {
			return ELSE(sc) // else
		}
	}
}

// EitherOr is a goal that behaves like the THIS Goal
// unless THIS fails, when it behaves like the THAT Goal.
func EitherOr(THIS, THAT Goal) Goal {
	return IfThenElse(THIS, THIS, THAT)

}
