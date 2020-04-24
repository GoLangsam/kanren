package kanren

import "testing"

func TestFreshKiwi(t *testing.T) {
	rel1 := func(fruit V) Goal {
		return Equal(
			NewSymbol("plum"),
			fruit,
		)
	}
	got := Fresh1(rel1).Try().String()
	want := "(((,~.0 . plum)))"
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
