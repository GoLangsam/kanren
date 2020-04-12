package kanren

import "testing"

import "github.com/GoLangsam/sexpr"

/*
scheme code:

	(run-goal 1
		(call/fresh kiwi
			(lambda (fruit)
				(== plum fruit)
			)
		)
	)
*/
func TestFreshKiwi(t *testing.T) {
	cf := CallFresh(func(fruit X) Goal {
		return Equal(
			sexpr.NewSymbol("plum"),
			fruit,
		)
	},
	)
	ss := cf(EmptyState())
	want := "((,v0 . plum) . 1)"
	got := ss.String()
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
