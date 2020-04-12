package kanren

import (
	"testing"
)

import micro "github.com/GoLangsam/kanren/internal/µ"
import "github.com/GoLangsam/sexpr"

func TestIfThenElseSuccess(t *testing.T) {
	ifte := IfThenElse(
		Success(),
		Equal(sexpr.NewSymbol("#f"), sexpr.NewVariable("y")),
		Equal(sexpr.NewSymbol("#t"), sexpr.NewVariable("y")),
	)
	ss := ifte(micro.EmptyState())
	got := ss.String()
	want := "(((,y . #f) . 0))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}

func TestIfThenElseFailure(t *testing.T) {
	ifte := IfThenElse(
		Failure(),
		Equal(sexpr.NewSymbol("#f"), sexpr.NewVariable("y")),
		Equal(sexpr.NewSymbol("#t"), sexpr.NewVariable("y")),
	)
	ss := ifte(micro.EmptyState())
	got := ss.String()
	want := "(((,y . #t) . 0))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}

func TestIfThenElseXIsTrue(t *testing.T) {
	ifte := IfThenElse(
		Equal(sexpr.NewSymbol("#t"), sexpr.NewVariable("x")),
		Equal(sexpr.NewSymbol("#f"), sexpr.NewVariable("y")),
		Equal(sexpr.NewSymbol("#t"), sexpr.NewVariable("y")),
	)
	ss := ifte(micro.EmptyState())
	got := ss.String()
	want := "(((,y . #f) (,x . #t) . 0))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}

func TestIfThenElseDisjoint(t *testing.T) {
	ifte := IfThenElse(
		Disjoint(
			Equal(sexpr.NewSymbol("#t"), sexpr.NewVariable("x")),
			Equal(sexpr.NewSymbol("#f"), sexpr.NewVariable("x")),
		),
		Equal(sexpr.NewSymbol("#f"), sexpr.NewVariable("y")),
		Equal(sexpr.NewSymbol("#t"), sexpr.NewVariable("y")),
	)
	ss := ifte(micro.EmptyState())
	got := ss.String()
	want := "(((,y . #f) (,x . #t) . 0) ((,y . #f) (,x . #f) . 0))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}
