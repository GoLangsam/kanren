package bind

func (bind *Ings) String() string {
	ss := make([]X, 0, len(bind.bound))

	for k, x := range bind.bound {
		ss = append(ss, cons(k, x))
	}

	return newList(ss...).String()
}
