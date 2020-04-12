package kanren

import "testing"

import "github.com/GoLangsam/sexpr"

func TestFreshKiwi(t *testing.T) {
	cf := CallFresh(func(fruit X) Goal {
		return Equal(
			sexpr.NewSymbol("plum"),
			fruit,
		)
	},
	)
	ss := cf(EmptyState())
	want := "(((,~.0 . plum)))"
	got := ss.String()
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}