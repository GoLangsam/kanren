package pipe

import "fmt"
import "strings"

// String blocks until stream is closed,
// consumes the stream
// and returns its strings representation.
func (s StreamOfStates) String() string {
	var b strings.Builder
	b.WriteString("(")
	for e, ok := s.Head(); ok; e, ok = s.Head() {
		//r e := range s {
		b.WriteString(e.String())
	}
	s.Drop()
	b.WriteString(")")
	return b.String()
}

// Printn prints up to n terms of a power series.
func (s StreamOfStates) Printn(n int) {
	defer fmt.Print("\n")

	for ; n > 0; n-- {
		ss, ok := s.Receive()
		if !ok {
			break
		}
		fmt.Print(ss.String())
		fmt.Print(" ")
	}
	s.Drop()
}
