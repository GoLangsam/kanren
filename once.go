package kanren

/*
Once is a goal that returns one successful state.

scheme code:

	(define (once g)
		(lambda (s)
			(let loop
				(
					(s1 (g s))
				)
				(cond
					(
						(null? s1)
						()
					)
					(
						(pair? s1)
						(cons (car s1) ())
					)
					(else
						(lambda ()
							(loop (s1))
						)
					)
				)
			)
		)
	)
*/

// Once is a goal that returns one successful state.
func Once(g Goal) Goal {
	return func(s S) StreamOfStates {
		return onceLoop(g(s))
	}
}

func onceLoop(ss StreamOfStates) StreamOfStates {
	if ss == nil {
		return mZero()
	}
	head, ok := ss.Head()
	if !ok {
		return mZero()
	}

	if head != nil {
		return Unit(head)
	}
	return Suspend(
		func() StreamOfStates {
			return onceLoop(ss)
		},
	)
}
