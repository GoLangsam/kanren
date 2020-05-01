package bind

import "github.com/GoLangsam/sexpr"

// =============================================================================

// V is a symbolic expression which represents nothing but a logic variable.
type V = *sexpr.Expression

// X represents a symbolic eXpression
type X = *sexpr.Expression

// =============================================================================

// for sMap
type key = string
type value = X

// isSMap documents the behaviour required from a substitution map.
type isSMap interface {
	Clone() sMap
	String() string

	Load(key) (value, bool)
	Store(key, value)
}

var _ isSMap = newMap() // assert wellbehaviour

// =============================================================================

// Ings represents bindings (or "substitutions" or "assignments"):
// Any logic variable can be bound to some symbolic expression -
// its current value.
//
// Use as `bind.Ings` (pun intended).
//
// The zero value is not useful - initialize with `bind.New()`.
type Ings struct {
	sMap sMap     // isSMap
	V    func() V // Factory for reified variables - returns a fresh variable with an autogenerated reified name.
}

// New returns a fresh and empty mapping of/for bind.Ings.
func New() Ings {
	return Ings{
		newMap(),
		Paginator(),
	}
}

// =============================================================================

// Clone returns a shallow copy.
func (bind Ings) Clone() Ings {
	return Ings{
		bind.sMap.Clone(),
		bind.V,
	}
}

// String returns a string of the symbolic map sorted by key.
func (bind Ings) String() string {
	return bind.sMap.String()
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
		bind.sMap.Store(v.Atom.Var.Name, x)
	}
	return
}

// =============================================================================

// Resolve the eXpression
// by chasing along the bindings recurring
// down to the first non-Variable eXpression
// or down to the first unbound eXpression
func (bind Ings) Resolve(x X) X {
	if !x.IsVariable() {
		return x
	}
	return bind.resolve(x)
}

// resolve is the recursive body
// where v.IsVariable is true.
func (bind Ings) resolve(v V) X {
	x, isBound := bind.sMap.Load(v.Atom.Var.Name)
	if !isBound {
		return v
	}
	if !x.IsVariable() {
		return x
	}
	return bind.resolve(x)
}

// =============================================================================

// Walk ... some call it `walkstar` or `walk*`
func (bind Ings) Walk(v V) X {
	x := bind.Resolve(v)
	if !x.IsPair() {
		return x
	}
	return sexpr.Cons(
		bind.Walk(x.Pair.Car),
		bind.Walk(x.Pair.Cdr),
	)
}

// =============================================================================

// Occurs reports whether v occurs in x.
func (bind Ings) Occurs(v V, x X) bool {
	x = bind.Resolve(x)

	switch {
	case v.IsVariable() && x.IsVariable():
		return v.Atom.Var.Name == x.Atom.Var.Name

	case !x.IsPair():
		return false
	}

	return bind.Occurs(v, x.Pair.Car) || bind.Occurs(v, x.Pair.Cdr)
}

// =============================================================================

// SafeBind attempts to bind v with x and fails
// if such bind would introduce a circle
// due to v occurs in x.
func (bind Ings) SafeBind(v V, x X) bool {
	if bind.Occurs(v, x) {
		return false
	}
	bind.Bind(v, x)
	return true
}

// =============================================================================

// Unify extends the bind.Ings with zero or more associations
// in an attempt to see whether the given eXpressions are equal
// and reports its success.
// Circular bindings imply failure (ok = false).
//
// Note: For improved fairness arguments are swapped upon recursion.
func (bind Ings) Unify(x, y X) bool {
	x = bind.Resolve(x)
	y = bind.Resolve(y)

	if x.IsVariable() && y.IsVariable() && x.Atom.Var.Name == y.Atom.Var.Name {
		return true
	}

	switch {
	case x.IsVariable():
		return bind.SafeBind(x, y)

	case y.IsVariable():
		return bind.SafeBind(y, x)

	case x.IsPair() && y.IsPair():

		unifiedCars := bind.Unify(y.Pair.Car, x.Pair.Car)
		if unifiedCars {
			return bind.Unify(y.Pair.Cdr, x.Pair.Cdr)
		}
		return false
	}
	return x.Equal(y)
}

// =============================================================================

// Reify ...
func (bind Ings) Reify(v V) X {
	x := bind.Walk(v)
	r := New().reify(x)
	return r.Walk(x)
}

func (bind Ings) reify(v X) Ings {
	x := bind.Resolve(v)

	switch {
	case x.IsVariable() && !isReifiedName(x.Atom.Var.Name):
		bind.Bind(x, bind.V()) // bind x to new fresh var
	case x.IsPair():
		bind = bind.reify(x.Pair.Car)
		bind = bind.reify(x.Pair.Cdr)
	}
	return bind
}

// =============================================================================
