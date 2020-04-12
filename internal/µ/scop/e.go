// package scop provides scop.Es(): scop.E
package scop

import "github.com/GoLangsam/kanren/internal/µ/bind"
import "github.com/GoLangsam/kanren/internal/µ/vari"

// E extends vari.Ables and their bind.Ings
// with Scope.
//
// Use as `scop.E` (pun intended).
//
// The zero value is not useful - initialize with `scop.Es()`.
type E struct {
	*vari.Able            // known variables
	*bind.Ings            // their current bindings
	up         *bind.Ings // outer bindings
}

// Ier creates fresh and empty reifier.
func Ier() *E {
	return &E{
		Able: vari.Ables(),
		Ings: bind.New(),
	}
}

// Extend provides an empty bind.Ings below the current one.
func (e *E) Extend() *E {
	return &E{
		Able: e.Able,
		Ings: bind.New(),
		up:   e.Ings,
	}
}
