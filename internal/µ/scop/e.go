// package scop provides scop.Es(): scop.E
package scop

import "github.com/GoLangsam/kanren/internal/µ/bind"

// E extends bind.Ings with Scope.
//
// Use as `scop.E` (pun intended).
//
// The zero value is not useful - initialize with `scop.Es()`.
type E struct {
	outer      *E // outer environment
	*bind.Ings    // current bindings
}

// Es creates fresh and empty scop.Es (pun intended).
func Es() *E {
	return &E{
		// outer: nil,
		Ings: bind.New(),
	}
}

// Extend provides an empty environment below the current one.
func (e *E) Extend() *E {
	return &E{
		outer: e,
		Ings:  bind.New(),
	}
}

func (e *E) Bound(v V) (value X, isBound bool) {
	value, isBound = e.Bound(v)
	if !isBound && e.outer != nil {
		value, isBound = e.outer.Bound(v)
	}
	return
}
