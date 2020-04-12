package kanren

import (
	"testing"
)

import "github.com/GoLangsam/sexpr"

func TestIfThenElseSuccess(t *testing.T) {
	s := EmptyState()
	y := s.Fresh("y")

	ifte := IfThenElse(
		Success(),
		Equal(sexpr.NewSymbol("#f"), y),
		Equal(sexpr.NewSymbol("#t"), y),
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
		Equal(sexpr.NewSymbol("#f"), y),
		Equal(sexpr.NewSymbol("#t"), y),
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
		Equal(sexpr.NewSymbol("#t"), x),
		Equal(sexpr.NewSymbol("#f"), y),
		Equal(sexpr.NewSymbol("#t"), y),
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
			Equal(sexpr.NewSymbol("#t"), x),
			Equal(sexpr.NewSymbol("#f"), x),
		),
		Equal(sexpr.NewSymbol("#f"), y),
		Equal(sexpr.NewSymbol("#t"), y),
	)
	ss := ifte(s)
	got := ss.String()
	want := "(((,y . #f) (,x . #t)) ((,y . #f) (,x . #f)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}
