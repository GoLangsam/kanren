package kanren

// IfThenElse is a goal that evaluates the THEN goal if the IF goal is successful,
// otherwise it evaluates the ELSE goal.
func IfThenElse(IF, THEN, ELSE Goal) Goal {
	return func(s S) StreamOfStates {
		head, ok := IF(s).Head()
		IF(s).Drop()

		if ok && head != nil { // && head.Len() != 0 {
			return THEN(s) // then
		} else {
			return ELSE(s) // else
		}
	}
}
