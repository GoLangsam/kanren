package µ

import (
	//	"strings"
	"testing"
)

func TestFailure(t *testing.T) {
	if got, want := Failure()(NewS()).String(), "()"; got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}

func TestSuccess(t *testing.T) {
	if got, want := Success()(NewS()).String(), "(())"; got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
