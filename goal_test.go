package kanren

import (
	//	"strings"
	"testing"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		u, v, want string
	}{
		{u: "#f", v: "#f", want: "(())"},
		{u: "#f", v: "#t", want: "()"},
	}
	for _, test := range tests {
		u, v, want := test.u, test.v, test.want
		t.Run("((== "+u+" "+v+") empty-s)", func(t *testing.T) {
			uexpr, err := Parse(u)
			if err != nil {
				t.Fatal(err)
			}
			vexpr, err := Parse(v)
			if err != nil {
				t.Fatal(err)
			}
			stream := Equal(uexpr, vexpr).Try()
			got := stream.String()
			if got != want {
				t.Fatalf("got %s want %s", got, want)
			}
		})
	}
}

/*
func TestNever(t *testing.T) {
	e := NewS()
	n := Never()(e)
	s, sok := n.Head()
	if s != nil {
		t.Fatalf("expected suspension")
	}
	if sok {
		t.Fatalf("expected immediate ending")
	}
}
*/

func TestDisjointO(t *testing.T) {
	e := NewS()
	x := e.Fresh("x")

	d := Disjoint(
		Equal(
			NewSymbol("olive"),
			x,
		),
		FAIL,
	)(e)
	s, sok := d.Head()
	got := s.String()
	want := "((,x . olive))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
	_ = sok

	s, sok = d.Head()
	if sok {
		t.Fatalf("expected never ending")
	}
}

func TestDisjoint2(t *testing.T) {
	e := NewS()
	x := e.Fresh("x")

	d := Disjoint(
		FAIL,
		Equal(
			NewSymbol("olive"),
			x,
		),
	)(e)
	s, sok := d.Head()
	got := s.String()
	want := "((,x . olive))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
	_ = sok
	s, sok = d.Head()
	if sok {
		t.Fatalf("expected only one")
	}
	s, sok = d.Head()
	if sok {
		t.Fatalf("expected no more")
	}
}

/*
func TestAlways(t *testing.T) {
	a := Always()(NewS())
	s, sok := a.Head()
	got := s.String()
	want := "(())"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
	s, sok = a.Head()
	if sok {
		t.Fatalf("expected never ending")
	}
}
*/

/* TODO: ap
func TestRunGoalAlways3(t *testing.T) {
	ss := Always()(NewS())
	if len(ss) != 3 {
		t.Fatalf("expected 3 got %d", len(ss))
	}
	sss := deriveFmaps(func(s *State) string {
		return s.String()
	}, ss)
	got := "(" + strings.Join(sss, " ") + ")"
	want := "((() . 0) (() . 0) (() . 0))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
*/

func TestRunGoalDisj2(t *testing.T) {
	e := NewS()
	x := e.Fresh("x")

	e1 := Equal(
		NewSymbol("olive"),
		x,
	)
	e2 := Equal(
		NewSymbol("oil"),
		x,
	)
	g := Disjoint(e1, e2)
	ss := g(e)
	l := 0
	for head, ok := ss.Head(); ok; head, ok = ss.Head() {
		_ = head
		l++
	}
	if l != 2 {
		t.Fatalf("expected 2, got %d", l)
	}
}

func TestRunGoalConj2NoResults(t *testing.T) {
	e := NewS()
	x := e.Fresh("x")

	e1 := Equal(
		NewSymbol("olive"),
		x,
	)
	e2 := Equal(
		NewSymbol("oil"),
		x,
	)
	g := Conjunction(e1, e2)
	ss := g(e)
	l := 0
	for _, ok := ss.Head(); ok; _, ok = ss.Head() {
		l++
	}
	if l != 0 {
		t.Fatalf("expected 0, got %d: %v", l, ss)
	}
}

func TestRunGoalConj2OneResults(t *testing.T) {
	e := NewS()
	x := e.Fresh("x")

	e1 := Equal(
		NewSymbol("olive"),
		x,
	)
	e2 := Equal(
		NewSymbol("olive"),
		x,
	)
	g := Conjunction(e1, e2)
	ss := g(e)
	first, _ := ss.Receive()
	got := first.String()
	want := "((,x . olive))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
	l := 0
	for _, ok := ss.Head(); ok; _, ok = ss.Head() {
		l++
		t.Fatalf("expected none, got %d: %v", l, ss)
	}
}
