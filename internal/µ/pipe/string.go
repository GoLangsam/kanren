package pipe

import "strings"

// String blocks until stream is closed,
// consumes the stream
// and returns its strings representation.
func (s StreamOfStates) String() string {
	var b strings.Builder
	b.WriteString("(")
	//r e := range s {
	for e, ok := s.Head(); ok; e, ok = s.Head() {
		b.WriteString(e.String())
	}
	b.WriteString(")")
	return b.String()
}
