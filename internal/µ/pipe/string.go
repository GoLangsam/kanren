package pipe

import "strings"

// Strings blocks until stream is closed,
// and returns its strings representation.
func (s StreamOfStates) String() string {
	if s == nil {
		return "<nil>"
	}

	var b strings.Builder
	b.WriteString("(")
	for e := range s {
		b.WriteString(e.String())
	}
	b.WriteString(")")
	return b.String()
}
