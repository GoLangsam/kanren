package bind

// bound represents the mapping of Variables to eXpressions/terms.
type bound map[V]X

func newBound() bound {
	return make(map[V]X)
}

func (bound bound) Clone() bound {
	clone := newBound()
	for k, v := range bound {
		clone[k] = v
	}
	return clone
}

// Bind binds x to v, so v is bound to x.
// Thus, (v . x) resembles a substitution pair.
//
// Bind is a noOp if v or x are nil or v is not a Variable.
//
// Note: Bind does not attempt to avoid circular bindings.
// Use Occurs to check beforehand.
func (bound bound) Bind(v V, x X) {
	if v.IsVariable() && x != nil {
		bound[v] = x
	}
	return
}

// Drop makes v unbound, reports whether v was bound, and
// returns the expression (if any) v was previously bound with.
func (bound bound) Drop(v V) (x X, wasBound bool) {
	x, wasBound = bound[v]
	if wasBound {
		delete(bound, v)
	}
	return
}

// Bound returns the expression to which v is bound, if any.
//
// This expression shall substitute the variable - so to say,
// which shall thus become substituted by this eXpression, its 'value' - so to say.
func (bound bound) Bound(v V) (value X, isBound bool) {
	value, isBound = bound[v]
	return
}

// =============================================================================

// Ings represents bindings (or "substitutions" or ""):
// any logic variable may be bound to some symbolic expression
// representing its current value.
//
// Use as `bind.Ings` (pun intended).
//
// The zero value is not useful - initialize with `bind.New()`.
type Ings struct {
	bound
}

// New creates fresh and empty mapping of/for bind.Ings and returns a pointer.
func New() *Ings {
	return &Ings{
		bound: newBound(),
	}
}

func (bind *Ings) Clone() *Ings {
	return &Ings{
		bound: bind.bound.Clone(),
	}
}

// =============================================================================

// Resolve the eXpression
// by chasing along the bindings recurring
// down to the first non-Variable eXpression
// or down to the first unbound eXpression
func (bind *Ings) Resolve(v X) X {
	if !v.IsVariable() {
		return v
	}
	x, isBound := bind.Bound(v)
	if !isBound {
		return v
	}
	return bind.Resolve(x)
}

// Walk ... some call it `walkstar` or `walk*`
func (bind *Ings) Walk(x X) X {
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
func (bind *Ings) Occurs(v V, x X) bool {
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
func (bind *Ings) exts(v V, x X) bool {
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
func (bind *Ings) Unify(x, y X) bool {
	x = bind.Resolve(x)
	y = bind.Resolve(y)

	vx, isXxVariable := x.AsVariable()
	vy, isYyVariable := y.AsVariable()

	if isXxVariable && isYyVariable && vx.Equal(vy) {
		return true
	}

	if isXxVariable {
		return bind.exts(x, y)
	}
	if isYyVariable {
		return bind.exts(y, x)
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
