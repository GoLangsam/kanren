package kanren

import (
	"testing"
)

func TestOnce(t *testing.T) {
	e := EmptyState()
	x := e.Fresh("x")
	y := e.Fresh("y")

	ifte := IfThenElse(
		Once(Disjoint(
			Equal(NewSymbol("#t"), x),
			Equal(NewSymbol("#f"), x),
		)),
		Equal(NewSymbol("#f"), y),
		Equal(NewSymbol("#t"), y),
	)
	ss := ifte(EmptyState())
	got := ss.String()
	want := "(((,y . #f) (,x . #t)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}
