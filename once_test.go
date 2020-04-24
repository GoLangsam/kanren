package kanren

import (
	"testing"
)

func TestOnce(t *testing.T) {
	e := NewS()
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
	ss := ifte(e)
	got := ss.String()
	//nt := "(((,x . #t) (,y . #f)))" WRONG!
	want := "(((,y . #f)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}
