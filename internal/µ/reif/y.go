// package reif provides a reif.Ier(): reif.Y
package reif

import "github.com/GoLangsam/kanren/internal/µ/bind"
import "github.com/GoLangsam/kanren/internal/µ/vari"

// Y extends variables and their bindings
// with Reify,
// which needs the ability to construct fresh variables on-the-fly.
//
// Use as `reif.Y` (pun intended).
//
// The zero value is not useful - initialize with `reif.Yer()`.
type Y struct {
	*vari.Able     // known variables
	*bind.Ings     // their current bindings
	count      int // used to create anonymus variables during reify
}

// Ier creates fresh and empty reifier.
func Ier() *Y {
	return &Y{
		Able: vari.Ables(),
		Ings: bind.New(),
	}
}

// Clone provides a clone of b.
// TODO: This is bullshit
func (y *Y) Clone() *Y {
	return &Y{
		Able: y.Able,
		Ings: y.Ings,
	}
}

// Reify ...
func (y *Y) Reify(x X) *Y {
	xx := y.Walk(x)
	u, isVariable := xx.AsVariable()
	switch {
	case isVariable: // bind u(=xx) to new fresh var
		v := y.Fresh(y.nextName())
		y.Bind(u, v.Expr())
	case xx.IsPair():
		y.Reify(xx.Car()).Reify(xx.Cdr())
	}
	return y
}
