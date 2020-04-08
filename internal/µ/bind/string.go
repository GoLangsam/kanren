package bind

func (bind *Ings) String() string {
	ss := make([]X, len(bind.bound))

	f := func(key V, val X) X {
		return cons(key.Expr(), val)
	}

	for k, x := range bind.bound {
		ss = append(ss, f(k, x))
	}

	return newList(ss...).String()
}
