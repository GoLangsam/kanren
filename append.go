package kanren

// AppendAWS is the relation: append(l, t) == out.
func AppendAWS(l, t, out X) Goal {
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
							AppendAWS(d, t, res),
						)
					})
				})
			}),
		)(s)
	}
}

// Append is the relation: append(aHead, aTail) == aList.
func Append(aHead, aTail, aList X) Goal {
	return func(s S) StreamOfStates {
		return Disjoint(
			Conjunction(
				Null(aHead),
				Equal(aTail, aList),
			),
			Fresh3(func(µHead, µTail, µList X) Goal {
				return Conjunction(
					Cons(µHead, µTail, aHead),
					Cons(µHead, µList, aList),
					Append(µTail, aTail, µList),
				)
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

// Cdr is the relation: Cdr(list) == tail.
func Cdr(list, tail X) Goal {
	return CallFresh(func(head X) Goal {
		return Equal(cons(head, tail), list)
	})
}
