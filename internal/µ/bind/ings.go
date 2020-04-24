package bind

import "github.com/GoLangsam/kanren/internal/µ/smap"

// SMap documents the behaviour required from a substitution map.
type SMap interface {
	Clone() smap.SMap
	String() string

	Load(V) (X, bool)
	Store(V, X)

	Delete(X) // only for tests - TODO: remove
}

var _ SMap = smap.New() // assert wellbehaviour

// Ings represents bindings (or "substitutions" or ""):
// any logic variable may be bound to some symbolic expression
// representing its current value.
//
// Use as `bind.Ings` (pun intended).
//
// The zero value is not useful - initialize with `bind.New()`.
type Ings struct {
	smap.SMap // SMap
}

// New creates fresh and empty mapping of/for bind.Ings and returns a pointer.
func New() Ings {
	return Ings{
		smap.New(),
	}
}

// Clone returns a shallow copy.
func (bind Ings) Clone() Ings {
	return Ings{
		bind.SMap.Clone(),
	}
}

// String returns a string of the symbolic map sorted by key.
func (bind Ings) String() string {
	return bind.SMap.String()
}

// =============================================================================

// Bind binds x to v, so v is bound to x.
// Thus, (v . x) resembles a substitution pair.
//
// Bind is a noOp if v or x are nil or v is not a Variable.
//
// Note: Bind does not attempt to avoid circular bindings.
// Use Occurs to check beforehand.
func (bind Ings) Bind(v V, x X) {
	if v.IsVariable() && x != nil {
		bind.Store(v, x)
	}
	return
}

// =============================================================================

// Resolve the eXpression
// by chasing along the bindings recurring
// down to the first non-Variable eXpression
// or down to the first unbound eXpression
func (bind Ings) Resolve(v X) X {
	if !v.IsVariable() {
		return v
	}
	x, isBound := bind.Load(v)
	if !isBound {
		return v
	}
	return bind.Resolve(x)
}

// Walk ... some call it `walkstar` or `walk*`
func (bind Ings) Walk(x X) X {
	x = bind.Resolve(x)
	if !x.IsPair() {
		return x
	}
	return cons(
		bind.Walk(x.Pair.Car),
		bind.Walk(x.Pair.Cdr),
	)
}

// Occurs reports whether v occurs in x.
func (bind Ings) Occurs(v V, x X) bool {
	x = bind.Resolve(x)
	u, isVariable := x.AsVariable()
	if isVariable {
		return u.Equal(v.Atom.Var)
	}
	if !x.IsPair() {
		return false
	}
	return bind.Occurs(v, x.Car()) || bind.Occurs(v, x.Cdr())
}

// exts attempts to bind v with x and fails
// if such bind would introduce a circle
// due to v occurs in x.
func (bind Ings) SafeBind(v V, x X) bool {
	if bind.Occurs(v, x) {
		return false
	}
	bind.Bind(v, x)
	return true
}

// Unify extends the bind.Ings with zero or more associations
// in an attempt to see whether the given eXpressions are equal
// and reports its success.
// Circular bindings imply failure (ok = false).
func (bind Ings) Unify(x, y X) bool {
	x = bind.Resolve(x)
	y = bind.Resolve(y)

	vx, isXxVariable := x.AsVariable()
	vy, isYyVariable := y.AsVariable()

	if isXxVariable && isYyVariable && vx.Equal(vy) {
		return true
	}

	if isXxVariable {
		return bind.SafeBind(x, y)
	}
	if isYyVariable {
		return bind.SafeBind(y, x)
	}

	if x.IsPair() && y.IsPair() {
		uniCars := bind.Unify(x.Car(), y.Car())
		if uniCars {
			return bind.Unify(x.Cdr(), y.Cdr())
		}
		return false
	}
	return x.Equal(y)
}
