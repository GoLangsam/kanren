package µ

import "fmt"
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
	ss := RunGoal(1,
		CallFresh(func(fruit V) Goal {
			return Equal(
				sexpr.NewSymbol("plum"),
				fruit.Expr(),
			)
		},
		),
	)
	if len(ss) != 1 {
		t.Fatalf("expected %d, but got %d results", 1, len(ss))
	}
	want := "((,v0 . plum) . 1)"
	//got := ss[0].String()
	got := fmt.Sprintf("%v", ss[0])
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
