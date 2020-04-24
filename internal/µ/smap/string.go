package smap

import "sort"

// String returns a string of the symbolic map sorted by key.
func (m SMap) String() string {
	ss := make([]X, 0, len(m))

	for k, x := range m {
		ss = append(ss, cons(k, x))
	}
	sort.Slice(ss, func(i, j int) bool { return (ss[i].Car().String() < ss[j].Car().String()) })

	return newList(ss...).String()
}

/*
	var out bytes.Buffer

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()

*/
