package pipe

import "strings"

// String blocks until stream is closed,
// consumes the stream
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
	s.Drop()
	b.WriteString(")")
	return b.String()
}
