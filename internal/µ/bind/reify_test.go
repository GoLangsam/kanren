package bind

import (
	"fmt"
	//	"strings"
	"testing"
)

/*
scheme code:

	(let
		(
			(a1 `(,x . (,u ,w ,y ,z ((ice) ,z))))
			(a2 `(,y . corn))
			(a3 `(,w . (,v ,u)))
		)
		(let
			(
				(s `(,a1 ,a2 ,a3))
			)
			(
				(reify x) s
			)
		)
	)
*/
func TestReify(t *testing.T) {
	a1 := "(,x . (,u ,w ,y ,z ((ice) ,z)))"
	a2 := "(,y . corn)"
	a3 := "(,w . (,v ,u))"
	s := "(" + a1 + " " + a2 + " " + a3 + ")"
	e, err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	if !e.IsPair() {
		t.Fatalf("expected list")
	}

	ss := New()
	x := ss.Fresh("x")
	y := ss.Fresh("y")
	w := ss.Fresh("w")
	ss.Bind(x, e.Car().Cdr())
	ss.Bind(y, e.Cdr().Car().Cdr())
	ss.Bind(w, e.Cdr().Cdr().Car().Cdr())

	gote := ss.Reify(x)
	got := fmt.Sprintf("%v", gote)
	want := "(,_0 (,_1 ,_0) corn ,_2 ((ice) ,_2))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}

/*
func TestNoReify(t *testing.T) {
	e1 := Equal(
		sexpr.NewSymbol("olive"),
		sexpr.NewVariable("x"),
	)
	e2 := Equal(
		sexpr.NewSymbol("oil"),
		sexpr.NewVariable("x"),
	)
	g := Disjoint(e1, e2)
	states := RunGoal(5, g)
	ss := make([]*sexpr.Expression, len(states))
	strs := make([]string, len(states))

	x := sexpr.NewVariable("x")
	r := reifyVarFromState(x)
	for i, s := range states {
		ss[i] = r(s)
		strs[i] = ss[i].String()
	}
	got := "(" + strings.Join(strs, " ") + ")"
	want := "(olive oil)"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
*/
