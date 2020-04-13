package kanren

import (
	"testing"
)

func TestIfThenElseSuccess(t *testing.T) {
	s := EmptyState()
	y := s.Fresh("y")

	ifte := IfThenElse(
		Success(),
		Equal(NewSymbol("#f"), y),
		Equal(NewSymbol("#t"), y),
	)
	ss := ifte(s)
	got := ss.String()
	want := "(((,y . #f)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}

func TestIfThenElseFailure(t *testing.T) {
	s := EmptyState()
	y := s.Fresh("y")

	ifte := IfThenElse(
		Failure(),
		Equal(NewSymbol("#f"), y),
		Equal(NewSymbol("#t"), y),
	)
	ss := ifte(s)
	got := ss.String()
	want := "(((,y . #t)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}

func TestIfThenElseXIsTrue(t *testing.T) {
	s := EmptyState()
	x := s.Fresh("x")
	y := s.Fresh("y")

	ifte := IfThenElse(
		Equal(NewSymbol("#t"), x),
		Equal(NewSymbol("#f"), y),
		Equal(NewSymbol("#t"), y),
	)
	ss := ifte(s)
	got := ss.String()
	want := "(((,x . #t) (,y . #f)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}

func TestIfThenElseDisjoint(t *testing.T) {
	s := EmptyState()
	x := s.Fresh("x")
	y := s.Fresh("y")

	ifte := IfThenElse(
		Disjoint(
			Equal(NewSymbol("#t"), x),
			Equal(NewSymbol("#f"), x),
		),
		Equal(NewSymbol("#f"), y),
		Equal(NewSymbol("#t"), y),
	)
	ss := ifte(s)
	got := ss.String()
	want := "(((,y . #f) (,x . #t)) ((,y . #f) (,x . #f)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}
