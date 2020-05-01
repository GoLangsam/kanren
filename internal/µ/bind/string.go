package bind

import "sort"
import "strings"

// String returns a string of the symbolic map sorted by key.
func (m sMap) String() string {
	keyS := make([]key, 0, len(m))

	for key, _ := range m {
		keyS = append(keyS, key)
	}
	sort.Slice(keyS, func(i, j int) bool { return (keyS[i] < keyS[j]) })

	var out strings.Builder

	beg := func() { out.WriteString("(") }
	end := func() { out.WriteString(")") }

	beg()
	for _, key := range keyS {
		beg()

		out.WriteString(",") // escape the variable name
		out.WriteString(string(key))
		out.WriteString(" . ")
		out.WriteString(m[key].String())

		end()
	}
	end()

	return out.String()
}
