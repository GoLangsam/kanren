// package reif provides a reif.Ier(): reif.Y
package reif

import "github.com/GoLangsam/kanren/internal/µ/bind"
import "github.com/GoLangsam/kanren/internal/µ/vari"

// Y extends vari.Ables and their bind.Ings
// with Reify,
// which needs the ability to construct fresh variables on-the-fly.
//
// Use as `reif.Y` (pun intended).
//
// The zero value is not useful - initialize with `reif.Ier()`.
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

// Clone provides a clone of y.
func (y *Y) Clone() *Y {
	return &Y{
		Able:  y.Able,
		Ings:  y.Ings.Clone(),
		count: y.count,
	}
}

func (y *Y) BindFresh(u V) *Y {
	yy := y.Clone()

	v := yy.Fresh(yy.nextName())
	yy.Bind(u, v)
	return yy
}

// Reify ...
func (y *Y) Reify(x X) *Y {
	s := y
	x = s.Walk(x)
	switch {
	case x.IsVariable(): // bind u(=xx) to new fresh var
		return s.BindFresh(x) // y.Bind(u, v.Expr())
	case x.IsPair():
		return s.Reify(x.Car()).Reify(x.Cdr())
	}
	return s
}
