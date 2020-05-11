package kanren

import (
	"testing"
)

/*
(let
	(
		(q (var 'q))
	)
	(map
		(reify q)
		(run-goal #f
			(call/fresh 'x
				(lambda (x)
					(call/fresh 'y
						(lambda (y)
							(conj
								(== `(,x ,y) q)
								(appendo x y `(cake & ice d t))
							)
						)
					)
				)
			)
		)
	)
)
*/

/*
 */

// results in all the combinations of two lists that when appended will result in (cake & ice d t)
func TestAppendAllCombinations(t *testing.T) {
	q := NewVariable("q")
	goal := Fresh2(func(x, y X) Goal {
		return Conjunction(
			Equal(
				//					cons(x, y),
				//					cons(x, cons(y, nil)),
				NewList(x, y),
				q,
			),
			Append(
				x,
				y,
				NewList(
					NewSymbol("cake"),
					NewSymbol("&"),
					NewSymbol("ice"),
					NewSymbol("d"),
					NewSymbol("t"),
				),
			),
		)
	})

	states := goal.Try()
	sexprs := states.Reify(q)
	exprS := []X{}

	for x := range sexprs {
		exprS = append(exprS, x)
		println("X=", x.String())
	}

	got := NewList(exprS...).String()
	want := "((() (cake & ice d t)) ((cake) (& ice d t)) ((cake &) (ice d t)) ((cake & ice) (d t)) ((cake & ice d) (t)) ((cake & ice d t) ()))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
	// Output:
	// X= (() (cake & ice d t))
	// X= ((cake) (& ice d t))
	// X= ((cake &) (ice d t))
	// X= ((cake & ice) (d t))
	// X= ((cake & ice d) (t))
	// X= ((cake & ice d t) ())
}

func TestAppendSingleList(t *testing.T) {
	q := NewVariable("q")
	goal := Append(
		cons(NewSymbol("a"), nil),
		cons(NewSymbol("b"), nil),
		q,
	)
	states := goal.Try()
	sexprs := states.Reify(q)
	exprS := []X{}

	for x := range sexprs {
		exprS = append(exprS, x)
		if len(exprS) > 3 {
			break
		}
	}

	got := NewList(exprS...).String()
	want := "((a b))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}

func TestAppendSingleAtom(t *testing.T) {
	q := NewVariable("q")
	goal := Append(
		cons(NewSymbol("a"), nil),
		NewSymbol("b"),
		q,
	)
	states := goal.Try()
	sexprs := states.Reify(q)
	exprS := []X{}

	for x := range sexprs {
		exprS = append(exprS, x)
		// if len(exprS) > 3 { break }
		// println("X=", x.String())
	}

	got := NewList(exprS...).String()
	want := "((a . b))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}

func TestAppendSingleAtoms(t *testing.T) {
	q := NewVariable("q")
	goal := Append(
		NewSymbol("a"),
		NewSymbol("b"),
		q,
	)
	states := goal.Try()
	sexprs := states.Reify(q)
	exprS := []X{}

	for x := range sexprs {
		exprS = append(exprS, x)
		// if len(exprS) > 3 { break }
		// println("X=", x.String())
	}

	got := NewList(exprS...).String()
	want := "()"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}

func TestCar(t *testing.T) {
	goal := IfThenElse(
		Car(
			NewList(
				NewSymbol("a"),
				NewSymbol("c"),
				NewSymbol("o"),
				NewSymbol("r"),
				NewSymbol("n"),
			),
			NewSymbol("a"),
		),
		Equal(NewSymbol("#t"), NewVariable("y")),
		Equal(NewSymbol("#f"), NewVariable("y")),
	)
	ss := goal.Try()
	got := ss.String()
	//nt := "(((,y . #t) (,v0 c o r n) . 1))" // WRONG
	want := "(((,y . #t)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}

func TestCdr(t *testing.T) {
	goal := IfThenElse(
		Cdr(
			NewList(
				NewSymbol("a"),
				NewSymbol("c"),
				NewSymbol("o"),
				NewSymbol("r"),
				NewSymbol("n"),
			),
			NewList(
				NewSymbol("c"),
				NewSymbol("o"),
				NewSymbol("r"),
				NewSymbol("n"),
			),
		),
		Equal(NewSymbol("#t"), NewVariable("y")),
		Equal(NewSymbol("#f"), NewVariable("y")),
	)
	ss := goal.Try()
	got := ss.String()
	//nt := "(((,y . #t) (,v0 c o r n) . 1))" // WRONG
	want := "(((,y . #t)))"
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}
