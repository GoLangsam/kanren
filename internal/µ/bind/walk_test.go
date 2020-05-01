package bind

import (
	"testing"
)

func TestResolve(t *testing.T) {
	zaxwyz := New()
	zaxwyz.Bind(zaxwyz.Fresh("z"), NewSymbol("a"))
	zaxwyz.Bind(zaxwyz.Fresh("x"), zaxwyz.Fresh("w"))
	zaxwyz.Bind(zaxwyz.Fresh("y"), zaxwyz.Fresh("z"))

	xyvxwx := New()
	xyvxwx.Bind(xyvxwx.Fresh("x"), xyvxwx.Fresh("y"))
	xyvxwx.Bind(xyvxwx.Fresh("v"), xyvxwx.Fresh("x"))
	xyvxwx.Bind(xyvxwx.Fresh("w"), xyvxwx.Fresh("x"))

	xbzywl := New()
	xbzywl.Bind(xbzywl.Fresh("x"), NewSymbol("b"))
	xbzywl.Bind(xbzywl.Fresh("z"), xbzywl.Fresh("y"))
	xbzywl.Bind(xbzywl.Fresh("w"), NewList(xbzywl.Fresh("x"), NewSymbol("e"), xbzywl.Fresh("z")))

	xezxyz := New()
	xezxyz.Bind(xezxyz.Fresh("x"), NewSymbol("e"))
	xezxyz.Bind(xezxyz.Fresh("z"), xezxyz.Fresh("x"))
	xezxyz.Bind(xezxyz.Fresh("y"), xezxyz.Fresh("z"))

	tests := []struct {
		u    string
		s    Ings
		want string
	}{
		{u: "z", s: zaxwyz, want: "a"},
		{u: "y", s: zaxwyz, want: "a"},
		{u: "x", s: zaxwyz, want: ",w"},
		{u: "x", s: xyvxwx, want: ",y"},
		{u: "v", s: xyvxwx, want: ",y"},
		{u: "w", s: xyvxwx, want: ",y"},
		{u: "w", s: xbzywl, want: "(,x e ,z)"},
		{u: "y", s: xezxyz, want: "e"},
	}
	for _, test := range tests {
		q, subs, want := test.u, test.s, test.want
		t.Run("(walk "+q+" "+subs.String()+")", func(t *testing.T) {
			v := subs.Fresh(q)
			got := subs.Resolve(v).String()
			if want != got {
				t.Fatalf("got %s want %s", got, want)
			}
		})
	}
}

func TestWalk(t *testing.T) {
	zaxwyz := New()
	zaxwyz.Bind(zaxwyz.Fresh("z"), NewSymbol("a"))
	zaxwyz.Bind(zaxwyz.Fresh("x"), zaxwyz.Fresh("w"))
	zaxwyz.Bind(zaxwyz.Fresh("y"), zaxwyz.Fresh("z"))

	xyvxwx := New()
	xyvxwx.Bind(xyvxwx.Fresh("x"), xyvxwx.Fresh("y"))
	xyvxwx.Bind(xyvxwx.Fresh("v"), xyvxwx.Fresh("x"))
	xyvxwx.Bind(xyvxwx.Fresh("w"), xyvxwx.Fresh("x"))

	xbzywl := New()
	xbzywl.Bind(xbzywl.Fresh("x"), NewSymbol("b"))
	xbzywl.Bind(xbzywl.Fresh("z"), xbzywl.Fresh("y"))
	xbzywl.Bind(xbzywl.Fresh("w"), NewList(xbzywl.Fresh("x"), NewSymbol("e"), xbzywl.Fresh("z")))

	xezxyz := New()
	xezxyz.Bind(xezxyz.Fresh("x"), NewSymbol("e"))
	xezxyz.Bind(xezxyz.Fresh("z"), xezxyz.Fresh("x"))
	xezxyz.Bind(xezxyz.Fresh("y"), xezxyz.Fresh("z"))

	tests := []struct {
		u    string
		s    Ings
		want string
	}{
		{u: "z", s: zaxwyz, want: "a"},
		{u: "y", s: zaxwyz, want: "a"},
		{u: "x", s: zaxwyz, want: ",w"},
		{u: "x", s: xyvxwx, want: ",y"},
		{u: "v", s: xyvxwx, want: ",y"},
		{u: "w", s: xyvxwx, want: ",y"},
		{u: "w", s: xbzywl, want: "(b e ,y)"},
		{u: "y", s: xezxyz, want: "e"},
	}
	for _, test := range tests {
		q, subs, want := test.u, test.s, test.want
		t.Run("(walk "+q+" "+subs.String()+")", func(t *testing.T) {
			v := subs.Fresh(q)
			got := subs.Walk(v).String()
			if want != got {
				t.Fatalf("got %s want %s", got, want)
			}
		})
	}
}
