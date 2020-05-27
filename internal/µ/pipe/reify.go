package pipe

import "github.com/GoLangsam/sexpr"

// =============================================================================

// V is a symbolic expression which represents nothing but a logic variable.
type V = *sexpr.Expression

// X represents a symbolic eXpression
type X = *sexpr.Expression

// =============================================================================

// Reify returns a channel to receive the reifications of v along s.
func (ss StreamOfStates) Reify(v V) chan X {
	cha := make(chan X)
	go func() {
		//r s := range ss {
		for s, ok := ss.Head(); ok; s, ok = ss.Head() {
			// println(v.String(), "=?", s.String())
			r := s.Reify(v)
			// println(v.String(), "=>", r.String())
			cha <- r
		}
		close(cha)
	}()

	return cha
}
