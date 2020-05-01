package kanren

import "testing"

func TestFreshKiwi(t *testing.T) {
	rel1 := func(fruit V) Goal {
		return Equal(
			NewSymbol("plum"),
			fruit,
		)
	}
	got := CallFresh(rel1).Try().String()
	want := "(((,_0 . plum)))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
