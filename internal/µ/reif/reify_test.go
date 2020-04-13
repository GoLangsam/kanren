package reif

import (
	"fmt"
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
	fmt.Printf("%v\n", e)
	ss := Ier()
	x := ss.Fresh("x")
	y := ss.Fresh("y")
	w := ss.Fresh("w")
	ss.Bind(x, e.Car().Cdr())
	ss.Bind(y, e.Cdr().Car().Cdr())
	ss.Bind(w, e.Cdr().Cdr().Car().Cdr())
	fmt.Println("ss:", ss)
	gote := ss.Reify(x)
	got := fmt.Sprintf("%v", gote)
	want := "(_0 (_1 _0) corn _2 ((ice) _2))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
