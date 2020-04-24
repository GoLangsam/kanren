package kanren

import (
	"testing"
)

func TestIfThenElseSuccess(t *testing.T) {
	s := NewS()
	y := s.Fresh("y")

	ifte := IfThenElse(
		GOAL,
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
	s := NewS()
	y := s.Fresh("y")

	ifte := IfThenElse(
		FAIL,
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
	s := NewS()
	x := s.Fresh("x")
	y := s.Fresh("y")

	ifte := IfThenElse(
		Equal(NewSymbol("#t"), x),
		Equal(NewSymbol("#t"), y),
		Equal(NewSymbol("#f"), y),
	)
	ss := ifte(s)
	got := ss.String()
	//nt := "(((,x . #t) (,y . #f)))" // WRONG: values of IF are irrelevant
	want := "(((,y . #t)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}

func TestIfThenElseDisjoint(t *testing.T) {
	s := NewS()
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
	//nt := "(((,y . #f) (,x . #t)) ((,y . #f) (,x . #f)))" // WRONG: values of IF are irrelevant
	want := "(((,y . #f)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}

func TestEitherOrDisjoint(t *testing.T) {
	s := NewS()
	x := s.Fresh("x")
	y := s.Fresh("y")

	ifte := EitherOr(
		Disjoint(
			Equal(NewSymbol("#t"), x),
			Equal(NewSymbol("#f"), x),
		),
		Equal(NewSymbol("#f"), y),
	)
	ss := ifte(s)
	got := ss.String()
	//nt := "(((,y . #f) (,x . #t)) ((,y . #f) (,x . #f)))" // TODO: this is wrong; values of IF are irrelevant
	want := "(((,x . #t))((,x . #f)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}
