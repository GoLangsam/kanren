package kanren

/*
Append is a goal that appends two lists into the third list.

scheme code:

	(define (appendo l t out)
		(lambda (s)
			(lambda ()
				(
					(disj
						(conj (nullo l) (== t out))
						(call/fresh 'a
							(lambda (a)
								(call/fresh 'd
									(lambda (d)
										(call/fresh 'res
											(lambda (res)
												(conj
													(conso a d l)
													(conj
														(conso a res out)
														(appendo d t res)
													)
												)
											)
										)
									)
								)
							)
						)
					)
				)
				s
			)
		)
	)
*/
func Append(l, t, out X) Goal {
	return func(s S) StreamOfStates {
		return Suspend(
			func() StreamOfStates {
				return Disjoint(
					Conjunction(
						Null(l),
						Equal(t, out),
					),
					CallFresh(func(a X) Goal {
						return CallFresh(func(d X) Goal {
							return CallFresh(func(res X) Goal {
								return Conjunction(
									Cons(a, d, l),
									Cons(a, res, out),
									Append(d, t, res),
								)
							})
						})
					}),
				)(s)
			},
		)
	}
}

// Null is a goal that checks if a list is null.
func Null(x X) Goal {
	return Equal(x, nil)
}

// Cons is a goal that conses the first two expressions into the third.
func Cons(a, d, p X) Goal {
	return func(s S) StreamOfStates {
		return Equal(cons(a, d), p)(s)
	}
}

// Car is a goal where the second parameter is the head of the list in the first parameter.
func Car(p, a X) Goal {
	return CallFresh(func(d X) Goal {
		return Equal(cons(a, d), p)
	})
}
