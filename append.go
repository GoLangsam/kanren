package kanren

// Append is the relation: append(l, t) == out.
func Append(l, t, out X) Goal {
	return func(s S) StreamOfStates {
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
	}
}

// Null is the relation: x == nil.
func Null(x X) Goal {
	return Equal(x, nil)
}

// Cons is the relation: Cons(car, cdr) == pair.
func Cons(car, cdr, pair X) Goal {
	return Equal(cons(car, cdr), pair)
}

// Car is the relation: Car(list) == head.
func Car(list, head X) Goal {
	return CallFresh(func(tail X) Goal {
		return Equal(cons(head, tail), list)
	})
}
