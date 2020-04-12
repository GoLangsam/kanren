package kanren

import (
	"testing"
)

import micro "github.com/GoLangsam/kanren/internal/µ"
import "github.com/GoLangsam/sexpr"

func TestOnce(t *testing.T) {
	ifte := IfThenElse(
		Once(Disjoint(
			Equal(sexpr.NewSymbol("#t"), sexpr.NewVariable("x")),
			Equal(sexpr.NewSymbol("#f"), sexpr.NewVariable("x")),
		)),
		Equal(sexpr.NewSymbol("#f"), sexpr.NewVariable("y")),
		Equal(sexpr.NewSymbol("#t"), sexpr.NewVariable("y")),
	)
	ss := ifte(micro.EmptyState())
	got := ss.String()
	want := "(((,y . #f) (,x . #t)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}
