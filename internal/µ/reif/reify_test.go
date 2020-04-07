package reif

import (
	"fmt"
	"testing"
)

import "github.com/GoLangsam/sexpr"

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
	e, err := sexpr.Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	if !e.IsPair() {
		t.Fatalf("expected list")
	}

	xv := sexpr.NewVariable("x")
	x, _ := xv.AsVariable()
	y, _ := sexpr.NewVariable("y").AsVariable()
	w, _ := sexpr.NewVariable("w").AsVariable()

	ss := Ier()
	ss.Bind(x, e.Car().Cdr())
	ss.Bind(y, e.Cdr().Car().Cdr())
	ss.Bind(w, e.Cdr().Cdr().Car().Cdr())

	gote := ss.Reify(xv)
	got := fmt.Sprintf("%v", gote)
	want := "(_0 (_1 _0) corn _2 ((ice) _2))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
