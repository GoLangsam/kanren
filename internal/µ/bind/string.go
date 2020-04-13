package bind

import "sort"

func (bind *Ings) String() string {
	ss := make([]X, 0, len(bind.bound))

	for k, x := range bind.bound {
		ss = append(ss, cons(k, x))
	}
	sort.Slice(ss, func(i, j int) bool { return (ss[i].Car().String() < ss[j].Car().String()) })

	return newList(ss...).String()
}
